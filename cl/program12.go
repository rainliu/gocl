// +build cl12 cl20

package cl

/*
#cgo CFLAGS: -I CL
#cgo !darwin LDFLAGS: -lOpenCL
#cgo darwin LDFLAGS: -framework OpenCL

#ifdef __APPLE__
#include "OpenCL/opencl.h"
#else
#include "CL/opencl.h"
#endif
#include <string.h>
#include <stdlib.h>

extern void go_prg_notify(cl_program program, void *user_data);
static void CL_CALLBACK c_prg_compile_notify(cl_program program, void *user_data) {
	go_prg_notify(program, user_data);
}

static cl_int CLCompileProgram(cl_program program,
							cl_uint num_devices,
							const cl_device_id *devices,
							const char *options,
							cl_uint num_input_headers,
							const cl_program *input_headers,
							const char **header_include_names,
							void *user_data){

    return clCompileProgram(program, num_devices, devices, options, num_input_headers, input_headers, header_include_names, c_prg_compile_notify, user_data);
}

extern void go_prg_link_notify(cl_program program, void *user_data);
static void CL_CALLBACK c_prg_link_notify(cl_program program, void *user_data) {
	go_prg_link_notify(program, user_data);
}

static cl_program CLLinkProgram(cl_context context,
							cl_uint num_devices,
							const cl_device_id *devices,
							const char *options,
							cl_uint num_input_programs,
							const cl_program *input_programs,
							void *user_data,
							cl_int *errcode_ret){

    return clLinkProgram(context, num_devices, devices, options, num_input_programs, input_programs, c_prg_link_notify, user_data, errcode_ret);
}
*/
import "C"

import "unsafe"

///////////////////////////////////////////////
//OpenCL 1.2
///////////////////////////////////////////////

var prg_link_notify map[unsafe.Pointer]CL_prg_notify

func init() {
	prg_link_notify = make(map[unsafe.Pointer]CL_prg_notify)
}

//export go_prg_link_notify
func go_prg_link_notify(program C.cl_program, user_data unsafe.Pointer) {
	var c_user_data []unsafe.Pointer
	c_user_data = *(*[]unsafe.Pointer)(user_data)
	prg_link_notify[c_user_data[1]](CL_program{program}, c_user_data[0])
}

func CLCreateProgramWithBuiltInKernels(context CL_context,
	num_devices CL_uint,
	devices []CL_device_id,
	kernel_names []byte,
	errcode_ret *CL_int) CL_program {
	if num_devices == 0 || len(devices) != int(num_devices) {
		if errcode_ret != nil {
			*errcode_ret = CL_INVALID_VALUE
		}
		return CL_program{nil}
	}

	var c_program C.cl_program
	var c_devices []C.cl_device_id
	var c_kernel_names *C.char
	var c_errcode_ret C.cl_int

	c_kernel_names = C.CString(string(kernel_names))
	defer C.free(unsafe.Pointer(c_kernel_names))

	c_devices = make([]C.cl_device_id, num_devices)
	for i := CL_uint(0); i < num_devices; i++ {
		c_devices[i] = devices[i].cl_device_id
	}

	c_program = C.clCreateProgramWithBuiltInKernels(context.cl_context,
		C.cl_uint(num_devices),
		&c_devices[0],
		c_kernel_names,
		&c_errcode_ret)

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_program{c_program}
}

