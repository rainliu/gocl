// +build cl11 cl12

package ocl

import (
	"errors"
	"gocl/cl"
	"unsafe"
)

type context1x interface {
	GetID() cl.CL_context
	GetInfo(param_name cl.CL_context_info) (interface{}, error)
	Retain() error
	Release() error

	CreateCommandQueue(device Device, properties []cl.CL_command_queue_properties) (CommandQueue, error)
	CreateBuffer(flags cl.CL_mem_flags, size cl.CL_size_t, host_ptr unsafe.Pointer) (Buffer, error)
	CreateEvent() (Event, error)
}

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
