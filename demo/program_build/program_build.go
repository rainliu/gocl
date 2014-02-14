package main

import (
	"fmt"
	"gocl/cl"
)

const (
	NUM_FILES = 2
)

var program_buffer = [NUM_FILES]string{"__kernel void good(__global float *a,\n" +
	"                   __global float *b,\n" +
	"                   __global float *c) {\n" +
	"   *c = *a + *b;\n" +
	"}",
	"__kernel void bad(__global float *a,\n" +
		"                   __global float *b,\n" +
		"                   __global float *c) {\n" +
		"   *c = *a + *b;\n" +
		"}",
}

func main() {

	/* Host/device data structures */
	var platform [1]cl.CL_platform_id
	var device [1]cl.CL_device_id
	var context cl.CL_context
	var i, err cl.CL_int

	/* Program data structures */
	var program cl.CL_program
	//FILE *program_handle;
	var program_log interface{}
	//const char *file_name[] = {PROGRAM_FILE_1, PROGRAM_FILE_2};
	options := "-cl-finite-math-only -cl-no-signed-zeros"
	var program_size [NUM_FILES]cl.CL_size_t
	var log_size cl.CL_size_t

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
	for i = 0; i < NUM_FILES; i++ {
		/*program_handle = fopen(file_name[i], "r");
		  if(program_handle == NULL) {
		     println("Couldn't find the program file");
		     return;
		  }
		  fseek(program_handle, 0, SEEK_END);
		  program_size[i] = ftell(program_handle);*/
		program_size[i] = cl.CL_size_t(len(program_buffer[i]))
		//rewind(program_handle);
		//program_buffer[i] = (char*)malloc(program_size[i]+1);
		//program_buffer[i][program_size[i]] = '\0';
		//fread(program_buffer[i], sizeof(char), program_size[i],
		//      program_handle);
		//fclose(program_handle);
	}

	/* Create a program containing all program content */
	program = cl.CLCreateProgramWithSource(context, NUM_FILES,
		program_buffer[:], program_size[:], &err)
	if err < 0 {
		println("Couldn't create the program")
		return
	}

	/* Build program */
	err = cl.CLBuildProgram(program, 1, device[:], options, nil, nil)
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

	/* Deallocate resources */
	//for(i=0; i<NUM_FILES; i++) {
	//   free(program_buffer[i]);
	//}
	cl.CLReleaseProgram(program)
	cl.CLReleaseContext(context)
}
