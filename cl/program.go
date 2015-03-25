// +build cl11 cl12 cl20

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
static void CL_CALLBACK c_prg_build_notify(cl_program program, void *user_data) {
	go_prg_notify(program, user_data);
}

static cl_int CLBuildProgram(cl_program program,
							cl_uint num_devices,
							const cl_device_id *devices,
							const char *options,
							void *user_data){

    return clBuildProgram(program, num_devices, devices, options, c_prg_build_notify, user_data);
}
*/
import "C"

import (
	"unsafe"
)

type CL_prg_notify func(program CL_program, user_data unsafe.Pointer)

var prg_notify map[C.cl_program]CL_prg_notify

func init() {
	prg_notify = make(map[C.cl_program]CL_prg_notify)
}

//export go_prg_notify
func go_prg_notify(program C.cl_program, user_data unsafe.Pointer) {
	prg_notify[program](CL_program{program}, user_data)
}

func CLCreateProgramWithSource(context CL_context,
	count CL_uint,
	strings [][]byte,
	lengths []CL_size_t,
	errcode_ret *CL_int) CL_program {

	if count == 0 || len(strings) != int(count) || len(lengths) != int(count) {
		if errcode_ret != nil {
			*errcode_ret = CL_INVALID_VALUE
		}
		return CL_program{nil}
	}

	for i := 0; i < int(count); i++ {
		if strings[i] == nil || lengths[i] == 0 {
			if errcode_ret != nil {
				*errcode_ret = CL_INVALID_VALUE
			}
			return CL_program{nil}
		}
	}

	var c_program C.cl_program
	var c_lengths []C.size_t
	var c_strings []*C.char
	var c_errcode_ret C.cl_int

	c_lengths = make([]C.size_t, count)
	c_strings = make([]*C.char, count)
	for i := 0; i < int(count); i++ {
		c_lengths[i] = C.size_t(lengths[i])
		c_strings[i] = C.CString(string(strings[i]))
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
		if errcode_ret != nil {
			*errcode_ret = CL_INVALID_VALUE
		}
		return CL_program{nil}
	}

	for i := 0; i < int(num_devices); i++ {
		if binaries[i] == nil || lengths[i] == 0 {
			if errcode_ret != nil {
				*errcode_ret = CL_INVALID_VALUE
			}
			return CL_program{nil}
		}
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
	options []byte,
	pfn_notify CL_prg_notify,
	user_data unsafe.Pointer) CL_int {

	if (num_devices == 0 && devices != nil) ||
		(num_devices != 0 && devices == nil) ||
		(pfn_notify == nil && user_data != nil) {
		return CL_INVALID_VALUE
	}

	var c_devices []C.cl_device_id
	var c_options *C.char
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

	if pfn_notify != nil {
		prg_notify[program.cl_program] = pfn_notify

		c_errcode_ret = C.CLBuildProgram(program.cl_program,
			C.cl_uint(num_devices),
			&c_devices[0],
			c_options,
			user_data)
	} else {
		c_errcode_ret = C.clBuildProgram(program.cl_program,
			C.cl_uint(num_devices),
			&c_devices[0],
			c_options,
			nil,
			nil)
	}

	return CL_int(c_errcode_ret)
}

func CLGetProgramInfo(program CL_program,
	param_name CL_program_info,
	param_value_size CL_size_t,
	param_value *interface{},
	param_value_size_ret *CL_size_t) CL_int {

	if (param_value_size == 0 || param_value == nil) && param_value_size_ret == nil {
		return CL_INVALID_VALUE
	} else {
		var c_param_value_size_ret C.size_t
		var c_errcode_ret C.cl_int

		if param_value_size == 0 || param_value == nil {
			c_errcode_ret = C.clGetProgramInfo(program.cl_program,
				C.cl_program_info(param_name),
				C.size_t(param_value_size),
				nil,
				&c_param_value_size_ret)
		} else {
			switch param_name {
			case CL_PROGRAM_SOURCE:

				value := make([]C.char, param_value_size)
				c_errcode_ret = C.clGetProgramInfo(program.cl_program,
					C.cl_program_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value[0]),
					&c_param_value_size_ret)

				*param_value = C.GoStringN(&value[0], C.int(c_param_value_size_ret-1))

			case CL_PROGRAM_REFERENCE_COUNT,
				CL_PROGRAM_NUM_DEVICES:

				var value C.cl_uint
				c_errcode_ret = C.clGetProgramInfo(program.cl_program,
					C.cl_program_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

			case CL_PROGRAM_CONTEXT:

				var value C.cl_context
				c_errcode_ret = C.clGetProgramInfo(program.cl_program,
					C.cl_program_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

			case CL_PROGRAM_DEVICES:
				var param C.cl_device_id
				length := int(C.size_t(param_value_size) / C.size_t(unsafe.Sizeof(param)))

				value1 := make([]C.cl_device_id, length)
				value2 := make([]CL_device_id, length)

				c_errcode_ret = C.clGetProgramInfo(program.cl_program,
					C.cl_program_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value1[0]),
					&c_param_value_size_ret)

				for i := 0; i < length; i++ {
					value2[i].cl_device_id = value1[i]
				}

				*param_value = value2

			case CL_PROGRAM_BINARY_SIZES:
				var param C.size_t
				length := int(C.size_t(param_value_size) / C.size_t(unsafe.Sizeof(param)))

				value1 := make([]C.size_t, length)
				value2 := make([]CL_size_t, length)

				c_errcode_ret = C.clGetProgramInfo(program.cl_program,
					C.cl_program_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value1[0]),
					&c_param_value_size_ret)

				for i := 0; i < length; i++ {
					value2[i] = CL_size_t(value1[i])
				}

				*param_value = value2

			case CL_PROGRAM_BINARIES:
				var param *C.uchar
				length := int(C.size_t(param_value_size) / C.size_t(unsafe.Sizeof(param)))

				value1 := make([]*C.uchar, length)
				value2 := make([]*CL_uchar, length)

				c_errcode_ret = C.clGetProgramInfo(program.cl_program,
					C.cl_program_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value1[0]),
					&c_param_value_size_ret)

				for i := 0; i < length; i++ {
					value2[i] = (*CL_uchar)(value1[i])
				}

				*param_value = value2

			default:
				return CL_INVALID_VALUE
			}
		}

		if param_value_size_ret != nil {
			*param_value_size_ret = CL_size_t(c_param_value_size_ret)
		}

		return CL_int(c_errcode_ret)
	}
}

