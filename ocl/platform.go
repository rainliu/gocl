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

func (this *Platform) Get() cl.CL_platform_id {
	return this.platform_id
}

func (this *Platform) IsHost() bool {
	return false
}

func (this *Platform) GetInfo(param_name cl.CL_platform_info) (string, error) {
	return "", nil
}

func (this *Platform) HasExtension(extension string) bool {
	return false
}

func (this *Platform) GetDevices() ([]Device, error) {
	return nil, nil
}

func GetPlatforms() ([]Platform, error) {
	var platforms []Platform
	var platformIDs []cl.CL_platform_id
	var num_platforms cl.CL_uint
	var err cl.CL_int

	err = cl.CLGetPlatformIDs(1, nil, &num_platforms)
	if err != cl.CL_SUCCESS {
		return nil, errors.New("GetPlatformIDs failure with errcode_ret " + string(err))
	}

	platformIDs = make([]cl.CL_platform_id, num_platforms)

	err = cl.CLGetPlatformIDs(num_platforms, platformIDs, nil)
	if err != cl.CL_SUCCESS {
		return nil, errors.New("GetPlatformIDs failure with errcode_ret " + string(err))
	}

	platforms = make([]Platform, num_platforms)

	for i := cl.CL_uint(0); i < num_platforms; i++ {
		platforms[i].platform_id = platformIDs[i]
	}

	return platforms, nil
}
