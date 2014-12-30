// +build cl11 cl12 cl20

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
		return nil, fmt.Errorf("GetInfo failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	}

	/* Access param data */
	if errCode = cl.CLGetProgramInfo(this.program_id, param_name, param_size, &param_value, nil); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetInfo failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	}

	return param_value, nil
}

func (this *program) Retain() error {
	if errCode := cl.CLRetainProgram(this.program_id); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("Retain failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	}
	return nil
}

func (this *program) Release() error {
	if errCode := cl.CLReleaseProgram(this.program_id); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("Release failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
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
		return fmt.Errorf("Build failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
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
		return nil, fmt.Errorf("GetBuildInfo failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	}

	/* Access param data */
	if errCode = cl.CLGetProgramBuildInfo(this.program_id, device.GetID(), param_name, param_size, &param_value, nil); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetBuildInfo failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	}

	return param_value, nil
}

func (this *program) CreateKernel(kernel_name []byte) (Kernel, error) {
	var errCode cl.CL_int

	if kernel_id := cl.CLCreateKernel(this.program_id, kernel_name, &errCode); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("CreateKernel failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	} else {
		return &kernel{kernel_id}, nil
	}
}

func (this *program) CreateKernels() ([]Kernel, error) {
	var kernels []Kernel
	var kernelIds []cl.CL_kernel
	var numKernels cl.CL_uint
	var errCode cl.CL_int

	/* Determine number of platforms */
	if errCode = cl.CLCreateKernelsInProgram(this.program_id, 0, nil, &numKernels); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("CreateKernels failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	}

	/* Access platforms */
	kernelIds = make([]cl.CL_kernel, numKernels)
	if errCode = cl.CLCreateKernelsInProgram(this.program_id, numKernels, kernelIds, nil); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("CreateKernels failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	}

	kernels = make([]Kernel, numKernels)
	for i := cl.CL_uint(0); i < numKernels; i++ {
		kernels[i] = &kernel{kernelIds[i]}
	}

	return kernels, nil
}
