// +build cl20

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
extern void go_svm_notify(cl_command_queue command_queue, cl_uint num_svm_pointers, void *svm_pointers[], void *user_data);
static void CL_CALLBACK c_svm_notify(cl_command_queue command_queue, cl_uint num_svm_pointers, void *svm_pointers[], void *user_data) {
	go_svm_notify(command_queue, num_svm_pointers, svm_pointers, user_data);
}

static cl_int CLEnqueueSVMFree(cl_command_queue command_queue,
					           cl_uint num_svm_pointers,
					           void *svm_pointers[],
					           void *user_data,
					           cl_uint num_events_in_wait_list,
					           const cl_event *event_wait_list,
					           cl_event *event){
	return clEnqueueSVMFree(command_queue, num_svm_pointers, svm_pointers, c_svm_notify, user_data, num_events_in_wait_list, event_wait_list, event);
}
*/
import "C"

import "unsafe"

type CL_svm_notify func(command_queue CL_command_queue, num_svm_pointers CL_uint, svm_pointers []unsafe.Pointer, user_data unsafe.Pointer)

var svm_notify map[unsafe.Pointer]CL_svm_notify

func init() {
	svm_notify = make(map[unsafe.Pointer]CL_svm_notify)
}

//export go_svm_notify
func go_svm_notify(command_queue C.cl_command_queue, num_svm_pointers C.cl_uint, svm_pointers *unsafe.Pointer, user_data unsafe.Pointer) {
	var c_user_data []unsafe.Pointer
	c_user_data = *(*[]unsafe.Pointer)(user_data)

	var c_svm_pointers []unsafe.Pointer
	c_svm_pointers = make([]unsafe.Pointer, num_svm_pointers)
	for i := 0; i < int(num_svm_pointers); i++ {
		c_svm_pointers[i] = unsafe.Pointer(uintptr(unsafe.Pointer(svm_pointers)) + uintptr(i)*unsafe.Sizeof(*svm_pointers))
	}
	svm_notify[c_user_data[1]](CL_command_queue{command_queue}, CL_uint(num_svm_pointers), c_svm_pointers, c_user_data[0])
}

func CLSVMAlloc(context CL_context,
	flags CL_mem_flags,
	size CL_size_t,
	alignment CL_uint) unsafe.Pointer {
	return unsafe.Pointer(C.clSVMAlloc(context.cl_context,
		C.cl_svm_mem_flags(flags),
		C.size_t(size),
		C.cl_uint(alignment)))
}

func CLSVMFree(context CL_context,
	svm_pointer unsafe.Pointer) {
	C.clSVMFree(context.cl_context, svm_pointer)
}

func CLEnqueueSVMFree(command_queue CL_command_queue,
	num_svm_pointers CL_uint,
	svm_pointers []unsafe.Pointer,
	pfn_notify CL_svm_notify,
	user_data unsafe.Pointer,
	num_events_in_wait_list CL_uint,
	event_wait_list []CL_event,
	event *CL_event) CL_int {

	if num_svm_pointers == 0 ||
		svm_pointers == nil ||
		int(num_svm_pointers) != len(svm_pointers) {
		return CL_INVALID_VALUE
	}

	if (num_events_in_wait_list == 0 && event_wait_list != nil) ||
		(num_events_in_wait_list != 0 && event_wait_list == nil) ||
		int(num_events_in_wait_list) != len(event_wait_list) {
		return CL_INVALID_EVENT_WAIT_LIST
	}

	for i := 0; i < len(svm_pointers); i++ {
		if svm_pointers[i] == nil {
			return CL_INVALID_VALUE
		}
	}
	var c_svm_pointers *unsafe.Pointer
	var c_event_wait_list []C.cl_event
	var c_event_wait_list_ptr *C.cl_event
	var c_event C.cl_event
	var c_errcode_ret C.cl_int

	if num_events_in_wait_list != 0 {
		c_event_wait_list = make([]C.cl_event, num_events_in_wait_list)
		for i := 0; i < int(num_events_in_wait_list); i++ {
			c_event_wait_list[i] = event_wait_list[i].cl_event
		}
		c_event_wait_list_ptr = &c_event_wait_list[0]
	} else {
		c_event_wait_list_ptr = nil
	}

	c_svm_pointers = &svm_pointers[0]

	if pfn_notify != nil {
		var c_user_data []unsafe.Pointer
		c_user_data = make([]unsafe.Pointer, 2)
		c_user_data[0] = user_data
		c_user_data[1] = unsafe.Pointer(&pfn_notify)

		svm_notify[c_user_data[1]] = pfn_notify

		c_errcode_ret = C.CLEnqueueSVMFree(command_queue.cl_command_queue,
			C.cl_uint(num_svm_pointers),
			c_svm_pointers,
			unsafe.Pointer(&c_user_data),
			C.cl_uint(num_events_in_wait_list),
			c_event_wait_list_ptr,
			&c_event)

	} else {
		c_errcode_ret = C.clEnqueueSVMFree(command_queue.cl_command_queue,
			C.cl_uint(num_svm_pointers),
			c_svm_pointers,
			nil,
			nil,
			C.cl_uint(num_events_in_wait_list),
			c_event_wait_list_ptr,
			&c_event)

	}

	if event != nil {
		event.cl_event = c_event
	}

	return CL_int(c_errcode_ret)
}

