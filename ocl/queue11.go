// +build cl11

package ocl

import (
	"fmt"
	"gocl/cl"
)

type CommandQueue interface {
	queue1x

	//cl11
	EnqueueMarker() (Event, error)
	EnqueueBarrier() error
	EnqueueWaitForEvents(event_wait_list []Event) error
}

func (this *command_queue) EnqueueMarker() (Event, error) {
	var event_id cl.CL_event
	if errCode := cl.CLEnqueueMarker(this.command_queue_id, &event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueMarker failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}

func (this *command_queue) EnqueueBarrier() error {
	if errCode := cl.CLEnqueueBarrier(this.command_queue_id); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("EnqueueBarrier failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return nil
	}
}

func (this *command_queue) EnqueueWaitForEvents(event_wait_list []Event) error {
	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode := cl.CLEnqueueWaitForEvents(this.command_queue_id, numEvents, events); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("EnqueueWaitForEvents failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return nil
	}
}
