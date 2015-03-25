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

func CLCreateSamplerWithProperties(context CL_context,
	properties []CL_sampler_properties,
	errcode_ret *CL_int) CL_sampler {
	var c_errcode_ret C.cl_int
	var c_sampler C.cl_sampler

	var c_properties []C.cl_sampler_properties
	var c_properties_ptr *C.cl_sampler_properties

	if properties != nil {
		c_properties = make([]C.cl_sampler_properties, len(properties))
		for i := 0; i < len(properties); i++ {
			c_properties[i] = C.cl_sampler_properties(properties[i])
		}
		c_properties_ptr = &c_properties[0]
	} else {
		c_properties_ptr = nil
	}

	c_sampler = C.clCreateSamplerWithProperties(context.cl_context,
		c_properties_ptr,
		&c_errcode_ret)

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_sampler{c_sampler}
}
