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

func CLCreateBuffer(context CL_context,
	flags CL_mem_flags,
	size CL_size_t,
	host_ptr unsafe.Pointer,
	errcode_ret *CL_int) CL_mem {

	var c_errcode_ret C.cl_int
	var c_memobj C.cl_mem

	c_memobj = C.clCreateBuffer(context.cl_context,
		C.cl_mem_flags(flags),
		C.size_t(size),
		host_ptr,
		&c_errcode_ret)

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_mem{c_memobj}
}

func CLCreateSubBuffer(buffer CL_mem,
	flags CL_mem_flags,
	buffer_create_type CL_buffer_create_type,
	buffer_create_info unsafe.Pointer,
	errcode_ret *CL_int) CL_mem {

	var c_errcode_ret C.cl_int
	var c_memobj C.cl_mem

	c_memobj = C.clCreateSubBuffer(buffer.cl_mem,
		C.cl_mem_flags(flags),
		C.cl_buffer_create_type(buffer_create_type),
		buffer_create_info,
		&c_errcode_ret)

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_mem{c_memobj}
}

func CLEnqueueReadBuffer(command_queue CL_command_queue,
	buffer CL_mem,
	blocking_read CL_bool,
	offset CL_size_t,
	cb CL_size_t,
	ptr unsafe.Pointer,
	num_events_in_wait_list CL_uint,
	event_wait_list []CL_event,
	event *CL_event) CL_int {

	if (num_events_in_wait_list == 0 && event_wait_list != nil) ||
		(num_events_in_wait_list != 0 && event_wait_list == nil) ||
		int(num_events_in_wait_list) != len(event_wait_list) {
		return CL_INVALID_EVENT_WAIT_LIST
	}

	var c_event C.cl_event
	var c_errcode_ret C.cl_int

	if num_events_in_wait_list != 0 {
		var c_event_wait_list []C.cl_event
		c_event_wait_list = make([]C.cl_event, num_events_in_wait_list)
		for i := 0; i < int(num_events_in_wait_list); i++ {
			c_event_wait_list[i] = event_wait_list[i].cl_event
		}

		c_errcode_ret = C.clEnqueueReadBuffer(command_queue.cl_command_queue,
			buffer.cl_mem,
			C.cl_bool(blocking_read),
			C.size_t(offset),
			C.size_t(cb),
			ptr,
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueReadBuffer(command_queue.cl_command_queue,
			buffer.cl_mem,
			C.cl_bool(blocking_read),
			C.size_t(offset),
			C.size_t(cb),
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

func CLEnqueueWriteBuffer(command_queue CL_command_queue,
	buffer CL_mem,
	blocking_write CL_bool,
	offset CL_size_t,
	cb CL_size_t,
	ptr unsafe.Pointer,
	num_events_in_wait_list CL_uint,
	event_wait_list []CL_event,
	event *CL_event) CL_int {

	if (num_events_in_wait_list == 0 && event_wait_list != nil) ||
		(num_events_in_wait_list != 0 && event_wait_list == nil) ||
		int(num_events_in_wait_list) != len(event_wait_list) {
		return CL_INVALID_EVENT_WAIT_LIST
	}

	var c_event C.cl_event
	var c_errcode_ret C.cl_int

	if num_events_in_wait_list != 0 {
		var c_event_wait_list []C.cl_event
		c_event_wait_list = make([]C.cl_event, num_events_in_wait_list)
		for i := 0; i < int(num_events_in_wait_list); i++ {
			c_event_wait_list[i] = event_wait_list[i].cl_event
		}

		c_errcode_ret = C.clEnqueueWriteBuffer(command_queue.cl_command_queue,
			buffer.cl_mem,
			C.cl_bool(blocking_write),
			C.size_t(offset),
			C.size_t(cb),
			ptr,
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueWriteBuffer(command_queue.cl_command_queue,
			buffer.cl_mem,
			C.cl_bool(blocking_write),
			C.size_t(offset),
			C.size_t(cb),
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

func CLEnqueueCopyBuffer(command_queue CL_command_queue,
	src_buffer CL_mem,
	dst_buffer CL_mem,
	src_offset CL_size_t,
	dst_offset CL_size_t,
	cb CL_size_t,
	num_events_in_wait_list CL_uint,
	event_wait_list []CL_event,
	event *CL_event) CL_int {

	if (num_events_in_wait_list == 0 && event_wait_list != nil) ||
		(num_events_in_wait_list != 0 && event_wait_list == nil) ||
		int(num_events_in_wait_list) != len(event_wait_list) {
		return CL_INVALID_EVENT_WAIT_LIST
	}

	var c_event C.cl_event
	var c_errcode_ret C.cl_int

	if num_events_in_wait_list != 0 {
		var c_event_wait_list []C.cl_event
		c_event_wait_list = make([]C.cl_event, num_events_in_wait_list)
		for i := 0; i < int(num_events_in_wait_list); i++ {
			c_event_wait_list[i] = event_wait_list[i].cl_event
		}

		c_errcode_ret = C.clEnqueueCopyBuffer(command_queue.cl_command_queue,
			src_buffer.cl_mem,
			dst_buffer.cl_mem,
			C.size_t(src_offset),
			C.size_t(dst_offset),
			C.size_t(cb),
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueCopyBuffer(command_queue.cl_command_queue,
			src_buffer.cl_mem,
			dst_buffer.cl_mem,
			C.size_t(src_offset),
			C.size_t(dst_offset),
			C.size_t(cb),
			0,
			nil,
			&c_event)
	}

	if event != nil {
		event.cl_event = c_event
	}

	return CL_int(c_errcode_ret)
}

func CLEnqueueReadBufferRect(command_queue CL_command_queue,
	buffer CL_mem,
	blocking_read CL_bool,
	buffer_origin [3]CL_size_t,
	host_origin [3]CL_size_t,
	region [3]CL_size_t,
	buffer_row_pitch CL_size_t,
	buffer_slice_pitch CL_size_t,
	host_row_pitch CL_size_t,
	host_slice_pitch CL_size_t,
	ptr unsafe.Pointer,
	num_events_in_wait_list CL_uint,
	event_wait_list []CL_event,
	event *CL_event) CL_int {

	if (num_events_in_wait_list == 0 && event_wait_list != nil) ||
		(num_events_in_wait_list != 0 && event_wait_list == nil) ||
		int(num_events_in_wait_list) != len(event_wait_list) {
		return CL_INVALID_EVENT_WAIT_LIST
	}

	var c_buffer_origin [3]C.size_t
	var c_host_origin [3]C.size_t
	var c_region [3]C.size_t

	for i := 0; i < 3; i++ {
		c_buffer_origin[i] = C.size_t(buffer_origin[i])
		c_host_origin[i] = C.size_t(host_origin[i])
		c_region[i] = C.size_t(region[i])
	}

	var c_event C.cl_event
	var c_errcode_ret C.cl_int

	if num_events_in_wait_list != 0 {
		var c_event_wait_list []C.cl_event
		c_event_wait_list = make([]C.cl_event, num_events_in_wait_list)
		for i := 0; i < int(num_events_in_wait_list); i++ {
			c_event_wait_list[i] = event_wait_list[i].cl_event
		}

		c_errcode_ret = C.clEnqueueReadBufferRect(command_queue.cl_command_queue,
			buffer.cl_mem,
			C.cl_bool(blocking_read),
			&c_buffer_origin[0],
			&c_host_origin[0],
			&c_region[0],
			C.size_t(buffer_row_pitch),
			C.size_t(buffer_slice_pitch),
			C.size_t(host_row_pitch),
			C.size_t(host_slice_pitch),
			ptr,
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueReadBufferRect(command_queue.cl_command_queue,
			buffer.cl_mem,
			C.cl_bool(blocking_read),
			&c_buffer_origin[0],
			&c_host_origin[0],
			&c_region[0],
			C.size_t(buffer_row_pitch),
			C.size_t(buffer_slice_pitch),
			C.size_t(host_row_pitch),
			C.size_t(host_slice_pitch),
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

func CLEnqueueWriteBufferRect(command_queue CL_command_queue,
	buffer CL_mem,
	blocking_write CL_bool,
	buffer_origin [3]CL_size_t,
	host_origin [3]CL_size_t,
	region [3]CL_size_t,
	buffer_row_pitch CL_size_t,
	buffer_slice_pitch CL_size_t,
	host_row_pitch CL_size_t,
	host_slice_pitch CL_size_t,
	ptr unsafe.Pointer,
	num_events_in_wait_list CL_uint,
	event_wait_list []CL_event,
	event *CL_event) CL_int {

	if (num_events_in_wait_list == 0 && event_wait_list != nil) ||
		(num_events_in_wait_list != 0 && event_wait_list == nil) ||
		int(num_events_in_wait_list) != len(event_wait_list) {
		return CL_INVALID_EVENT_WAIT_LIST
	}

	var c_buffer_origin [3]C.size_t
	var c_host_origin [3]C.size_t
	var c_region [3]C.size_t

	for i := 0; i < 3; i++ {
		c_buffer_origin[i] = C.size_t(buffer_origin[i])
		c_host_origin[i] = C.size_t(host_origin[i])
		c_region[i] = C.size_t(region[i])
	}

	var c_event C.cl_event
	var c_errcode_ret C.cl_int

	if num_events_in_wait_list != 0 {
		var c_event_wait_list []C.cl_event
		c_event_wait_list = make([]C.cl_event, num_events_in_wait_list)
		for i := 0; i < int(num_events_in_wait_list); i++ {
			c_event_wait_list[i] = event_wait_list[i].cl_event
		}

		c_errcode_ret = C.clEnqueueWriteBufferRect(command_queue.cl_command_queue,
			buffer.cl_mem,
			C.cl_bool(blocking_write),
			&c_buffer_origin[0],
			&c_host_origin[0],
			&c_region[0],
			C.size_t(buffer_row_pitch),
			C.size_t(buffer_slice_pitch),
			C.size_t(host_row_pitch),
			C.size_t(host_slice_pitch),
			ptr,
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueWriteBufferRect(command_queue.cl_command_queue,
			buffer.cl_mem,
			C.cl_bool(blocking_write),
			&c_buffer_origin[0],
			&c_host_origin[0],
			&c_region[0],
			C.size_t(buffer_row_pitch),
			C.size_t(buffer_slice_pitch),
			C.size_t(host_row_pitch),
			C.size_t(host_slice_pitch),
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

func CLEnqueueCopyBufferRect(command_queue CL_command_queue,
	src_buffer CL_mem,
	dst_buffer CL_mem,
	src_origin [3]CL_size_t,
	dst_origin [3]CL_size_t,
	region [3]CL_size_t,
	src_row_pitch CL_size_t,
	src_slice_pitch CL_size_t,
	dst_row_pitch CL_size_t,
	dst_slice_pitch CL_size_t,
	num_events_in_wait_list CL_uint,
	event_wait_list []CL_event,
	event *CL_event) CL_int {

	if (num_events_in_wait_list == 0 && event_wait_list != nil) ||
		(num_events_in_wait_list != 0 && event_wait_list == nil) ||
		int(num_events_in_wait_list) != len(event_wait_list) {
		return CL_INVALID_EVENT_WAIT_LIST
	}

	var c_src_origin [3]C.size_t
	var c_dst_origin [3]C.size_t
	var c_region [3]C.size_t

	for i := 0; i < 3; i++ {
		c_src_origin[i] = C.size_t(src_origin[i])
		c_dst_origin[i] = C.size_t(dst_origin[i])
		c_region[i] = C.size_t(region[i])
	}

	var c_event C.cl_event
	var c_errcode_ret C.cl_int

	if num_events_in_wait_list != 0 {
		var c_event_wait_list []C.cl_event
		c_event_wait_list = make([]C.cl_event, num_events_in_wait_list)
		for i := 0; i < int(num_events_in_wait_list); i++ {
			c_event_wait_list[i] = event_wait_list[i].cl_event
		}

		c_errcode_ret = C.clEnqueueCopyBufferRect(command_queue.cl_command_queue,
			src_buffer.cl_mem,
			dst_buffer.cl_mem,
			&c_src_origin[0],
			&c_dst_origin[0],
			&c_region[0],
			C.size_t(src_row_pitch),
			C.size_t(src_slice_pitch),
			C.size_t(dst_row_pitch),
			C.size_t(dst_slice_pitch),
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueCopyBufferRect(command_queue.cl_command_queue,
			src_buffer.cl_mem,
			dst_buffer.cl_mem,
			&c_src_origin[0],
			&c_dst_origin[0],
			&c_region[0],
			C.size_t(src_row_pitch),
			C.size_t(src_slice_pitch),
			C.size_t(dst_row_pitch),
			C.size_t(dst_slice_pitch),
			0,
			nil,
			&c_event)
	}

	if event != nil {
		event.cl_event = c_event
	}

	return CL_int(c_errcode_ret)
}

func CLEnqueueMapBuffer(command_queue CL_command_queue,
	buffer CL_mem,
	blocking_map CL_bool,
	map_flags CL_map_flags,
	offset CL_size_t,
	cb CL_size_t,
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

	var c_event C.cl_event
	var c_errcode_ret C.cl_int
	var c_ptr_ret unsafe.Pointer

	if num_events_in_wait_list != 0 {
		var c_event_wait_list []C.cl_event
		c_event_wait_list = make([]C.cl_event, num_events_in_wait_list)
		for i := 0; i < int(num_events_in_wait_list); i++ {
			c_event_wait_list[i] = event_wait_list[i].cl_event
		}

		c_ptr_ret = C.clEnqueueMapBuffer(command_queue.cl_command_queue,
			buffer.cl_mem,
			C.cl_bool(blocking_map),
			C.cl_map_flags(map_flags),
			C.size_t(offset),
			C.size_t(cb),
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event,
			&c_errcode_ret)
	} else {
		c_ptr_ret = C.clEnqueueMapBuffer(command_queue.cl_command_queue,
			buffer.cl_mem,
			C.cl_bool(blocking_map),
			C.cl_map_flags(map_flags),
			C.size_t(offset),
			C.size_t(cb),
			0,
			nil,
			&c_event,
			&c_errcode_ret)
	}

	if event != nil {
		event.cl_event = c_event
	}

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return c_ptr_ret
}
