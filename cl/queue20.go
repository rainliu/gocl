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

func CLCreateCommandQueueWithProperties(context CL_context,
	device CL_device_id,
	properties []CL_command_queue_properties,
	errcode_ret *CL_int) CL_command_queue {
	var c_errcode_ret C.cl_int
	var c_command_queue C.cl_command_queue

	var c_properties []C.cl_command_queue_properties
	var c_properties_ptr *C.cl_command_queue_properties

	if properties != nil {
		c_properties = make([]C.cl_command_queue_properties, len(properties))
		for i := 0; i < len(properties); i++ {
			c_properties[i] = C.cl_command_queue_properties(properties[i])
		}
		c_properties_ptr = &c_properties[0]
	} else {
		c_properties_ptr = nil
	}

	c_command_queue = C.clCreateCommandQueueWithProperties(context.cl_context,
		device.cl_device_id,
		c_properties_ptr,
		&c_errcode_ret)

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_command_queue{c_command_queue}
}
