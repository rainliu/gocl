// +build cl12

package ocl

import (
	"fmt"
	"gocl/cl"
)

type Kernel interface {
	kernel1x

	//cl12
	GetArgInfo(arg_index cl.CL_uint,
		param_name cl.CL_kernel_arg_info) (interface{}, error)
}

func (this *kernel) GetArgInfo(arg_index cl.CL_uint,
	param_name cl.CL_kernel_arg_info) (interface{}, error) {
	/* param data */
	var param_value interface{}
	var param_size cl.CL_size_t
	var errCode cl.CL_int

	/* Find size of param data */
	if errCode = cl.CLGetKernelArgInfo(this.kernel_id, arg_index, param_name, 0, nil, &param_size); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetArgInfo failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}

	/* Access param data */
	if errCode = cl.CLGetKernelArgInfo(this.kernel_id, arg_index, param_name, param_size, &param_value, nil); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetArgInfo failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}

	return param_value, nil
}
