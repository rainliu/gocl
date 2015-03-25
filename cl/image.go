// +build cl11 cl12 cl20

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

import (
	"unsafe"
)

func CLGetSupportedImageFormats(context CL_context,
	flags CL_mem_flags,
	image_type CL_mem_object_type,
	num_entries CL_uint,
	image_formats []CL_image_format,
	num_image_formats *CL_uint) CL_int {

	if (num_entries == 0 || image_formats == nil) && num_image_formats == nil {
		return CL_INVALID_VALUE
	} else {
		var c_num_image_formats C.cl_uint
		var c_errcode_ret C.cl_int

		if num_entries == 0 || image_formats == nil {
			c_errcode_ret = C.clGetSupportedImageFormats(context.cl_context,
				C.cl_mem_flags(flags),
				C.cl_mem_object_type(image_type),
				C.cl_uint(num_entries),
				nil,
				&c_num_image_formats)
		} else {
			c_image_formats := make([]C.cl_image_format, len(image_formats))
			c_errcode_ret = C.clGetSupportedImageFormats(context.cl_context,
				C.cl_mem_flags(flags),
				C.cl_mem_object_type(image_type),
				C.cl_uint(num_entries),
				&c_image_formats[0],
				&c_num_image_formats)
			if c_errcode_ret == C.CL_SUCCESS {
				for i := 0; i < len(image_formats); i++ {
					image_formats[i].Image_channel_data_type = CL_channel_type(c_image_formats[i].image_channel_data_type)
					image_formats[i].Image_channel_order = CL_channel_order(c_image_formats[i].image_channel_order)
				}
			}
		}

		if num_image_formats != nil {
			*num_image_formats = CL_uint(c_num_image_formats)
		}

		return CL_int(c_errcode_ret)
	}
}

func CLEnqueueMapImage(command_queue CL_command_queue,
	image CL_mem,
	blocking_map CL_bool,
	map_flags CL_map_flags,
	origin [3]CL_size_t,
	region [3]CL_size_t,
	image_row_pitch *CL_size_t,
	image_slice_pitch *CL_size_t,
	num_events_in_wait_list CL_uint,
	event_wait_list []CL_event,
	event *CL_event,
	errcode_ret *CL_int) unsafe.Pointer {

	if (num_events_in_wait_list == 0 && event_wait_list != nil) ||
		(num_events_in_wait_list != 0 && event_wait_list == nil) ||
		int(num_events_in_wait_list) != len(event_wait_list) {

		if errcode_ret != nil {
			*errcode_ret = CL_INVALID_EVENT_WAIT_LIST
		}
		return nil
	}

	var c_origin, c_region [3]C.size_t
	var c_image_row_pitch, c_image_slice_pitch C.size_t
	var c_event C.cl_event
	var c_errcode_ret C.cl_int
	var c_ptr_ret unsafe.Pointer

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

		c_ptr_ret = C.clEnqueueMapImage(command_queue.cl_command_queue,
			image.cl_mem,
			C.cl_bool(blocking_map),
			C.cl_map_flags(map_flags),
			&c_origin[0],
			&c_region[0],
			&c_image_row_pitch,
			&c_image_slice_pitch,
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event,
			&c_errcode_ret)
	} else {
		c_ptr_ret = C.clEnqueueMapImage(command_queue.cl_command_queue,
			image.cl_mem,
			C.cl_bool(blocking_map),
			C.cl_map_flags(map_flags),
			&c_origin[0],
			&c_region[0],
			&c_image_row_pitch,
			&c_image_slice_pitch,
			0,
			nil,
			&c_event,
			&c_errcode_ret)
	}

	if image_row_pitch != nil {
		*image_row_pitch = CL_size_t(c_image_row_pitch)
	}
	if image_slice_pitch != nil {
		*image_slice_pitch = CL_size_t(c_image_slice_pitch)
	}

	if event != nil {
		event.cl_event = c_event
	}
	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return c_ptr_ret
}

