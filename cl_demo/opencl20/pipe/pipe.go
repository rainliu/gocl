// This program demo OpenCL 2.0 Pipe feature
package main

import (
	"fmt"
	"gocl/cl"
	"gocl/cl_demo/utils"
	"unsafe"
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
	// STEP 5: Create device buffers
	//-----------------------------------------------------
	/* root node of the binary tree */
	var svmRoot *node

	/* svm buffer for binary tree */
	var svmTreeBuf unsafe.Pointer

	/* svm buffer for search keys */
	var svmSearchBuf unsafe.Pointer

	// initialize any device/SVM memory here.
	svmTreeBuf = cl.CLSVMAlloc(context,
		cl.CL_MEM_READ_WRITE,
		cl.CL_size_t(NUMBER_OF_NODES*unsafe.Sizeof(node)),
		0)
	if nil == svmTreeBuf {
		println("clSVMAlloc(svmTreeBuf) failed.")
		return
	}
	defer cl.CLSVMFree(context, svmTreeBuf)

	svmSearchBuf = cl.CLSVMAlloc(context,
		cl.CL_MEM_READ_WRITE,
		cl.CL_size_t(NUMBER_OF_SEARCH_KEY*unsafe.Sizeof(searchKey)),
		0)
	if nil == svmSearchBuf {
		println("clSVMAlloc(svmSearchBuf) failed.")
		return
	}
	defer cl.CLSVMFree(context, svmSearchBuf)

	//create the binary tree and set the root
	svmRoot = cpuCreateBinaryTree(cmdQueue, svmTreeBuf)

	//initialize search keys
	cpuInitSearchKeys(cmdQueue, svmSearchBuf)

	/* if voice is not deliberately muzzled, shout parameters */
	fmt.Printf("-------------------------------------------------------------------------\n")
	fmt.Printf("Searching %d keys in a BST having %d Nodes...", NUMBER_OF_SEARCH_KEY, NUMBER_OF_NODES)
	fmt.Printf("-------------------------------------------------------------------------\n")

	//-----------------------------------------------------
	// STEP 6: Create and compile the program
	//-----------------------------------------------------
	programSource, programeSize := utils.Load_programsource("bst.cl")

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
	status = cl.CLBuildProgram(program,
		numDevices,
		devices,
		nil,
		nil,
		nil)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLBuildProgram")

	//-----------------------------------------------------
	// STEP 7: Create the kernel
	//-----------------------------------------------------
	var kernel cl.CL_kernel

	// Use clCreateKernel() to create a kernel
	kernel = cl.CLCreateKernel(program, []byte("bst_kernel"), &status)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLCreateKernel")
	defer cl.CLReleaseKernel(kernel)

	//-----------------------------------------------------
	// STEP 8: Set the kernel arguments
	//-----------------------------------------------------
	// Associate the input and output buffers with the
	// kernel
	// using clSetKernelArg()
	// Set appropriate arguments to the kernel
	status = cl.CLSetKernelArgSVMPointer(kernel,
		0,
		svmTreeBuf)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clSetKernelArgSVMPointer(svmTreeBuf)")

	status = cl.CLSetKernelArgSVMPointer(kernel,
		1,
		svmSearchBuf)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clSetKernelArgSVMPointer(svmSearchBuf)")

	//-----------------------------------------------------
	// STEP 9: Configure the work-item structure
	//-----------------------------------------------------
	// Define an index space (global work size) of work
	// items for
	// execution. A workgroup size (local work size) is not
	// required,
	// but can be used.
	var localWorkSize [1]cl.CL_size_t
	var kernelWorkGroupSize interface{}
	status = cl.CLGetKernelWorkGroupInfo(kernel,
		devices[0],
		cl.CL_KERNEL_WORK_GROUP_SIZE,
		cl.CL_size_t(unsafe.Sizeof(localWorkSize[0])),
		&kernelWorkGroupSize,
		NULL)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLGetKernelWorkGroupInfo")
	localWorkSize[0] = kernelWorkGroupSize.(cl.CL_size_t)

	var globalWorkSize [1]cl.CL_size_t
	globalWorkSize[0] = NUMBER_OF_SEARCH_KEY

	//-----------------------------------------------------
	// STEP 10: Enqueue the kernel for execution
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
		localWorkSize[:],
		0,
		nil,
		nil)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "CLEnqueueNDRangeKernel")

	status = cl.CLFlush(cmdQueue)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clFlush")

	status = cl.CLFinish(cmdQueue)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clFinish")

	//-----------------------------------------------------
	// STEP 11: Verify the results
	//-----------------------------------------------------
	// reference implementation
	svmBinaryTreeCPUReference()

	// compare the results and see if they match
	pass := compare(cmdQueue, searchKey)
	if pass {
		println("Passed!")
	} else {
		println("Failed!")
	}
}
