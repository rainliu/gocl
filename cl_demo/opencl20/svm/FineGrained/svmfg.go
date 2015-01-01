package main

import (
	"fmt"
	"gocl/cl"
	"gocl/cl_demo/utils"
	"math/rand"
	"unsafe"
)

// Array of the structures defined below is built and populated
// with the random values on the host.
// Then it is traversed in the OpenCL kernel on the device.
type Element struct {
	internal *cl.CL_float //points to the "value" of another Element from the same array
	external *cl.CL_float //points to the entry in a separate array of floating-point values
	value    cl.CL_float
}

var sampleElement Element
var sampleFloat cl.CL_float

func svmbasic(size cl.CL_size_t,
	context cl.CL_context,
	queue cl.CL_command_queue,
	kernel cl.CL_kernel) {
	// Prepare input data as follows.
	// Build two arrays:
	//     - an array that consists of the Element structures
	//       (refer to svmbasic.h for the structure definition)
	//     - an array that consists of the float values
	//
	// Each structure of the first array has the following pointers:
	//     - 'internal', which points to a 'value' field of another entry
	//       of the same array.
	//     - 'external', which points to a float value from the the
	//       second array.
	//
	// Pointers are set randomly. The structures do not reflect any real usage
	// scenario, but are illustrative for a simple device-side traversal.
	//
	//        Array of Element                        Array of floats
	//           structures
	//
	//    ||====================||
	//    ||    .............   ||                   ||============||
	//    ||    .............   ||<-----+            || .......... ||
	//    ||====================||      |            ||    float   ||
	//    ||   float* internal--||------+            ||    float   ||
	//    ||   float* external--||------------------>||    float   ||
	//    ||   float value <----||------+            || .......... ||
	//    ||====================||      |            || .......... ||
	//    ||    .............   ||      |            ||    float   ||
	//    ||    .............   ||      |            ||    float   ||
	//    ||====================||      |            ||    float   ||
	//    ||====================||      |            ||    float   ||
	//    ||   float* internal--||------+            ||    float   ||
	//    ||   float* external--||------------------>||    float   ||
	//    ||   float value      ||                   ||    float   ||
	//    ||====================||                   ||    float   ||
	//    ||    .............   ||                   || .......... ||
	//    ||    .............   ||                   ||============||
	//    ||====================||
	//
	// The two arrays are created independently and are used to illustrate
	// two new OpenCL 2.0 API functions:
	//    - the array of Element structures is passed to the kernel as a
	//      kernel argument with the clSetKernelArgSVMPointer function
	//    - the array of floats is used by the kernel indirectly, and this
	//      dependency should be also specified with the clSetKernelExecInfo
	//      function prior to the kernel execution

	var err cl.CL_int

	// To enable host & device code to share pointer to the same address space
	// the arrays should be allocated as SVM memory. Use the clSVMAlloc function
	// to allocate SVM memory.
	//
	// Optionally, this function allows specifying alignment in bytes as its
	// last argument. As this basic example doesn't require any _special_ alignment,
	// the following code illustrates requesting default alignment via passing
	// zero value.

	inputElements := cl.CLSVMAlloc(context, // the context where this memory is supposed to be used
		cl.CL_MEM_READ_ONLY|cl.CL_MEM_SVM_FINE_GRAIN_BUFFER,
		size*cl.CL_size_t(unsafe.Sizeof(sampleElement)), // amount of memory to allocate (in bytes)
		0) // alignment in bytes (0 means default)
	if nil == inputElements {
		println("Cannot allocate SVM memory with clSVMAlloc: it returns null pointer. You might be out of memory.")
		return
	}
	defer cl.CLSVMFree(context, inputElements)

	inputFloats := cl.CLSVMAlloc(context, // the context where this memory is supposed to be used
		cl.CL_MEM_READ_ONLY|cl.CL_MEM_SVM_FINE_GRAIN_BUFFER,
		size*cl.CL_size_t(unsafe.Sizeof(sampleFloat)), // amount of memory to allocate (in bytes)
		0) // alignment in bytes (0 means default)
	if nil == inputFloats {
		println("Cannot allocate SVM memory with clSVMAlloc: it returns null pointer. You might be out of memory.")
		return
	}
	defer cl.CLSVMFree(context, inputFloats)

	// The OpenCL kernel uses the aforementioned input arrays to compute
	// values for the output array.

	output := cl.CLSVMAlloc(context, // the context where this memory is supposed to be used
		cl.CL_MEM_WRITE_ONLY|cl.CL_MEM_SVM_FINE_GRAIN_BUFFER,
		size*cl.CL_size_t(unsafe.Sizeof(sampleFloat)), // amount of memory to allocate (in bytes)
		0) // alignment in bytes (0 means default)
	defer cl.CLSVMFree(context, output)

	if nil == output {
		println("Cannot allocate SVM memory with clSVMAlloc: it returns null pointer. You might be out of memory.")
		return
	}

	// Note: in the coarse-grained SVM, mapping of inputElement and inputFloats is
	// needed to do the following initialization. While here, in the fine-grained SVM,
	// it is not necessary.

	// Populate data-structures with initial data.
	r := rand.New(rand.NewSource(99))

	for i := cl.CL_size_t(0); i < size; i++ {
		inputElement := (*Element)(unsafe.Pointer(uintptr(inputElements) + uintptr(i)*unsafe.Sizeof(sampleElement)))
		inputFloat := (*cl.CL_float)(unsafe.Pointer(uintptr(inputFloats) + uintptr(i)*unsafe.Sizeof(sampleFloat)))
		randElement := (*Element)(unsafe.Pointer(uintptr(inputElements) + uintptr(r.Intn(int(size)))*unsafe.Sizeof(sampleElement)))
		randFloat := (*cl.CL_float)(unsafe.Pointer(uintptr(inputFloats) + uintptr(r.Intn(int(size)))*unsafe.Sizeof(sampleFloat)))

		inputElement.internal = &(randElement.value)
		inputElement.external = randFloat
		inputElement.value = cl.CL_float(i)
		*inputFloat = cl.CL_float(i + size)
	}

	// Note: in the coarse-grained SVM, unmapping of inputElement and inputFloats is
	// needed before scheduling the kernel for execution. While here, in the fine-grained SVM,
	// it is not necessary.

	// Pass arguments to the kernel.
	// According to the OpenCL 2.0 specification, you need to use a special
	// function to pass a pointer from SVM memory to kernel.

	err = cl.CLSetKernelArgSVMPointer(kernel, 0, inputElements)
	utils.CHECK_STATUS(err, cl.CL_SUCCESS, "CLSetKernelArgSVMPointer")

	err = cl.CLSetKernelArgSVMPointer(kernel, 1, output)
	utils.CHECK_STATUS(err, cl.CL_SUCCESS, "CLSetKernelArgSVMPointer")

	// For buffer based SVM (both coarse- and fine-grain) if one SVM buffer
	// points to memory allocated in another SVM buffer, such allocations
	// should be passed to the kernel via clSetKernelExecInfo.

	err = cl.CLSetKernelExecInfo(kernel,
		cl.CL_KERNEL_EXEC_INFO_SVM_PTRS,
		cl.CL_size_t(unsafe.Sizeof(inputFloats)),
		unsafe.Pointer(&inputFloats))
	utils.CHECK_STATUS(err, cl.CL_SUCCESS, "CLSetKernelExecInfo")

	// Run the kernel.
	println("Running kernel...")

	var globalWorkSize [1]cl.CL_size_t
	globalWorkSize[0] = size

	err = cl.CLEnqueueNDRangeKernel(queue,
		kernel,
		1,
		nil,
		globalWorkSize[:],
		nil,
		0,
		nil,
		nil)
	utils.CHECK_STATUS(err, cl.CL_SUCCESS, "CLEnqueueNDRangeKernel")

	// Note: In the fine-grained SVM, after enqueuing the kernel above, the host application is
	// not blocked from accessing SVM allocations that were passed to the kernel. The host
	// can access the same regions of SVM memory as does the kernel if the kernel and the host
	// read/modify different bytes. If one side (host or device) needs to modify the same bytes
	// that are simultaniously read/modified by another side, atomics operations are usually
	// required to maintain sufficient memory consistency. This sample doesn't use this possibility
	// and the host just waits in clFinish below until the kernel is finished.
	err = cl.CLFinish(queue)
	utils.CHECK_STATUS(err, cl.CL_SUCCESS, "CLFinish")

	println(" DONE.")

	// Validate output state for correctness.
	// Compare: in the coarse-grained SVM case you need to map the output.
	// Here it is not needed.

	println("Checking correctness of the output buffer...")
	for i := cl.CL_size_t(0); i < size; i++ {
		inputElement := (*Element)(unsafe.Pointer(uintptr(inputElements) + uintptr(i)*unsafe.Sizeof(sampleElement)))
		outputFloat := (*cl.CL_float)(unsafe.Pointer(uintptr(output) + uintptr(i)*unsafe.Sizeof(sampleFloat)))
		expectedValue := *(inputElement.internal) + *(inputElement.external)
		if *outputFloat != expectedValue {
			println(" FAILED.")
			fmt.Printf("Mismatch at position %d, read %f, expected %f\n", i, *outputFloat, expectedValue)
			return
		}
	}
	println(" PASSED.")
}

