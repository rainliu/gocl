package ocl

import (
	"errors"
	"gocl/cl"
)

type Platform struct {
	platform_id cl.CL_platform_id
}

func NewPlatform(platform_id cl.CL_platform_id) *Platform {
	this := &Platform{}
	this.platform_id = platform_id
	return this
}

func (this *Platform) GetID() cl.CL_platform_id {
	return this.platform_id
}

func (this *Platform) GetInfo(param_name cl.CL_platform_info) (string, error) {
	/* param data */
	var param_value interface{}
	var param_size cl.CL_size_t
	var errCode cl.CL_int

	/* Find size of param data */
	if errCode = cl.CLGetPlatformInfo(this.platform_id, param_name, 0, nil, &param_size); errCode != cl.CL_SUCCESS {
		return "", errors.New("GetInfo failure with errcode_ret " + string(errCode))
	}

	/* Access param data */
	if errCode = cl.CLGetPlatformInfo(this.platform_id, param_name, param_size, &param_value, nil); errCode != cl.CL_SUCCESS {
		return "", errors.New("GetInfo failure with errcode_ret " + string(errCode))
	}

	return param_value.(string), nil
}

func (this *Platform) GetDevices(deviceType cl.CL_device_type) ([]Device, error) {
	var devices []Device
	var deviceIds []cl.CL_device_id
	var numDevices cl.CL_uint
	var errCode cl.CL_int

	/* Determine number of connected devices */
	if errCode = cl.CLGetDeviceIDs(this.platform_id, deviceType, 0, nil, &numDevices); errCode != cl.CL_SUCCESS {
		return nil, errors.New("GetDevices failure with errcode_ret " + string(errCode))
	}

	/* Access connected devices */
	deviceIds = make([]cl.CL_device_id, numDevices)
	if errCode = cl.CLGetDeviceIDs(this.platform_id, deviceType, numDevices, deviceIds, nil); errCode != cl.CL_SUCCESS {
		return nil, errors.New("GetDevices failure with errcode_ret " + string(errCode))
	}

	devices = make([]Device, numDevices)
	for i := cl.CL_uint(0); i < numDevices; i++ {
		devices[i].device_id = deviceIds[i]
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
		return nil, errors.New("GetPlatforms failure with errcode_ret " + string(errCode))
	}

	/* Access platforms */
	platformIds = make([]cl.CL_platform_id, numPlatforms)
	if errCode = cl.CLGetPlatformIDs(numPlatforms, platformIds, nil); errCode != cl.CL_SUCCESS {
		return nil, errors.New("GetPlatforms failure with errcode_ret " + string(errCode))
	}

	platforms = make([]Platform, numPlatforms)
	for i := cl.CL_uint(0); i < numPlatforms; i++ {
		platforms[i].platform_id = platformIds[i]
	}

	return platforms, nil
}
