package main

import (
	"fmt"
	"gocl/cl"
	"gocl/cl_demo/utils"
	"unsafe"
)

const PROGRAM_FILE = "atomic.cl"

var KERNEL_FUNC = []byte("atomic")

func main() {

	/* OpenCL data structures */
	var device []cl.CL_device_id
	var context cl.CL_context
	var queue cl.CL_command_queue
	var program *cl.CL_program
	var kernel cl.CL_kernel
	var err cl.CL_int

	var offset, global_size, local_size [1]cl.CL_size_t

	/* Data and events */
	var data [2]cl.CL_int
	var data_buffer cl.CL_mem

	/* Create a device and context */
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

	/* Create a buffer to hold data */
	data_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_WRITE_ONLY,
		cl.CL_size_t(unsafe.Sizeof(data[0]))*2, nil, &err)
	if err < 0 {
		println("Couldn't create a buffer")
		return
	}

	/* Create kernel argument */
	err = cl.CLSetKernelArg(kernel, 0, cl.CL_size_t(unsafe.Sizeof(data_buffer)), unsafe.Pointer(&data_buffer))
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
	offset[0] = 0
	global_size[0] = 8
	local_size[0] = 4
	err = cl.CLEnqueueNDRangeKernel(queue, kernel, 1, offset[:], global_size[:], local_size[:], 0, nil, nil)
	if err < 0 {
		println("Couldn't enqueue the kernel")
		return
	}

	/* Read the buffer */
	err = cl.CLEnqueueReadBuffer(queue, data_buffer, cl.CL_TRUE, 0,
		cl.CL_size_t(unsafe.Sizeof(data[0]))*2, unsafe.Pointer(&data[0]), 0, nil, nil)
	if err < 0 {
		println("Couldn't read the buffer")
		return
	}

	fmt.Printf("Increment: %d\n", data[0])
	fmt.Printf("Atomic increment: %d\n", data[1])

	/* Deallocate resources */
	cl.CLReleaseMemObject(data_buffer)
	cl.CLReleaseKernel(kernel)
	cl.CLReleaseCommandQueue(queue)
	cl.CLReleaseProgram(*program)
	cl.CLReleaseContext(context)
}