func CLGetProgramBuildInfo(program CL_program,
	device CL_device_id,
	param_name CL_program_build_info,
	param_value_size CL_size_t,
	param_value *interface{},
	param_value_size_ret *CL_size_t) CL_int {

	if (param_value_size == 0 || param_value == nil) && param_value_size_ret == nil {
		return CL_INVALID_VALUE
	} else {
		var c_param_value_size_ret C.size_t
		var c_errcode_ret C.cl_int

		if param_value_size == 0 || param_value == nil {
			c_errcode_ret = C.clGetProgramBuildInfo(program.cl_program,
				device.cl_device_id,
				C.cl_program_build_info(param_name),
				C.size_t(param_value_size),
				nil,
				&c_param_value_size_ret)
		} else {
			switch param_name {
			case CL_PROGRAM_BUILD_STATUS:
				var value C.cl_build_status

				c_errcode_ret = C.clGetProgramBuildInfo(program.cl_program,
					device.cl_device_id,
					C.cl_program_build_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_build_status(value)

			case CL_PROGRAM_BUILD_OPTIONS,
				CL_PROGRAM_BUILD_LOG:

				value := make([]C.char, param_value_size)
				c_errcode_ret = C.clGetProgramBuildInfo(program.cl_program,
					device.cl_device_id,
					C.cl_program_build_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value[0]),
					&c_param_value_size_ret)

				*param_value = C.GoStringN(&value[0], C.int(c_param_value_size_ret-1))

			case CL_PROGRAM_BINARY_TYPE:
				var value C.cl_program_binary_type

				c_errcode_ret = C.clGetProgramBuildInfo(program.cl_program,
					device.cl_device_id,
					C.cl_program_build_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_program_binary_type(value)

			case CL_PROGRAM_BUILD_GLOBAL_VARIABLE_TOTAL_SIZE:
				var value C.size_t

				c_errcode_ret = C.clGetProgramBuildInfo(program.cl_program,
					device.cl_device_id,
					C.cl_program_build_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_size_t(value)

			default:
				return CL_INVALID_VALUE
			}
		}

		if param_value_size_ret != nil {
			*param_value_size_ret = CL_size_t(c_param_value_size_ret)
		}

		return CL_int(c_errcode_ret)
	}
}
