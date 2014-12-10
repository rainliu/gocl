// +build cl12

package ocl

import (
	"errors"
	"gocl/cl"
)

type CommandQueue interface {
	queue1x

	//cl12
	EnqueueMarkerWithWaitList (event_wait_list []Event) (Event, error)
	EnqueueBarrierWithWaitList(event_wait_list []Event) (Event, error)
}

func (this *command_queue) EnqueueMarkerWithWaitList(event_wait_list []Event) (Event, error) {		
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i:= cl.CL_uint(0); i<numEvents; i++{
		events[i] = event_wait_list[i].GetID()
	}

	if errCode := cl.CLEnqueueMarkerWithWaitList(this.command_queue_id, numEvents, events, &event_id); errCode != cl.CL_SUCCESS {
		return nil, errors.New("EnqueueMarkerWithWaitList failure with errcode_ret " + string(errCode))
	} else {
		return &event{event_id}, nil
	}

}

func (this *command_queue) EnqueueBarrierWithWaitList(event_wait_list []Event) (Event, error) {
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i:= cl.CL_uint(0); i<numEvents; i++{
		events[i] = event_wait_list[i].GetID()
	}

	if errCode := cl.CLEnqueueBarrierWithWaitList(this.command_queue_id, numEvents, events, &event_id); errCode != cl.CL_SUCCESS {
		return nil, errors.New("EnqueueBarrierWithWaitList failure with errcode_ret " + string(errCode))
	} else {
		return &event{event_id}, nil
	}
}