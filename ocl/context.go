// +build cl11 cl12

package ocl

import (
	"errors"
	"gocl/cl"
	"unsafe"
)

type Context interface {
	GetInfo(param_name cl.CL_context_info) (interface{}, error)
	Retain() error
	Release() error
}

type context struct {
	context_id cl.CL_context
}

func CreateContext(properties []cl.CL_context_properties,
	devices []Device,
	pfn_notify cl.CL_ctx_notify,
	user_data unsafe.Pointer) (Context, error) {
	//var errCode cl.CL_int

	return nil, nil
}

func CreateContextFromType(properties []cl.CL_context_properties,
	device_type cl.CL_device_type,
	pfn_notify cl.CL_ctx_notify,
	user_data unsafe.Pointer) (Context, error) {
	//var errCode cl.CL_int

	return nil, nil
}

func (this *context) GetInfo(param_name cl.CL_context_info) (interface{}, error) {
	/* param data */
	var param_value interface{}
	var param_size cl.CL_size_t
	var errCode cl.CL_int

	/* Find size of param data */
	if errCode = cl.CLGetContextInfo(this.context_id, param_name, 0, nil, &param_size); errCode != cl.CL_SUCCESS {
		return "", errors.New("GetInfo failure with errcode_ret " + string(errCode))
	}

	/* Access param data */
	if errCode = cl.CLGetContextInfo(this.context_id, param_name, param_size, &param_value, nil); errCode != cl.CL_SUCCESS {
		return "", errors.New("GetInfo failure with errcode_ret " + string(errCode))
	}

	return param_value, nil
}

func (this *context) Retain() error {
	if errCode := cl.CLRetainContext(this.context_id); errCode != cl.CL_SUCCESS {
		return errors.New("Retain failure with errcode_ret " + string(errCode))
	}
	return nil
}

func (this *context) Release() error {
	if errCode := cl.CLReleaseContext(this.context_id); errCode != cl.CL_SUCCESS {
		return errors.New("Release failure with errcode_ret " + string(errCode))
	}
	return nil
}
