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

#include "PipeProducerConsumerKernels.hpp"
#include <cmath>

int 
PIPE_PCK::setupPIPE_PCK()
{
  localDevHist = (cl_int *)malloc(MAX_HIST_BINS*sizeof(cl_int));
  CHECK_ALLOCATION(localDevHist, "failed to allocate memory (devHist)");

  memset(localDevHist,0,MAX_HIST_BINS*sizeof(cl_int));

  cpuHist = (cl_int *)malloc(MAX_HIST_BINS*sizeof(cl_int));
  CHECK_ALLOCATION(cpuHist, "failed to allocate memory (cpuHist)");

  memset(cpuHist,0,MAX_HIST_BINS*sizeof(cl_int));

  return SDK_SUCCESS;
}

int
PIPE_PCK::setupCL()
{
    cl_int status = CL_SUCCESS;
    cl_device_type dType;

    if(sampleArgs->deviceType.compare("cpu") == 0)
    {
        dType = CL_DEVICE_TYPE_CPU;
    }
    else //deviceType = "gpu"
    {
        dType = CL_DEVICE_TYPE_GPU;
        if(sampleArgs->isThereGPU() == false)
        {
            std::cout << "GPU not found. Falling back to CPU device" << std::endl;
            dType = CL_DEVICE_TYPE_CPU;
        }
    }

    /*
     * Have a look at the available platforms and pick either
     * the AMD one if available or a reasonable default.
     */
    cl_platform_id platform = NULL;
    int retValue = getPlatform(platform, sampleArgs->platformId,
                               sampleArgs->isPlatformEnabled());
    CHECK_ERROR(retValue, SDK_SUCCESS, "sampleCommon::getPlatform() failed");

    // Display available devices.
    retValue = displayDevices(platform, dType);
    CHECK_ERROR(retValue, SDK_SUCCESS, "displayDevices() failed");

    // If we could find our platform, use it. Otherwise use just available platform.

    cl_context_properties cps[3] =
    {
        CL_CONTEXT_PLATFORM,
        (cl_context_properties)platform,
        0
    };

    context = clCreateContextFromType(
                  cps,
                  dType,
                  NULL,
                  NULL,
                  &status);
    CHECK_OPENCL_ERROR( status, "clCreateContextFromType failed.");

    // getting device on which to run the sample
    status = getDevices(context, &devices, sampleArgs->deviceId,
                        sampleArgs->isDeviceIdEnabled());
    CHECK_ERROR(status, SDK_SUCCESS, "getDevices() failed");

	status = deviceInfo.setDeviceInfo(devices[sampleArgs->deviceId]);
    CHECK_OPENCL_ERROR(status, "deviceInfo.setDeviceInfo failed");

	//Check OpenCL 2.x compliance
	bool checkOCLversion = deviceInfo.checkOpenCL2_XCompatibility();

	if (!checkOCLversion)
	{
		OPENCL_EXPECTED_ERROR("Unsupported device! Required CL_DEVICE_OPENCL_C_VERSION 2.0 or higher");
	}

    {
      //The block is to move the declaration of prop closer to its use
      cl_queue_properties prop[] = {0};
      for (int i = 0; i < MAX_COMMAND_QUEUE; ++i)
	{
	  commandQueue[i] = clCreateCommandQueueWithProperties
	    (context,
	     devices[sampleArgs->deviceId],
	     prop,
	     &status);
	  CHECK_OPENCL_ERROR(status, "clCreateCommandQueue failed.");
	}
    }
    
    //Set device info of given cl_device_id
    retValue = deviceInfo.setDeviceInfo(devices[sampleArgs->deviceId]);
    CHECK_ERROR(retValue, SDK_SUCCESS, "SDKDeviceInfo::setDeviceInfo() failed");

    //Create and initialize memory objects
    rngPipe = clCreatePipe(context,
			   CL_MEM_READ_WRITE,
			   szPipePkt,
			   szPipe,
			   NULL,
			   &status);
    CHECK_OPENCL_ERROR(status, "clCreatePipe failed.");

    devHist = clCreateBuffer(context,
			     CL_MEM_READ_WRITE|CL_MEM_COPY_HOST_PTR,
			     MAX_HIST_BINS*sizeof(cl_int),
			     localDevHist,
			     &status);
    CHECK_OPENCL_ERROR(status, "clCreateBuffer failed.");      

    //create a CL program using the kernel source
    buildProgramData buildData;
    buildData.kernelName 
      = std::string("PipeProducerConsumerKernels_Kernels.cl");
    buildData.devices 
      = devices;
    buildData.deviceId 
      = sampleArgs->deviceId;
    buildData.flagsStr 
      = std::string("");

    if(sampleArgs->isLoadBinaryEnabled())
      {
	buildData.binaryName = std::string(sampleArgs->loadBinary.c_str());
      }
    
    if(sampleArgs->isComplierFlagsSpecified())
      {
        buildData.flagsFileName = std::string(sampleArgs->flags.c_str());
      }

    retValue = buildOpenCLProgram(program, context, buildData);
    CHECK_ERROR(retValue, SDK_SUCCESS, "buildOpenCLProgram() failed");

    // producer kernel
    produceKernel = clCreateKernel(
				   program,
				   "pipe_producer",
				   &status);
    CHECK_OPENCL_ERROR(status, "clCreateKernel failed(produceKernel).");

    status =  kernelInfo.setKernelWorkGroupInfo(produceKernel,
              devices[sampleArgs->deviceId]);
    CHECK_ERROR(status, SDK_SUCCESS, "setKernelWorkGroupInfo() failed");

    // consumer kernel
    consumeKernel = clCreateKernel(
				   program,
				   "pipe_consumer",
				   &status);
    CHECK_OPENCL_ERROR(status, "clCreateKernel failed(consumeKernel).");

    status =  kernelInfo.setKernelWorkGroupInfo(consumeKernel,
              devices[sampleArgs->deviceId]);
    CHECK_ERROR(status, SDK_SUCCESS, "setKernelWorkGroupInfo() failed");

    return SDK_SUCCESS;
}

