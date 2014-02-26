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

func CLCreateKernel(program CL_program,
	kernel_name string,
	errcode_ret *CL_int) CL_kernel {

	var c_errcode_ret C.cl_int
	var c_kernel C.cl_kernel

	var c_kernel_name *C.char

	c_kernel_name = C.CString(kernel_name)
	defer C.free(unsafe.Pointer(c_kernel_name))

	c_kernel = C.clCreateKernel(program.cl_program,
		c_kernel_name, &c_errcode_ret)

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_kernel{c_kernel}
}

func CLCreateKernelsInProgram(program CL_program,
	num_kernels CL_uint,
	kernels []CL_kernel,
	num_kernels_ret *CL_uint) CL_int {

	if (num_kernels == 0 || num_kernels != CL_uint(len(kernels))) && num_kernels_ret == nil {
		return CL_INVALID_VALUE
	} else {
		var c_ret C.cl_int
		var c_num_kernels_ret C.cl_uint

		if num_kernels == 0 || kernels == nil {
			c_ret = C.clCreateKernelsInProgram(program.cl_program,
				0, nil, &c_num_kernels_ret)
		} else {
			var c_kernels []C.cl_kernel

			c_kernels = make([]C.cl_kernel, num_kernels)

			c_ret = C.clCreateKernelsInProgram(program.cl_program,
				C.cl_uint(num_kernels),
				&c_kernels[0],
				&c_num_kernels_ret)

			if c_ret == C.CL_SUCCESS {
				for i := 0; i < int(num_kernels); i++ {
					kernels[i].cl_kernel = c_kernels[i]
				}
			}
		}
		if num_kernels_ret != nil {
			*num_kernels_ret = CL_uint(c_num_kernels_ret)
		}

		return CL_int(c_ret)
	}
}

func CLRetainKernel(kernel CL_kernel) CL_int {
	return CL_int(C.clRetainKernel(kernel.cl_kernel))
}

func CLReleaseKernel(kernel CL_kernel) CL_int {
	return CL_int(C.clReleaseKernel(kernel.cl_kernel))
}

func CLSetKernelArg(kernel CL_kernel,
	arg_index CL_uint,
	arg_size CL_size_t,
	arg_value unsafe.Pointer) CL_int {
	/*
		var c_ptr unsafe.Pointer

		switch value := arg_value.(type) {
		case *CL_mem:
			c_ptr = unsafe.Pointer(&value.cl_mem)
		case *CL_sampler:
			c_ptr = unsafe.Pointer(&value.cl_sampler)
		case float32:
			f := C.float(value)
			c_ptr = unsafe.Pointer(&f)
		default:
			return CL_INVALID_VALUE
		}
	*/
	return CL_int(C.clSetKernelArg(kernel.cl_kernel, C.cl_uint(arg_index), C.size_t(arg_size), arg_value))
}

