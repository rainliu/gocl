package cl

/*
#cgo CFLAGS: -I CL
#cgo LDFLAGS: -lOpenCL

#include "CL/opencl.h"
#include <stdlib.h>
#include <string.h>

extern void go_prg_notify(cl_program program, void *user_data);
static void CL_CALLBACK c_prg_notify(cl_program program, void *user_data) {
	go_prg_notify(program, user_data);
}

static cl_int CLBuildProgram(cl_program program,
							cl_uint num_devices,
							const cl_device_id *devices,
							const char *options,
							void *user_data){

    return clBuildProgram(program, num_devices, devices, options, c_prg_notify, user_data);
}
*/
import "C"

import (
	"unsafe"
)

type CL_prg_notify func(program CL_program, user_data interface{})

var prg_notify CL_prg_notify

//export go_prg_notify
func go_prg_notify(program C.cl_program, user_data unsafe.Pointer) {
	prg_notify(CL_program{program}, user_data)
}

func CLCreateProgramWithSource(context CL_context,
	count CL_uint,
	strings []string,
	lengths []CL_size_t,
	errcode_ret *CL_int) CL_program {

	if count == 0 || len(strings) != int(count) {
		*errcode_ret = CL_INVALID_VALUE
		return CL_program{nil}
	}

	var c_program C.cl_program
	var c_lengths []C.size_t
	var c_strings []*C.char
	var c_errcode_ret C.cl_int

	c_lengths = make([]C.size_t, count)
	c_strings = make([]*C.char, count)
	for i := CL_uint(0); i < count; i++ {
		c_lengths[i] = C.size_t(lengths[i])
		c_strings[i] = C.CString(strings[i])
		defer C.free(unsafe.Pointer(c_strings[i]))
	}

	c_program = C.clCreateProgramWithSource(context.cl_context,
		C.cl_uint(count),
		&c_strings[0],
		&c_lengths[0],
		&c_errcode_ret)

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_program{c_program}
}

func CLCreateProgramWithBinary(context CL_context,
	num_devices CL_uint,
	devices []CL_device_id,
	lengths []CL_size_t,
	binaries [][]byte,
	binary_status []CL_int,
	errcode_ret *CL_int) CL_program {

	if num_devices == 0 ||
		len(devices) != int(num_devices) ||
		len(lengths) != int(num_devices) ||
		len(binaries) != int(num_devices) ||
		len(binary_status) != int(num_devices) {
		*errcode_ret = CL_INVALID_VALUE
		return CL_program{nil}
	}

	var c_program C.cl_program
	var c_devices []C.cl_device_id
	var c_lengths []C.size_t
	var c_binaries []*C.uchar
	var c_status []C.cl_int
	var c_errcode_ret C.cl_int

	c_devices = make([]C.cl_device_id, num_devices)
	c_lengths = make([]C.size_t, num_devices)
	c_binaries = make([]*C.uchar, num_devices)
	c_status = make([]C.cl_int, num_devices)
	for i := CL_uint(0); i < num_devices; i++ {
		c_devices[i] = devices[i].cl_device_id
		c_lengths[i] = C.size_t(lengths[i])
		c_binaries[i] = (*C.uchar)(C.malloc(c_lengths[i]))
		C.memcpy(unsafe.Pointer(c_binaries[i]), unsafe.Pointer(&binaries[i][0]), c_lengths[i])
		defer C.free(unsafe.Pointer(c_binaries[i]))
	}

	c_program = C.clCreateProgramWithBinary(context.cl_context,
		C.cl_uint(num_devices),
		&c_devices[0],
		&c_lengths[0],
		&c_binaries[0],
		&c_status[0],
		&c_errcode_ret)

	for i := CL_uint(0); i < num_devices; i++ {
		binary_status[i] = CL_int(c_status[i])
	}

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_program{c_program}
}

func CLRetainProgram(program CL_program) CL_int {
	return CL_int(C.clRetainProgram(program.cl_program))
}

func CLReleaseProgram(program CL_program) CL_int {
	return CL_int(C.clReleaseProgram(program.cl_program))
}

func CLBuildProgram(program CL_program,
	num_devices CL_uint,
	devices []CL_device_id,
	options string,
	pfn_notify CL_prg_notify,
	user_data interface{}) CL_int {

	if num_devices == 0 || len(devices) != int(num_devices) {
		return CL_INVALID_VALUE
	}

	var c_devices []C.cl_device_id
	var c_errcode_ret C.cl_int
	var c_options *C.char

	c_devices = make([]C.cl_device_id, len(devices))
	for i := 0; i < len(devices); i++ {
		c_devices[i] = C.cl_device_id(devices[i].cl_device_id)
	}
	c_options = C.CString(options)
	defer C.free(unsafe.Pointer(c_options))

	if pfn_notify != nil {
		prg_notify = pfn_notify

		c_errcode_ret = C.CLBuildProgram(program.cl_program,
			C.cl_uint(num_devices),
			&c_devices[0],
			c_options,
			unsafe.Pointer(&user_data))

	} else {
		prg_notify = nil

		c_errcode_ret = C.clBuildProgram(program.cl_program,
			C.cl_uint(num_devices),
			&c_devices[0],
			c_options,
			nil,
			nil)

	}

	return CL_int(c_errcode_ret)
}