int
PIPE_PCK::runCLKernels()
{
    cl_int status;

    // Set appropriate arguments to the kernel
    status = clSetKernelArg(produceKernel, 
    			    0, 
    			    sizeof(cl_mem),
    			    (void *)(&rngPipe));

    CHECK_OPENCL_ERROR(status, "clSetKernelArg failed.(rngPipe)");

    status = clSetKernelArg(produceKernel, 
    			    1, 
    			    sizeof(cl_int),
    			    (void *)(&pipePktPerThread));

    CHECK_OPENCL_ERROR(status, "clSetKernelArg failed.(pipePktPerThread)");

    status = clSetKernelArg(produceKernel, 
    			    2, 
    			    sizeof(cl_int),
    			    (void *)(&seed));

    CHECK_OPENCL_ERROR(status, "clSetKernelArg failed.(seed)");

    status = clSetKernelArg(produceKernel, 
    			    3, 
    			    sizeof(cl_int),
    			    (void *)(&rngType));
    CHECK_OPENCL_ERROR(status, "clSetKernelArg failed.(rngType)");

    // Enqueue both the kernels.
    size_t globalThreads[] = {producerGlobalSize};
    size_t localThreads[]  = {producerGroupSize};

    cl_event produceEvt;
    status = clEnqueueNDRangeKernel(
                 commandQueue[0],
                 produceKernel,
                 1,
                 NULL,
                 globalThreads,
                 localThreads,
                 0,
                 NULL,
                 &produceEvt);
    CHECK_OPENCL_ERROR(status, "clEnqueueNDRangeKernel failed.");

    /* 
       launch consumer kernel only after producer has finished.
       This is done to avoid concurrent kernels execution as the
       memory consistency of pipe is guaranteed only across 
       synchronization points.
    */

    status = clWaitForEvents(1,&produceEvt);
    CHECK_OPENCL_ERROR(status, "clWaitForEvents(produceEvt) failed.");

    status = clSetKernelArg(consumeKernel, 
			    0, 
			    sizeof(cl_mem),
			    (void *)(&rngPipe));
    CHECK_OPENCL_ERROR(status, "clSetKernelArg failed.(rngPipe)");

    status = clSetKernelArg(consumeKernel, 
			    1, 
			    sizeof(cl_mem),
			    (void *)(&devHist));
    CHECK_OPENCL_ERROR(status, "clSetKernelArg failed.(devHist)");

    status = clSetKernelArg(consumeKernel, 
			    2, 
			    sizeof(cl_float),
			    (void *)(&histMin));
    CHECK_OPENCL_ERROR(status, "clSetKernelArg failed.(histMin)");

    status = clSetKernelArg(consumeKernel, 
			    3, 
			    sizeof(cl_float),
			    (void *)(&histMax));
    CHECK_OPENCL_ERROR(status, "clSetKernelArg failed.(histMax)");

    globalThreads[0] = consumerGlobalSize;
    localThreads[0]  = consumerGroupSize;

    cl_event consumeEvt;
    status = clEnqueueNDRangeKernel(
                 commandQueue[1],
                 consumeKernel,
                 1,
                 NULL,
                 globalThreads,
                 localThreads,
                 0,
                 NULL,
                 &consumeEvt);
    CHECK_OPENCL_ERROR(status, "clEnqueueNDRangeKernel failed.");

    status = clFlush(commandQueue[0]);
    CHECK_OPENCL_ERROR(status, "clFlush failed(0).");

    status = clFlush(commandQueue[1]);
    CHECK_OPENCL_ERROR(status, "clFlush failed(1).");

    //wait for kernels to finish
    status = clFinish(commandQueue[0]);
    CHECK_OPENCL_ERROR(status, "clFinish failed(0).");

    status = clFinish(commandQueue[1]);
    CHECK_OPENCL_ERROR(status, "clFinish failed(1).");


    //copy the data back to host buffer
    cl_event readEvt;
    status = clEnqueueReadBuffer(commandQueue[1],
				 devHist,
				 CL_TRUE,
				 0,
				 (MAX_HIST_BINS)*sizeof(cl_int),
				 (void *)localDevHist,
				 0,
				 NULL,
				 &readEvt);
    CHECK_OPENCL_ERROR(status, "clEnqueueReadBuffer failed.");

    return SDK_SUCCESS;
}

