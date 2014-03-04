package main

import (
	"fmt"
	"gocl/cl"
	"os"
)

func main() {

	/* Host/device data structures */
	var platform [1]cl.CL_platform_id
	var device [1]cl.CL_device_id
	var context cl.CL_context
	var queue cl.CL_command_queue
	var err cl.CL_int

	/* Program/kernel data structures */
	var program cl.CL_program
	var program_buffer [1][]byte
	var program_log interface{}
	var program_size [1]cl.CL_size_t
	var log_size cl.CL_size_t
	var kernel cl.CL_kernel

	/* Access the first installed platform */
	err = cl.CLGetPlatformIDs(1, platform[:], nil)
	if err < 0 {
		println("Couldn't find any platforms")
		return
	}

	/* Access the first GPU/CPU */
	err = cl.CLGetDeviceIDs(platform[0], cl.CL_DEVICE_TYPE_GPU, 1, device[:], nil)
	if err == cl.CL_DEVICE_NOT_FOUND {
		err = cl.CLGetDeviceIDs(platform[0], cl.CL_DEVICE_TYPE_CPU, 1, device[:], nil)
	}
	if err < 0 {
		println("Couldn't find any devices")
		return
	}

	/* Create a context */
	context = cl.CLCreateContext(nil, 1, device[:], nil, nil, &err)
	if err < 0 {
		println("Couldn't create a context")
		return
	}

	/* Read each program file and place content into buffer array */
	program_handle, err1 := os.Open("blank.cl")
	if err1 != nil {
		println("Couldn't find the program file")
		return
	}
	defer program_handle.Close()

	fi, err2 := program_handle.Stat()
	if err2 != nil {
		println("Couldn't find the program stat")
		return
	}
	program_size[0] = cl.CL_size_t(fi.Size())
	program_buffer[0] = make([]byte, program_size[0])
	read_size, err3 := program_handle.Read(program_buffer[0])
	if err3 != nil || cl.CL_size_t(read_size) != program_size[0] {
		println("read file error or file size wrong")
		return
	}

	/* Create program from file */
	program = cl.CLCreateProgramWithSource(context, 1,
		program_buffer[:], program_size[:], &err)
	if err < 0 {
		println("Couldn't create the program")
		return
	}

	/* Build program */
	err = cl.CLBuildProgram(program, 1, device[:], nil, nil, nil)
	if err < 0 {
		/* Find size of log and print to std output */
		cl.CLGetProgramBuildInfo(program, device[0], cl.CL_PROGRAM_BUILD_LOG,
			0, nil, &log_size)
		//program_log = (char*) malloc(log_size+1);
		//program_log[log_size] = '\0';
		cl.CLGetProgramBuildInfo(program, device[0], cl.CL_PROGRAM_BUILD_LOG,
			log_size, &program_log, nil)
		fmt.Printf("%s\n", program_log)
		//free(program_log);
		return
	}

	/* Create the kernel */
	kernel = cl.CLCreateKernel(program, []byte("blank"), &err)
	if err < 0 {
		println("Couldn't create the kernel")
		return
	}

	/* Create the command queue */
	queue = cl.CLCreateCommandQueue(context, device[0], 0, &err)
	if err < 0 {
		println("Couldn't create the command queue")
		return
	}

	/* Enqueue the kernel execution command */
	err = cl.CLEnqueueTask(queue, kernel, 0, nil, nil)
	if err < 0 {
		println("Couldn't enqueue the kernel execution command")
		return
	} else {
		println("Successfully queued kernel.\n")
	}

	/* Deallocate resources */
	cl.CLReleaseCommandQueue(queue)
	cl.CLReleaseKernel(kernel)
	cl.CLReleaseProgram(program)
	cl.CLReleaseContext(context)
}
