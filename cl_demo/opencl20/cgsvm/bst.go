// This program demo OpenCL 2.0 Coarse-Grain SVM memory feature
package main

import (
	"fmt"
	"gocl/cl"
	"gocl/cl_demo/utils"
	"math/rand"
	"unsafe"
)

const (
	NUMBER_OF_NODES      = 1024 * 1024
	NUMBER_OF_SEARCH_KEY = NUMBER_OF_NODES / 4
)

/* binary tree node definition */
type node struct {
	value cl.CL_int
	left  *node
	right *node
}

var sampleNode node

/* search keys */
type searchKey struct {
	key        cl.CL_int
	oclNode    *node
	nativeNode *node
}

var sampleKey searchKey

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
		cl.CL_size_t(NUMBER_OF_NODES*unsafe.Sizeof(sampleNode)),
		0)
	if nil == svmTreeBuf {
		println("clSVMAlloc(svmTreeBuf) failed.")
		return
	}
	defer cl.CLSVMFree(context, svmTreeBuf)

	svmSearchBuf = cl.CLSVMAlloc(context,
		cl.CL_MEM_READ_WRITE,
		cl.CL_size_t(NUMBER_OF_SEARCH_KEY*unsafe.Sizeof(sampleKey)),
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
	fmt.Printf("Searching %d keys in a BST having %d Nodes...\n", NUMBER_OF_SEARCH_KEY, NUMBER_OF_NODES)
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
		nil)
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
	svmBinaryTreeCPUReference(cmdQueue,
		svmRoot,
		svmTreeBuf,
		svmSearchBuf)

	// compare the results and see if they match
	pass := svmCompareResults(cmdQueue, svmSearchBuf)
	if pass {
		println("Passed!")
	} else {
		println("Failed!")
	}
}

/**
 * cpuCreateBinaryTree()
 * creates a tree from the data in "svmTreeBuf". If this is NULL returns NULL
 * else returns root of the tree.
 **/
func cpuCreateBinaryTree(commandQueue cl.CL_command_queue,
	svmTreeBuf unsafe.Pointer) *node {
	var root *node
	var status cl.CL_int

	// reserve svm space for CPU update
	status = cl.CLEnqueueSVMMap(commandQueue,
		cl.CL_TRUE, //blocking call
		cl.CL_MAP_WRITE_INVALIDATE_REGION,
		svmTreeBuf,
		cl.CL_size_t(NUMBER_OF_NODES*unsafe.Sizeof(sampleNode)),
		0,
		nil,
		nil)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clEnqueueSVMMap(svmTreeBuf)")

	//init node and make bt
	root = cpuMakeBinaryTree(svmTreeBuf)

	status = cl.CLEnqueueSVMUnmap(commandQueue,
		svmTreeBuf,
		0,
		nil,
		nil)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clEnqueueSVMUnmap(svmTreeBuf)")

	return root
}

func cpuMakeBinaryTree(svmTreeBuf unsafe.Pointer) *node {
	var root *node
	var nextData *node
	var nextNode *node
	var insertedFlag bool

	r := rand.New(rand.NewSource(99))

	// initialize nodes
	for i := 0; i < NUMBER_OF_NODES; i++ {
		nextData = (*node)(unsafe.Pointer(uintptr(svmTreeBuf) + uintptr(i)*unsafe.Sizeof(sampleNode)))

		// allocate a random value to node
		nextData.value = cl.CL_int(r.Int())
		// all pointers are null
		nextData.left = nil
		nextData.right = nil
	}

	// allocate first node to root
	root = (*node)(svmTreeBuf)

	// iterative tree insert
	for i := 1; i < NUMBER_OF_NODES; i++ {
		nextData = (*node)(unsafe.Pointer(uintptr(svmTreeBuf) + uintptr(i)*unsafe.Sizeof(sampleNode)))
		nextNode = root
		insertedFlag = false

		for false == insertedFlag {
			if nextData.value <= nextNode.value {
				// move left
				if nil == nextNode.left {
					nextNode.left = nextData
					insertedFlag = true
				} else {
					nextNode = nextNode.left
				}
			} else {
				// move right
				if nil == nextNode.right {
					nextNode.right = nextData
					insertedFlag = true
				} else {
					nextNode = nextNode.right
				}
			}
		}
	}

	return root
}

