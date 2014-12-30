// +build cl11 cl12 cl20

package ocl

import (
	"gocl/cl"
	"unsafe"
)

type Event interface {
	GetID() cl.CL_event
	GetInfo(param_name cl.CL_event_info) (interface{}, error)
	Retain() error
	Release() error

	SetStatus(execution_status cl.CL_int) error
	SetCallback(command_exec_callback_type cl.CL_int, pfn_notify cl.CL_evt_notify, user_data unsafe.Pointer) error
	GetProfilingInfo(param_name cl.CL_profiling_info) (interface{}, error)
}

type Memory interface {
	GetID() cl.CL_mem
	GetInfo(param_name cl.CL_mem_info) (interface{}, error)
	Retain() error
	Release() error

	SetCallback(pfn_notify cl.CL_mem_notify,
		user_data unsafe.Pointer) error
	EnqueueUnmap(queue CommandQueue,
		mapped_ptr unsafe.Pointer,
		event_wait_list []Event) (Event, error)
}

type Platform interface {
	GetID() cl.CL_platform_id
	GetInfo(param_name cl.CL_platform_info) (interface{}, error)
	GetDevices(deviceType cl.CL_device_type) ([]Device, error)

	UnloadCompiler() error
}

type Sampler interface {
	GetID() cl.CL_sampler
	GetInfo(param_name cl.CL_sampler_info) (interface{}, error)
	Retain() error
	Release() error
}
