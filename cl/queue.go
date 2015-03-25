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
*/
import "C"

import (
	"unsafe"
)

func CLRetainCommandQueue(command_queue CL_command_queue) CL_int {
	return CL_int(C.clRetainCommandQueue(command_queue.cl_command_queue))
}

func CLReleaseCommandQueue(command_queue CL_command_queue) CL_int {
	return CL_int(C.clReleaseCommandQueue(command_queue.cl_command_queue))
}

func CLGetCommandQueueInfo(command_queue CL_command_queue,
	param_name CL_command_queue_info,
	param_value_size CL_size_t,
	param_value *interface{},
	param_value_size_ret *CL_size_t) CL_int {

	if (param_value_size == 0 || param_value == nil) && param_value_size_ret == nil {
		return CL_INVALID_VALUE
	} else {
		var c_param_value_size_ret C.size_t
		var c_errcode_ret C.cl_int

		if param_value_size == 0 || param_value == nil {
			c_errcode_ret = C.clGetCommandQueueInfo(command_queue.cl_command_queue,
				C.cl_command_queue_info(param_name),
				C.size_t(param_value_size),
				nil,
				&c_param_value_size_ret)
		} else {
			switch param_name {
			case CL_QUEUE_CONTEXT:
				var value C.cl_context

				c_errcode_ret = C.clGetCommandQueueInfo(command_queue.cl_command_queue,
					C.cl_command_queue_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_context{value}

			case CL_QUEUE_DEVICE:
				var value C.cl_device_id

				c_errcode_ret = C.clGetCommandQueueInfo(command_queue.cl_command_queue,
					C.cl_command_queue_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_device_id{value}

			case CL_QUEUE_REFERENCE_COUNT,
				CL_QUEUE_SIZE:
				var value C.cl_uint

				c_errcode_ret = C.clGetCommandQueueInfo(command_queue.cl_command_queue,
					C.cl_command_queue_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_uint(value)

			case CL_QUEUE_PROPERTIES:
				var value C.cl_command_queue_properties

				c_errcode_ret = C.clGetCommandQueueInfo(command_queue.cl_command_queue,
					C.cl_command_queue_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_command_queue_properties(value)
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

func CLFlush(command_queue CL_command_queue) CL_int {
	return CL_int(C.clFlush(command_queue.cl_command_queue))
}

func CLFinish(command_queue CL_command_queue) CL_int {
	return CL_int(C.clFinish(command_queue.cl_command_queue))
}