func CLEnqueueCopyImageToBuffer(command_queue CL_command_queue,
	src_image CL_mem,
	dst_buffer CL_mem,
	src_origin [3]CL_size_t,
	region [3]CL_size_t,
	dst_offset CL_size_t,
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
		c_origin[i] = C.size_t(src_origin[i])
		c_region[i] = C.size_t(region[i])
	}
	if num_events_in_wait_list != 0 {
		var c_event_wait_list []C.cl_event
		c_event_wait_list = make([]C.cl_event, num_events_in_wait_list)
		for i := 0; i < int(num_events_in_wait_list); i++ {
			c_event_wait_list[i] = event_wait_list[i].cl_event
		}

		c_errcode_ret = C.clEnqueueCopyImageToBuffer(command_queue.cl_command_queue,
			src_image.cl_mem,
			dst_buffer.cl_mem,
			&c_origin[0],
			&c_region[0],
			C.size_t(dst_offset),
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueCopyImageToBuffer(command_queue.cl_command_queue,
			src_image.cl_mem,
			dst_buffer.cl_mem,
			&c_origin[0],
			&c_region[0],
			C.size_t(dst_offset),
			0,
			nil,
			&c_event)
	}

	if event != nil {
		event.cl_event = c_event
	}

	return CL_int(c_errcode_ret)
}

