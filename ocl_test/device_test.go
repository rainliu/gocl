package ocl_test

import (
	"gocl/cl"
	"gocl/ocl"
	"testing"
)

func TestDevice(t *testing.T) {
	/* Host/device data structures */
	var platforms []ocl.Platform
	var devices []ocl.Device
	var err error

	/* Param data */
	var param_value interface{}

	/* Identify a platform */
	if platforms, err = ocl.GetPlatforms(); err != nil {
		t.Errorf(err.Error())
		return
	}

	/* Determine connected devices */
	if devices, err = platforms[0].GetDevices(cl.CL_DEVICE_TYPE_ALL); err != nil {
		t.Errorf(err.Error())
		return
	} else {
		t.Logf("Number of device: %d\n", len(devices))
	}

	/* Obtain data for each connected device */
	for i := 0; i < len(devices); i++ {
		if param_value, err = devices[i].GetInfo(cl.CL_DEVICE_NAME); err != nil {
			t.Errorf("Failed to find OpenCL device info %s.\n", "CL_DEVICE_NAME")
			return
		} else {
			t.Logf("\t%s:\t %v\n", "CL_DEVICE_NAME", param_value)
		}

		if param_value, err = devices[i].GetInfo(cl.CL_DEVICE_TYPE); err != nil {
			t.Errorf("Failed to find OpenCL device info %s.\n", "CL_DEVICE_TYPE")
			return
		} else {
			switch param_value.(cl.CL_device_type) {
			case cl.CL_DEVICE_TYPE_CPU:
				t.Logf("\t%s:\t %v\n", "CL_DEVICE_TYPE", "CL_DEVICE_TYPE_CPU")
			case cl.CL_DEVICE_TYPE_GPU:
				t.Logf("\t%s:\t %v\n", "CL_DEVICE_TYPE", "CL_DEVICE_TYPE_GPU")
			}

		}
	}
}
