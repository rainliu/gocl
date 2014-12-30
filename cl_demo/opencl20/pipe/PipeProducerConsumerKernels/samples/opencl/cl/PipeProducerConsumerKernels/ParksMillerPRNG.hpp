/**********************************************************************
Copyright ©2014 Advanced Micro Devices, Inc. All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

   Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
   Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or
 other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
 DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
 OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
 NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
********************************************************************/

/***
 * Host side implementation of Parks Miller PRNG
 ***/

#ifndef _PARKS_MILLER_PRNG_H_
#define _PARKS_MILLER_PRNG_H_

#include <CL/cl.h>
#include "ParksMillerPRNGConst.hpp"

/***
 * BEGIN - CLASS PM_PRNG   
 ***/

class PM_PRNG
{
public:
  cl_int        iv[MAX_NTAB];
  cl_int        seed;

  /* constructor and destructor */
  PM_PRNG(cl_int _seed = 0);

  ~PM_PRNG();

  /* initialization */
  void rngInit(cl_int lSeed);

  /* generation */
  cl_int  rngPM(cl_int prn, cl_int ch);

};

PM_PRNG::PM_PRNG(int _seed)
{
  seed = _seed;
};

PM_PRNG::~PM_PRNG()
{
};

void PM_PRNG::rngInit(cl_int lSeed)
{
  cl_int j;
  cl_int k;

  //xor seed with the mask to avoid its value being zero.
  lSeed ^= MASK;

  //fill the shuffling buffer.
  iv[0] = lSeed;

  for(j = 1; j < MAX_NTAB; j++)			
    {
      k = lSeed/IQ;
      lSeed = IA * (lSeed - k * IQ) - IR * k;
      
      if(lSeed < 0)
	lSeed += IM;

      iv[j] = lSeed;
    }
}

cl_int PM_PRNG::rngPM(cl_int prn, cl_int ch)
{
  cl_int j;
  cl_int k;
  cl_int nrn;

  j   = prn/NDIV;
  nrn = iv[j + ch*NTAB];

  k = prn / IQ;
  prn = IA * (prn - k * IQ) - IR * k;
  
  if(prn < 0)
    prn += IM;
  
  iv[j +ch*NTAB] = prn;

  return nrn;
}

/***
 * END -  CLASS PM_PRNG   
 ***/

/* uniform to gaussian */
cl_float2 boxMuller(cl_float2 u)
{
  cl_float2 g;

  cl_float r = sqrt(-2 * log(u.x));   
  cl_float theta = (cl_float)(2.0 * PI * u.y);   
  
  g.x = (cl_float)(r * sin(theta));
  g.y = (cl_float)(r * cos(theta));

  return g;   
}

#endif
