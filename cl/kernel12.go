// +build cl12 cl20

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

///////////////////////////////////////////////
//OpenCL 1.2
///////////////////////////////////////////////

func CLGetKernelArgInfo(kernel CL_kernel,
	arg_index CL_uint,
	param_name CL_kernel_arg_info,
	param_value_size CL_size_t,
	param_value *interface{},
	param_value_size_ret *CL_size_t) CL_int {

	if (param_value_size == 0 || param_value == nil) && param_value_size_ret == nil {
		return CL_INVALID_VALUE
	} else {
		var c_param_value_size_ret C.size_t
		var c_errcode_ret C.cl_int

		if param_value_size == 0 || param_value == nil {
			c_errcode_ret = C.clGetKernelArgInfo(kernel.cl_kernel,
				C.cl_uint(arg_index),
				C.cl_kernel_arg_info(param_name),
				C.size_t(param_value_size),
				nil,
				&c_param_value_size_ret)
		} else {
			switch param_name {

			case CL_KERNEL_ARG_ADDRESS_QUALIFIER:

				var value C.cl_kernel_arg_address_qualifier
				c_errcode_ret = C.clGetKernelArgInfo(kernel.cl_kernel,
					C.cl_uint(arg_index),
					C.cl_kernel_arg_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_kernel_arg_address_qualifier(value)

			case CL_KERNEL_ARG_ACCESS_QUALIFIER:

				var value C.cl_kernel_arg_access_qualifier
				c_errcode_ret = C.clGetKernelArgInfo(kernel.cl_kernel,
					C.cl_uint(arg_index),
					C.cl_kernel_arg_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_kernel_arg_access_qualifier(value)

			case CL_KERNEL_ARG_TYPE_NAME,
				CL_KERNEL_ARG_NAME:

				value := make([]C.char, param_value_size)
				c_errcode_ret = C.clGetKernelInfo(kernel.cl_kernel,
					C.cl_kernel_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value[0]),
					&c_param_value_size_ret)

				*param_value = C.GoStringN(&value[0], C.int(c_param_value_size_ret-1))

			case CL_KERNEL_ARG_TYPE_QUALIFIER:

				var value C.cl_kernel_arg_type_qualifier
				c_errcode_ret = C.clGetKernelArgInfo(kernel.cl_kernel,
					C.cl_uint(arg_index),
					C.cl_kernel_arg_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_kernel_arg_type_qualifier(value)

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
