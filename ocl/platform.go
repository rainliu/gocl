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

func (this *Platform) GetInfo(param_name cl.CL_platform_info) (string, error) {
	/* Extension data */
	var param_value interface{}
	var param_size cl.CL_size_t
	var errCode cl.CL_int

	/* Find size of extension data */
	errCode = cl.CLGetPlatformInfo(this.platform_id,
		param_name, 0, nil, &param_size)
	if errCode != cl.CL_SUCCESS {
		return "", errors.New("GetInfo failure with errcode_ret " + string(errCode))
	}

	/* Access extension data */
	errCode = cl.CLGetPlatformInfo(this.platform_id,
		param_name, param_size, &param_value, nil)
	if errCode != cl.CL_SUCCESS {
		return "", errors.New("GetInfo failure with errcode_ret " + string(errCode))
	}

	return param_value.(string), nil
}

func (this *Platform) GetID() cl.CL_platform_id {
	return this.platform_id
}

func (this *Platform) GetDevices(device_type cl.CL_device_type) ([]Device, error) {
	return nil, nil
}

func GetPlatforms() ([]Platform, error) {
	var platforms []Platform
	var platformIds []cl.CL_platform_id
	var numPlatforms cl.CL_uint
	var errCode cl.CL_int

	errCode = cl.CLGetPlatformIDs(0, nil, &numPlatforms)
	if errCode != cl.CL_SUCCESS {
		return nil, errors.New("GetPlatformIDs failure with errcode_ret " + string(errCode))
	}

	platformIds = make([]cl.CL_platform_id, numPlatforms)

	errCode = cl.CLGetPlatformIDs(numPlatforms, platformIds, nil)
	if errCode != cl.CL_SUCCESS {
		return nil, errors.New("GetPlatformIDs failure with errcode_ret " + string(errCode))
	}

	platforms = make([]Platform, numPlatforms)

	for i := cl.CL_uint(0); i < numPlatforms; i++ {
		platforms[i].platform_id = platformIds[i]
	}

	return platforms, nil
}
