// +build cl11 cl12

package ocl

import (
	"errors"
	"gocl/cl"
	"unsafe"
)

type Buffer interface {
	Memory

	CreateSubBuffer(flags cl.CL_mem_flags, buffer_create_type cl.CL_buffer_create_type, buffer_create_info unsafe.Pointer) (Buffer, error)
}

type buffer struct {
	memory
}

func (this *buffer) CreateSubBuffer(flags cl.CL_mem_flags,
	buffer_create_type cl.CL_buffer_create_type,
	buffer_create_info unsafe.Pointer) (Buffer, error) {
	var errCode cl.CL_int

	if memory_id := cl.CLCreateSubBuffer(this.memory_id, flags, buffer_create_type, buffer_create_info, &errCode); errCode != cl.CL_SUCCESS {
		return nil, errors.New("CreateSubBuffer failure with errcode_ret " + string(errCode))
	} else {
		return &buffer{memory{memory_id}}, nil
	}
}
