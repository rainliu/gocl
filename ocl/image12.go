// +build cl12

package ocl

import (
	"fmt"
	"gocl/cl"
	"unsafe"
)

type Image interface {
	image1x

	//cl12
	EnqueueFill(queue CommandQueue,
		fill_color unsafe.Pointer,
		origin [3]cl.CL_size_t,
		region [3]cl.CL_size_t,
		event_wait_list []Event) (Event, error)
}

func (this *image) EnqueueFill(queue CommandQueue,
	fill_color unsafe.Pointer,
	origin [3]cl.CL_size_t,
	region [3]cl.CL_size_t,
	event_wait_list []Event) (Event, error) {
	var errCode cl.CL_int
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode = cl.CLEnqueueFillImage(queue.GetID(),
		this.memory_id,
		fill_color,
		origin,
		region,
		numEvents,
		events,
		&event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueFill failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}
