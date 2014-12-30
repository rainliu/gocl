// +build cl11 cl12 cl20

package ocl

import (
	"fmt"
	"gocl/cl"
)

type command_queue struct {
	command_queue_id cl.CL_command_queue
}

func (this *command_queue) GetID() cl.CL_command_queue {
	return this.command_queue_id
}

func (this *command_queue) GetInfo(param_name cl.CL_command_queue_info) (interface{}, error) {
	/* param data */
	var param_value interface{}
	var param_size cl.CL_size_t
	var errCode cl.CL_int

	/* Find size of param data */
	if errCode = cl.CLGetCommandQueueInfo(this.command_queue_id, param_name, 0, nil, &param_size); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetInfo failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	}

	/* Access param data */
	if errCode = cl.CLGetCommandQueueInfo(this.command_queue_id, param_name, param_size, &param_value, nil); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetInfo failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	}

	return param_value, nil
}

func (this *command_queue) Retain() error {
	if errCode := cl.CLRetainCommandQueue(this.command_queue_id); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("Retain failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	}
	return nil
}

func (this *command_queue) Release() error {
	if errCode := cl.CLReleaseCommandQueue(this.command_queue_id); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("Release failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	}
	return nil
}

func (this *command_queue) Flush() error {
	if errCode := cl.CLFlush(this.command_queue_id); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("Flush failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	}
	return nil
}

func (this *command_queue) Finish() error {
	if errCode := cl.CLFinish(this.command_queue_id); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("Finish failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	}
	return nil
}

func (this *command_queue) EnqueueCopyBuffer(src_buffer Buffer,
	dst_buffer Buffer,
	src_offset cl.CL_size_t,
	dst_offset cl.CL_size_t,
	cb cl.CL_size_t,
	event_wait_list []Event) (Event, error) {
	var errCode cl.CL_int
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode = cl.CLEnqueueCopyBuffer(this.command_queue_id,
		src_buffer.GetID(),
		dst_buffer.GetID(),
		src_offset,
		dst_offset,
		cb,
		numEvents,
		events,
		&event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueCopyBuffer failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}

func (this *command_queue) EnqueueCopyBufferRect(src_buffer Buffer,
	dst_buffer Buffer,
	src_origin [3]cl.CL_size_t,
	dst_origin [3]cl.CL_size_t,
	region [3]cl.CL_size_t,
	src_row_pitch cl.CL_size_t,
	src_slice_pitch cl.CL_size_t,
	dst_row_pitch cl.CL_size_t,
	dst_slice_pitch cl.CL_size_t,
	event_wait_list []Event) (Event, error) {
	var errCode cl.CL_int
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode = cl.CLEnqueueCopyBufferRect(this.command_queue_id,
		src_buffer.GetID(),
		dst_buffer.GetID(),
		src_origin,
		dst_origin,
		region,
		src_row_pitch,
		src_slice_pitch,
		dst_row_pitch,
		dst_slice_pitch,
		numEvents,
		events,
		&event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueCopyBufferRect failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}

func (this *command_queue) EnqueueCopyImage(src_image Image,
	dst_image Image,
	src_origin [3]cl.CL_size_t,
	dst_origin [3]cl.CL_size_t,
	region [3]cl.CL_size_t,
	event_wait_list []Event) (Event, error) {
	var errCode cl.CL_int
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode = cl.CLEnqueueCopyImage(this.command_queue_id,
		src_image.GetID(),
		dst_image.GetID(),
		src_origin,
		dst_origin,
		region,
		numEvents,
		events,
		&event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueCopyImage failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}

func (this *command_queue) EnqueueCopyImageToBuffer(src_image Image,
	dst_buffer Buffer,
	src_origin [3]cl.CL_size_t,
	region [3]cl.CL_size_t,
	dst_offset cl.CL_size_t,
	event_wait_list []Event) (Event, error) {
	var errCode cl.CL_int
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode = cl.CLEnqueueCopyImageToBuffer(this.command_queue_id,
		src_image.GetID(),
		dst_buffer.GetID(),
		src_origin,
		region,
		dst_offset,
		numEvents,
		events,
		&event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueCopyImageToBuffer failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}

func (this *command_queue) EnqueueCopyBufferToImage(src_buffer Buffer,
	dst_image Image,
	src_offset cl.CL_size_t,
	dst_origin [3]cl.CL_size_t,
	region [3]cl.CL_size_t,
	event_wait_list []Event) (Event, error) {
	var errCode cl.CL_int
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode = cl.CLEnqueueCopyBufferToImage(this.command_queue_id,
		src_buffer.GetID(),
		dst_image.GetID(),
		src_offset,
		dst_origin,
		region,
		numEvents,
		events,
		&event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueCopyBufferToImage failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}
