// This program demo OpenCL 2.0 Pipe feature
package main

import (
	"fmt"
	"gocl/cl"
	"gocl/cl_demo/utils"
	"unsafe"
)

const (
	//constants for urng
	IA   = 16807                 // a
	IM   = 2147483647            // m
	AM   = 1.0 / cl.CL_float(IM) // 1/m - To calculate floating point result
	IQ   = 127773
	IR   = 2836
	NTAB = 16
	NDIV = (1 + (IM-1)/NTAB)
	EPS  = 1.2e-7
	RMAX = (1.0 - EPS)

	PRNG_CHANNELS = 256
	MAX_NTAB      = (PRNG_CHANNELS * NTAB)

	//mathematical constants
	PI = 3.14159265359

	//arbitrary mask to xor with seed. required to avoid seed to rng being zero.
	MASK = 0x2F24EA81
)

//random varible type. currently uniform and gaussion are supported.
const (
	RV_UNIFORM  cl.CL_int = 0
	RV_GAUSSIAN cl.CL_int = 1
)

const (
	MAX_COMMAND_QUEUE   = 2
	PIPE_PKT_PER_THREAD = 8

	PRODUCER_GLOBAL_SIZE = (PRNG_CHANNELS)
	PRODUCER_GROUP_SIZE  = 256

	CONSUMER_GLOBAL_SIZE = (PIPE_PKT_PER_THREAD * PRODUCER_GLOBAL_SIZE)
	CONSUMER_GROUP_SIZE  = 256

	PIPE_SIZE = (CONSUMER_GLOBAL_SIZE)

	SEED = 0x53AA45

	COMP_TOL = 1.0

	//maximum bins in histogram. this should be multiple of wave-front size.
	MAX_HIST_BINS = 256
)