func CLEnqueueSVMMemcpy(command_queue CL_command_queue,
	blocking_copy CL_bool,
	dst_ptr unsafe.Pointer,
	src_ptr unsafe.Pointer,
	size CL_size_t,
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

		c_errcode_ret = C.clEnqueueSVMMemcpy(command_queue.cl_command_queue,
			C.cl_bool(blocking_copy),
			dst_ptr,
			src_ptr,
			C.size_t(size),
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueSVMMemcpy(command_queue.cl_command_queue,
			C.cl_bool(blocking_copy),
			dst_ptr,
			src_ptr,
			C.size_t(size),
			0,
			nil,
			&c_event)
	}

	if event != nil {
		event.cl_event = c_event
	}

	return CL_int(c_errcode_ret)
}

func CLEnqueueSVMMemFill(command_queue CL_command_queue,
	svm_ptr unsafe.Pointer,
	pattern unsafe.Pointer,
	pattern_size CL_size_t,
	size CL_size_t,
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

		c_errcode_ret = C.clEnqueueSVMMemFill(command_queue.cl_command_queue,
			svm_ptr,
			pattern,
			C.size_t(pattern_size),
			C.size_t(size),
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueSVMMemFill(command_queue.cl_command_queue,
			svm_ptr,
			pattern,
			C.size_t(pattern_size),
			C.size_t(size),
			0,
			nil,
			&c_event)
	}

	if event != nil {
		event.cl_event = c_event
	}

	return CL_int(c_errcode_ret)
}

func CLEnqueueSVMMap(command_queue CL_command_queue,
	blocking_map CL_bool,
	map_flags CL_map_flags,
	svm_ptr unsafe.Pointer,
	size CL_size_t,
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

		c_errcode_ret = C.clEnqueueSVMMap(command_queue.cl_command_queue,
			C.cl_bool(blocking_map),
			C.cl_map_flags(map_flags),
			svm_ptr,
			C.size_t(size),
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueSVMMap(command_queue.cl_command_queue,
			C.cl_bool(blocking_map),
			C.cl_map_flags(map_flags),
			svm_ptr,
			C.size_t(size),
			0,
			nil,
			&c_event)
	}

	if event != nil {
		event.cl_event = c_event
	}

	return CL_int(c_errcode_ret)
}

func CLEnqueueSVMUnmap(command_queue CL_command_queue,
	svm_ptr unsafe.Pointer,
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

		c_errcode_ret = C.clEnqueueSVMUnmap(command_queue.cl_command_queue,
			svm_ptr,
			C.cl_uint(num_events_in_wait_list),
			&c_event_wait_list[0],
			&c_event)
	} else {
		c_errcode_ret = C.clEnqueueSVMUnmap(command_queue.cl_command_queue,
			svm_ptr,
			0,
			nil,
			&c_event)
	}

	if event != nil {
		event.cl_event = c_event
	}

	return CL_int(c_errcode_ret)
}
