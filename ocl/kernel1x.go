// +build cl11 cl12

package ocl

import (
	"fmt"
	"gocl/cl"
	"unsafe"
)

type kernel1x interface {
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

	//cl1x only, not in cl20
	EnqueueTask(queue CommandQueue,
		event_wait_list []Event) (Event, error)
}

func (this *kernel) EnqueueTask(queue CommandQueue,
	event_wait_list []Event) (Event, error) {
	var errCode cl.CL_int
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode = cl.CLEnqueueTask(queue.GetID(),
		this.kernel_id,
		numEvents,
		events,
		&event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueTask failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}
