package main

import (
	"gocl/cl"
	"os"
	"testing"
)

func TestQueue(t *testing.T) {

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
		t.Errorf("Couldn't find any platforms")
	}

	/* Access the first GPU/CPU */
	err = cl.CLGetDeviceIDs(platform[0], cl.CL_DEVICE_TYPE_GPU, 1, device[:], nil)
	if err == cl.CL_DEVICE_NOT_FOUND {
		err = cl.CLGetDeviceIDs(platform[0], cl.CL_DEVICE_TYPE_CPU, 1, device[:], nil)
	}
	if err < 0 {
		t.Errorf("Couldn't find any devices")
	}

	/* Create a context */
	context = cl.CLCreateContext(nil, 1, device[:], nil, nil, &err)
	if err < 0 {
		t.Errorf("Couldn't create a context")
	}

	/* Read each program file and place content into buffer array */
	program_handle, err1 := os.Open("blank.cl")
	if err1 != nil {
		t.Errorf("Couldn't find the program file")
	}
	defer program_handle.Close()

	fi, err2 := program_handle.Stat()
	if err2 != nil {
		t.Errorf("Couldn't find the program stat")
	}
	program_size[0] = cl.CL_size_t(fi.Size())
	program_buffer[0] = make([]byte, program_size[0])
	read_size, err3 := program_handle.Read(program_buffer[0])
	if err3 != nil || cl.CL_size_t(read_size) != program_size[0] {
		t.Errorf("read file error or file size wrong")
	}

	/* Create program from file */
	program = cl.CLCreateProgramWithSource(context, 1,
		program_buffer[:], program_size[:], &err)
	if err < 0 {
		t.Errorf("Couldn't create the program")
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
		t.Errorf("%s\n", program_log)
		//free(program_log);
	}

	/* Create the kernel */
	kernel = cl.CLCreateKernel(program, []byte("blank"), &err)
	if err < 0 {
		t.Errorf("Couldn't create the kernel")
	}

	/* Create the command queue */
	queue = cl.CLCreateCommandQueue(context, device[0], 0, &err)
	if err < 0 {
		t.Errorf("Couldn't create the command queue")
	}

	/* Enqueue the kernel execution command */
	err = cl.CLEnqueueTask(queue, kernel, 0, nil, nil)
	if err < 0 {
		t.Errorf("Couldn't enqueue the kernel execution command")
	} else {
		t.Logf("Successfully queued kernel.\n")
	}

	/* Deallocate resources */
	cl.CLReleaseCommandQueue(queue)
	cl.CLReleaseKernel(kernel)
	cl.CLReleaseProgram(program)
	cl.CLReleaseContext(context)
}
