// +build opencl1.1 opencl1.2

package cl

/*
#cgo CFLAGS: -I CL
#cgo LDFLAGS: -lOpenCL

#include "CL/opencl.h"

extern void go_evt_notify(cl_event event, cl_int event_command_exec_status, void *user_data);
static void CL_CALLBACK c_evt_notify(cl_event event, cl_int event_command_exec_status, void *user_data) {
	go_evt_notify(event, event_command_exec_status, user_data);
}

static cl_int CLSetEventCallback(	cl_event event,
									cl_int command_exec_callback_type,
									void *user_data){
    return clSetEventCallback(event, command_exec_callback_type, c_evt_notify, user_data);
}
*/
import "C"

import (
	"unsafe"
)

type CL_evt_notify func(event CL_event, event_command_exec_status CL_int, user_data unsafe.Pointer)

var evt_notify CL_evt_notify

//export go_evt_notify
func go_evt_notify(event C.cl_event, event_command_exec_status C.cl_int, user_data unsafe.Pointer) {
	evt_notify(CL_event{event}, CL_int(event_command_exec_status), user_data)
}

func CLCreateUserEvent(context CL_context,
	errcode_ret *CL_int) CL_event {

	var c_event C.cl_event
	var c_errcode_ret C.cl_int

	c_event = C.clCreateUserEvent(context.cl_context, &c_errcode_ret)

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_event{c_event}
}

func CLSetUserEventStatus(event CL_event,
	execution_status CL_int) CL_int {
	return CL_int(C.clSetUserEventStatus(event.cl_event, C.cl_int(execution_status)))
}

func CLWaitForEvents(num_events CL_uint,
	event_list []CL_event) CL_int {

	if num_events == 0 || int(num_events) != len(event_list) {
		return CL_INVALID_VALUE
	}

	var c_event_list []C.cl_event
	c_event_list = make([]C.cl_event, len(event_list))
	for i := 0; i < len(event_list); i++ {
		c_event_list[i] = event_list[i].cl_event
	}

	return CL_int(C.clWaitForEvents(C.cl_uint(num_events), &c_event_list[0]))
}

func CLGetEventInfo(event CL_event,
	param_name CL_event_info,
	param_value_size CL_size_t,
	param_value *interface{},
	param_value_size_ret *CL_size_t) CL_int {

	var ret C.cl_int

	if (param_value_size == 0 || param_value == nil) && param_value_size_ret == nil {
		ret = C.clGetEventInfo(event.cl_event,
			C.cl_event_info(param_name),
			0,
			nil,
			nil)
	} else {
		var size_ret C.size_t

		if param_value_size == 0 || param_value == nil {
			ret = C.clGetEventInfo(event.cl_event,
				C.cl_event_info(param_name),
				0,
				nil,
				&size_ret)
		} else {
			switch param_name {
			case CL_EVENT_COMMAND_QUEUE:
				var value C.cl_command_queue
				ret = C.clGetEventInfo(event.cl_event,
					C.cl_event_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&size_ret)

				*param_value = CL_command_queue{value}

			case CL_EVENT_COMMAND_TYPE:
				var value C.cl_command_type
				ret = C.clGetEventInfo(event.cl_event,
					C.cl_event_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&size_ret)

				*param_value = CL_command_type(value)

			case CL_EVENT_CONTEXT:
				var value C.cl_context
				ret = C.clGetEventInfo(event.cl_event,
					C.cl_event_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&size_ret)

				*param_value = CL_context{value}

			case CL_EVENT_REFERENCE_COUNT:
				var value C.cl_uint
				ret = C.clGetEventInfo(event.cl_event,
					C.cl_event_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&size_ret)

				*param_value = CL_uint(value)

			case CL_EVENT_COMMAND_EXECUTION_STATUS:
				var value C.cl_int
				ret = C.clGetEventInfo(event.cl_event,
					C.cl_event_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&size_ret)

				*param_value = CL_int(value)
			default:
				return CL_INVALID_VALUE
			}
		}

		if param_value_size_ret != nil {
			*param_value_size_ret = CL_size_t(size_ret)
		}
	}

	return CL_int(ret)
}

func CLSetEventCallback(event CL_event,
	command_exec_callback_type CL_int,
	pfn_notify CL_evt_notify,
	user_data unsafe.Pointer) CL_int {

	if pfn_notify != nil {
		evt_notify = pfn_notify

		return CL_int(C.CLSetEventCallback(event.cl_event,
			C.cl_int(command_exec_callback_type),
			user_data))
	} else {
		evt_notify = nil

		return CL_int(C.clSetEventCallback(event.cl_event,
			C.cl_int(command_exec_callback_type),
			nil,
			nil))
	}
}

func CLRetainEvent(event CL_event) CL_int {
	return CL_int(C.clRetainEvent(event.cl_event))
}

func CLReleaseEvent(event CL_event) CL_int {
	return CL_int(C.clReleaseEvent(event.cl_event))
}

func CLGEtEventProfilingInfo(event CL_event,
	param_name CL_profiling_info,
	param_value_size CL_size_t,
	param_value *interface{},
	param_value_size_ret *CL_size_t) CL_int {

	var ret C.cl_int

	if (param_value_size == 0 || param_value == nil) && param_value_size_ret == nil {
		ret = C.clGetEventProfilingInfo(event.cl_event,
			C.cl_profiling_info(param_name),
			0,
			nil,
			nil)
	} else {
		var size_ret C.size_t

		if param_value_size == 0 || param_value == nil {
			ret = C.clGetEventProfilingInfo(event.cl_event,
				C.cl_profiling_info(param_name),
				0,
				nil,
				&size_ret)
		} else {
			switch param_name {
			case CL_PROFILING_COMMAND_QUEUED,
				CL_PROFILING_COMMAND_SUBMIT,
				CL_PROFILING_COMMAND_START,
				CL_PROFILING_COMMAND_END:

				var value C.cl_ulong
				ret = C.clGetEventProfilingInfo(event.cl_event,
					C.cl_profiling_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&size_ret)

				*param_value = CL_ulong(value)
			default:
				return CL_INVALID_VALUE
			}
		}

		if param_value_size_ret != nil {
			*param_value_size_ret = CL_size_t(size_ret)
		}
	}

	return CL_int(ret)
}