void
PIPE_PCK::sampleCPUReference()
{
  PM_PRNG   pmPRNG;
  cl_int2   irn[PRNG_CHANNELS];
  cl_float2 frn[PRNG_CHANNELS];
  cl_float2 grn[PRNG_CHANNELS];

  float     binWidth;

  // Initialize the prng
  pmPRNG.rngInit(seed);

  // Put starting values for each channel 
  for (int ch = 0; ch < PRNG_CHANNELS; ++ch)
    {
      irn[ch].x = irn[ch].y = (ch +1)*(ch +1);
    }

  //compute binWidth
  binWidth = (histMax - histMin)/(float)(MAX_HIST_BINS);

  for(int pkt =0; pkt < pipePktPerThread; ++pkt)
    {
      // Generate random numbers
      for (int ch = 0; ch < PRNG_CHANNELS; ++ch)
	{
	  irn[ch].x = pmPRNG.rngPM(irn[ch].y, ch);      
	  irn[ch].y = pmPRNG.rngPM(irn[ch].x, ch);      

	  frn[ch].x = irn[ch].x *AM;
	  if(frn[ch].x > RMAX) 
	    frn[ch].x = (cl_float)RMAX;
	  
	  frn[ch].y = irn[ch].y *AM;
	  if(frn[ch].y > RMAX) 
	    frn[ch].y = (cl_float)RMAX;
	  
	  if(rngType == RV_GAUSSIAN)
	    {
	      grn[ch] = boxMuller(frn[ch]);
	    }
	  else
	    {
	      grn[ch] = frn[ch];
	    }

	}

      // host side histogram
      for (int ch = 0; ch < PRNG_CHANNELS; ++ch)
	{
	  float rmin      = histMin;
	  float rmax      = rmin + binWidth;
	  int   found     = 0;

	  for(int bindex = 0; (bindex < MAX_HIST_BINS) && (found != 2); bindex++)
	    {

	      if ((grn[ch].x >= rmin) && (grn[ch].x < rmax))
		{
		  found += 1;
		  cpuHist[bindex] += 1;
		}

	      if ((grn[ch].y >= rmin) && (grn[ch].y < rmax))
		{
		  found += 1;
		  cpuHist[bindex] += 1;
		}
	      
	      rmin = rmax;
	      rmax = rmin + binWidth;
	    }
	}      
    }
}