func CLCompileProgram(program CL_program,
	num_devices CL_uint,
	devices []CL_device_id,
	options []byte,
	num_input_headers CL_uint,
	input_headers []CL_program,
	header_include_names [][]byte,
	pfn_notify CL_prg_notify,
	user_data unsafe.Pointer) CL_int {

	if (num_devices == 0 && devices != nil) ||
		(num_devices != 0 && devices == nil) ||
		(int(num_input_headers) != len(input_headers)) ||
		(int(num_input_headers) != len(header_include_names)) ||
		(pfn_notify == nil && user_data != nil) {
		return CL_INVALID_VALUE
	}

	var c_devices []C.cl_device_id
	var c_options *C.char
	var c_input_headers []C.cl_program
	var c_header_include_names []*C.char
	var c_errcode_ret C.cl_int

	c_devices = make([]C.cl_device_id, len(devices))
	for i := 0; i < len(devices); i++ {
		c_devices[i] = C.cl_device_id(devices[i].cl_device_id)
	}
	if options != nil {
		c_options = C.CString(string(options))
		defer C.free(unsafe.Pointer(c_options))
	} else {
		c_options = nil
	}

	c_input_headers = make([]C.cl_program, num_input_headers)
	c_header_include_names = make([]*C.char, num_input_headers)
	for i := CL_uint(0); i < num_input_headers; i++ {
		c_input_headers[i] = input_headers[i].cl_program
		c_header_include_names[i] = C.CString(string(header_include_names[i]))
		defer C.free(unsafe.Pointer(c_header_include_names[i]))
	}

	if pfn_notify != nil {
		prg_notify[program.cl_program] = pfn_notify

		c_errcode_ret = C.CLCompileProgram(program.cl_program,
			C.cl_uint(num_devices),
			&c_devices[0],
			c_options,
			C.cl_uint(num_input_headers),
			&c_input_headers[0],
			&c_header_include_names[0],
			user_data)
	} else {
		c_errcode_ret = C.clCompileProgram(program.cl_program,
			C.cl_uint(num_devices),
			&c_devices[0],
			c_options,
			C.cl_uint(num_input_headers),
			&c_input_headers[0],
			&c_header_include_names[0],
			nil,
			nil)
	}

	return CL_int(c_errcode_ret)
}

func CLLinkProgram(context CL_context,
	num_devices CL_uint,
	devices []CL_device_id,
	options []byte,
	num_input_programs CL_uint,
	input_programs []CL_program,
	pfn_notify CL_prg_notify,
	user_data unsafe.Pointer,
	errcode_ret *CL_int) CL_program {

	if (num_devices == 0 && devices != nil) ||
		(num_devices != 0 && devices == nil) ||
		(int(num_input_programs) != len(input_programs)) ||
		(pfn_notify == nil && user_data != nil) {
		if errcode_ret != nil {
			*errcode_ret = CL_INVALID_VALUE
		}
		return CL_program{nil}
	}

	var c_devices []C.cl_device_id
	var c_options *C.char
	var c_input_programs []C.cl_program
	var c_errcode_ret C.cl_int
	var c_program_ret C.cl_program

	c_devices = make([]C.cl_device_id, len(devices))
	for i := 0; i < len(devices); i++ {
		c_devices[i] = C.cl_device_id(devices[i].cl_device_id)
	}
	if options != nil {
		c_options = C.CString(string(options))
		defer C.free(unsafe.Pointer(c_options))
	} else {
		c_options = nil
	}

	c_input_programs = make([]C.cl_program, num_input_programs)
	for i := CL_uint(0); i < num_input_programs; i++ {
		c_input_programs[i] = input_programs[i].cl_program
	}

	if pfn_notify != nil {
		var c_user_data []unsafe.Pointer
		c_user_data = make([]unsafe.Pointer, 2)
		c_user_data[0] = user_data
		c_user_data[1] = unsafe.Pointer(&pfn_notify)

		prg_link_notify[c_user_data[1]] = pfn_notify

		c_program_ret = C.CLLinkProgram(context.cl_context,
			C.cl_uint(num_devices),
			&c_devices[0],
			c_options,
			C.cl_uint(num_input_programs),
			&c_input_programs[0],
			unsafe.Pointer(&c_user_data),
			&c_errcode_ret)
	} else {
		c_program_ret = C.clLinkProgram(context.cl_context,
			C.cl_uint(num_devices),
			&c_devices[0],
			c_options,
			C.cl_uint(num_input_programs),
			&c_input_programs[0],
			nil,
			nil,
			&c_errcode_ret)
	}

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_program{c_program_ret}
}

func CLUnloadPlatformCompiler(platform CL_platform_id) CL_int {
	return CL_int(C.clUnloadPlatformCompiler(platform.cl_platform_id))
}
