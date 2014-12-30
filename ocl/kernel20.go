// +build cl20

package ocl

import (
	"gocl/cl"
	"unsafe"
)

type kernel20 interface {
	GetID() cl.CL_kernel
	GetInfo(param_name cl.CL_kernel_info) (interface{}, error)
	Retain() error
	Release() error

	SetArg(arg_index cl.CL_uint,
		arg_size cl.CL_size_t,
		arg_value unsafe.Pointer) error
	GetWorkGroupInfo(device Device,
		param_name cl.CL_kernel_work_group_info) (interface{}, error)
	EnqueueNDRange(queue CommandQueue,
		work_dim cl.CL_uint,
		global_work_offset []cl.CL_size_t,
		global_work_size []cl.CL_size_t,
		local_work_size []cl.CL_size_t,
		event_wait_list []Event) (Event, error)

	//cl20
}

//TODO
//func CLSetKernelArgSVMPointer
//func CLSetKernelExecInfo