int
PIPE_PCK::initialize()
{
    cl_int status = 0;
    // Call base class Initialize to get default configuration
    if(sampleArgs->initialize() != SDK_SUCCESS)
    {
        return SDK_FAILURE;
    }

    //iteration option
    Option* iteration_option = new Option;
    CHECK_ALLOCATION(iteration_option, "Memory Allocation error.\n");

    iteration_option->_sVersion = "i";
    iteration_option->_lVersion = "iterations";
    iteration_option->_description = "Number of iterations to execute kernel";
    iteration_option->_type = CA_ARG_INT;
    iteration_option->_value = &iterations;

    sampleArgs->AddOption(iteration_option);

    delete iteration_option;

    //pipe size option
    Option* sz_pipe_option = new Option;
    CHECK_ALLOCATION(sz_pipe_option, "Memory Allocation error.\n");

    sz_pipe_option->_sVersion = "s";
    sz_pipe_option->_lVersion = "pipeSize";
    sz_pipe_option->_description = "Pipe size";
    sz_pipe_option->_type = CA_ARG_INT;
    sz_pipe_option->_value = &szPipe;

    sampleArgs->AddOption(sz_pipe_option);

    delete sz_pipe_option;

    //pipe size option
    Option* seed_option = new Option;
    CHECK_ALLOCATION(seed_option, "Memory Allocation error.\n");

    seed_option->_sVersion = "b";
    seed_option->_lVersion = "seed";
    seed_option->_description = "Seed for random number generator";
    seed_option->_type = CA_ARG_INT;
    seed_option->_value = &seed;

    sampleArgs->AddOption(seed_option);

    delete seed_option;
    
    return SDK_SUCCESS;
}


int
PIPE_PCK::setup()
{
    cl_int status = 0;


    if(szPipe % PRNG_CHANNELS)
      szPipe = (szPipe/PRNG_CHANNELS)*PRNG_CHANNELS + PRNG_CHANNELS;

    consumerGlobalSize = szPipe;
    pipePktPerThread   = szPipe/PRNG_CHANNELS;

    // create and initialize timers
    int timer = sampleTimer->createTimer();
    sampleTimer->resetTimer(timer);
    sampleTimer->startTimer(timer);

    if(setupPIPE_PCK() != SDK_SUCCESS)
      {
	return SDK_FAILURE;
      }

    if(setupCL() != SDK_SUCCESS)
      {
        return SDK_FAILURE;
      }

    sampleTimer->stopTimer(timer);
    // Compute setup time
    setupTime = (double)(sampleTimer->readTimer(timer));

    return SDK_SUCCESS;
}

int
PIPE_PCK::run()
{
    cl_int status = 0;

    // warm up run
    if(runCLKernels() != SDK_SUCCESS)
        {
            return SDK_FAILURE;
        }

    if(!sampleArgs->quiet)
      {
	std::cout << "Executing kernel for " << iterations
		  << " iterations" <<std::endl;
	std::cout << "-------------------------------------------" << std::endl;
      }

    // create and initialize timers
    int timer = sampleTimer->createTimer();
    sampleTimer->resetTimer(timer);
    sampleTimer->startTimer(timer);

    for(int i = 0; i < iterations; i++)
    {
        // Set kernel arguments and run kernel
        if(runCLKernels() != SDK_SUCCESS)
        {
            return SDK_FAILURE;
        }
    }

    sampleTimer->stopTimer(timer);

    // Compute kernel time
    kernelTime = (double)(sampleTimer->readTimer(timer)) / iterations;

    return SDK_SUCCESS;
}

