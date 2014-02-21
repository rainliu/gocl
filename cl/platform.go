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

func CLGetPlatformIDs(num_entries CL_uint,
	platforms []CL_platform_id,
	num_platforms *CL_uint) CL_int {

	var ret C.cl_int

	if (num_entries == 0 || platforms == nil) && num_platforms == nil {
		ret = C.clGetPlatformIDs(0,
			(*C.cl_platform_id)(nil),
			(*C.cl_uint)(nil))
	} else {
		var num C.cl_uint

		if num_entries == 0 || platforms == nil {
			ret = C.clGetPlatformIDs(0,
				(*C.cl_platform_id)(nil),
				&num)
		} else {
			platforms_id := make([]C.cl_platform_id, len(platforms))
			ret = C.clGetPlatformIDs(C.cl_uint(num_entries),
				&platforms_id[0],
				&num)
			if ret == C.CL_SUCCESS {
				for i := 0; i < len(platforms); i++ {
					platforms[i].cl_platform_id = platforms_id[i]
				}
			}
		}

		if num_platforms != nil {
			*num_platforms = CL_uint(num)
		}
	}

	return CL_int(ret)
}

func CLGetPlatformInfo(platform CL_platform_id,
	param_name CL_platform_info,
	param_value_size CL_size_t,
	param_value *interface{},
	param_value_size_ret *CL_size_t) CL_int {

	var ret C.cl_int

	if (param_value_size == 0 || param_value == nil) && param_value_size_ret == nil {
		ret = C.clGetPlatformInfo(platform.cl_platform_id,
			C.cl_platform_info(param_name),
			0,
			nil,
			nil)
	} else {
		var size_ret C.size_t

		if param_value_size == 0 || param_value == nil {
			ret = C.clGetPlatformInfo(platform.cl_platform_id,
				C.cl_platform_info(param_name),
				0,
				nil,
				&size_ret)
		} else {
			switch param_name {
			case CL_PLATFORM_PROFILE,
				CL_PLATFORM_VERSION,
				CL_PLATFORM_NAME,
				CL_PLATFORM_VENDOR,
				CL_PLATFORM_EXTENSIONS:

				value := make([]C.char, param_value_size)
				ret = C.clGetPlatformInfo(platform.cl_platform_id,
					C.cl_platform_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value[0]),
					&size_ret)

				*param_value = C.GoStringN(&value[0], C.int(size_ret-1))
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
