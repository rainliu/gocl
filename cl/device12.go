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

///////////////////////////////////////////////
//OpenCL 1.2
///////////////////////////////////////////////
func CLCreateSubDevices(in_device CL_device_id,
	properties []CL_device_partition_property,
	num_devices CL_uint,
	out_devices []CL_device_id,
	num_devices_ret *CL_uint) CL_int {

	if (num_devices == 0 && out_devices != nil) || (out_devices == nil && num_devices_ret == nil) {
		return CL_INVALID_VALUE
	} else {
		var c_properties []C.cl_device_partition_property
		var c_properties_ptr *C.cl_device_partition_property

		if properties != nil {
			c_properties = make([]C.cl_device_partition_property, len(properties))
			for i := 0; i < len(properties); i++ {
				c_properties[i] = C.cl_device_partition_property(properties[i])
			}
			c_properties_ptr = &c_properties[0]
		} else {
			c_properties_ptr = nil
		}

		var c_errcode_ret C.cl_int
		var c_num_devices_ret C.cl_uint

		if out_devices == nil {
			c_errcode_ret = C.clCreateSubDevices(in_device.cl_device_id,
				c_properties_ptr,
				C.cl_uint(num_devices),
				nil,
				&c_num_devices_ret)
		} else {
			c_out_devices := make([]C.cl_device_id, len(out_devices))
			c_errcode_ret = C.clCreateSubDevices(in_device.cl_device_id,
				c_properties_ptr,
				C.cl_uint(num_devices),
				&c_out_devices[0],
				&c_num_devices_ret)
			if c_errcode_ret == C.CL_SUCCESS {
				for i := 0; i < len(out_devices); i++ {
					out_devices[i].cl_device_id = c_out_devices[i]
				}
			}
		}

		if num_devices_ret != nil {
			*num_devices_ret = CL_uint(c_num_devices_ret)
		}
		return CL_int(c_errcode_ret)
	}
}

func CLRetainDevice(device CL_device_id) CL_int {
	return CL_int(C.clRetainDevice(device.cl_device_id))
}

func CLReleaseDevice(device CL_device_id) CL_int {
	return CL_int(C.clReleaseDevice(device.cl_device_id))
}
