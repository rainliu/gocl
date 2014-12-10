// +build cl11 cl12

package ocl

import (
	"errors"
	"gocl/cl"
)

type command_queue struct {
	command_queue_id cl.CL_command_queue
}

func (this *command_queue) GetID() cl.CL_command_queue {
	return this.command_queue_id
}

func (this *command_queue) GetInfo(param_name cl.CL_command_queue_info) (interface{}, error) {
	/* param data */
	var param_value interface{}
	var param_size cl.CL_size_t
	var errCode cl.CL_int

	/* Find size of param data */
	if errCode = cl.CLGetCommandQueueInfo(this.command_queue_id, param_name, 0, nil, &param_size); errCode != cl.CL_SUCCESS {
		return nil, errors.New("GetInfo failure with errcode_ret " + string(errCode))
	}

	/* Access param data */
	if errCode = cl.CLGetCommandQueueInfo(this.command_queue_id, param_name, param_size, &param_value, nil); errCode != cl.CL_SUCCESS {
		return nil, errors.New("GetInfo failure with errcode_ret " + string(errCode))
	}

	return param_value, nil
}

func (this *command_queue) Retain() error {
	if errCode := cl.CLRetainCommandQueue(this.command_queue_id); errCode != cl.CL_SUCCESS {
		return errors.New("Retain failure with errcode_ret " + string(errCode))
	}
	return nil
}

func (this *command_queue) Release() error {
	if errCode := cl.CLReleaseCommandQueue(this.command_queue_id); errCode != cl.CL_SUCCESS {
		return errors.New("Release failure with errcode_ret " + string(errCode))
	}
	return nil
}
