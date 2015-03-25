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
import "unsafe"

func CLGetPlatformIDs(num_entries CL_uint,
	platforms []CL_platform_id,
	num_platforms *CL_uint) CL_int {

	if (num_entries == 0 && platforms != nil) || (num_platforms == nil && platforms == nil) {
		return CL_INVALID_VALUE
	} else {
		var c_errcode_ret C.cl_int
		var c_num_platforms C.cl_uint

		if platforms == nil {
			c_errcode_ret = C.clGetPlatformIDs(C.cl_uint(num_entries),
				nil,
				&c_num_platforms)
		} else {
			c_platforms := make([]C.cl_platform_id, len(platforms))
			c_errcode_ret = C.clGetPlatformIDs(C.cl_uint(num_entries),
				&c_platforms[0],
				&c_num_platforms)
			if c_errcode_ret == C.CL_SUCCESS {
				for i := 0; i < len(platforms); i++ {
					platforms[i].cl_platform_id = c_platforms[i]
				}
			}
		}

		if num_platforms != nil {
			*num_platforms = CL_uint(c_num_platforms)
		}
		return CL_int(c_errcode_ret)
	}
}

func CLGetPlatformInfo(platform CL_platform_id,
	param_name CL_platform_info,
	param_value_size CL_size_t,
	param_value *interface{},
	param_value_size_ret *CL_size_t) CL_int {

	if (param_value_size == 0 || param_value == nil) && param_value_size_ret == nil {
		return CL_INVALID_VALUE
	} else {
		var c_param_value_size_ret C.size_t
		var c_errcode_ret C.cl_int

		if param_value_size == 0 || param_value == nil {
			c_errcode_ret = C.clGetPlatformInfo(platform.cl_platform_id,
				C.cl_platform_info(param_name),
				C.size_t(param_value_size),
				nil,
				&c_param_value_size_ret)
		} else {
			switch param_name {
			case CL_PLATFORM_PROFILE,
				CL_PLATFORM_VERSION,
				CL_PLATFORM_NAME,
				CL_PLATFORM_VENDOR,
				CL_PLATFORM_EXTENSIONS:

				value := make([]C.char, param_value_size)
				c_errcode_ret = C.clGetPlatformInfo(platform.cl_platform_id,
					C.cl_platform_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value[0]),
					&c_param_value_size_ret)

				*param_value = C.GoStringN(&value[0], C.int(c_param_value_size_ret-1))
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
