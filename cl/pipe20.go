// +build cl20

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
*/
import "C"

import "unsafe"

func CLCreatePipe(context CL_context,
	flags CL_mem_flags,
	pipe_packet_size CL_uint,
	pipe_max_packets CL_uint,
	properties []CL_pipe_properties,
	errcode_ret *CL_int) CL_mem {

	var c_errcode_ret C.cl_int
	var c_memobj C.cl_mem

	var c_properties []C.cl_pipe_properties
	var c_properties_ptr *C.cl_pipe_properties

	if properties != nil {
		c_properties = make([]C.cl_pipe_properties, len(properties))
		for i := 0; i < len(properties); i++ {
			c_properties[i] = C.cl_pipe_properties(properties[i])
		}
		c_properties_ptr = &c_properties[0]
	} else {
		c_properties_ptr = nil
	}

	c_memobj = C.clCreatePipe(context.cl_context,
		C.cl_mem_flags(flags),
		C.cl_uint(pipe_packet_size),
		C.cl_uint(pipe_max_packets),
		c_properties_ptr,
		&c_errcode_ret)

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_mem{c_memobj}
}

func CLGetPipeInfo(pipe CL_mem,
	param_name CL_pipe_info,
	param_value_size CL_size_t,
	param_value *interface{},
	param_value_size_ret *CL_size_t) CL_int {

	if (param_value_size == 0 || param_value == nil) && param_value_size_ret == nil {
		return CL_INVALID_VALUE
	} else {
		var c_param_value_size_ret C.size_t
		var c_errcode_ret C.cl_int

		if param_value_size == 0 || param_value == nil {
			c_errcode_ret = C.clGetPipeInfo(pipe.cl_mem,
				C.cl_pipe_info(param_name),
				C.size_t(param_value_size),
				nil,
				&c_param_value_size_ret)
		} else {
			switch param_name {
			case CL_PIPE_PACKET_SIZE,
				CL_PIPE_MAX_PACKETS:
				var value C.cl_uint
				c_errcode_ret = C.clGetPipeInfo(pipe.cl_mem,
					C.cl_pipe_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_uint(value)
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
