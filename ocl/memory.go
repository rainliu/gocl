// +build cl11 cl12

package ocl

import (
	"errors"
	"gocl/cl"
	"unsafe"
)

type Memory interface {
	GetID() cl.CL_mem
	GetInfo(param_name cl.CL_mem_info) (interface{}, error)
	Retain() error
	Release() error

	SetDestructorCallback(pfn_notify cl.CL_mem_notify, user_data unsafe.Pointer) error

	//to be fix CL_event
	EnqueueUnmap(queue CommandQueue, mapped_ptr unsafe.Pointer, event_wait_list []cl.CL_event) (cl.CL_event, error)
}

type memory struct {
	memory_id cl.CL_mem
}

func (this *memory) GetID() cl.CL_mem {
	return this.memory_id
}

func (this *memory) GetInfo(param_name cl.CL_mem_info) (interface{}, error) {
	/* param data */
	var param_value interface{}
	var param_size cl.CL_size_t
	var errCode cl.CL_int

	/* Find size of param data */
	if errCode = cl.CLGetMemObjectInfo(this.memory_id, param_name, 0, nil, &param_size); errCode != cl.CL_SUCCESS {
		return nil, errors.New("GetInfo failure with errcode_ret " + string(errCode))
	}

	/* Access param data */
	if errCode = cl.CLGetMemObjectInfo(this.memory_id, param_name, param_size, &param_value, nil); errCode != cl.CL_SUCCESS {
		return nil, errors.New("GetInfo failure with errcode_ret " + string(errCode))
	}

	return param_value, nil
}

func (this *memory) Retain() error {
	if errCode := cl.CLRetainMemObject(this.memory_id); errCode != cl.CL_SUCCESS {
		return errors.New("Retain failure with errcode_ret " + string(errCode))
	}
	return nil
}

func (this *memory) Release() error {
	if errCode := cl.CLReleaseMemObject(this.memory_id); errCode != cl.CL_SUCCESS {
		return errors.New("Release failure with errcode_ret " + string(errCode))
	}
	return nil
}

func (this *memory) SetDestructorCallback(pfn_notify cl.CL_mem_notify, user_data unsafe.Pointer) error {
	if errCode := cl.CLSetMemObjectDestructorCallback(this.memory_id, pfn_notify, user_data); errCode != cl.CL_SUCCESS {
		return errors.New("SetDestructorCallback failure with errcode_ret " + string(errCode))
	} else {
		return nil
	}
}

func (this *memory) EnqueueUnmap(queue CommandQueue,
	mapped_ptr unsafe.Pointer,
	event_wait_list []cl.CL_event) (cl.CL_event, error) {
	var event cl.CL_event
	return event, nil
}