func CLEnqueueCopyBufferToImage(command_queue CL_command_queue,
	src_buffer CL_mem,
	dst_image CL_mem,
	src_offset CL_size_t,
	dst_origin [3]CL_size_t,
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
		c_origin[i] = C.size_t(dst_origin[i])
		c_region[i] = C.size_t(region[i])
	}
	if num_events_in_wait_list != 0 {
		var c_event_wait_list []C.cl_event
		c_event_wait_list = make([]C.cl_event, num_events_in_wait_list)
		for i := 0; i < int(num_events_in_wait_list); i++ {
			c_event_wait_list[i] = event_wait_list[i].cl_event
		}

		c_errcode_ret = C.clEnqueueCopyBufferToImage(command_queue.cl_command_queue,
			src_buffer.cl_mem,
			dst_image.cl_mem,
			C.size_t(src_offset),
			&c_origin[0],
			&c_region[0],
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueCopyBufferToImage(command_queue.cl_command_queue,
			src_buffer.cl_mem,
			dst_image.cl_mem,
			C.size_t(src_offset),
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

func CLEnqueueReadImage(command_queue CL_command_queue,
	image CL_mem,
	blocking_read CL_bool,
	origin [3]CL_size_t,
	region [3]CL_size_t,
	row_pitch CL_size_t,
	slice_pitch CL_size_t,
	ptr unsafe.Pointer,
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

		c_errcode_ret = C.clEnqueueReadImage(command_queue.cl_command_queue,
			image.cl_mem,
			C.cl_bool(blocking_read),
			&c_origin[0],
			&c_region[0],
			C.size_t(row_pitch),
			C.size_t(slice_pitch),
			ptr,
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueReadImage(command_queue.cl_command_queue,
			image.cl_mem,
			C.cl_bool(blocking_read),
			&c_origin[0],
			&c_region[0],
			C.size_t(row_pitch),
			C.size_t(slice_pitch),
			ptr,
			0,
			nil,
			&c_event)
	}

	if event != nil {
		event.cl_event = c_event
	}

	return CL_int(c_errcode_ret)
}

func CLEnqueueWriteImage(command_queue CL_command_queue,
	image CL_mem,
	blocking_write CL_bool,
	origin [3]CL_size_t,
	region [3]CL_size_t,
	row_pitch CL_size_t,
	slice_pitch CL_size_t,
	ptr unsafe.Pointer,
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

		c_errcode_ret = C.clEnqueueWriteImage(command_queue.cl_command_queue,
			image.cl_mem,
			C.cl_bool(blocking_write),
			&c_origin[0],
			&c_region[0],
			C.size_t(row_pitch),
			C.size_t(slice_pitch),
			ptr,
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueWriteImage(command_queue.cl_command_queue,
			image.cl_mem,
			C.cl_bool(blocking_write),
			&c_origin[0],
			&c_region[0],
			C.size_t(row_pitch),
			C.size_t(slice_pitch),
			ptr,
			0,
			nil,
			&c_event)
	}

	if event != nil {
		event.cl_event = c_event
	}

	return CL_int(c_errcode_ret)
}

func CLEnqueueCopyImage(command_queue CL_command_queue,
	src_image CL_mem,
	dst_image CL_mem,
	src_origin [3]CL_size_t,
	dst_origin [3]CL_size_t,
	region [3]CL_size_t,
	num_events_in_wait_list CL_uint,
	event_wait_list []CL_event,
	event *CL_event) CL_int {

	if (num_events_in_wait_list == 0 && event_wait_list != nil) ||
		(num_events_in_wait_list != 0 && event_wait_list == nil) ||
		int(num_events_in_wait_list) != len(event_wait_list) {
		return CL_INVALID_EVENT_WAIT_LIST
	}

	var c_src_origin, c_dst_origin, c_region [3]C.size_t
	var c_event C.cl_event
	var c_errcode_ret C.cl_int

	for i := 0; i < 3; i++ {
		c_src_origin[i] = C.size_t(src_origin[i])
		c_dst_origin[i] = C.size_t(dst_origin[i])
		c_region[i] = C.size_t(region[i])
	}
	if num_events_in_wait_list != 0 {
		var c_event_wait_list []C.cl_event
		c_event_wait_list = make([]C.cl_event, num_events_in_wait_list)
		for i := 0; i < int(num_events_in_wait_list); i++ {
			c_event_wait_list[i] = event_wait_list[i].cl_event
		}

		c_errcode_ret = C.clEnqueueCopyImage(command_queue.cl_command_queue,
			src_image.cl_mem,
			dst_image.cl_mem,
			&c_src_origin[0],
			&c_dst_origin[0],
			&c_region[0],
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueCopyImage(command_queue.cl_command_queue,
			src_image.cl_mem,
			dst_image.cl_mem,
			&c_src_origin[0],
			&c_dst_origin[0],
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

func CLGetImageInfo(image CL_mem,
	param_name CL_image_info,
	param_value_size CL_size_t,
	param_value *interface{},
	param_value_size_ret *CL_size_t) CL_int {

	if (param_value_size == 0 || param_value == nil) && param_value_size_ret == nil {
		return CL_INVALID_VALUE
	} else {
		var c_param_value_size_ret C.size_t
		var c_errcode_ret C.cl_int

		if param_value_size == 0 || param_value == nil {
			c_errcode_ret = C.clGetImageInfo(image.cl_mem,
				C.cl_image_info(param_name),
				C.size_t(param_value_size),
				nil,
				&c_param_value_size_ret)
		} else {
			switch param_name {
			case CL_IMAGE_FORMAT:

				var value C.cl_image_format
				c_errcode_ret = C.clGetImageInfo(image.cl_mem,
					C.cl_image_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_image_format{CL_channel_order(value.image_channel_order),
					CL_channel_type(value.image_channel_data_type)}

			case CL_IMAGE_ELEMENT_SIZE,
				CL_IMAGE_ROW_PITCH,
				CL_IMAGE_SLICE_PITCH,
				CL_IMAGE_HEIGHT,
				CL_IMAGE_WIDTH,
				CL_IMAGE_DEPTH,
				CL_IMAGE_ARRAY_SIZE:

				var value C.size_t
				c_errcode_ret = C.clGetImageInfo(image.cl_mem,
					C.cl_image_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_size_t(value)

			case CL_IMAGE_BUFFER:

				var value C.cl_mem
				c_errcode_ret = C.clGetImageInfo(image.cl_mem,
					C.cl_image_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_mem{value}

			case CL_IMAGE_NUM_MIP_LEVELS,
				CL_IMAGE_NUM_SAMPLES:
				var value C.cl_uint
				c_errcode_ret = C.clGetImageInfo(image.cl_mem,
					C.cl_image_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_uint(value)

			default:
				return CL_INVALID_VALUE
			}
		}

		if param_value_size_ret != nil {
			*param_value_size_ret = CL_size_t(c_param_value_size_ret)
		}

		return CL_int(c_errcode_ret)
	}
}
