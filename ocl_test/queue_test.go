package ocl_test

import (
	"gocl/cl"
	"gocl/ocl"
	"testing"
)

func TestQueue(t *testing.T) {
	/* Host/device data structures */
	var platforms []ocl.Platform
	var devices []ocl.Device
	var context ocl.Context
	var queue ocl.CommandQueue
	var err error

	var ref_count interface{}

	/* Identify a platform */
	if platforms, err = ocl.GetPlatforms(); err != nil {
		t.Errorf(err.Error())
		return
	}

	/* Determine connected devices */
	if devices, err = platforms[0].GetDevices(cl.CL_DEVICE_TYPE_GPU); err != nil {
		if devices, err = platforms[0].GetDevices(cl.CL_DEVICE_TYPE_CPU); err != nil {
			t.Errorf(err.Error())
			return
		}
	}

	/* Create the context */
	if context, err = devices[0].CreateContext(nil, nil, nil); err != nil {
		t.Errorf(err.Error())
		return
	}
	defer context.Release()

	/* Create the command queue */
	if queue, err = context.CreateCommandQueue(devices[0], nil); err != nil {
		t.Errorf(err.Error())
		return
	}
	defer queue.Release()

	/* Get the reference count */
	if ref_count, err = queue.GetInfo(cl.CL_QUEUE_REFERENCE_COUNT); err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Logf("Initial reference count: %d\n", ref_count.(cl.CL_uint))

	/* Update and display the reference count */
	queue.Retain()
	if ref_count, err = queue.GetInfo(cl.CL_QUEUE_REFERENCE_COUNT); err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Logf("Reference count: %d\n", ref_count.(cl.CL_uint))

	queue.Release()
	if ref_count, err = queue.GetInfo(cl.CL_QUEUE_REFERENCE_COUNT); err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Logf("Reference count: %d\n", ref_count.(cl.CL_uint))

	// /* Program/kernel data structures */
	// var program cl.CL_program
	// var program_buffer [1][]byte
	// var program_log interface{}
	// var program_size [1]cl.CL_size_t
	// var log_size cl.CL_size_t
	// var kernel cl.CL_kernel

	// /* Read each program file and place content into buffer array */
	// program_handle, err1 := os.Open("blank.cl")
	// if err1 != nil {
	// 	t.Errorf("Couldn't find the program file")
	// }
	// defer program_handle.Close()

	// fi, err2 := program_handle.Stat()
	// if err2 != nil {
	// 	t.Errorf("Couldn't find the program stat")
	// }
	// program_size[0] = cl.CL_size_t(fi.Size())
	// program_buffer[0] = make([]byte, program_size[0])
	// read_size, err3 := program_handle.Read(program_buffer[0])
	// if err3 != nil || cl.CL_size_t(read_size) != program_size[0] {
	// 	t.Errorf("read file error or file size wrong")
	// }

	//  Create program from file
	// program = cl.CLCreateProgramWithSource(context, 1,
	// 	program_buffer[:], program_size[:], &err)
	// if err < 0 {
	// 	t.Errorf("Couldn't create the program")
	// }

	// /* Build program */
	// err = cl.CLBuildProgram(program, 1, device[:], nil, nil, nil)
	// if err < 0 {
	// 	/* Find size of log and print to std output */
	// 	cl.CLGetProgramBuildInfo(program, device[0], cl.CL_PROGRAM_BUILD_LOG,
	// 		0, nil, &log_size)
	// 	//program_log = (char*) malloc(log_size+1);
	// 	//program_log[log_size] = '\0';
	// 	cl.CLGetProgramBuildInfo(program, device[0], cl.CL_PROGRAM_BUILD_LOG,
	// 		log_size, &program_log, nil)
	// 	t.Errorf("%s\n", program_log)
	// 	//free(program_log);
	// }

	// /* Create the kernel */
	// kernel = cl.CLCreateKernel(program, []byte("blank"), &err)
	// if err < 0 {
	// 	t.Errorf("Couldn't create the kernel")
	// }

	// /* Enqueue the kernel execution command */
	// err = cl.CLEnqueueTask(queue, kernel, 0, nil, nil)
	// if err < 0 {
	// 	t.Errorf("Couldn't enqueue the kernel execution command")
	// } else {
	// 	t.Logf("Successfully queued kernel.\n")
	// }

	//cl.CLReleaseKernel(kernel)
	//cl.CLReleaseProgram(program)
}
