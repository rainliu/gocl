/**********************************************************************
Copyright ©2014 Advanced Micro Devices, Inc. All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or
 other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
 DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
 OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
 NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
********************************************************************/

//constants for urng
#define IA (16807)    		    // a
#define IM (2147483647) 		    // m
#define AM (1.0f/(float)IM)         // 1/m - To calculate floating point result
#define IQ (127773) 
#define IR (2836)
#define NTAB (16)
#define NDIV (1 + (IM - 1)/ NTAB)
#define EPS 1.2e-7
#define RMAX (1.0f - EPS)

#define PRNG_CHANNELS        256

#define MAX_NTAB ((PRNG_CHANNELS)*(NTAB))

//mathematical constants
#define PI  3.14159265359 

//arbitrary mask to xor with seed. required to avoid seed to rng being zero. 
#define MASK 0x2F24EA81      

//maximum bins in histogram. this should be multiple of wave-front size.
#define MAX_HIST_BINS   256

//random varible type. currently uniform and gaussion are supported.
enum RV_TYPE
  {
    RV_UNIFORM,
    RV_GAUSSIAN
  };

/***
 * rng_div:
 * Only god know why normal "/" is not working on GPU
 ***/
int rng_div(int dv, int dr)
{
  int k = 0;

  while(dv > dr)
    {
      k++;
      dv -= dr;
    }

  return k;
}

/***
 * rng_init:
 * takes [seed] and uses Parks-Miller random number generation to fill up
 * schuffle [iv].
 ***/
void  rng_init(int seed, __local int *iv)
{
  int j;
  int k;
  int t;

  //xor seed with the mask to avoid its value being zero.
  seed ^= MASK;

  //fill the shuffling buffer.
  iv[0] = seed;

  for(j = 1; j < MAX_NTAB; j++)			
    {
      //k = seed/IQ;
      k = rng_div(seed,IQ);

      seed = IA * (seed - k * IQ) - IR * k;
      
      if(seed < 0)
	seed += IM;

      iv[j] = seed;
    }
}

/****
 * rng_pm:
 * takes previous rv [prn], and schuffle [iv] and outputs next rv.
 ****/
int rng_pm(int prn, __local int *iv)
{
  int j;
  int k;
  int nrn;

  j   = prn/NDIV;
  nrn = iv[j];

  //k = prn / IQ;
  k = rng_div(prn,IQ);

  prn = IA * (prn - k * IQ) - IR * k;
  
  if(prn < 0)
    prn += IM;
  
  iv[j] = prn;

  return nrn;
}

/***
 * box_muller:
 * takes two uniform rv [unifrom] between [0..1] and outputs two gaussion 
 * rv ~ N(0,1).
 ***/

float2 box_muller(float2 uniform)   
{   
  float r = sqrt(-2 * log(uniform.x));   
  float theta = 2 * PI * uniform.y;   
  return (float2)(r * sin(theta), r * cos(theta));   
}  


/**
 * pipe_producer:
 * produces [pkt_per_thread] random numbers per thread. takes [seed] as seed to 
 * random number generator. [rng_type] defines the type
 * of random number. pushes random numbers to pipe [rng_pipe] to be used by consumer 
 * kernel.
 **/
