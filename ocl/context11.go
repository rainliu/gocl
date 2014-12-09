// +build cl11

package ocl

import (
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
}
