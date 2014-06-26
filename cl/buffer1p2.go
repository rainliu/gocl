// +build CL12

package cl

/*
#cgo CFLAGS: -I CL
#cgo LDFLAGS: -lOpenCL

#include "CL/opencl.h"
*/
import "C"
import "unsafe"

///////////////////////////////////////////////
//OpenCL 1.2
///////////////////////////////////////////////

func CLEnqueueFillBuffer(command_queue CL_command_queue,
	buffer CL_mem,
	pattern unsafe.Pointer,
	pattern_size CL_size_t,
	offset CL_size_t,
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

		c_errcode_ret = C.clEnqueueFillBuffer(command_queue.cl_command_queue,
			buffer.cl_mem,
			pattern,
			C.size_t(pattern_size),
			C.size_t(offset),
			C.size_t(cb),
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueFillBuffer(command_queue.cl_command_queue,
			buffer.cl_mem,
			pattern,
			C.size_t(pattern_size),
			C.size_t(offset),
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
