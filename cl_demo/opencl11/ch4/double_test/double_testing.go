package main

import (
	"fmt"
	"gocl/cl"
	"gocl/cl_demo/utils"
	"strings"
	"unsafe"
)

const PROGRAM_FILE = "double_test.cl"

var KERNEL_FUNC = []byte("double_test")

func main() {

	/* OpenCL data structures */
	var device []cl.CL_device_id
	var context cl.CL_context
	var queue cl.CL_command_queue
	var program *cl.CL_program
	var kernel cl.CL_kernel
	var err cl.CL_int

	/* Data and buffers */
	var a float32 = 6.0
	var b float32 = 2.0
	var result float32
	var a_buffer, b_buffer, output_buffer cl.CL_mem

	/* Extension data */
	var sizeofuint cl.CL_uint
	var addr_data interface{}
	var ext_data interface{}
	fp64_ext := "cl_khr_fp64"
	var ext_size cl.CL_size_t
	var options []byte

	/* Create a device and context */
	device = utils.Create_device()
	context = cl.CLCreateContext(nil, 1, device[:], nil, nil, &err)
	if err < 0 {
		println("Couldn't create a context")
		return
	}

	/* Obtain the device data */
	if cl.CLGetDeviceInfo(device[0], cl.CL_DEVICE_ADDRESS_BITS,
		cl.CL_size_t(unsafe.Sizeof(sizeofuint)), &addr_data, nil) < 0 {
		println("Couldn't read extension data")
		return
	}
	fmt.Printf("Address width: %v\n", addr_data.(cl.CL_uint))

	/* Define "FP_64" option if doubles are supported */
	cl.CLGetDeviceInfo(device[0], cl.CL_DEVICE_EXTENSIONS,
		0, nil, &ext_size)
	// ext_data = (char*)malloc(ext_size + 1);
	// ext_data[ext_size] = '\0';
	cl.CLGetDeviceInfo(device[0], cl.CL_DEVICE_EXTENSIONS,
		ext_size, &ext_data, nil)
	if strings.Contains(ext_data.(string), fp64_ext) {
		fmt.Printf("The %s extension is supported.\n", fp64_ext)
		options = []byte("-DFP_64 ")
	} else {
		fmt.Printf("The %s extension is not supported. %s\n", fp64_ext, ext_data.(string))
	}

	/* Build the program and create the kernel */
	program = utils.Build_program(context, device[:], PROGRAM_FILE, options)
	kernel = cl.CLCreateKernel(*program, KERNEL_FUNC, &err)
	if err < 0 {
		println("Couldn't create a kernel")
		return
	}

	/* Create CL buffers to hold input and output data */
	a_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_READ_ONLY|
		cl.CL_MEM_COPY_HOST_PTR, cl.CL_size_t(unsafe.Sizeof(a)), unsafe.Pointer(&a), &err)
	if err < 0 {
		println("Couldn't create a memory object")
		return
	}

	b_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_READ_ONLY|
		cl.CL_MEM_COPY_HOST_PTR, cl.CL_size_t(unsafe.Sizeof(b)), unsafe.Pointer(&b), nil)
	output_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_WRITE_ONLY,
		cl.CL_size_t(unsafe.Sizeof(b)), nil, nil)

	/* Create kernel arguments */
	err = cl.CLSetKernelArg(kernel, 0, cl.CL_size_t(unsafe.Sizeof(a_buffer)), unsafe.Pointer(&a_buffer))
	if err < 0 {
		println("Couldn't set a kernel argument")
		return
	}
	cl.CLSetKernelArg(kernel, 1, cl.CL_size_t(unsafe.Sizeof(b_buffer)), unsafe.Pointer(&b_buffer))
	cl.CLSetKernelArg(kernel, 2, cl.CL_size_t(unsafe.Sizeof(output_buffer)), unsafe.Pointer(&output_buffer))

	/* Create a command queue */
	queue = cl.CLCreateCommandQueue(context, device[0], 0, &err)
	if err < 0 {
		println("Couldn't create a command queue")
		return
	}

	/* Enqueue kernel */
	err = cl.CLEnqueueTask(queue, kernel, 0, nil, nil)
	if err < 0 {
		println("Couldn't enqueue the kernel")
		return
	}

	/* Read and print the result */
	err = cl.CLEnqueueReadBuffer(queue, output_buffer, cl.CL_TRUE, 0,
		cl.CL_size_t(unsafe.Sizeof(result)), unsafe.Pointer(&result), 0, nil, nil)
	if err < 0 {
		println("Couldn't read the output buffer")
		return
	}
	fmt.Printf("The kernel result is %f\n", result)

	/* Deallocate resources */
	cl.CLReleaseMemObject(a_buffer)
	cl.CLReleaseMemObject(b_buffer)
	cl.CLReleaseMemObject(output_buffer)
	cl.CLReleaseKernel(kernel)
	cl.CLReleaseCommandQueue(queue)
	cl.CLReleaseProgram(*program)
	cl.CLReleaseContext(context)
}
