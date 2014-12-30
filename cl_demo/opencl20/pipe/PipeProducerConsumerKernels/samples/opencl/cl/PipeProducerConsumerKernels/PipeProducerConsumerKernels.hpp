/**********************************************************************
Copyright ©2014 Advanced Micro Devices, Inc. All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

•   Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
•   Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or
 other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
 DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
 OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
 NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
********************************************************************/

#ifndef _PIPE_PRODUCER_CONSUMER_H_
#define _PIPE_PRODUCER_CONSUMER_H_

#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
#include <string.h>
#include "CLUtil.hpp"

#include "ParksMillerPRNG.hpp"

#define SAMPLE_VERSION				"AMD-APP-SDK-v2.9.1.1"
#define OCL_COMIPLER_FLAGS			"PipeProducerConsumerKernels_OclFlags.txt"

#define MAX_COMMAND_QUEUE           2
#define PIPE_PKT_PER_THREAD         8

#define PRODUCER_GLOBAL_SIZE        (PRNG_CHANNELS)
#define PRODUCER_GROUP_SIZE         256

#define CONSUMER_GLOBAL_SIZE        (PIPE_PKT_PER_THREAD*PRODUCER_GLOBAL_SIZE)
#define CONSUMER_GROUP_SIZE         256

#define PIPE_SIZE                   (CONSUMER_GLOBAL_SIZE)
 
#define SEED                        0x53AA45

#define COMP_TOL                    1.0 

//maximum bins in histogram. this should be multiple of wave-front size.
#define MAX_HIST_BINS               256

using namespace appsdk;

class PIPE_PCK
{
  /**< time taken to setup OpenCL resources and building kernel */
  cl_double setupTime;                

  /**< time taken to run kernel and read result back */
  cl_double kernelTime;               

  /**< CL context */
  cl_context context;
                 
  /**< CL device list */
  cl_device_id *devices;              

  /**< CL command queue */
  cl_command_queue commandQueue[MAX_COMMAND_QUEUE];
      
  /**< CL program  */
  cl_program program;                 

  /**< CL kernel */
  cl_kernel produceKernel;
  cl_kernel consumeKernel;
                   
  /**< producer Work-group size */
  size_t producerGroupSize;                  

  /**< consumer Work-group size */
  size_t consumerGroupSize;                  

  /**< producer global size */
  size_t producerGlobalSize;                  

  /**< consumer global size */
  size_t consumerGlobalSize;                  

  /**< Number of iterations for kernel execution */
  int iterations;                     

  /**< Structure to store device information*/           
  SDKDeviceInfo deviceInfo;

  /**< Structure to store kernel related info */
  KernelWorkGroupInfo kernelInfo;

  /**< SDKTimer object */
  SDKTimer *sampleTimer;      

  /**< size of the pipe */
  cl_uint   szPipe;

  /**< size of the pipe packet */
  cl_uint  szPipePkt;

  /**< number of pipe batches */
  cl_int   pipePktPerThread;

  /**< pipe structure to connect two kernels */
  cl_mem   rngPipe;
  cl_mem   devHist;

  /**< seed for rng */
  cl_int   seed;

  /**< histogram size */
  cl_float histMin;
  cl_float histMax;

  /**< local histograms */
  cl_int   *localDevHist;
  cl_int   *cpuHist;

  /**< rng type */
  cl_int   rngType;
public:

  /**< CLCommand argument class */
  CLCommandArgs   *sampleArgs;   

  /**
   * Constructor
   * Initialize member variables
   */
  PIPE_PCK()
  {
    sampleArgs  = new CLCommandArgs();
    sampleTimer = new SDKTimer();

    sampleArgs->sampleVerStr = SAMPLE_VERSION;
    sampleArgs->flags        = OCL_COMIPLER_FLAGS;

    producerGroupSize  = PRODUCER_GROUP_SIZE;
    producerGlobalSize = PRODUCER_GLOBAL_SIZE;

    consumerGroupSize  = CONSUMER_GROUP_SIZE;
    consumerGlobalSize = CONSUMER_GLOBAL_SIZE;

    rngType            = RV_GAUSSIAN;

    iterations         = 1;
    seed               = SEED;

    if(rngType == RV_UNIFORM)
      {
	histMin            = 0.0;
	histMax            = 1.0;
      }
    else
      {
	histMin            = -10.0;
	histMax            =  10.0;
      }

    szPipe           = PIPE_SIZE;
    szPipePkt        = sizeof(cl_float2);
    pipePktPerThread = PIPE_PKT_PER_THREAD;
  }
  
  ~PIPE_PCK()
  {
  }
  
  /**
   * Allocate image memory and Load bitmap file
   * @return SDK_SUCCESS on success and SDK_FAILURE on failure
   */
  int setupPIPE_PCK();
  
  /**
   * OpenCL related initialisations.
   * Set up Context, Device list, Command Queue, Memory buffers
   * Build CL kernel program executable
   * @return SDK_SUCCESS on success and SDK_FAILURE on failure
   */
  int setupCL();
  
  /**
   * Set values for kernels' arguments, enqueue calls to the kernels
   * on to the command queue, wait till end of kernel execution.
   * Get kernel start and end time if timing is enabled
   * @return SDK_SUCCESS on success and SDK_FAILURE on failure
   */
  int runCLKernels();
  
  /**
   * Reference CPU implementation for performance comparison
   */
  void sampleCPUReference();
  
  /**
   * Override from SDKSample. Print sample stats.
   */
  void printStats();
  
  /**
   * Override from SDKSample. Initialize
   * command line parser, add custom options
   * @return SDK_SUCCESS on success and SDK_FAILURE on failure
   */
  int initialize();
  
  /**
   * Override from SDKSample, adjust width and height
   * of execution domain, perform all sample setup
   * @return SDK_SUCCESS on success and SDK_FAILURE on failure
   */
  int setup();
  
  /**
   * Override from SDKSample
   * Run OpenCL Sobel Filter
   * @return SDK_SUCCESS on success and SDK_FAILURE on failure
   */
  int run();
  
  /**
   * Override from SDKSample
   * Cleanup memory allocations
   * @return SDK_SUCCESS on success and SDK_FAILURE on failure
   */
  int cleanup();
  
  /**
   * Override from SDKSample
   * Verify against reference implementation
   * @return SDK_SUCCESS on success and SDK_FAILURE on failure
   */
  int verifyResults();
};

#endif 