func main() {
	// Use this to check the output of each API call
	var status cl.CL_int

	//-----------------------------------------------------
	// STEP 1: Discover and initialize the platforms
	//-----------------------------------------------------
	var numPlatforms cl.CL_uint

	// Use clGetPlatformIDs() to retrieve the number of
	// platforms
	status = cl.CLGetPlatformIDs(0, nil, &numPlatforms)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLGetPlatformIDs")

	// Allocate enough space for each platform
	platforms := make([]cl.CL_platform_id, numPlatforms)

	// Fill in platforms with clGetPlatformIDs()
	status = cl.CLGetPlatformIDs(numPlatforms, platforms, nil)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLGetPlatformIDs")

	//-----------------------------------------------------
	// STEP 2: Discover and initialize the GPU devices
	//-----------------------------------------------------
	var numDevices cl.CL_uint

	// Use clGetDeviceIDs() to retrieve the number of
	// devices present
	status = cl.CLGetDeviceIDs(platforms[0],
		cl.CL_DEVICE_TYPE_GPU,
		0,
		nil,
		&numDevices)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLGetDeviceIDs")

	// Allocate enough space for each device
	devices := make([]cl.CL_device_id, numDevices)

	// Fill in devices with clGetDeviceIDs()
	status = cl.CLGetDeviceIDs(platforms[0],
		cl.CL_DEVICE_TYPE_GPU,
		numDevices,
		devices,
		nil)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLGetDeviceIDs")

	var caps cl.CL_device_svm_capabilities
	var caps_value interface{}

	status = cl.CLGetDeviceInfo(
		devices[0],
		cl.CL_DEVICE_SVM_CAPABILITIES,
		cl.CL_size_t(unsafe.Sizeof(caps)),
		&caps_value,
		nil)
	caps = caps_value.(cl.CL_device_svm_capabilities)

	// Coarse-grained buffer SVM should be available on any OpenCL 2.0 device.
	// So it is either not an OpenCL 2.0 device or it must support coarse-grained buffer SVM:
	if !(status == cl.CL_SUCCESS && (caps&cl.CL_DEVICE_SVM_FINE_GRAIN_BUFFER) != 0) {
		fmt.Printf("Cannot detect fine-grained buffer SVM capabilities on the device. The device seemingly doesn't support fine-grained buffer SVM. caps=%x\n", caps)
		println("")
		return
	}

	//-----------------------------------------------------
	// STEP 3: Create a context
	//-----------------------------------------------------
	// Create a context using clCreateContext() and
	// associate it with the devices
	context := cl.CLCreateContext(nil,
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
	// Create a command queue using clCreateCommandQueueWithProperties(),
	// and associate it with the device you want to execute
	queue := cl.CLCreateCommandQueueWithProperties(context,
		devices[0],
		nil,
		&status)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLCreateCommandQueueWithProperties")
	defer cl.CLReleaseCommandQueue(queue)

	//-----------------------------------------------------
	// STEP 5: Create and compile the program
	//-----------------------------------------------------
	programSource, programeSize := utils.Load_programsource("svmfg.cl")

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
		var log interface{}
		var log_size cl.CL_size_t
		/* Find size of log and print to std output */
		cl.CLGetProgramBuildInfo(program, devices[0], cl.CL_PROGRAM_BUILD_LOG, 0, nil, &log_size)
		cl.CLGetProgramBuildInfo(program, devices[0], cl.CL_PROGRAM_BUILD_LOG, log_size, &log, nil)
		fmt.Printf("%s\n", log)
		return
	}
	//utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLBuildProgram")

	//-----------------------------------------------------
	// STEP 7: Create the kernel
	//-----------------------------------------------------
	// Use clCreateKernel() to create a kernel
	kernel := cl.CLCreateKernel(program,
		[]byte("svmbasic"),
		&status)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLCreateKernel")
	defer cl.CLReleaseKernel(kernel)

	// Then call the main sample routine - resource allocations, OpenCL kernel
	// execution, and so on.
	svmbasic(1024*1024, context, queue, kernel)

	// All resource deallocations happen in defer.
}
