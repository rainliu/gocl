// +build cl11 cl12

package ocl

import (
	"fmt"
	"gocl/cl"
)

type Platform interface {
	GetID() cl.CL_platform_id
	GetInfo(param_name cl.CL_platform_info) (interface{}, error)
	GetDevices(deviceType cl.CL_device_type) ([]Device, error)

	UnloadCompiler() error
}

type platform struct {
	platform_id cl.CL_platform_id
}

func (this *platform) GetID() cl.CL_platform_id {
	return this.platform_id
}

func (this *platform) GetInfo(param_name cl.CL_platform_info) (interface{}, error) {
	/* param data */
	var param_value interface{}
	var param_size cl.CL_size_t
	var errCode cl.CL_int

	/* Find size of param data */
	if errCode = cl.CLGetPlatformInfo(this.platform_id, param_name, 0, nil, &param_size); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetInfo failure with errcode_ret %d", errCode)
	}

	/* Access param data */
	if errCode = cl.CLGetPlatformInfo(this.platform_id, param_name, param_size, &param_value, nil); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetInfo failure with errcode_ret %d", errCode)
	}

	return param_value, nil
}

func (this *platform) GetDevices(deviceType cl.CL_device_type) ([]Device, error) {
	var devices []Device
	var deviceIds []cl.CL_device_id
	var numDevices cl.CL_uint
	var errCode cl.CL_int

	/* Determine number of connected devices */
	if errCode = cl.CLGetDeviceIDs(this.platform_id, deviceType, 0, nil, &numDevices); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetDevices failure with errcode_ret %d", errCode)
	}

	/* Access connected devices */
	deviceIds = make([]cl.CL_device_id, numDevices)
	if errCode = cl.CLGetDeviceIDs(this.platform_id, deviceType, numDevices, deviceIds, nil); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetDevices failure with errcode_ret %d", errCode)
	}

	devices = make([]Device, numDevices)
	for i := cl.CL_uint(0); i < numDevices; i++ {
		devices[i] = &device{deviceIds[i]}
	}

	return devices, nil
}

func GetPlatforms() ([]Platform, error) {
	var platforms []Platform
	var platformIds []cl.CL_platform_id
	var numPlatforms cl.CL_uint
	var errCode cl.CL_int

	/* Determine number of platforms */
	if errCode = cl.CLGetPlatformIDs(0, nil, &numPlatforms); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetPlatforms failure with errcode_ret %d", errCode)
	}

	/* Access platforms */
	platformIds = make([]cl.CL_platform_id, numPlatforms)
	if errCode = cl.CLGetPlatformIDs(numPlatforms, platformIds, nil); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetPlatforms failure with errcode_ret %d", errCode)
	}

	platforms = make([]Platform, numPlatforms)
	for i := cl.CL_uint(0); i < numPlatforms; i++ {
		platforms[i] = &platform{platformIds[i]}
	}

	return platforms, nil
}
