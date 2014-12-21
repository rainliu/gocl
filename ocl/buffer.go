// +build cl11 cl12 cl20

package ocl

import (
	"fmt"
	"gocl/cl"
	"unsafe"
)

type buffer struct {
	memory
}

func (this *buffer) CreateSubBuffer(flags cl.CL_mem_flags,
	buffer_create_type cl.CL_buffer_create_type,
	buffer_create_info unsafe.Pointer) (Buffer, error) {
	var errCode cl.CL_int

	if memory_id := cl.CLCreateSubBuffer(this.memory_id,
		flags,
		buffer_create_type,
		buffer_create_info,
		&errCode); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("CreateSubBuffer failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &buffer{memory{memory_id}}, nil
	}
}

func (this *buffer) EnqueueRead(queue CommandQueue,
	blocking_read cl.CL_bool,
	offset cl.CL_size_t,
	cb cl.CL_size_t,
	ptr unsafe.Pointer,
	event_wait_list []Event) (Event, error) {
	var errCode cl.CL_int
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode = cl.CLEnqueueReadBuffer(queue.GetID(),
		this.memory_id,
		blocking_read,
		offset,
		cb,
		ptr,
		numEvents,
		events,
		&event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueRead failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}

func (this *buffer) EnqueueWrite(queue CommandQueue,
	blocking_write cl.CL_bool,
	offset cl.CL_size_t,
	cb cl.CL_size_t,
	ptr unsafe.Pointer,
	event_wait_list []Event) (Event, error) {
	var errCode cl.CL_int
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode = cl.CLEnqueueWriteBuffer(queue.GetID(),
		this.memory_id,
		blocking_write,
		offset,
		cb,
		ptr,
		numEvents,
		events,
		&event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueWrite failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}

func (this *buffer) EnqueueReadRect(queue CommandQueue,
	blocking_read cl.CL_bool,
	buffer_origin [3]cl.CL_size_t,
	host_origin [3]cl.CL_size_t,
	region [3]cl.CL_size_t,
	buffer_row_pitch cl.CL_size_t,
	buffer_slice_pitch cl.CL_size_t,
	host_row_pitch cl.CL_size_t,
	host_slice_pitch cl.CL_size_t,
	ptr unsafe.Pointer,
	event_wait_list []Event) (Event, error) {
	var errCode cl.CL_int
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode = cl.CLEnqueueReadBufferRect(queue.GetID(),
		this.memory_id,
		blocking_read,
		buffer_origin,
		host_origin,
		region,
		buffer_row_pitch,
		buffer_slice_pitch,
		host_row_pitch,
		host_slice_pitch,
		ptr,
		numEvents,
		events,
		&event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueReadRect failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}

func (this *buffer) EnqueueWriteRect(queue CommandQueue,
	blocking_write cl.CL_bool,
	buffer_origin [3]cl.CL_size_t,
	host_origin [3]cl.CL_size_t,
	region [3]cl.CL_size_t,
	buffer_row_pitch cl.CL_size_t,
	buffer_slice_pitch cl.CL_size_t,
	host_row_pitch cl.CL_size_t,
	host_slice_pitch cl.CL_size_t,
	ptr unsafe.Pointer,
	event_wait_list []Event) (Event, error) {
	var errCode cl.CL_int
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode = cl.CLEnqueueWriteBufferRect(queue.GetID(),
		this.memory_id,
		blocking_write,
		buffer_origin,
		host_origin,
		region,
		buffer_row_pitch,
		buffer_slice_pitch,
		host_row_pitch,
		host_slice_pitch,
		ptr,
		numEvents,
		events,
		&event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueWriteRect failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}

func (this *buffer) EnqueueMap(queue CommandQueue,
	blocking_map cl.CL_bool,
	map_flags cl.CL_map_flags,
	offset cl.CL_size_t,
	cb cl.CL_size_t,
	event_wait_list []Event) (unsafe.Pointer, Event, error) {
	var errCode cl.CL_int
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if mapped_ptr := cl.CLEnqueueMapBuffer(queue.GetID(),
		this.memory_id,
		blocking_map,
		map_flags,
		offset,
		cb,
		numEvents,
		events,
		&event_id,
		&errCode); errCode != cl.CL_SUCCESS {
		return nil, nil, fmt.Errorf("EnqueueMap failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return mapped_ptr, &event{event_id}, nil
	}
}
