// +build cl11 cl12

package cl

/*
#cgo CFLAGS: -I CL
#cgo !darwin LDFLAGS: -lOpenCL
#cgo darwin LDFLAGS: -framework OpenCL

#include "CL/opencl.h"
*/
import "C"

import (
	"unsafe"
)

func CLCreateSampler(context CL_context,
	normalized_coords CL_bool,
	addressing_mode CL_addressing_mode,
	filter_mode CL_filter_mode,
	errcode_ret *CL_int) CL_sampler {

	var c_errcode_ret C.cl_int
	var c_sampler C.cl_sampler

	c_sampler = C.clCreateSampler(context.cl_context,
		C.cl_bool(normalized_coords),
		C.cl_addressing_mode(addressing_mode),
		C.cl_filter_mode(filter_mode),
		&c_errcode_ret)

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_sampler{c_sampler}
}

func CLRetainSampler(sampler CL_sampler) CL_int {
	return CL_int(C.clRetainSampler(sampler.cl_sampler))
}

func CLReleaseSampler(sampler CL_sampler) CL_int {
	return CL_int(C.clReleaseSampler(sampler.cl_sampler))
}

func CLGetSamplerInfo(sampler CL_sampler,
	param_name CL_sampler_info,
	param_value_size CL_size_t,
	param_value *interface{},
	param_value_size_ret *CL_size_t) CL_int {

	if (param_value_size == 0 || param_value == nil) && param_value_size_ret == nil {
		return CL_INVALID_VALUE
	} else {
		var c_param_value_size_ret C.size_t
		var c_errcode_ret C.cl_int

		if param_value_size == 0 || param_value == nil {
			c_errcode_ret = C.clGetSamplerInfo(sampler.cl_sampler,
				C.cl_sampler_info(param_name),
				C.size_t(param_value_size),
				nil,
				&c_param_value_size_ret)
		} else {
			switch param_name {
			case CL_SAMPLER_REFERENCE_COUNT:

				var value C.cl_uint
				c_errcode_ret = C.clGetSamplerInfo(sampler.cl_sampler,
					C.cl_sampler_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_uint(value)
			case CL_SAMPLER_CONTEXT:

				var value C.cl_context
				c_errcode_ret = C.clGetSamplerInfo(sampler.cl_sampler,
					C.cl_sampler_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_context{value}
			case CL_SAMPLER_FILTER_MODE:

				var value C.cl_filter_mode
				c_errcode_ret = C.clGetSamplerInfo(sampler.cl_sampler,
					C.cl_sampler_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_filter_mode(value)
			case CL_SAMPLER_ADDRESSING_MODE:

				var value C.cl_addressing_mode
				c_errcode_ret = C.clGetSamplerInfo(sampler.cl_sampler,
					C.cl_sampler_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_addressing_mode(value)
			case CL_SAMPLER_NORMALIZED_COORDS:

				var value C.cl_bool
				c_errcode_ret = C.clGetSamplerInfo(sampler.cl_sampler,
					C.cl_sampler_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_bool(value)
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
