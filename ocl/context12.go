// +build cl12

package ocl

import (
	"errors"
	"gocl/cl"
	"unsafe"
)

type Context interface {
	GetID() cl.CL_context
	GetInfo(param_name cl.CL_context_info) (interface{}, error)
	Retain() error
	Release() error

	CreateCommandQueue(device Device, properties []cl.CL_command_queue_properties) (CommandQueue, error)
	CreateBuffer(flags cl.CL_mem_flags, size cl.CL_size_t, host_ptr unsafe.Pointer) (Buffer, error)
	CreateImage(flags cl.CL_mem_flags, image_format *cl.CL_image_format, image_desc *cl.CL_image_desc, host_ptr unsafe.Pointer) (Image, error)
}

func (this *context) CreateImage(flags cl.CL_mem_flags,
	image_format *cl.CL_image_format,
	image_desc *cl.CL_image_desc,
	host_ptr unsafe.Pointer) (Image, error) {
	var errCode cl.CL_int

	if memory_id := cl.CLCreateImage(this.context_id, flags, image_format, image_desc, host_ptr, &errCode); errCode != cl.CL_SUCCESS {
		return nil, errors.New("CreateImage failure with errcode_ret " + string(errCode))
	} else {
		return &image{memory{memory_id}}, nil
	}
}
