package cl_test

import (
	"gocl/cl"
	"os"
	"testing"
)

const (
	NUM_FILES = 2
)

func TestProgram(t *testing.T) {

	/* Host/device data structures */
	var platform [1]cl.CL_platform_id
	var device [1]cl.CL_device_id
	var context cl.CL_context
	var i, err cl.CL_int

	/* Program data structures */
	var program cl.CL_program
	var program_buffer [NUM_FILES][]byte
	var program_log interface{}
	var file_name = []string{"bad.cl", "good.cl"}
	options := "-cl-finite-math-only -cl-no-signed-zeros"
	var program_size [NUM_FILES]cl.CL_size_t
	var log_size cl.CL_size_t

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
	for i = 0; i < NUM_FILES; i++ {
		program_handle, err := os.Open(file_name[i])
		if err != nil {
			t.Errorf("Couldn't find the program file")
		}
		defer program_handle.Close()

		fi, err2 := program_handle.Stat()
		if err2 != nil {
			t.Errorf("Couldn't find the program stat")
		}
		program_size[i] = cl.CL_size_t(fi.Size())
		program_buffer[i] = make([]byte, program_size[i])
		read_size, err3 := program_handle.Read(program_buffer[i])
		if err3 != nil || cl.CL_size_t(read_size) != program_size[i] {
			t.Errorf("read file error or file size wrong")
		}
	}

	/* Create a program containing all program content */
	program = cl.CLCreateProgramWithSource(context, NUM_FILES,
		program_buffer[:], program_size[:], &err)
	if err < 0 {
		t.Errorf("Couldn't create the program")
	}

	/* Build program */
	err = cl.CLBuildProgram(program, 1, device[:], []byte(options), nil, nil)
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

	/* Deallocate resources */
	//for(i=0; i<NUM_FILES; i++) {
	//   free(program_buffer[i]);
	//}
	cl.CLReleaseProgram(program)
	cl.CLReleaseContext(context)
}
