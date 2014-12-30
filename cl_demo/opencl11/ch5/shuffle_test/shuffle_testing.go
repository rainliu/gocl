package main

import (
	"fmt"
	"gocl/cl"
	"gocl/cl_demo/utils"
	"unsafe"
)

const PROGRAM_FILE = "shuffle_test.cl"

var KERNEL_FUNC = []byte("shuffle_test")

func main() {

	/* OpenCL data structures */
	var device []cl.CL_device_id
	var context cl.CL_context
	var queue cl.CL_command_queue
	var program *cl.CL_program
	var kernel cl.CL_kernel
	var err cl.CL_int

	/* Data and buffers */
	var shuffle1 [8]float32
	var shuffle2 [16]byte
	var shuffle1_buffer, shuffle2_buffer cl.CL_mem

	/* Create a context */
	device = utils.Create_device()
	context = cl.CLCreateContext(nil, 1, device[:], nil, nil, &err)
	if err < 0 {
		println("Couldn't create a context")
		return
	}

	/* Build the program and create a kernel */
	program = utils.Build_program(context, device[:], PROGRAM_FILE, nil)
	kernel = cl.CLCreateKernel(*program, KERNEL_FUNC, &err)
	if err < 0 {
		println("Couldn't create a kernel")
		return
	}

	/* Create a write-only buffer to hold the output data */
	shuffle1_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_WRITE_ONLY,
		cl.CL_size_t(unsafe.Sizeof(shuffle1)), nil, &err)
	if err < 0 {
		println("Couldn't create a buffer")
		return
	}
	shuffle2_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_WRITE_ONLY,
		cl.CL_size_t(unsafe.Sizeof(shuffle2)), nil, &err)

	/* Create kernel argument */
	err = cl.CLSetKernelArg(kernel, 0, cl.CL_size_t(unsafe.Sizeof(shuffle1_buffer)), unsafe.Pointer(&shuffle1_buffer))
	if err < 0 {
		println("Couldn't set a kernel argument")
		return
	}
	cl.CLSetKernelArg(kernel, 1, cl.CL_size_t(unsafe.Sizeof(shuffle2_buffer)), unsafe.Pointer(&shuffle2_buffer))

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
	err = cl.CLEnqueueReadBuffer(queue, shuffle1_buffer, cl.CL_TRUE, 0,
		cl.CL_size_t(unsafe.Sizeof(shuffle1)), unsafe.Pointer(&shuffle1), 0, nil, nil)
	if err < 0 {
		println("Couldn't read the buffer")
		return
	}
	cl.CLEnqueueReadBuffer(queue, shuffle2_buffer, cl.CL_TRUE, 0,
		cl.CL_size_t(unsafe.Sizeof(shuffle2)), unsafe.Pointer(&shuffle2), 0, nil, nil)

	fmt.Printf("Shuffle1: ")
	for i := 0; i < 7; i++ {
		fmt.Printf("%.2f, ", shuffle1[i])
	}
	fmt.Printf("%.2f\n", shuffle1[7])

	fmt.Printf("Shuffle2: ")
	for i := 0; i < 16; i++ {
		fmt.Printf("%c", shuffle2[i])
	}
	fmt.Printf("\n")

	/* Deallocate resources */
	cl.CLReleaseMemObject(shuffle1_buffer)
	cl.CLReleaseMemObject(shuffle2_buffer)
	cl.CLReleaseKernel(kernel)
	cl.CLReleaseCommandQueue(queue)
	cl.CLReleaseProgram(*program)
	cl.CLReleaseContext(context)
}
