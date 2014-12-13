// +build cl11 cl12

package ocl

import (
	"fmt"
	"gocl/cl"
	"unsafe"
)

type program struct {
	program_id cl.CL_program
}

func (this *program) GetID() cl.CL_program {
	return this.program_id
}

func (this *program) GetInfo(param_name cl.CL_program_info) (interface{}, error) {
	/* param data */
	var param_value interface{}
	var param_size cl.CL_size_t
	var errCode cl.CL_int

	/* Find size of param data */
	if errCode = cl.CLGetProgramInfo(this.program_id, param_name, 0, nil, &param_size); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetInfo failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}

	/* Access param data */
	if errCode = cl.CLGetProgramInfo(this.program_id, param_name, param_size, &param_value, nil); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetInfo failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}

	return param_value, nil
}

func (this *program) Retain() error {
	if errCode := cl.CLRetainProgram(this.program_id); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("Retain failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}
	return nil
}

func (this *program) Release() error {
	if errCode := cl.CLReleaseProgram(this.program_id); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("Release failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}
	return nil
}

func (this *program) Build(devices []Device,
	options []byte,
	pfn_notify cl.CL_prg_notify,
	user_data unsafe.Pointer) error {

	numDevices := cl.CL_uint(len(devices))
	deviceIds := make([]cl.CL_device_id, numDevices)
	for i := cl.CL_uint(0); i < numDevices; i++ {
		deviceIds[i] = devices[i].GetID()
	}

	if errCode := cl.CLBuildProgram(this.program_id, numDevices, deviceIds, options, pfn_notify, user_data); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("Build failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}
	return nil
}

func (this *program) GetBuildInfo(device Device, param_name cl.CL_program_build_info) (interface{}, error) {
	/* param data */
	var param_value interface{}
	var param_size cl.CL_size_t
	var errCode cl.CL_int

	/* Find size of param data */
	if errCode = cl.CLGetProgramBuildInfo(this.program_id, device.GetID(), param_name, 0, nil, &param_size); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetBuildInfo failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}

	/* Access param data */
	if errCode = cl.CLGetProgramBuildInfo(this.program_id, device.GetID(), param_name, param_size, &param_value, nil); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetBuildInfo failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}

	return param_value, nil
}
