package ocl_test

import (
	"gocl/cl"
	"gocl/ocl"
	"os"
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
	devices = devices[0:1]

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

	/* Program/kernel data structures */
	var program ocl.Program
	var program_size [1]cl.CL_size_t
	var program_buffer [1][]byte
	var program_log interface{}

	/* Read each program file and place content into buffer array */
	program_handle, err1 := os.Open("blank.cl")
	if err1 != nil {
		t.Errorf(err1.Error())
		return
	}
	defer program_handle.Close()

	fi, err2 := program_handle.Stat()
	if err2 != nil {
		t.Errorf(err2.Error())
		return
	}
	program_size[0] = cl.CL_size_t(fi.Size())
	program_buffer[0] = make([]byte, program_size[0])
	read_size, err3 := program_handle.Read(program_buffer[0])
	if err3 != nil || cl.CL_size_t(read_size) != program_size[0] {
		t.Errorf("read file error or file size wrong")
		return
	}

	// Create program from file
	if program, err = context.CreateProgramWithSource(1, program_buffer[:], program_size[:]); err != nil {
		t.Errorf(err.Error())
		return
	}
	defer program.Release()

	/* Build program */
	if err = program.Build(devices, nil, nil, nil); err != nil {
		t.Errorf(err.Error())
		/* Find size of log and print to std output */
		if program_log, err = program.GetBuildInfo(devices[0], cl.CL_PROGRAM_BUILD_LOG); err != nil {
			t.Errorf(err.Error())
		} else {
			t.Errorf("%s\n", program_log.(string))
		}
		return
	}

	//var kernel cl.CL_kernel
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

}
