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
 * Constants for Parks-Miller PRNG.
 ***/

#ifndef _PARKS_MILLER_PRNG_CONST_H_
#define _PARKS_MILLER_PRNG_CONST_H_

//constants for urng
#define IA 16807    		    // a
#define IM 2147483647 		    // m
#define AM (1.0f/(float)IM)         // 1/m - To calculate floating point result
#define IQ 127773 
#define IR 2836
#define NTAB 16
#define NDIV (1 + (IM - 1)/ NTAB)
#define EPS 1.2e-7
#define RMAX (1.0f - EPS)

#define PRNG_CHANNELS    256
#define MAX_NTAB         (PRNG_CHANNELS*NTAB)

//mathematical constants
#define PI   3.14159265359

//arbitrary mask to xor with seed. required to avoid seed to rng being zero. 
#define MASK 0x2F24EA81      

//random varible type. currently uniform and gaussion are supported.
enum RV_TYPE
  {
    RV_UNIFORM,
    RV_GAUSSIAN
  };

#endif
