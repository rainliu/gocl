// +build cl11 cl12

package ocl

import (
	"errors"
	"gocl/cl"
)

type Device struct {
	device_id cl.CL_device_id
}

func NewDevice(device_id cl.CL_device_id) *Device {
	this := &Device{}
	this.device_id = device_id
	return this
}

func (this *Device) GetID() cl.CL_device_id {
	return this.device_id
}

func (this *Device) GetInfo(param_name cl.CL_device_info) (interface{}, error) {
	/* param data */
	var param_value interface{}
	var param_size cl.CL_size_t
	var errCode cl.CL_int

	/* Find size of param data */
	if errCode = cl.CLGetDeviceInfo(this.device_id, param_name, 0, nil, &param_size); errCode != cl.CL_SUCCESS {
		return "", errors.New("GetInfo failure with errcode_ret " + string(errCode))
	}

	/* Access param data */
	if errCode = cl.CLGetDeviceInfo(this.device_id, param_name, param_size, &param_value, nil); errCode != cl.CL_SUCCESS {
		return "", errors.New("GetInfo failure with errcode_ret " + string(errCode))
	}

	return param_value, nil
}
