// +build cl12 cl20

package cl

/*
#cgo CFLAGS: -I CL
#cgo !darwin LDFLAGS: -lOpenCL
#cgo darwin LDFLAGS: -framework OpenCL

#ifdef __APPLE__
#include "OpenCL/opencl.h"
#else
#include "CL/opencl.h"
#endif
 */
import "C"
import "unsafe"

///////////////////////////////////////////////
//OpenCL 1.2
///////////////////////////////////////////////

func CLCreateImage(context CL_context,
	flags CL_mem_flags,
	image_format *CL_image_format,
	image_desc *CL_image_desc,
	host_ptr unsafe.Pointer,
	errcode_ret *CL_int) CL_mem {

	var c_image_format C.cl_image_format
	var c_image_desc C.cl_image_desc
	var c_errcode_ret C.cl_int
	var c_image C.cl_mem

	c_image_format.image_channel_order = C.cl_channel_order(image_format.Image_channel_order)
	c_image_format.image_channel_data_type = C.cl_channel_type(image_format.Image_channel_data_type)

	c_image_desc.image_type = C.cl_mem_object_type(image_desc.Image_type)
	c_image_desc.image_width = C.size_t(image_desc.Image_width)
	c_image_desc.image_height = C.size_t(image_desc.Image_height)
	c_image_desc.image_depth = C.size_t(image_desc.Image_depth)
	c_image_desc.image_array_size = C.size_t(image_desc.Image_array_size)
	c_image_desc.image_row_pitch = C.size_t(image_desc.Image_row_pitch)
	c_image_desc.image_slice_pitch = C.size_t(image_desc.Image_slice_pitch)
	c_image_desc.num_mip_levels = C.cl_uint(image_desc.Num_mip_levels)
	c_image_desc.num_samples = C.cl_uint(image_desc.Num_samples)
	c_image_desc.buffer = image_desc.Buffer.cl_mem

	c_image = C.clCreateImage(context.cl_context,
		C.cl_mem_flags(flags),
		&c_image_format,
		&c_image_desc,
		host_ptr,
		&c_errcode_ret)

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_mem{c_image}
}

func CLEnqueueFillImage(command_queue CL_command_queue,
	image CL_mem,
	fill_color unsafe.Pointer,
	origin [3]CL_size_t,
	region [3]CL_size_t,
	num_events_in_wait_list CL_uint,
	event_wait_list []CL_event,
	event *CL_event) CL_int {

	if (num_events_in_wait_list == 0 && event_wait_list != nil) ||
		(num_events_in_wait_list != 0 && event_wait_list == nil) ||
		int(num_events_in_wait_list) != len(event_wait_list) {
		return CL_INVALID_EVENT_WAIT_LIST
	}

	var c_origin, c_region [3]C.size_t
	var c_event C.cl_event
	var c_errcode_ret C.cl_int
	for i := 0; i < 3; i++ {
		c_origin[i] = C.size_t(origin[i])
		c_region[i] = C.size_t(region[i])
	}

	if num_events_in_wait_list != 0 {
		var c_event_wait_list []C.cl_event
		c_event_wait_list = make([]C.cl_event, num_events_in_wait_list)
		for i := 0; i < int(num_events_in_wait_list); i++ {
			c_event_wait_list[i] = event_wait_list[i].cl_event
		}

		c_errcode_ret = C.clEnqueueFillImage(command_queue.cl_command_queue,
			image.cl_mem,
			fill_color,
			&c_origin[0],
			&c_region[0],
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueFillImage(command_queue.cl_command_queue,
			image.cl_mem,
			fill_color,
			&c_origin[0],
			&c_region[0],
			0,
			nil,
			&c_event)
	}

	if event != nil {
		event.cl_event = c_event
	}

	return CL_int(c_errcode_ret)
}
