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

///////////////////////////////////////////////
//OpenCL 1.2
///////////////////////////////////////////////

func CLEnqueueMigrateMemObjects(command_queue CL_command_queue,
	num_mem_objects CL_uint,
	mem_objects []CL_mem,
	flags CL_mem_migration_flags,
	num_events_in_wait_list CL_uint,
	event_wait_list []CL_event,
	event *CL_event) CL_int {

	if num_mem_objects == 0 || mem_objects == nil || int(num_mem_objects) != len(mem_objects) {
		return CL_INVALID_VALUE
	}

	if (num_events_in_wait_list == 0 && event_wait_list != nil) ||
		(num_events_in_wait_list != 0 && event_wait_list == nil) ||
		int(num_events_in_wait_list) != len(event_wait_list) {
		return CL_INVALID_EVENT_WAIT_LIST
	}

	var c_mem_objects []C.cl_mem
	var c_event C.cl_event
	var c_errcode_ret C.cl_int

	c_mem_objects = make([]C.cl_mem, len(mem_objects))
	for i := 0; i < len(mem_objects); i++ {
		c_mem_objects[i] = mem_objects[i].cl_mem
	}

	if num_events_in_wait_list != 0 {
		var c_event_wait_list []C.cl_event
		c_event_wait_list = make([]C.cl_event, num_events_in_wait_list)
		for i := 0; i < int(num_events_in_wait_list); i++ {
			c_event_wait_list[i] = event_wait_list[i].cl_event
		}

		c_errcode_ret = C.clEnqueueMigrateMemObjects(command_queue.cl_command_queue,
			C.cl_uint(num_mem_objects),
			&c_mem_objects[0],
			C.cl_mem_migration_flags(flags),
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueMigrateMemObjects(command_queue.cl_command_queue,
			C.cl_uint(num_mem_objects),
			&c_mem_objects[0],
			C.cl_mem_migration_flags(flags),
			0,
			nil,
			&c_event)
	}

	if event != nil {
		event.cl_event = c_event
	}

	return CL_int(c_errcode_ret)
}
