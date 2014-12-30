// This program implements a vector addition using OpenCL
package main

import (
    "unsafe"
    "gocl/cl"
    "gocl/cl_demo/utils"
)

func main() {
    // This code executes on the OpenCL host
    
    // Host data
    var size cl.CL_int
    var A []cl.CL_int;//input array
    var B []cl.CL_int;//input array
    var C []cl.CL_int;//output array
    
    // Elements in each array
    const elements = cl.CL_size_t(2048);   
    
    // Compute the size of the data 
    datasize := cl.CL_size_t(unsafe.Sizeof(size))*elements;

    // Allocate space for input/output data
    A = make([]cl.CL_int, datasize);
    B = make([]cl.CL_int, datasize);
    C = make([]cl.CL_int, datasize);
    // Initialize the input data
    for i := cl.CL_int(0); i < cl.CL_int(elements); i++ {
        A[i] = i;
        B[i] = i;
    }

    // Use this to check the output of each API call
    var status cl.CL_int;  
     
    //-----------------------------------------------------
    // STEP 1: Discover and initialize the platforms
    //-----------------------------------------------------
    
    var numPlatforms cl.CL_uint
    var platforms []cl.CL_platform_id;
    
    // Use clGetPlatformIDs() to retrieve the number of 
    // platforms
    status = cl.CLGetPlatformIDs(0, nil, &numPlatforms);
 
    // Allocate enough space for each platform
    platforms = make([]cl.CL_platform_id, numPlatforms);
 
    // Fill in platforms with clGetPlatformIDs()
    status = cl.CLGetPlatformIDs(numPlatforms, platforms, nil);
    if status!=cl.CL_SUCCESS{
        println("CLGetPlatformIDs status!=cl.CL_SUCCESS")
        return;
    }
    //-----------------------------------------------------
    // STEP 2: Discover and initialize the devices
    //----------------------------------------------------- 
    
    var numDevices cl.CL_uint;
    var devices []cl.CL_device_id;

    // Use clGetDeviceIDs() to retrieve the number of 
    // devices present
    status = cl.CLGetDeviceIDs(platforms[0], 
        cl.CL_DEVICE_TYPE_ALL, 
        0, 
        nil, 
        &numDevices);
    if status!=cl.CL_SUCCESS{
        println("CLGetDeviceIDs status!=cl.CL_SUCCESS")
        return;
    }

    // Allocate enough space for each device
    devices = make([]cl.CL_device_id, numDevices);

    // Fill in devices with clGetDeviceIDs()
    status = cl.CLGetDeviceIDs(platforms[0], 
        cl.CL_DEVICE_TYPE_ALL,        
        numDevices, 
        devices, 
        nil);
    if status!=cl.CL_SUCCESS{
        println("CLGetDeviceIDs status!=cl.CL_SUCCESS")
        return;
    }
    //-----------------------------------------------------
    // STEP 3: Create a context
    //----------------------------------------------------- 
    
    var context cl.CL_context;

    // Create a context using clCreateContext() and 
    // associate it with the devices
    context = cl.CLCreateContext(nil, 
        numDevices, 
        devices, 
        nil, 
        nil, 
        &status);
    if status!=cl.CL_SUCCESS{
        println("CLCreateContext status!=cl.CL_SUCCESS")
        return;
    }
    //-----------------------------------------------------
    // STEP 4: Create a command queue
    //----------------------------------------------------- 
    
    var cmdQueue cl.CL_command_queue; 

    // Create a command queue using clCreateCommandQueue(),
    // and associate it with the device you want to execute 
    // on
    cmdQueue = cl.CLCreateCommandQueue(context, 
        devices[0], 
        0, 
        &status);
    if status!=cl.CL_SUCCESS{
        println("CLCreateCommandQueue status!=cl.CL_SUCCESS")
        return;
    }
    //-----------------------------------------------------
    // STEP 5: Create device buffers
    //----------------------------------------------------- 
    
    var bufferA cl.CL_mem;  // Input array on the device
    var bufferB cl.CL_mem;  // Input array on the device
    var bufferC cl.CL_mem;  // Output array on the device

    // Use clCreateBuffer() to create a buffer object (d_A) 
    // that will contain the data from the host array A
    bufferA = cl.CLCreateBuffer(context, 
        cl.CL_MEM_READ_ONLY,                         
        datasize, 
        nil, 
        &status);
    if status!=cl.CL_SUCCESS{
        println("CLCreateBuffer status!=cl.CL_SUCCESS")
        return;
    }
    // Use clCreateBuffer() to create a buffer object (d_B)
    // that will contain the data from the host array B
    bufferB = cl.CLCreateBuffer(context, 
        cl.CL_MEM_READ_ONLY,                         
        datasize, 
        nil, 
        &status);
    if status!=cl.CL_SUCCESS{
        println("CLCreateBuffer status!=cl.CL_SUCCESS")
        return;
    }
    // Use clCreateBuffer() to create a buffer object (d_C) 
    // with enough space to hold the output data
    bufferC = cl.CLCreateBuffer(context, 
        cl.CL_MEM_WRITE_ONLY,                 
        datasize, 
        nil, 
        &status);
    if status!=cl.CL_SUCCESS{
        println("CLCreateBuffer status!=cl.CL_SUCCESS")
        return;
    }
    //-----------------------------------------------------
    // STEP 6: Write host data to device buffers
    //----------------------------------------------------- 
    
    // Use clEnqueueWriteBuffer() to write input array A to
    // the device buffer bufferA
    status = cl.CLEnqueueWriteBuffer(cmdQueue, 
        bufferA, 
        cl.CL_FALSE, 
        0, 
        datasize,                         
        unsafe.Pointer(&A[0]), 
        0, 
        nil, 
        nil);
    if status!=cl.CL_SUCCESS{
        println("CLEnqueueWriteBuffer status!=cl.CL_SUCCESS")
        return;
    }    
    // Use clEnqueueWriteBuffer() to write input array B to 
    // the device buffer bufferB
    status = cl.CLEnqueueWriteBuffer(cmdQueue, 
        bufferB, 
        cl.CL_FALSE, 
        0, 
        datasize,                                  
        unsafe.Pointer(&B[0]), 
        0, 
        nil, 
        nil);
    if status!=cl.CL_SUCCESS{
        println("CLEnqueueWriteBuffer status!=cl.CL_SUCCESS")
        return;
    }
    //-----------------------------------------------------
    // STEP 7: Create and compile the program
    //----------------------------------------------------- 
    programSource, programeSize := ch0.Load_programsource("chapter2.cl");
    
    // Create a program using clCreateProgramWithSource()
    program := cl.CLCreateProgramWithSource(context, 
        1, 
        programSource[:],                                 
        programeSize[:], 
        &status);
    if status!=cl.CL_SUCCESS{
        println("CLCreateProgramWithSource status!=cl.CL_SUCCESS")
        return;
    }
    // Build (compile) the program for the devices with
    // clBuildProgram()
    status = cl.CLBuildProgram(program, 
        numDevices, 
        devices, 
        nil, 
        nil, 
        nil);
    if status!=cl.CL_SUCCESS{
        println("CLBuildProgram status!=cl.CL_SUCCESS")
        return;
    }   
    //-----------------------------------------------------
    // STEP 8: Create the kernel
    //----------------------------------------------------- 

    var kernel cl.CL_kernel;

    // Use clCreateKernel() to create a kernel from the 
    // vector addition function (named "vecadd")
    kernel = cl.CLCreateKernel(program, []byte("vecadd"), &status);
    if status!=cl.CL_SUCCESS{
        println("CLCreateKernel status!=cl.CL_SUCCESS")
        return;
    }
    //-----------------------------------------------------
    // STEP 9: Set the kernel arguments
    //----------------------------------------------------- 
    
    // Associate the input and output buffers with the 
    // kernel 
    // using clSetKernelArg()
    status  = cl.CLSetKernelArg(kernel, 
        0, 
        cl.CL_size_t(unsafe.Sizeof(bufferA)), 
        unsafe.Pointer(&bufferA));
    status |= cl.CLSetKernelArg(kernel, 
        1, 
        cl.CL_size_t(unsafe.Sizeof(bufferB)), 
        unsafe.Pointer(&bufferB));
    status |= cl.CLSetKernelArg(kernel, 
        2, 
        cl.CL_size_t(unsafe.Sizeof(bufferC)), 
        unsafe.Pointer(&bufferC));
    if status!=cl.CL_SUCCESS{
        println("CLSetKernelArg status!=cl.CL_SUCCESS")
        return;
    }
    //-----------------------------------------------------
    // STEP 10: Configure the work-item structure
    //----------------------------------------------------- 
    
    // Define an index space (global work size) of work 
    // items for 
    // execution. A workgroup size (local work size) is not 
    // required, 
    // but can be used.
    var globalWorkSize [1]cl.CL_size_t;    
    // There are 'elements' work-items 
    globalWorkSize[0] = elements;

    //-----------------------------------------------------
    // STEP 11: Enqueue the kernel for execution
    //----------------------------------------------------- 
    
    // Execute the kernel by using 
    // clEnqueueNDRangeKernel().
    // 'globalWorkSize' is the 1D dimension of the 
    // work-items
    status = cl.CLEnqueueNDRangeKernel(cmdQueue, 
        kernel, 
        1, 
        nil, 
        globalWorkSize[:], 
        nil, 
        0, 
        nil, 
        nil);
    if status!=cl.CL_SUCCESS{
        println("CLEnqueueNDRangeKernel status!=cl.CL_SUCCESS")
        return;
    }
    //-----------------------------------------------------
    // STEP 12: Read the output buffer back to the host
    //----------------------------------------------------- 
    
    // Use clEnqueueReadBuffer() to read the OpenCL output  
    // buffer (bufferC) 
    // to the host output array (C)
    cl.CLEnqueueReadBuffer(cmdQueue, 
        bufferC, 
        cl.CL_TRUE, 
        0, 
        datasize, 
        unsafe.Pointer(&C[0]), 
        0, 
        nil, 
        nil);
    if status!=cl.CL_SUCCESS{
        println("CLEnqueueReadBuffer status!=cl.CL_SUCCESS")
        return;
    }
    // Verify the output
    result := true;
    for i := cl.CL_int(0); i < cl.CL_int(elements); i++ {
        if C[i] != i+i {
            result = false;
            break;
        }
    }
    if result {
        println("Output is correct\n");
    } else {
        println("Output is incorrect\n");
    }

    //-----------------------------------------------------
    // STEP 13: Release OpenCL resources
    //----------------------------------------------------- 
    
    // Free OpenCL resources
    cl.CLReleaseKernel(kernel);
    cl.CLReleaseProgram(program);
    cl.CLReleaseCommandQueue(cmdQueue);
    cl.CLReleaseMemObject(bufferA);
    cl.CLReleaseMemObject(bufferB);
    cl.CLReleaseMemObject(bufferC);
    cl.CLReleaseContext(context);
}
