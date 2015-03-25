// +build cl11

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

///////////////////////////////////////////////
//OpenCL 1.1
///////////////////////////////////////////////

func CLEnqueueMarker(command_queue CL_command_queue, event *CL_event) CL_int {
	var c_event C.cl_event
	var c_errcode_ret C.cl_int

	c_errcode_ret = C.clEnqueueMarker(command_queue.cl_command_queue, &c_event)

	if event != nil {
		event.cl_event = c_event
	}

	return CL_int(c_errcode_ret)
}

func CLEnqueueBarrier(command_queue CL_command_queue) CL_int {
	return CL_int(C.clEnqueueBarrier(command_queue.cl_command_queue))
}

func CLEnqueueWaitForEvents(command_queue CL_command_queue,
	num_events_in_wait_list CL_uint,
	event_wait_list []CL_event) CL_int {

	if (num_events_in_wait_list == 0 && event_wait_list != nil) ||
		(num_events_in_wait_list != 0 && event_wait_list == nil) ||
		int(num_events_in_wait_list) != len(event_wait_list) {
		return CL_INVALID_EVENT_WAIT_LIST
	}

	var c_errcode_ret C.cl_int

	if num_events_in_wait_list != 0 {
		var c_event_wait_list []C.cl_event
		c_event_wait_list = make([]C.cl_event, num_events_in_wait_list)
		for i := 0; i < int(num_events_in_wait_list); i++ {
			c_event_wait_list[i] = event_wait_list[i].cl_event
		}

		c_errcode_ret = C.clEnqueueWaitForEvents(command_queue.cl_command_queue,
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0])
	} else {
		c_errcode_ret = C.clEnqueueWaitForEvents(command_queue.cl_command_queue,
			0,
			nil)
	}

	return CL_int(c_errcode_ret)
}
