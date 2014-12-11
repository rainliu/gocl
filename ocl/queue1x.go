// +build cl11 cl12

package ocl

import (
	"gocl/cl"
)

type queue1x interface {
	GetID() cl.CL_command_queue
	GetInfo(param_name cl.CL_command_queue_info) (interface{}, error)
	Retain() error
	Release() error
	Flush() error
	Finish() error

	EnqueueCopyBuffer(src_buffer Buffer,
		dst_buffer Buffer,
		src_offset cl.CL_size_t,
		dst_offset cl.CL_size_t,
		cb cl.CL_size_t,
		event_wait_list []Event) (Event, error)
	EnqueueCopyBufferRect(src_buffer Buffer,
		dst_buffer Buffer,
		src_origin [3]cl.CL_size_t,
		dst_origin [3]cl.CL_size_t,
		region [3]cl.CL_size_t,
		src_row_pitch cl.CL_size_t,
		src_slice_pitch cl.CL_size_t,
		dst_row_pitch cl.CL_size_t,
		dst_slice_pitch cl.CL_size_t,
		event_wait_list []Event) (Event, error)
}
