// +build cl12

package ocl

import (
	"fmt"
	"gocl/cl"
)

type CommandQueue interface {
	queue1x

	//cl12
	EnqueueMarkerWithWaitList(event_wait_list []Event) (Event, error)
	EnqueueBarrierWithWaitList(event_wait_list []Event) (Event, error)
	EnqueueMigrateMemObjects(mem_objects []Memory, flags cl.CL_mem_migration_flags, event_wait_list []Event) (Event, error)
}

func (this *command_queue) EnqueueMarkerWithWaitList(event_wait_list []Event) (Event, error) {
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode := cl.CLEnqueueMarkerWithWaitList(this.command_queue_id, numEvents, events, &event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueMarkerWithWaitList failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}

func (this *command_queue) EnqueueBarrierWithWaitList(event_wait_list []Event) (Event, error) {
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode := cl.CLEnqueueBarrierWithWaitList(this.command_queue_id, numEvents, events, &event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueBarrierWithWaitList failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}

func (this *command_queue) EnqueueMigrateMemObjects(mem_objects []Memory, flags cl.CL_mem_migration_flags, event_wait_list []Event) (Event, error) {
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	numMemorys := cl.CL_uint(len(mem_objects))
	memorys := make([]cl.CL_mem, numMemorys)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		memorys[i] = mem_objects[i].GetID()
	}

	if errCode := cl.CLEnqueueMigrateMemObjects(this.command_queue_id, numMemorys, memorys, flags, numEvents, events, &event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueMigrateMemObjects failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}
