package main

import (
	"fmt"
	"gocl/cl"
	"gocl/cl_demo/utils"
	"unsafe"
)

const PROGRAM_FILE = "hello_kernel.cl"

var KERNEL_FUNC = []byte("hello_kernel")

func main() {

	/* OpenCL data structures */
	var device []cl.CL_device_id
	var context cl.CL_context
	var queue cl.CL_command_queue
	var program *cl.CL_program
	var kernel cl.CL_kernel
	var err cl.CL_int

	/* Data and buffers */
	var msg [16]byte
	var msg_buffer cl.CL_mem

	/* Create a device and context */
	device = utils.Create_device()
	context = cl.CLCreateContext(nil, 1, device[:], nil, nil, &err)
	if err < 0 {
		println("Couldn't create a context")
		return
	}

	/* Build a program and create a kernel */
	program = utils.Build_program(context, device[:], PROGRAM_FILE, nil)
	kernel = cl.CLCreateKernel(*program, KERNEL_FUNC, &err)
	if err < 0 {
		println("Couldn't create a kernel")
		return
	}

	/* Create a buffer to hold the output data */
	msg_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_WRITE_ONLY,
		cl.CL_size_t(unsafe.Sizeof(msg)), nil, &err)
	if err < 0 {
		println("Couldn't create a buffer")
		return
	}

	/* Create kernel argument */
	err = cl.CLSetKernelArg(kernel, 0, cl.CL_size_t(unsafe.Sizeof(msg_buffer)), unsafe.Pointer(&msg_buffer))
	if err < 0 {
		println("Couldn't set a kernel argument")
		return
	}

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
	err = cl.CLEnqueueReadBuffer(queue, msg_buffer, cl.CL_TRUE, 0,
		cl.CL_size_t(unsafe.Sizeof(msg)), unsafe.Pointer(&msg[0]), 0, nil, nil)
	if err < 0 {
		println("Couldn't read the output buffer")
		return
	}
	fmt.Printf("Kernel output: %s\n", msg)

	/* Deallocate resources */
	cl.CLReleaseMemObject(msg_buffer)
	cl.CLReleaseKernel(kernel)
	cl.CLReleaseCommandQueue(queue)
	cl.CLReleaseProgram(*program)
	cl.CLReleaseContext(context)
}