func cpuInitSearchKeys(commandQueue cl.CL_command_queue,
	svmSearchBuf unsafe.Pointer) {
	var nextData *searchKey
	var status cl.CL_int

	status = cl.CLEnqueueSVMMap(commandQueue,
		cl.CL_TRUE, //blocking call
		cl.CL_MAP_WRITE_INVALIDATE_REGION,
		svmSearchBuf,
		cl.CL_size_t(NUMBER_OF_SEARCH_KEY*unsafe.Sizeof(sampleKey)),
		0,
		nil,
		nil)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clEnqueueSVMMap(svmSearchBuf)")

	r := rand.New(rand.NewSource(999))

	// initialize nodes
	for i := 0; i < NUMBER_OF_SEARCH_KEY; i++ {
		nextData = (*searchKey)(unsafe.Pointer(uintptr(svmSearchBuf) + uintptr(i)*unsafe.Sizeof(sampleKey)))
		// allocate a random value to node
		nextData.key = cl.CL_int(r.Int())
		// all pointers are null
		nextData.oclNode = nil
		nextData.nativeNode = nil
	}

	status = cl.CLEnqueueSVMUnmap(commandQueue,
		svmSearchBuf,
		0,
		nil,
		nil)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clEnqueueSVMUnmap(svmSearchBuf)")
}

func svmBinaryTreeCPUReference(commandQueue cl.CL_command_queue,
	svmRoot *node,
	svmTreeBuf unsafe.Pointer,
	svmSearchBuf unsafe.Pointer) {
	var status cl.CL_int

	/* reserve svm buffers for cpu usage */
	status = cl.CLEnqueueSVMMap(commandQueue,
		cl.CL_TRUE, //blocking call
		cl.CL_MAP_READ,
		svmTreeBuf,
		cl.CL_size_t(NUMBER_OF_NODES*unsafe.Sizeof(sampleNode)),
		0,
		nil,
		nil)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clEnqueueSVMMap(svmTreeBuf)")

	status = cl.CLEnqueueSVMMap(commandQueue,
		cl.CL_TRUE, //blocking call
		cl.CL_MAP_WRITE,
		svmSearchBuf,
		cl.CL_size_t(NUMBER_OF_SEARCH_KEY*unsafe.Sizeof(sampleKey)),
		0,
		nil,
		nil)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clEnqueueSVMMap(svmSearchBuf)")

	for i := 0; i < NUMBER_OF_SEARCH_KEY; i++ {
		/* search tree */
		searchNode := svmRoot
		currKey := (*searchKey)(unsafe.Pointer(uintptr(svmSearchBuf) + uintptr(i)*unsafe.Sizeof(sampleKey)))

		for nil != searchNode {
			if currKey.key == searchNode.value {
				/* rejoice on finding key */
				currKey.nativeNode = searchNode
				searchNode = nil
			} else if currKey.key < searchNode.value {
				/* move left */
				searchNode = searchNode.left
			} else {
				/* move right */
				searchNode = searchNode.right
			}
		}
	}

	status = cl.CLEnqueueSVMUnmap(commandQueue,
		svmSearchBuf,
		0,
		nil,
		nil)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clEnqueueSVMUnmap(svmSearchBuf)")

	status = cl.CLEnqueueSVMUnmap(commandQueue,
		svmTreeBuf,
		0,
		nil,
		nil)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clEnqueueSVMUnmap(svmTreeBuf)")
}

func svmCompareResults(commandQueue cl.CL_command_queue,
	svmSearchBuf unsafe.Pointer) bool {
	var compare_status bool
	var status cl.CL_int

	status = cl.CLEnqueueSVMMap(commandQueue,
		cl.CL_TRUE, //blocking call
		cl.CL_MAP_WRITE_INVALIDATE_REGION,
		svmSearchBuf,
		cl.CL_size_t(NUMBER_OF_SEARCH_KEY*unsafe.Sizeof(sampleKey)),
		0,
		nil,
		nil)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clEnqueueSVMMap(svmSearchBuf)")

	compare_status = true
	for i := 0; i < NUMBER_OF_SEARCH_KEY; i++ {
		currKey := (*searchKey)(unsafe.Pointer(uintptr(svmSearchBuf) + uintptr(i)*unsafe.Sizeof(sampleKey)))

		/* compare OCL and native nodes */
		if currKey.oclNode != currKey.nativeNode {
			compare_status = false
			break
		}
	}

	status = cl.CLEnqueueSVMUnmap(commandQueue,
		svmSearchBuf,
		0,
		nil,
		nil)
	utils.CHECK_STATUS(status, cl.CL_SUCCESS, "clEnqueueSVMUnmap(svmSearchBuf)")

	return compare_status
}
