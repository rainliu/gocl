// +build cl11 cl12

package ocl

import (
	"gocl/cl"
	"unsafe"
)

type memory1x interface {
	GetID() cl.CL_mem
	GetInfo(param_name cl.CL_mem_info) (interface{}, error)
	Retain() error
	Release() error
	SetCallback(pfn_notify cl.CL_mem_notify, user_data unsafe.Pointer) error

	//to be fix CL_event
	EnqueueUnmap(queue CommandQueue, mapped_ptr unsafe.Pointer, event_wait_list []cl.CL_event) (cl.CL_event, error)
}