func main() {
	// Use this to check the output of each API call
	var status cl.CL_int

	//-----------------------------------------------------
	// STEP 1: Discover and initialize the platforms
	//-----------------------------------------------------
	var numPlatforms cl.CL_uint
	var platforms []cl.CL_platform_id

	// Use clGetPlatformIDs() to retrieve the number of
	// platforms
	status = cl.CLGetPlatformIDs(0, nil, &numPlatforms)

	// Allocate enough space for each platform
	platforms = make([]cl.CL_platform_id, numPlatforms)

	// Fill in platforms with clGetPlatformIDs()
	status = cl.CLGetPlatformIDs(numPlatforms, platforms, nil)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLGetPlatformIDs")

	//-----------------------------------------------------
	// STEP 2: Discover and initialize the GPU devices
	//-----------------------------------------------------
	var numDevices cl.CL_uint
	var devices []cl.CL_device_id

	// Use clGetDeviceIDs() to retrieve the number of
	// devices present
	status = cl.CLGetDeviceIDs(platforms[0],
		cl.CL_DEVICE_TYPE_GPU,
		0,
		nil,
		&numDevices)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLGetDeviceIDs")

	// Allocate enough space for each device
	devices = make([]cl.CL_device_id, numDevices)

	// Fill in devices with clGetDeviceIDs()
	status = cl.CLGetDeviceIDs(platforms[0],
		cl.CL_DEVICE_TYPE_GPU,
		numDevices,
		devices,
		nil)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLGetDeviceIDs")

	//-----------------------------------------------------
	// STEP 3: Create a context
	//-----------------------------------------------------
	var context cl.CL_context

	// Create a context using clCreateContext() and
	// associate it with the devices
	context = cl.CLCreateContext(nil,
		numDevices,
		devices,
		nil,
		nil,
		&status)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLCreateContext")
	defer cl.CLReleaseContext(context)

	//-----------------------------------------------------
	// STEP 4: Create a command queue
	//-----------------------------------------------------
	var commandQueue [MAX_COMMAND_QUEUE]cl.CL_command_queue

	// Create a command queue using clCreateCommandQueueWithProperties(),
	// and associate it with the device you want to execute
	for i := 0; i < MAX_COMMAND_QUEUE; i++ {
		commandQueue[i] = cl.CLCreateCommandQueueWithProperties(context,
			devices[0],
			nil,
			&status)
		utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLCreateCommandQueueWithProperties")
		defer cl.CLReleaseCommandQueue(commandQueue[i])
	}

	//-----------------------------------------------------
	// STEP 5: Create device buffers
	//-----------------------------------------------------
	producerGroupSize := cl.CL_size_t(PRODUCER_GROUP_SIZE)
	producerGlobalSize := cl.CL_size_t(PRODUCER_GLOBAL_SIZE)

	consumerGroupSize := cl.CL_size_t(CONSUMER_GROUP_SIZE)
	consumerGlobalSize := cl.CL_size_t(CONSUMER_GLOBAL_SIZE)

	var samplePipePkt [2]cl.CL_float
	szPipe := cl.CL_uint(PIPE_SIZE)
	szPipePkt := cl.CL_uint(unsafe.Sizeof(samplePipePkt))
	if szPipe%PRNG_CHANNELS != 0 {
		szPipe = (szPipe/PRNG_CHANNELS)*PRNG_CHANNELS + PRNG_CHANNELS
	}
	consumerGlobalSize = cl.CL_size_t(szPipe)
	pipePktPerThread := cl.CL_int(szPipe) / PRNG_CHANNELS
	seed := cl.CL_int(SEED)
	rngType := cl.CL_int(RV_GAUSSIAN)
	var histMin cl.CL_float
	var histMax cl.CL_float
	if rngType == cl.CL_int(RV_UNIFORM) {
		histMin = 0.0
		histMax = 1.0
	} else {
		histMin = -10.0
		histMax = 10.0
	}

	localDevHist := make([]cl.CL_int, MAX_HIST_BINS)
	cpuHist := make([]cl.CL_int, MAX_HIST_BINS)

	//Create and initialize memory objects
	rngPipe := cl.CLCreatePipe(context,
		cl.CL_MEM_READ_WRITE,
		szPipePkt,
		szPipe,
		nil,
		&status)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clCreatePipe")

	devHist := cl.CLCreateBuffer(context,
		cl.CL_MEM_READ_WRITE|cl.CL_MEM_COPY_HOST_PTR,
		MAX_HIST_BINS*cl.CL_size_t(unsafe.Sizeof(localDevHist[0])),
		unsafe.Pointer(&localDevHist[0]),
		&status)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clCreateBuffer")

	//-----------------------------------------------------
	// STEP 6: Create and compile the program
	//-----------------------------------------------------
	programSource, programeSize := utils.Load_programsource("pipe.cl")

	// Create a program using clCreateProgramWithSource()
	program := cl.CLCreateProgramWithSource(context,
		1,
		programSource[:],
		programeSize[:],
		&status)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLCreateProgramWithSource")
	defer cl.CLReleaseProgram(program)

	// Build (compile) the program for the devices with
	// clBuildProgram()
	options := "-cl-std=CL2.0"
	status = cl.CLBuildProgram(program,
		numDevices,
		devices,
		[]byte(options),
		nil,
		nil)
	if status != cl.CL_SUCCESS {
		var program_log interface{}
		var log_size cl.CL_size_t

		/* Find size of log and print to std output */
		cl.CLGetProgramBuildInfo(program, devices[0], cl.CL_PROGRAM_BUILD_LOG,
			0, nil, &log_size)
		cl.CLGetProgramBuildInfo(program, devices[0], cl.CL_PROGRAM_BUILD_LOG,
			log_size, &program_log, nil)
		fmt.Printf("%s\n", program_log)
		return
	}
	//utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLBuildProgram")

	//-----------------------------------------------------
	// STEP 7: Create the kernel
	//-----------------------------------------------------
	// Use clCreateKernel() to create a kernel
	produceKernel := cl.CLCreateKernel(program, []byte("pipe_producer"), &status)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLCreateKernel")
	defer cl.CLReleaseKernel(produceKernel)

	consumeKernel := cl.CLCreateKernel(program, []byte("pipe_consumer"), &status)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLCreateKernel")
	defer cl.CLReleaseKernel(consumeKernel)

	//-----------------------------------------------------
	// STEP 8: Set the kernel arguments
	//-----------------------------------------------------
	// Associate the input and output buffers with the
	// kernel
	// using clSetKernelArg()
	// Set appropriate arguments to the kernel
	status = cl.CLSetKernelArg(produceKernel,
		0,
		cl.CL_size_t(unsafe.Sizeof(rngPipe)),
		unsafe.Pointer(&rngPipe))

	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clSetKernelArg(rngPipe)")

	status = cl.CLSetKernelArg(produceKernel,
		1,
		cl.CL_size_t(unsafe.Sizeof(pipePktPerThread)),
		unsafe.Pointer(&pipePktPerThread))
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clSetKernelArg(pipePktPerThread)")

	status = cl.CLSetKernelArg(produceKernel,
		2,
		cl.CL_size_t(unsafe.Sizeof(seed)),
		unsafe.Pointer(&seed))
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clSetKernelArg(seed)")

	status = cl.CLSetKernelArg(produceKernel,
		3,
		cl.CL_size_t(unsafe.Sizeof(rngType)),
		unsafe.Pointer(&rngType))
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clSetKernelArg(rngType)")

	//-----------------------------------------------------
	// STEP 9: Configure the work-item structure
	//-----------------------------------------------------
	// Define an index space (global work size) of work
	// items for
	// execution. A workgroup size (local work size) is not
	// required,
	// but can be used.
	// Enqueue both the kernels.
	var globalThreads = []cl.CL_size_t{producerGlobalSize}
	var localThreads = []cl.CL_size_t{producerGroupSize}

	//-----------------------------------------------------
	// STEP 10: Enqueue the kernel for execution
	//-----------------------------------------------------
	// Execute the kernel by using
	// clEnqueueNDRangeKernel().
	// 'globalWorkSize' is the 1D dimension of the
	// work-items
	var produceEvt [1]cl.CL_event
	status = cl.CLEnqueueNDRangeKernel(commandQueue[0],
		produceKernel,
		1,
		nil,
		globalThreads,
		localThreads,
		0,
		nil,
		&produceEvt[0])
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clEnqueueNDRangeKernel")

	/*
	   launch consumer kernel only after producer has finished.
	   This is done to avoid concurrent kernels execution as the
	   memory consistency of pipe is guaranteed only across
	   synchronization points.
	*/
	status = cl.CLWaitForEvents(1, produceEvt[:])
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clWaitForEvents(produceEvt)")

	//-----------------------------------------------------
	// STEP 8: Set the kernel arguments
	//-----------------------------------------------------
	// Associate the input and output buffers with the
	// kernel
	// using clSetKernelArg()
	// Set appropriate arguments to the kernel
	status = cl.CLSetKernelArg(consumeKernel,
		0,
		cl.CL_size_t(unsafe.Sizeof(rngPipe)),
		unsafe.Pointer(&rngPipe))
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clSetKernelArg(rngPipe)")

	status = cl.CLSetKernelArg(consumeKernel,
		1,
		cl.CL_size_t(unsafe.Sizeof(devHist)),
		unsafe.Pointer(&devHist))
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clSetKernelArg(devHist)")

	status = cl.CLSetKernelArg(consumeKernel,
		2,
		cl.CL_size_t(unsafe.Sizeof(histMin)),
		unsafe.Pointer(&histMin))
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clSetKernelArg(histMin)")

	status = cl.CLSetKernelArg(consumeKernel,
		3,
		cl.CL_size_t(unsafe.Sizeof(histMax)),
		unsafe.Pointer(&histMax))
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clSetKernelArg(histMax)")

	//-----------------------------------------------------
	// STEP 9: Configure the work-item structure
	//-----------------------------------------------------
	// Define an index space (global work size) of work
	// items for
	// execution. A workgroup size (local work size) is not
	// required,
	// but can be used.
	globalThreads[0] = consumerGlobalSize
	localThreads[0] = consumerGroupSize

	//-----------------------------------------------------
	// STEP 10: Enqueue the kernel for execution
	//-----------------------------------------------------
	// Execute the kernel by using
	// clEnqueueNDRangeKernel().
	// 'globalWorkSize' is the 1D dimension of the
	// work-items
	var consumeEvt [1]cl.CL_event
	status = cl.CLEnqueueNDRangeKernel(
		commandQueue[1],
		consumeKernel,
		1,
		nil,
		globalThreads,
		localThreads,
		0,
		nil,
		&consumeEvt[0])
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clEnqueueNDRangeKernel")

	status = cl.CLFlush(commandQueue[0])
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clFlush(0)")

	status = cl.CLFlush(commandQueue[1])
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clFlush(1)")

	//wait for kernels to finish
	status = cl.CLFinish(commandQueue[0])
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clFinish(0)")

	status = cl.CLFinish(commandQueue[1])
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clFinish(1)")

	//-----------------------------------------------------
	// STEP 11: Read the output buffer back to the host
	//-----------------------------------------------------
	// Use clEnqueueReadBuffer() to read the OpenCL output
	// buffer (bufferC)
	// to the host output array (C)
	//copy the data back to host buffer
	var readEvt cl.CL_event
	status = cl.CLEnqueueReadBuffer(commandQueue[1],
		devHist,
		cl.CL_TRUE,
		0,
		(MAX_HIST_BINS)*cl.CL_size_t(unsafe.Sizeof(localDevHist[0])),
		unsafe.Pointer(&localDevHist[0]),
		0,
		nil,
		&readEvt)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clEnqueueReadBuffer")

	//-----------------------------------------------------
	// STEP 12: Verify the results
	//-----------------------------------------------------
	//Find the tolerance limit
	fTol := (float32)(CONSUMER_GLOBAL_SIZE) * (float32)(COMP_TOL) / (float32)(100.0)
	iTol := (int)(fTol)
	if iTol == 0 {
		iTol = 1
	}

	//CPU side histogram computation
	CPUReference(seed, pipePktPerThread, rngType, cpuHist, histMax, histMin)

	//Compare
	for bin := 0; bin < MAX_HIST_BINS; bin++ {
		diff := int(localDevHist[bin] - cpuHist[bin])

		if diff < 0 {
			diff = -diff
		}
		if diff > iTol {
			println("Failed!")
			return
		}
	}

	println("Passed!")
}

