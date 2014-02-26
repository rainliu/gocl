// +build opencl1.1 opencl1.2

package cl

/*
#cgo CFLAGS: -I CL
#cgo LDFLAGS: -lOpenCL

#include "CL/opencl.h"
*/
import "C"

import (
	"unsafe"
)

func CLCreateCommandQueue(context CL_context,
	device CL_device_id,
	properties CL_command_queue_properties,
	errcode_ret *CL_int) CL_command_queue {
	var c_errcode_ret C.cl_int
	var c_command_queue C.cl_command_queue

	c_command_queue = C.clCreateCommandQueue(context.cl_context,
		device.cl_device_id,
		C.cl_command_queue_properties(properties),
		&c_errcode_ret)

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_command_queue{c_command_queue}
}

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

	var ret C.cl_int

	if (param_value_size == 0 || param_value == nil) && param_value_size_ret == nil {
		ret = C.clGetCommandQueueInfo(command_queue.cl_command_queue,
			C.cl_command_queue_info(param_name),
			0,
			nil,
			nil)
	} else {
		var size_ret C.size_t

		if param_value_size == 0 || param_value == nil {
			ret = C.clGetCommandQueueInfo(command_queue.cl_command_queue,
				C.cl_command_queue_info(param_name),
				0,
				nil,
				&size_ret)
		} else {
			switch param_name {
			case CL_QUEUE_CONTEXT:
				var value C.cl_context

				ret = C.clGetCommandQueueInfo(command_queue.cl_command_queue,
					C.cl_command_queue_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&size_ret)

				*param_value = CL_context{value}

			case CL_QUEUE_DEVICE:
				var value C.cl_device_id

				ret = C.clGetCommandQueueInfo(command_queue.cl_command_queue,
					C.cl_command_queue_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&size_ret)

				*param_value = CL_device_id{value}

			case CL_QUEUE_REFERENCE_COUNT:
				//CL_QUEUE_SIZE: //OPENCL 2.0
				var value C.cl_uint

				ret = C.clGetCommandQueueInfo(command_queue.cl_command_queue,
					C.cl_command_queue_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&size_ret)

				*param_value = CL_uint(value)

			case CL_QUEUE_PROPERTIES:
				var value C.cl_command_queue_properties

				ret = C.clGetCommandQueueInfo(command_queue.cl_command_queue,
					C.cl_command_queue_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&size_ret)

				*param_value = CL_command_queue_properties(value)
			default:
				return CL_INVALID_VALUE
			}
		}

		if param_value_size_ret != nil {
			*param_value_size_ret = CL_size_t(size_ret)
		}

	}

	return CL_int(ret)
}

func CLFlush(command_queue CL_command_queue) CL_int {
	return CL_int(C.clFlush(command_queue.cl_command_queue))
}

func CLFinish(command_queue CL_command_queue) CL_int {
	return CL_int(C.clFinish(command_queue.cl_command_queue))
}