__kernel void pipe_producer(__write_only pipe float2 rng_pipe, 
			    int   pkt_per_thread,
			    int   seed,
			    int   rng_type)
{
  float2 ufrn;
  float2 gfrn;
  int2   irn;
  int    iter;
  int    lflag;

  __local  int    iv[MAX_NTAB];


  int    lid  = get_local_id(0);
  int    szgr = get_local_size(0);
  
  //initialize random number generator.
  if (lid == 0)
    {
      rng_init(seed, iv);
    }
  work_group_barrier(CLK_LOCAL_MEM_FENCE|CLK_GLOBAL_MEM_FENCE);

  iter   = 0;
  irn.x  = (lid +1)*(lid +1);
  irn.y  = (lid +1)*(lid +1);

  while(iter < pkt_per_thread)
    {
	  //reserve space in pipe for writing random numbers.
      reserve_id_t rid = work_group_reserve_write_pipe(rng_pipe, szgr);
      
     
      if(is_valid_reserve_id(rid))
	{
	  //draw two unifromly distributed rv.
          irn.x = rng_pm(irn.y, (iv + lid*NTAB));
	  work_group_barrier(CLK_LOCAL_MEM_FENCE);

	  irn.y = rng_pm(irn.x, (iv + lid*NTAB));
	  work_group_barrier(CLK_LOCAL_MEM_FENCE);

	  ufrn.x = (float)(irn.x)*AM;
	  if (ufrn.x > RMAX)
	    ufrn.x = RMAX;

	  ufrn.y = (float)(irn.y)*AM;
	  if (ufrn.y > RMAX)
	    ufrn.y = RMAX;

	  //If gaussian distribution is requested, apply box muller transform.
	  if(rng_type == RV_GAUSSIAN)
	    {
	      gfrn = box_muller(ufrn);
	    }
	  else
	    {
	      gfrn = ufrn;
	    }

	  //write into pipe.
	  write_pipe(rng_pipe,rid,lid, &gfrn);
	  work_group_commit_write_pipe(rng_pipe, rid);
	}
      
      work_group_barrier(CLK_GLOBAL_MEM_FENCE);

      iter += 1;
    }
}

/**
 * pipe_consumer:
 * reads random variables from [rng_pipe] and produces their histogram in [hist]. the
 * histogram is generated for rv between hist_min and hist_max and any other rv out of
 * this range is ignored. the number of bins in histogram are given by MAX_HIST_BINS.
 **/

__kernel void pipe_consumer(__read_only pipe float2          rng_pipe,
			    __global         int             *hist,
			                     float           hist_min,
			                     float           hist_max)
{
  int          bindex;
  int          found, freq;
  int          lap, laps;
  
  float2       rn;
  float        bin_width;
  float        rmin,rmax;
  
  __local   int  lhist[MAX_HIST_BINS];
  
  int lid  = get_local_id(0);
  int gid  = get_global_id(0);
  int grid = get_group_id(0);
  int szgr = get_local_size(0);
  int szgl = get_global_size(0);

  //reset hist
  laps          = (MAX_HIST_BINS) /szgr;

  if(grid == 0)
    {
      for(lap = 0; lap < laps; ++lap)
	{
	  bindex = lid + lap*szgr;
	  hist[bindex] = 0;
	}
    }

  work_group_barrier(CLK_GLOBAL_MEM_FENCE);

  //reserve pipe for reading
  reserve_id_t rid = work_group_reserve_read_pipe(rng_pipe, szgr);
      
      
  if(is_valid_reserve_id(rid))
    {
      //read random number from the pipe.
      read_pipe(rng_pipe,rid,lid, &rn);
      work_group_commit_read_pipe(rng_pipe, rid);
    }
  
  //each work-group generates local histogram
  bin_width = (hist_max - hist_min)/(float)(MAX_HIST_BINS);

  rmin      = hist_min;
  rmax      = rmin + bin_width;
  found     = 0;

  for(bindex = 0; bindex < MAX_HIST_BINS; bindex++)
    {
      if ((rn.x >= rmin) && (rn.x < rmax))
	{
	  found += 1;
	}

      if ((rn.y >= rmin) && (rn.y < rmax))
	{
	  found += 1;
	}

      work_group_barrier(CLK_LOCAL_MEM_FENCE);
      
      freq  = work_group_reduce_add(found);
      if (lid == 0)
	{
	  lhist[bindex] = freq;
	}

      work_group_barrier(CLK_LOCAL_MEM_FENCE);

      rmin  = rmax;
      rmax  = rmin + bin_width;
      found = 0;
    }
  
  work_group_barrier(CLK_LOCAL_MEM_FENCE);

  //add all local histograms to global histogram
  for(lap = 0; lap < laps; ++lap)
    {
      bindex = lid + lap*szgr;
      atomic_add((volatile __global int *)(hist + bindex), lhist[bindex]);
    }

  work_group_barrier(CLK_GLOBAL_MEM_FENCE);
}
