package main

import (
	"gocl/cl"
)

// Array of the structures defined below is built and populated
// with the random values on the host.
// Then it is traversed in the OpenCL kernel on the device.
type Element struct {
	internal *cl.CL_float //points to the "value" of another Element from the same array
	external *cl.CL_float //points to the entry in a separate array of floating-point values
	value    cl.CL_float
}

/*
func svmbasic (size_t size,
    cl_context context,
    cl_command_queue queue,
    cl_kernel kernel){
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

    cl_int err = CL_SUCCESS;

    // To enable host & device code to share pointer to the same address space
    // the arrays should be allocated as SVM memory. Use the clSVMAlloc function
    // to allocate SVM memory.
    //
    // Optionally, this function allows specifying alignment in bytes as its
    // last argument. As this basic example doesn't require any _special_ alignment,
    // the following code illustrates requesting default alignment via passing
    // zero value.

    Element* inputElements =
        (Element*)clSVMAlloc(
            context,                // the context where this memory is supposed to be used
            CL_MEM_READ_ONLY,
            size*sizeof(Element),   // amount of memory to allocate (in bytes)
            0                       // alignment in bytes (0 means default)
        );

    float* inputFloats =
        (float*)clSVMAlloc(
            context,                // the context where this memory is supposed to be used
            CL_MEM_READ_ONLY,
            size*sizeof(float),     // amount of memory to allocate (in bytes)
            0                       // alignment in bytes (0 means default)
        );

    // The OpenCL kernel uses the aforementioned input arrays to compute
    // values for the output array.

    float* output =
        (float*)clSVMAlloc(
            context,                // the context where this memory is supposed to be used
            CL_MEM_WRITE_ONLY,
            size*sizeof(float),     // amount of memory to allocate (in bytes)
            0                       // alignment in bytes (0 means default)
    );

    if(!inputElements || !inputFloats || !output)
    {
        throw Error(
            "Cannot allocate SVM memory with clSVMAlloc: "
            "it returns null pointer. "
            "You might be out of memory."
        );
    }

    // In the coarse-grained buffer SVM model, only one OpenCL device (or
    // host) can have ownership for writing to the buffer. Specifically, host
    // explicitly requests the ownership by mapping/unmapping the SVM buffer.
    //
    // So to fill the input SVM buffers on the host, you need to map them to have
    // access from the host program.
    //
    // The following two map calls are required in case of coarse-grained SVM only.

    err = clEnqueueSVMMap(
        queue,
        CL_TRUE,       // blocking map
        CL_MAP_WRITE,
        inputElements,
        sizeof(Element)*size,
        0, 0, 0
    );
    SAMPLE_CHECK_ERRORS(err);

    err = clEnqueueSVMMap(
        queue,
        CL_TRUE,       // blocking map
        CL_MAP_WRITE,
        inputFloats,
        sizeof(float)*size,
        0, 0, 0
    );
    SAMPLE_CHECK_ERRORS(err);

    // Populate data-structures with initial data.

    for (size_t i = 0;  i < size;  i++)
    {
        inputElements[i].internal = &(inputElements[rand_index(size)].value);
        inputElements[i].external = &(inputFloats[rand_index(size)]);
        inputElements[i].value = float(i);
        inputFloats[i] = float(i + size);
    }

    // The following two unmap calls are required in case of coarse-grained SVM only

    err = clEnqueueSVMUnmap(
        queue,
        inputElements,
        0, 0, 0
    );
    SAMPLE_CHECK_ERRORS(err);

    err = clEnqueueSVMUnmap(
        queue,
        inputFloats,
        0, 0, 0
    );
    SAMPLE_CHECK_ERRORS(err);

    // Pass arguments to the kernel.
    // According to the OpenCL 2.0 specification, you need to use a special
    // function to pass a pointer from SVM memory to kernel.

    err = clSetKernelArgSVMPointer(kernel, 0, inputElements);
    SAMPLE_CHECK_ERRORS(err);

    err = clSetKernelArgSVMPointer(kernel, 1, output);
    SAMPLE_CHECK_ERRORS(err);

    // For buffer based SVM (both coarse- and fine-grain) if one SVM buffer
    // points to memory allocated in another SVM buffer, such allocations
    // should be passed to the kernel via clSetKernelExecInfo.

    err = clSetKernelExecInfo(
        kernel,
        CL_KERNEL_EXEC_INFO_SVM_PTRS,
        sizeof(inputFloats),
        &inputFloats
    );
    SAMPLE_CHECK_ERRORS(err);

    // Run the kernel.
    cout << "Running kernel..." << flush;

    err = clEnqueueNDRangeKernel(
        queue,
        kernel,
        1,
        0, &size, 0,
        0, 0, 0
    );
    SAMPLE_CHECK_ERRORS(err);

    // Map the output SVM buffer to read the results.
    // Mapping is required for coarse-grained SVM only.

    err = clEnqueueSVMMap(
        queue,
        CL_TRUE,       // blocking map
        CL_MAP_READ,
        output,
        sizeof(float)*size,
        0, 0, 0
    );
    SAMPLE_CHECK_ERRORS(err);

    cout << " DONE.\n";

    // Validate output state for correctness.

    cout << "Checking correctness of the output buffer..." << flush;
    for(size_t i = 0; i < size; i++)
    {
        float expectedValue = *(inputElements[i].internal) + *(inputElements[i].external);
        if(output[i] != expectedValue)
        {
            cout << " FAILED.\n";

            cerr
                << "Mismatch at position " << i
                << ", read " << output[i]
                << ", expected " << expectedValue << "\n";

            throw Error("Validation failed");
        }
    }
    cout << " PASSED.\n";

    err = clEnqueueSVMUnmap(
        queue,
        output,
        0, 0, 0
    );
    SAMPLE_CHECK_ERRORS(err);

    err = clFinish(queue);
    SAMPLE_CHECK_ERRORS(err);

    // Release all SVM buffers and exit.

    clSVMFree(context, output);
    clSVMFree(context, inputFloats);
    clSVMFree(context, inputElements);
}


bool checkSVMAvailability (cl_device_id device)
{
    cl_device_svm_capabilities caps;

    cl_int err = clGetDeviceInfo(
        device,
        CL_DEVICE_SVM_CAPABILITIES,
        sizeof(cl_device_svm_capabilities),
        &caps,
        0
    );

    // Coarse-grained buffer SVM should be available on any OpenCL 2.0 device.
    // So it is either not an OpenCL 2.0 device or it must support coarse-grained buffer SVM:
    return err == CL_SUCCESS && (caps & CL_DEVICE_SVM_COARSE_GRAIN_BUFFER);
}
*/

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

	var caps cl.CL_device_svm_capabilities
	var caps_value interface{}

	status = cl.CLGetDeviceInfo(
		devices[0],
		cl.CL_DEVICE_SVM_CAPABILITIES,
		cl.CL_size_t(unsafe.Sizeof(caps)),
		&caps_value,
		0)
	caps = caps_value.(cl.CL_device_svm_capabilities)

	// Coarse-grained buffer SVM should be available on any OpenCL 2.0 device.
	// So it is either not an OpenCL 2.0 device or it must support coarse-grained buffer SVM:
	if !(status == cl.CL_SUCCESS && (caps & cl.CL_DEVICE_SVM_COARSE_GRAIN_BUFFER)) {
		println("Cannot detect SVM capabilities of the device.")
		println("The device seemingly doesn't support SVM.")
		return
	}
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
	var cmdQueue cl.CL_command_queue

	// Create a command queue using clCreateCommandQueueWithProperties(),
	// and associate it with the device you want to execute
	cmdQueue = cl.CLCreateCommandQueueWithProperties(context,
		devices[0],
		nil,
		&status)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLCreateCommandQueueWithProperties")
	defer cl.CLReleaseCommandQueue(cmdQueue)

	//-----------------------------------------------------
	// STEP 5: Create and compile the program
	//-----------------------------------------------------
	programSource, programeSize := utils.Load_programsource("svmcg.cl")

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
	kernel := cl.CLCreateKernel(program, []byte("svmbasic"), &status)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLCreateKernel")
	defer cl.CLReleaseKernel(kernel)

	// Then call the main sample routine - resource allocations, OpenCL kernel
	// execution, and so on.
	//svmbasic(size.getValue(), oclobjects.context, oclobjects.queue, executable.kernel)

	// All resource deallocations happen in defer.
}
