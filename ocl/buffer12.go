// +build cl12

package ocl

import (
	"fmt"
	"gocl/cl"
	"unsafe"
)

type Buffer interface {
	buffer1x

	//cl12
	EnqueueFill(queue CommandQueue,
		pattern unsafe.Pointer,
		pattern_size cl.CL_size_t,
		offset cl.CL_size_t,
		cb cl.CL_size_t,
		event_wait_list []Event) (Event, error)
}

func (this *buffer) EnqueueFill(queue CommandQueue,
	pattern unsafe.Pointer,
	pattern_size cl.CL_size_t,
	offset cl.CL_size_t,
	cb cl.CL_size_t,
	event_wait_list []Event) (Event, error) {
	var errCode cl.CL_int
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode = cl.CLEnqueueFillBuffer(queue.GetID(),
		this.memory_id,
		pattern,
		pattern_size,
		offset,
		cb,
		numEvents,
		events,
		&event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueFill failure with errcode_ret %d", errCode)
	} else {
		return &event{event_id}, nil
	}
}
