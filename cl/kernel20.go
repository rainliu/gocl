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

func CLSetKernelArgSVMPointer(kernel CL_kernel,
	arg_index CL_uint,
	arg_value unsafe.Pointer) CL_int {
	return CL_int(C.clSetKernelArgSVMPointer(kernel.cl_kernel,
		C.cl_uint(arg_index),
		arg_value))
}

func CLSetKernelExecInfo(kernel CL_kernel,
	param_name CL_kernel_exec_info,
	param_value_size CL_size_t,
	param_value unsafe.Pointer) CL_int {
	return CL_int(C.clSetKernelExecInfo(kernel.cl_kernel,
		C.cl_kernel_exec_info(param_name),
		C.size_t(param_value_size),
		param_value))
}
