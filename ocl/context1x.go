// +build cl11 cl12

package ocl

import (
	"errors"
	"gocl/cl"
)

func (this *context) CreateCommandQueue(device Device,
	properties []cl.CL_command_queue_properties) (CommandQueue, error) {
	var property cl.CL_command_queue_properties
	var errCode cl.CL_int

	if properties == nil {
		property = 0
	} else {
		property = properties[0]
	}

	if command_queue_id := cl.CLCreateCommandQueue(this.context_id, device.GetID(), property, &errCode); errCode != cl.CL_SUCCESS {
		return nil, errors.New("CreateCommandQueue failure with errcode_ret " + string(errCode))
	} else {
		return &command_queue{command_queue_id}, nil
	}
}