func CLGetKernelInfo(kernel CL_kernel,
	param_name CL_kernel_info,
	param_value_size CL_size_t,
	param_value *interface{},
	param_value_size_ret *CL_size_t) CL_int {

	var ret C.cl_int

	if (param_value_size == 0 || param_value == nil) && param_value_size_ret == nil {
		ret = C.clGetKernelInfo(kernel.cl_kernel,
			C.cl_kernel_info(param_name),
			0,
			nil,
			nil)
	} else {
		var size_ret C.size_t

		if param_value_size == 0 || param_value == nil {
			ret = C.clGetKernelInfo(kernel.cl_kernel,
				C.cl_kernel_info(param_name),
				0,
				nil,
				&size_ret)
		} else {
			switch param_name {

			case CL_KERNEL_FUNCTION_NAME:

				value := make([]C.char, param_value_size)
				ret = C.clGetKernelInfo(kernel.cl_kernel,
					C.cl_kernel_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value[0]),
					&size_ret)

				*param_value = C.GoStringN(&value[0], C.int(size_ret-1))

			case CL_KERNEL_NUM_ARGS,
				CL_KERNEL_REFERENCE_COUNT:

				var value C.cl_uint
				ret = C.clGetKernelInfo(kernel.cl_kernel,
					C.cl_kernel_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&size_ret)

				*param_value = CL_uint(value)

			case CL_KERNEL_CONTEXT:
				var value C.cl_context
				ret = C.clGetKernelInfo(kernel.cl_kernel,
					C.cl_kernel_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&size_ret)

				*param_value = CL_context{value}

			case CL_KERNEL_PROGRAM:
				var value C.cl_program
				ret = C.clGetKernelInfo(kernel.cl_kernel,
					C.cl_kernel_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&size_ret)

				*param_value = CL_program{value}

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

func CLGetKernelWorkGroupInfo(kernel CL_kernel,
	device CL_device_id,
	param_name CL_kernel_work_group_info,
	param_value_size CL_size_t,
	param_value *interface{},
	param_value_size_ret *CL_size_t) CL_int {

	var ret C.cl_int

	if (param_value_size == 0 || param_value == nil) && param_value_size_ret == nil {
		ret = C.clGetKernelWorkGroupInfo(kernel.cl_kernel,
			device.cl_device_id,
			C.cl_kernel_work_group_info(param_name),
			0,
			nil,
			nil)
	} else {
		var size_ret C.size_t

		if param_value_size == 0 || param_value == nil {
			ret = C.clGetKernelWorkGroupInfo(kernel.cl_kernel,
				device.cl_device_id,
				C.cl_kernel_work_group_info(param_name),
				0,
				nil,
				&size_ret)
		} else {
			switch param_name {

			case CL_KERNEL_GLOBAL_WORK_SIZE,
				CL_KERNEL_WORK_GROUP_SIZE,
				CL_KERNEL_PREFERRED_WORK_GROUP_SIZE_MULTIPLE:

				var value C.size_t
				ret = C.clGetKernelWorkGroupInfo(kernel.cl_kernel,
					device.cl_device_id,
					C.cl_kernel_work_group_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&size_ret)

				*param_value = CL_size_t(value)

			case CL_KERNEL_LOCAL_MEM_SIZE,
				CL_KERNEL_PRIVATE_MEM_SIZE:
				var value C.cl_ulong
				ret = C.clGetKernelWorkGroupInfo(kernel.cl_kernel,
					device.cl_device_id,
					C.cl_kernel_work_group_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&size_ret)

				*param_value = CL_ulong(value)

			case CL_KERNEL_COMPILE_WORK_GROUP_SIZE:
				var value [3]C.size_t
				ret = C.clGetKernelWorkGroupInfo(kernel.cl_kernel,
					device.cl_device_id,
					C.cl_kernel_work_group_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value[0]),
					&size_ret)

				*param_value = [3]CL_size_t{CL_size_t(value[0]), CL_size_t(value[1]), CL_size_t(value[2])}

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

func CLEnqueueNDRangeKernel(command_queue CL_command_queue,
	kernel CL_kernel,
	work_dim CL_uint,
	global_work_offset []CL_size_t,
	global_work_size []CL_size_t,
	local_work_size []CL_size_t,
	num_events_in_wait_list CL_uint,
	event_wait_list []CL_event,
	event *CL_event) CL_int {

	if work_dim > 3 || work_dim < 1 {
		return CL_INVALID_WORK_DIMENSION
	}

	if num_events_in_wait_list != 0 && len(event_wait_list) != int(num_events_in_wait_list) {
		return CL_INVALID_VALUE
	}

	var c_global_work_offset_ptr, c_global_work_size_ptr, c_local_work_size_ptr *C.size_t
	var c_global_work_offset, c_global_work_size, c_local_work_size []C.size_t
	var c_event C.cl_event
	var c_ret C.cl_int

	if global_work_offset != nil {
		c_global_work_offset = make([]C.size_t, len(global_work_offset))
		for i := 0; i < len(global_work_offset); i++ {
			c_global_work_offset[i] = C.size_t(global_work_offset[i])
		}
		c_global_work_offset_ptr = &c_global_work_offset[0]
	} else {
		c_global_work_offset_ptr = nil
	}

	if global_work_size != nil {
		c_global_work_size = make([]C.size_t, len(global_work_size))
		for i := 0; i < len(global_work_size); i++ {
			c_global_work_size[i] = C.size_t(global_work_size[i])
		}
		c_global_work_size_ptr = &c_global_work_size[0]
	} else {
		c_global_work_size_ptr = nil
	}

	if local_work_size != nil {
		c_local_work_size = make([]C.size_t, len(local_work_size))
		for i := 0; i < len(local_work_size); i++ {
			c_local_work_size[i] = C.size_t(local_work_size[i])
		}
		c_local_work_size_ptr = &c_local_work_size[0]
	} else {
		c_local_work_size_ptr = nil
	}

	if num_events_in_wait_list != 0 {
		var c_event_wait_list []C.cl_event
		c_event_wait_list = make([]C.cl_event, num_events_in_wait_list)
		for i := 0; i < int(num_events_in_wait_list); i++ {
			c_event_wait_list[i] = event_wait_list[i].cl_event
		}
		c_ret = C.clEnqueueNDRangeKernel(command_queue.cl_command_queue,
			kernel.cl_kernel,
			C.cl_uint(work_dim),
			c_global_work_offset_ptr,
			c_global_work_size_ptr,
			c_local_work_size_ptr,
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_ret = C.clEnqueueNDRangeKernel(command_queue.cl_command_queue,
			kernel.cl_kernel,
			C.cl_uint(work_dim),
			c_global_work_offset_ptr,
			c_global_work_size_ptr,
			c_local_work_size_ptr,
			0,
			nil,
			&c_event)
	}

	if event != nil {
		event.cl_event = c_event
	}

	return CL_int(c_ret)
}

func CLEnqueueTask(command_queue CL_command_queue,
	kernel CL_kernel,
	num_events_in_wait_list CL_uint,
	event_wait_list []CL_event,
	event *CL_event) CL_int {

	if num_events_in_wait_list != 0 && len(event_wait_list) != int(num_events_in_wait_list) {
		return CL_INVALID_VALUE
	}

	var c_event C.cl_event
	var c_ret C.cl_int

	if num_events_in_wait_list != 0 {
		var c_event_wait_list []C.cl_event
		c_event_wait_list = make([]C.cl_event, num_events_in_wait_list)
		for i := 0; i < int(num_events_in_wait_list); i++ {
			c_event_wait_list[i] = event_wait_list[i].cl_event
		}

		c_ret = C.clEnqueueTask(command_queue.cl_command_queue,
			kernel.cl_kernel,
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_ret = C.clEnqueueTask(command_queue.cl_command_queue,
			kernel.cl_kernel,
			0,
			nil,
			&c_event)
	}

	if event != nil {
		event.cl_event = c_event
	}

	return CL_int(c_ret)
}
