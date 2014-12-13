// +build cl11 cl12

package ocl

import (
	"fmt"
	"gocl/cl"
	"unsafe"
)

type image struct {
	memory
}

func (this *image) GetImageInfo(param_name cl.CL_image_info) (interface{}, error) {
	/* param data */
	var param_value interface{}
	var param_size cl.CL_size_t
	var errCode cl.CL_int

	/* Find size of param data */
	if errCode = cl.CLGetImageInfo(this.memory_id, param_name, 0, nil, &param_size); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetImageInfo failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}

	/* Access param data */
	if errCode = cl.CLGetImageInfo(this.memory_id, param_name, param_size, &param_value, nil); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetImageInfo failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}

	return param_value, nil
}

func (this *image) EnqueueRead(queue CommandQueue,
	blocking_read cl.CL_bool,
	origin [3]cl.CL_size_t,
	region [3]cl.CL_size_t,
	row_pitch cl.CL_size_t,
	slice_pitch cl.CL_size_t,
	ptr unsafe.Pointer,
	event_wait_list []Event) (Event, error) {
	var errCode cl.CL_int
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode = cl.CLEnqueueReadImage(queue.GetID(),
		this.memory_id,
		blocking_read,
		origin,
		region,
		row_pitch,
		slice_pitch,
		ptr,
		numEvents,
		events,
		&event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueRead failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}

func (this *image) EnqueueWrite(queue CommandQueue,
	blocking_write cl.CL_bool,
	origin [3]cl.CL_size_t,
	region [3]cl.CL_size_t,
	row_pitch cl.CL_size_t,
	slice_pitch cl.CL_size_t,
	ptr unsafe.Pointer,
	event_wait_list []Event) (Event, error) {
	var errCode cl.CL_int
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode = cl.CLEnqueueWriteImage(queue.GetID(),
		this.memory_id,
		blocking_write,
		origin,
		region,
		row_pitch,
		slice_pitch,
		ptr,
		numEvents,
		events,
		&event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueWrite failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}

func (this *image) EnqueueMap(queue CommandQueue,
	blocking_map cl.CL_bool,
	map_flags cl.CL_map_flags,
	origin [3]cl.CL_size_t,
	region [3]cl.CL_size_t,
	image_row_pitch *cl.CL_size_t,
	image_slice_pitch *cl.CL_size_t,
	event_wait_list []Event) (unsafe.Pointer, Event, error) {
	var errCode cl.CL_int
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if mapped_ptr := cl.CLEnqueueMapImage(queue.GetID(),
		this.memory_id,
		blocking_map,
		map_flags,
		origin,
		region,
		image_row_pitch,
		image_slice_pitch,
		numEvents,
		events,
		&event_id,
		&errCode); errCode != cl.CL_SUCCESS {
		return nil, nil, fmt.Errorf("EnqueueMap failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return mapped_ptr, &event{event_id}, nil
	}
}