func CPUReference(seed cl.CL_int,
	pipePktPerThread cl.CL_int,
	rngType cl.CL_int,
	cpuHist []cl.CL_int,
	histMax, histMin cl.CL_float) {
	var pmPRNG PM_PRNG
	var irn [PRNG_CHANNELS][2]cl.CL_int
	var frn [PRNG_CHANNELS][2]cl.CL_float
	var grn [PRNG_CHANNELS][2]cl.CL_float

	var binWidth cl.CL_float

	// Initialize the prng
	pmPRNG.rngInit(seed)

	// Put starting values for each channel
	for ch := cl.CL_int(0); ch < cl.CL_int(PRNG_CHANNELS); ch++ {
		irn[ch][0] = (ch + 1) * (ch + 1)
		irn[ch][1] = (ch + 1) * (ch + 1)
	}

	//compute binWidth
	binWidth = (histMax - histMin) / cl.CL_float(MAX_HIST_BINS)

	for pkt := cl.CL_int(0); pkt < pipePktPerThread; pkt++ {
		// Generate random numbers
		for ch := 0; ch < PRNG_CHANNELS; ch++ {
			irn[ch][0] = pmPRNG.rngPM(irn[ch][1], cl.CL_int(ch))
			irn[ch][1] = pmPRNG.rngPM(irn[ch][0], cl.CL_int(ch))

			frn[ch][0] = cl.CL_float(irn[ch][0]) * AM
			if frn[ch][0] > RMAX {
				frn[ch][0] = cl.CL_float(RMAX)
			}
			frn[ch][1] = cl.CL_float(irn[ch][1]) * AM
			if frn[ch][1] > RMAX {
				frn[ch][1] = cl.CL_float(RMAX)
			}

			if rngType == RV_GAUSSIAN {
				grn[ch] = boxMuller(frn[ch])
			} else {
				grn[ch] = frn[ch]
			}

		}

		// host side histogram
		for ch := 0; ch < PRNG_CHANNELS; ch++ {
			rmin := histMin
			rmax := rmin + binWidth
			found := 0

			for bindex := 0; (bindex < MAX_HIST_BINS) && (found != 2); bindex++ {
				if (grn[ch][0] >= rmin) && (grn[ch][0] < rmax) {
					found += 1
					cpuHist[bindex] += 1
				}

				if (grn[ch][1] >= rmin) && (grn[ch][1] < rmax) {
					found += 1
					cpuHist[bindex] += 1
				}

				rmin = rmax
				rmax = rmin + binWidth
			}
		}
	}
}
