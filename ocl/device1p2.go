// +build cl12

package ocl

import (
	"errors"
	"gocl/cl"
)

func (this *Device) CreateSubDevices(properties []cl.CL_device_partition_property) ([]Device, error) {
	var numDevices cl.CL_uint
	var deviceIds []cl.CL_device_id
	var devices []Device
	var errCode cl.CL_int

	/* Determine number of connected devices */
	if errCode = cl.CLCreateSubDevices(this.device_id, properties, 0, nil, &numDevices); errCode != cl.CL_SUCCESS {
		return nil, errors.New("CreateSubDevices failure with errcode_ret " + string(errCode))
	}

	/* Access connected devices */
	deviceIds = make([]cl.CL_device_id, numDevices)
	if errCode = cl.CLCreateSubDevices(this.device_id, properties, numDevices, deviceIds, nil); errCode != cl.CL_SUCCESS {
		return nil, errors.New("CreateSubDevices failure with errcode_ret " + string(errCode))
	}

	devices = make([]Device, numDevices)
	for i := cl.CL_uint(0); i < numDevices; i++ {
		devices[i].device_id = deviceIds[i]
	}

	return devices, nil
}

func (this *Device) Retain() error {
	if errCode := cl.CLRetainDevice(this.device_id); errCode != cl.CL_SUCCESS {
		return errors.New("Retain failure with errcode_ret " + string(errCode))
	}
	return nil
}

func (this *Device) Release() error {
	if errCode := cl.CLReleaseDevice(this.device_id); errCode != cl.CL_SUCCESS {
		return errors.New("Release failure with errcode_ret " + string(errCode))
	}
	return nil
}