int
PIPE_PCK::cleanup()
{

    // Releases OpenCL resources (Context, Memory etc.
    cl_int status;

    status = clReleaseKernel(produceKernel);
    CHECK_OPENCL_ERROR(status, "clReleaseKernel failed.");

    status = clReleaseKernel(consumeKernel);
    CHECK_OPENCL_ERROR(status, "clReleaseKernel failed.");

    status = clReleaseProgram(program);
    CHECK_OPENCL_ERROR(status, "clReleaseProgram failed.");

    status = clReleaseMemObject(rngPipe);
    CHECK_OPENCL_ERROR(status, "clReleaseMemObject failed.");

    status = clReleaseMemObject(devHist);
    CHECK_OPENCL_ERROR(status, "clReleaseMemObject failed.");

    for (int i = 0; i < MAX_COMMAND_QUEUE; ++i)
      {
	status = clReleaseCommandQueue(commandQueue[i]);
	CHECK_OPENCL_ERROR(status, "clReleaseCommandQueue failed.");
      }

    status = clReleaseContext(context);
    CHECK_OPENCL_ERROR(status, "clReleaseContext failed.");

    // release program resources (input memory etc.)
    if(localDevHist)
      free(localDevHist);

    if(cpuHist)
      free(cpuHist);

    return SDK_SUCCESS;
}

int
PIPE_PCK::verifyResults()
{
    if(sampleArgs->verify)
    {
      //Find the tolerance limit
      float fTol = (float)(CONSUMER_GLOBAL_SIZE)*(float)(COMP_TOL)/(float)100.0;
      int   iTol = (int)(fTol);
      if (iTol == 0)
	iTol = 1;
      
      //CPU side histogram computation
      sampleCPUReference();

      //Compare
      for(int bin = 0; bin < MAX_HIST_BINS; ++bin)
	{
	  int diff = localDevHist[bin] - cpuHist[bin];

	  if (diff < 0)
	    diff = -diff;

	  if(diff > iTol)
	    {
	      std::cout << "Failed! \n" << std::endl;
	      return SDK_SUCCESS;
	    }
	}

      std::cout << "Passed! \n" << std::endl;
    }
    return SDK_SUCCESS;
}

void
PIPE_PCK::printStats()
{
    std::string strArray[3] =
    {
      "Total Pipe Packets",
      "Setup Time(sec)",
      "(Kernel + Transfer)Time(sec)"
    };
    std::string stats[3];

    sampleTimer->totalTime = setupTime + kernelTime;

    stats[0] = toString(szPipe,std::dec);
    stats[1] = toString(setupTime, std::dec);
    stats[2] = toString(kernelTime, std::dec);

    if(sampleArgs->timing)
    {
        printStatistics(strArray, stats, 3);
    }

}

int
main(int argc, char * argv[])
{
    cl_int   status = 0;
    PIPE_PCK clPIPE_PCK;

    if(clPIPE_PCK.initialize() != SDK_SUCCESS)
    {
        return SDK_FAILURE;
    }

    if(clPIPE_PCK.sampleArgs->parseCommandLine(argc, argv) != SDK_SUCCESS)
    {
        return SDK_FAILURE;
    }

    if(clPIPE_PCK.setup() != SDK_SUCCESS)
    {
        return SDK_FAILURE;
    }

    if(clPIPE_PCK.run() != SDK_SUCCESS)
    {
        return SDK_FAILURE;
    }

    if(clPIPE_PCK.verifyResults() != SDK_SUCCESS)
    {
        return SDK_FAILURE;
    }

    if(clPIPE_PCK.cleanup() != SDK_SUCCESS)
    {
        return SDK_FAILURE;
    }

    clPIPE_PCK.printStats();

    return SDK_SUCCESS;
}
