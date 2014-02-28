package main

import (
	"gocl/cl"
	"os"
	"testing"
	"unsafe"
)

func TestMatvec(t *testing.T) {
	/* Host/device data structures */
	var platform [1]cl.CL_platform_id
	var device [1]cl.CL_device_id
	var context cl.CL_context
	var queue cl.CL_command_queue
	var i, err cl.CL_int

	/* Program/kernel data structures */
	var program cl.CL_program
	var program_buffer [1][]byte
	var program_log interface{}
	var program_size [1]cl.CL_size_t
	var log_size cl.CL_size_t
	var kernel cl.CL_kernel

	/* Data and buffers */
	var mat [16]float32
	var vec, result [4]float32
	var correct = [4]float32{0.0, 0.0, 0.0, 0.0}
	var mat_buff, vec_buff, res_buff cl.CL_mem

	/* Initialize data to be processed by the kernel */
	for i = 0; i < 16; i++ {
		mat[i] = float32(i) * 2.0
	}

	for i = 0; i < 4; i++ {
		vec[i] = float32(i) * 3.0
		correct[0] += mat[i] * vec[i]
		correct[1] += mat[i+4] * vec[i]
		correct[2] += mat[i+8] * vec[i]
		correct[3] += mat[i+12] * vec[i]
	}

	/* Identify a platform */
	err = cl.CLGetPlatformIDs(1, platform[:], nil)
	if err < 0 {
		t.Errorf("Couldn't find any platforms")
	}

	/* Access a device */
	err = cl.CLGetDeviceIDs(platform[0], cl.CL_DEVICE_TYPE_GPU, 1, device[:], nil)
	if err < 0 {
		err = cl.CLGetDeviceIDs(platform[0], cl.CL_DEVICE_TYPE_CPU, 1, device[:], nil)
		if err < 0 {
			t.Errorf("Couldn't find any devices")
		}
	}

	/* Create the context */
	context = cl.CLCreateContext(nil, 1, device[:], nil, nil, &err)
	if err < 0 {
		t.Errorf("Couldn't create a context")
	}

	/* Read program file and place content into buffer */
	program_handle, err1 := os.Open("matvec.cl")
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

	/* Create a program containing all program content */
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

	/* Create kernel for the mat_vec_mult function */
	kernel = cl.CLCreateKernel(program, []byte("matvec_mult"), &err)
	if err < 0 {
		t.Errorf("Couldn't create the kernel")
		return
	}

	/* Create CL buffers to hold input and output data */
	mat_buff = cl.CLCreateBuffer(context, cl.CL_MEM_READ_ONLY|
		cl.CL_MEM_COPY_HOST_PTR, cl.CL_size_t(unsafe.Sizeof(mat)), unsafe.Pointer(&mat[0]), &err)
	if err < 0 {
		t.Errorf("Couldn't create a buffer object")
		return
	}
	vec_buff = cl.CLCreateBuffer(context, cl.CL_MEM_READ_ONLY|
		cl.CL_MEM_COPY_HOST_PTR, cl.CL_size_t(unsafe.Sizeof(vec)), unsafe.Pointer(&vec[0]), nil)
	res_buff = cl.CLCreateBuffer(context, cl.CL_MEM_WRITE_ONLY,
		cl.CL_size_t(unsafe.Sizeof(result)), nil, nil)

	/* Create kernel arguments from the CL buffers */
	err = cl.CLSetKernelArg(kernel, 0, cl.CL_size_t(unsafe.Sizeof(mat_buff)), unsafe.Pointer(&mat_buff))
	if err < 0 {
		t.Errorf("Couldn't set the kernel argument")
		return
	}
	cl.CLSetKernelArg(kernel, 1, cl.CL_size_t(unsafe.Sizeof(vec_buff)), unsafe.Pointer(&vec_buff))
	cl.CLSetKernelArg(kernel, 2, cl.CL_size_t(unsafe.Sizeof(res_buff)), unsafe.Pointer(&res_buff))

	/* Create a CL command queue for the device*/
	queue = cl.CLCreateCommandQueue(context, device[0], 0, &err)
	if err < 0 {
		t.Errorf("Couldn't create the command queue, errcode=%d\n", err)
		return
	}

	/* Enqueue the command queue to the device */
	var work_units_per_kernel = [1]cl.CL_size_t{4} /* 4 work-units per kernel */
	err = cl.CLEnqueueNDRangeKernel(queue, kernel, 1, nil, work_units_per_kernel[:],
		nil, 0, nil, nil)
	if err < 0 {
		t.Errorf("Couldn't enqueue the kernel execution command, errcode=%d\n", err)
		return
	}

	/* Read the result */
	err = cl.CLEnqueueReadBuffer(queue, res_buff, cl.CL_TRUE, 0, cl.CL_size_t(unsafe.Sizeof(result)),
		unsafe.Pointer(&result[0]), 0, nil, nil)
	if err < 0 {
		t.Errorf("Couldn't enqueue the read buffer command")
		return
	}

	/* Test the result */
	if (result[0] == correct[0]) && (result[1] == correct[1]) &&
		(result[2] == correct[2]) && (result[3] == correct[3]) {
		t.Logf("Matrix-vector multiplication successful.")
	} else {
		t.Errorf("Matrix-vector multiplication unsuccessful.")
	}

	/* Deallocate resources */
	cl.CLReleaseMemObject(mat_buff)
	cl.CLReleaseMemObject(vec_buff)
	cl.CLReleaseMemObject(res_buff)
	cl.CLReleaseKernel(kernel)
	cl.CLReleaseCommandQueue(queue)
	cl.CLReleaseProgram(program)
	cl.CLReleaseContext(context)
}
