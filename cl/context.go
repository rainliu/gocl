package cl

/*
#cgo CFLAGS: -I CL
#cgo LDFLAGS: -lOpenCL

#include "CL/opencl.h"

extern void go_ctx_notify(char *errinfo, void *private_info, int cb, void *user_data);
static void CL_CALLBACK c_ctx_notify(const char *errinfo, const void *private_info, size_t cb, void *user_data) {
	go_ctx_notify((char *)errinfo, (void *)private_info, cb, user_data);
}

static cl_context CLCreateContext(	const cl_context_properties *  	properties,
					                cl_uint                  		num_devices,
					                const cl_device_id *     		devices,
					                void *                   		user_data,
					                cl_int *                 		errcode_ret){

    return clCreateContext(properties, num_devices, devices, c_ctx_notify, user_data, errcode_ret);
}

static cl_context CLCreateContextFromType(	const cl_context_properties *  	properties,
					                		cl_device_type     				device_type,
					                		void *                   		user_data,
					                		cl_int *                 		errcode_ret){

    return clCreateContextFromType(properties, device_type, c_ctx_notify, user_data, errcode_ret);
}
*/
import "C"

import (
	"unsafe"
)

type CL_ctx_notify func(errinfo string, private_info interface{}, cb int, user_data unsafe.Pointer)

var ctx_notify CL_ctx_notify

//export go_ctx_notify
func go_ctx_notify(errinfo *C.char, private_info unsafe.Pointer, cb C.int, user_data unsafe.Pointer) {
	ctx_notify(C.GoString(errinfo), private_info, int(cb), user_data)
}

func CLCreateContext(properties []CL_context_properties,
	num_devices CL_uint,
	devices []CL_device_id,
	pfn_notify CL_ctx_notify,
	user_data unsafe.Pointer,
	errcode_ret *CL_int) CL_context {

	var c_properties []C.cl_context_properties
	var c_devices []C.cl_device_id
	var c_properties_ptr *C.cl_context_properties
	var c_devices_ptr *C.cl_device_id
	var c_errcode_ret C.cl_int
	var c_context C.cl_context

	if properties != nil && len(properties) > 0 {
		c_properties = make([]C.cl_context_properties, len(properties))
		for i := 0; i < len(properties); i++ {
			c_properties[i] = C.cl_context_properties(properties[i])
		}
		c_properties_ptr = &c_properties[0]
	} else {
		c_properties_ptr = nil
	}

	if devices != nil && len(devices) > 0 {
		c_devices = make([]C.cl_device_id, len(devices))
		for i := 0; i < len(devices); i++ {
			c_devices[i] = C.cl_device_id(devices[i].cl_device_id)
		}
		c_devices_ptr = &c_devices[0]
	} else {
		c_devices_ptr = nil
	}

	if pfn_notify != nil {
		ctx_notify = pfn_notify

		c_context = C.CLCreateContext(c_properties_ptr,
			C.cl_uint(len(c_devices)),
			c_devices_ptr,
			user_data,
			&c_errcode_ret)

	} else {
		ctx_notify = nil

		c_context = C.clCreateContext(c_properties_ptr,
			C.cl_uint(len(c_devices)),
			c_devices_ptr,
			nil,
			nil,
			&c_errcode_ret)

	}

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_context{c_context}
}

func CLCreateContextFromType(properties []CL_context_properties,
	device_type CL_device_type,
	pfn_notify CL_ctx_notify,
	user_data unsafe.Pointer,
	errcode_ret *CL_int) CL_context {

	var c_properties []C.cl_context_properties
	var c_properties_ptr *C.cl_context_properties
	var c_errcode_ret C.cl_int
	var c_context C.cl_context

	if properties != nil && len(properties) > 0 {
		c_properties = make([]C.cl_context_properties, len(properties))
		for i := 0; i < len(properties); i++ {
			c_properties[i] = C.cl_context_properties(properties[i])
		}
		c_properties_ptr = &c_properties[0]
	} else {
		c_properties_ptr = nil
	}

	if pfn_notify != nil {
		ctx_notify = pfn_notify

		c_context = C.CLCreateContextFromType(c_properties_ptr,
			C.cl_device_type(device_type),
			user_data,
			&c_errcode_ret)

	} else {
		ctx_notify = nil

		c_context = C.clCreateContextFromType(c_properties_ptr,
			C.cl_device_type(device_type),
			nil,
			nil,
			&c_errcode_ret)

	}

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_context{c_context}
}

func CLGetContextInfo(context CL_context,
	param_name CL_context_info,
	param_value_size CL_size_t,
	param_value *interface{},
	param_value_size_ret *CL_size_t) CL_int {

	var ret C.cl_int

	if (param_value_size == 0 || param_value == nil) && param_value_size_ret == nil {
		ret = C.clGetContextInfo(context.cl_context,
			C.cl_context_info(param_name),
			0,
			nil,
			nil)
	} else {
		var size_ret C.size_t

		if param_value_size == 0 || param_value == nil {
			ret = C.clGetContextInfo(context.cl_context,
				C.cl_context_info(param_name),
				0,
				nil,
				&size_ret)
		} else {
			switch param_name {
			case CL_CONTEXT_REFERENCE_COUNT,
				CL_CONTEXT_NUM_DEVICES:

				var value C.cl_uint
				ret = C.clGetContextInfo(context.cl_context,
					C.cl_context_info(param_name),
					C.size_t(unsafe.Sizeof(value)),
					unsafe.Pointer(&value),
					&size_ret)

				*param_value = CL_uint(value)

			case CL_CONTEXT_DEVICES:
				var s_device C.cl_device_id
				num_devices := int(C.size_t(param_value_size) / C.size_t(unsafe.Sizeof(s_device)))
				c_devices := make([]C.cl_device_id, num_devices)

				ret = C.clGetContextInfo(context.cl_context,
					C.cl_context_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&c_devices[0]),
					&size_ret)

				devices := make([]CL_device_id, num_devices)
				for i := 0; i < num_devices; i++ {
					devices[i].cl_device_id = c_devices[i]
				}

				*param_value = devices

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

func CLRetainContext(context CL_context) CL_int {
	return CL_int(C.clRetainContext(context.cl_context))
}

func CLReleaseContext(context CL_context) CL_int {
	return CL_int(C.clReleaseContext(context.cl_context))
}
