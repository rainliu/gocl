// +build cl11 cl12

package ocl

import (
	"gocl/cl"
	"unsafe"
)

type buffer1x interface {
	Memory

	CreateSubBuffer(flags cl.CL_mem_flags,
		buffer_create_type cl.CL_buffer_create_type,
		buffer_create_info unsafe.Pointer) (Buffer, error)

	EnqueueMap(queue CommandQueue,
		blocking_map cl.CL_bool,
		map_flags cl.CL_map_flags,
		offset cl.CL_size_t,
		cb cl.CL_size_t,
		event_wait_list []Event) (unsafe.Pointer, Event, error)
	EnqueueRead(queue CommandQueue,
		blocking_read cl.CL_bool,
		offset cl.CL_size_t,
		cb cl.CL_size_t,
		ptr unsafe.Pointer,
		event_wait_list []Event) (Event, error)
	EnqueueWrite(queue CommandQueue,
		blocking_write cl.CL_bool,
		offset cl.CL_size_t,
		cb cl.CL_size_t,
		ptr unsafe.Pointer,
		event_wait_list []Event) (Event, error)

	EnqueueReadRect(queue CommandQueue,
		blocking_read cl.CL_bool,
		buffer_origin [3]cl.CL_size_t,
		host_origin [3]cl.CL_size_t,
		region [3]cl.CL_size_t,
		buffer_row_pitch cl.CL_size_t,
		buffer_slice_pitch cl.CL_size_t,
		host_row_pitch cl.CL_size_t,
		host_slice_pitch cl.CL_size_t,
		ptr unsafe.Pointer,
		event_wait_list []Event) (Event, error)
	EnqueueWriteRect(queue CommandQueue,
		blocking_write cl.CL_bool,
		buffer_origin [3]cl.CL_size_t,
		host_origin [3]cl.CL_size_t,
		region [3]cl.CL_size_t,
		buffer_row_pitch cl.CL_size_t,
		buffer_slice_pitch cl.CL_size_t,
		host_row_pitch cl.CL_size_t,
		host_slice_pitch cl.CL_size_t,
		ptr unsafe.Pointer,
		event_wait_list []Event) (Event, error)
}
