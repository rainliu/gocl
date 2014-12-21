// +build cl11 cl12 cl20

package ocl

import (
	"fmt"
	"gocl/cl"
	"unsafe"
)

type kernel struct {
	kernel_id cl.CL_kernel
}

func (this *kernel) GetID() cl.CL_kernel {
	return this.kernel_id
}

func (this *kernel) GetInfo(param_name cl.CL_kernel_info) (interface{}, error) {
	/* param data */
	var param_value interface{}
	var param_size cl.CL_size_t
	var errCode cl.CL_int

	/* Find size of param data */
	if errCode = cl.CLGetKernelInfo(this.kernel_id, param_name, 0, nil, &param_size); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetInfo failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}

	/* Access param data */
	if errCode = cl.CLGetKernelInfo(this.kernel_id, param_name, param_size, &param_value, nil); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetInfo failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}

	return param_value, nil
}

func (this *kernel) Retain() error {
	if errCode := cl.CLRetainKernel(this.kernel_id); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("Retain failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}
	return nil
}

func (this *kernel) Release() error {
	if errCode := cl.CLReleaseKernel(this.kernel_id); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("Release failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}
	return nil
}

func (this *kernel) SetArg(arg_index cl.CL_uint,
	arg_size cl.CL_size_t,
	arg_value unsafe.Pointer) error {
	if errCode := cl.CLSetKernelArg(this.kernel_id, arg_index, arg_size, arg_value); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("SetArg failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}
	return nil
}

func (this *kernel) GetWorkGroupInfo(device Device,
	param_name cl.CL_kernel_work_group_info) (interface{}, error) {
	/* param data */
	var param_value interface{}
	var param_size cl.CL_size_t
	var errCode cl.CL_int

	/* Find size of param data */
	if errCode = cl.CLGetKernelWorkGroupInfo(this.kernel_id, device.GetID(), param_name, 0, nil, &param_size); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetWorkGroupInfo failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}

	/* Access param data */
	if errCode = cl.CLGetKernelWorkGroupInfo(this.kernel_id, device.GetID(), param_name, param_size, &param_value, nil); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetWorkGroupInfo failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}

	return param_value, nil
}

func (this *kernel) EnqueueNDRange(queue CommandQueue,
	work_dim cl.CL_uint,
	global_work_offset []cl.CL_size_t,
	global_work_size []cl.CL_size_t,
	local_work_size []cl.CL_size_t,
	event_wait_list []Event) (Event, error) {
	var errCode cl.CL_int
	var event_id cl.CL_event

	numEvents := cl.CL_uint(len(event_wait_list))
	events := make([]cl.CL_event, numEvents)
	for i := cl.CL_uint(0); i < numEvents; i++ {
		events[i] = event_wait_list[i].GetID()
	}

	if errCode = cl.CLEnqueueNDRangeKernel(queue.GetID(),
		this.kernel_id,
		work_dim,
		global_work_offset,
		global_work_size,
		local_work_size,
		numEvents,
		events,
		&event_id); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("EnqueueNDRange failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}
