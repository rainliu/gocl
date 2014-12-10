// +build cl12

package ocl

import (
	//"errors"
	//"gocl/cl"
)

type CommandQueue interface {
	queue1x
}

/*
func (this *command_queue) EnqueueMarkerWithWaitList(num_events_in_wait_list CL_uint,
	event_wait_list []CL_event,
	event *CL_event) CL_int {

}

func (this *command_queue) EnqueueBarrierWithWaitList(num_events_in_wait_list CL_uint,
	event_wait_list []CL_event,
	event *CL_event) CL_int {
}*/