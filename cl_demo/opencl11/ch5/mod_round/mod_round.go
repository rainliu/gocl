package main

import (
	"fmt"
	"gocl/cl"
	"gocl/cl_demo/utils"
	"unsafe"
)

const PROGRAM_FILE = "mod_round.cl"

var KERNEL_FUNC = []byte("mod_round")

func main() {

	/* OpenCL data structures */
	var device []cl.CL_device_id
	var context cl.CL_context
	var queue cl.CL_command_queue
	var program *cl.CL_program
	var kernel cl.CL_kernel
	var err cl.CL_int

	/* Data and buffers */
	var mod_input = [2]float32{317.0, 23.0}
	var mod_output [2]float32
	var round_input = [4]float32{-6.5, -3.5, 3.5, 6.5}
	var round_output [20]float32
	var mod_input_buffer, mod_output_buffer,
		round_input_buffer, round_output_buffer cl.CL_mem

	/* Create a context */
	device = utils.Create_device()
	context = cl.CLCreateContext(nil, 1, device[:], nil, nil, &err)
	if err < 0 {
		println("Couldn't create a context")
		return
	}

	/* Build the program and create a kernel */
	program = utils.Build_program(context, device, PROGRAM_FILE, nil)
	kernel = cl.CLCreateKernel(*program, KERNEL_FUNC, &err)
	if err < 0 {
		println("Couldn't create a kernel")
		return
	}

	/* Create buffers to hold input/output data */
	mod_input_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_READ_ONLY|cl.CL_MEM_COPY_HOST_PTR,
		cl.CL_size_t(unsafe.Sizeof(mod_input)), unsafe.Pointer(&mod_input[0]), &err)
	if err < 0 {
		println("Couldn't create a buffer")
		return
	}
	mod_output_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_WRITE_ONLY,
		cl.CL_size_t(unsafe.Sizeof(mod_output)), nil, nil)
	round_input_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_READ_ONLY|cl.CL_MEM_COPY_HOST_PTR,
		cl.CL_size_t(unsafe.Sizeof(round_input)), unsafe.Pointer(&round_input[0]), nil)
	round_output_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_WRITE_ONLY,
		cl.CL_size_t(unsafe.Sizeof(round_output)), nil, nil)

	/* Create kernel argument */
	err = cl.CLSetKernelArg(kernel, 0, cl.CL_size_t(unsafe.Sizeof(mod_input_buffer)), unsafe.Pointer(&mod_input_buffer))
	if err < 0 {
		println("Couldn't set a kernel argument")
		return
	}
	cl.CLSetKernelArg(kernel, 1, cl.CL_size_t(unsafe.Sizeof(mod_output_buffer)), unsafe.Pointer(&mod_output_buffer))
	cl.CLSetKernelArg(kernel, 2, cl.CL_size_t(unsafe.Sizeof(round_input_buffer)), unsafe.Pointer(&round_input_buffer))
	cl.CLSetKernelArg(kernel, 3, cl.CL_size_t(unsafe.Sizeof(round_output_buffer)), unsafe.Pointer(&round_output_buffer))

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

	/* Read the results */
	err = cl.CLEnqueueReadBuffer(queue, mod_output_buffer, cl.CL_TRUE, 0,
		cl.CL_size_t(unsafe.Sizeof(mod_output)), unsafe.Pointer(&mod_output), 0, nil, nil)
	if err < 0 {
		println("Couldn't read the buffer")
		return
	}
	cl.CLEnqueueReadBuffer(queue, round_output_buffer, cl.CL_TRUE, 0,
		cl.CL_size_t(unsafe.Sizeof(round_output)), unsafe.Pointer(&round_output), 0, nil, nil)

	/* Display data */
	fmt.Printf("fmod(%.1f, %.1f)      = %.1f\n", mod_input[0], mod_input[1], mod_output[0])
	fmt.Printf("remainder(%.1f, %.1f) = %.1f\n\n", mod_input[0], mod_input[1], mod_output[1])

	fmt.Printf("Rounding input: %.1f %.1f %.1f %.1f\n",
		round_input[0], round_input[1], round_input[2], round_input[3])
	fmt.Printf("rint:  %.1f, %.1f, %.1f, %.1f\n",
		round_output[0], round_output[1], round_output[2], round_output[3])
	fmt.Printf("round: %.1f, %.1f, %.1f, %.1f\n",
		round_output[4], round_output[5], round_output[6], round_output[7])
	fmt.Printf("ceil:  %.1f, %.1f, %.1f, %.1f\n",
		round_output[8], round_output[9], round_output[10], round_output[11])
	fmt.Printf("floor: %.1f, %.1f, %.1f, %.1f\n",
		round_output[12], round_output[13], round_output[14], round_output[15])
	fmt.Printf("trunc: %.1f, %.1f, %.1f, %.1f\n",
		round_output[16], round_output[17], round_output[18], round_output[19])

	/* Deallocate resources */
	cl.CLReleaseMemObject(mod_input_buffer)
	cl.CLReleaseMemObject(mod_output_buffer)
	cl.CLReleaseMemObject(round_input_buffer)
	cl.CLReleaseMemObject(round_output_buffer)
	cl.CLReleaseKernel(kernel)
	cl.CLReleaseCommandQueue(queue)
	cl.CLReleaseProgram(*program)
	cl.CLReleaseContext(context)
}
