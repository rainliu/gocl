// +build cl11 cl12

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
