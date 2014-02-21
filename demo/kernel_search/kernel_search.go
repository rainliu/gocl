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
	var err cl.CL_int

	/* Program data structures */
	var program cl.CL_program
	var program_buffer [1][]byte
	var program_log interface{}
	var program_size [1]cl.CL_size_t
	var log_size cl.CL_size_t
	var kernels []cl.CL_kernel
	var found bool
	var i, num_kernels cl.CL_uint

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
	program_handle, err1 := os.Open("test.cl")
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

	/* Create a program containing all program content */
	program = cl.CLCreateProgramWithSource(context, 1,
		program_buffer[:], program_size[:], &err)
	if err < 0 {
		println("Couldn't create the program")
		return
	}

	/* Build program */
	err = cl.CLBuildProgram(program, 1, device[:], "", nil, nil)
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

	/* Find out how many kernels are in the source file */
	err = cl.CLCreateKernelsInProgram(program, 0, nil, &num_kernels)
	if err < 0 {
		println("Couldn't find any kernels")
		return
	} else {
		fmt.Printf("num_kernels = %d\n", num_kernels)
	}

	/* Create a kernel for each function */
	kernels = make([]cl.CL_kernel, num_kernels)
	err = cl.CLCreateKernelsInProgram(program, num_kernels, kernels, nil)
	if err < 0 {
		println("Couldn't create kernels")
		return
	}

	/* Search for the named kernel */
	for i = 0; i < num_kernels; i++ {
		var kernel_name_size cl.CL_size_t
		var kernel_name interface{}

		err = cl.CLGetKernelInfo(kernels[i], cl.CL_KERNEL_FUNCTION_NAME,
			0, nil, &kernel_name_size)
		if err < 0 {
			fmt.Printf("Couldn't get kernel size of name, errcode=%d\n", err)
			return
		}
		err = cl.CLGetKernelInfo(kernels[i], cl.CL_KERNEL_FUNCTION_NAME,
			kernel_name_size, &kernel_name, nil)
		if err < 0 {
			fmt.Printf("Couldn't get kernel info of name, errcode=%d\n", err)
			return
		}
		if kernel_name.(string) == "mult" {
			found = true
			fmt.Printf("Found mult kernel at index %d.\n", i)
			break
		}
	}
	if !found {
		fmt.Printf("Not found mult kernel\n")
	}

	for i = 0; i < num_kernels; i++ {
		cl.CLReleaseKernel(kernels[i])
	}

	cl.CLReleaseProgram(program)
	cl.CLReleaseContext(context)
}
