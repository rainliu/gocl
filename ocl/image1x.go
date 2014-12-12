// +build cl11 cl12

package ocl

import (
	"gocl/cl"
	"unsafe"
)

type image1x interface {
	Memory

	GetImageInfo(param_name cl.CL_image_info) (interface{}, error)

	EnqueueMap(queue CommandQueue,
		blocking_map cl.CL_bool,
		map_flags cl.CL_map_flags,
		origin [3]cl.CL_size_t,
		region [3]cl.CL_size_t,
		image_row_pitch *cl.CL_size_t,
		image_slice_pitch *cl.CL_size_t,
		event_wait_list []Event) (unsafe.Pointer, Event, error)
	EnqueueRead(queue CommandQueue,
		blocking_read cl.CL_bool,
		origin [3]cl.CL_size_t,
		region [3]cl.CL_size_t,
		row_pitch cl.CL_size_t,
		slice_pitch cl.CL_size_t,
		ptr unsafe.Pointer,
		event_wait_list []Event) (Event, error)
	EnqueueWrite(queue CommandQueue,
		blocking_write cl.CL_bool,
		origin [3]cl.CL_size_t,
		region [3]cl.CL_size_t,
		row_pitch cl.CL_size_t,
		slice_pitch cl.CL_size_t,
		ptr unsafe.Pointer,
		event_wait_list []Event) (Event, error)
}
