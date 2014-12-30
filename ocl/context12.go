// +build cl12

package ocl

import (
	"fmt"
	"gocl/cl"
	"unsafe"
)

func (this *context) CreateImage(flags cl.CL_mem_flags,
	image_format *cl.CL_image_format,
	image_desc *cl.CL_image_desc,
	host_ptr unsafe.Pointer) (Image, error) {
	var errCode cl.CL_int

	if memory_id := cl.CLCreateImage(this.context_id,
		flags,
		image_format,
		image_desc,
		host_ptr,
		&errCode); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("CreateImage failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &image{memory{memory_id}}, nil
	}
}

func (this *context) CreateProgramWithBuiltInKernels(devices []Device,
	kernel_names []byte) (Program, error) {
	var errCode cl.CL_int

	numDevices := cl.CL_uint(len(devices))
	deviceIds := make([]cl.CL_device_id, numDevices)
	for i := cl.CL_uint(0); i < numDevices; i++ {
		deviceIds[i] = devices[i].GetID()
	}

	if program_id := cl.CLCreateProgramWithBuiltInKernels(this.context_id, numDevices, deviceIds, kernel_names, &errCode); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("CreateProgramWithBuiltInKernels failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &program{program_id}, nil
	}
}

func (this *context) LinkProgram(devices []Device,
	options []byte,
	input_programs []Program,
	pfn_notify cl.CL_prg_notify,
	user_data unsafe.Pointer) (Program, error) {
	var errCode cl.CL_int

	numDevices := cl.CL_uint(len(devices))
	deviceIds := make([]cl.CL_device_id, numDevices)
	for i := cl.CL_uint(0); i < numDevices; i++ {
		deviceIds[i] = devices[i].GetID()
	}

	numInputPrograms := cl.CL_uint(len(input_programs))
	inputPrograms := make([]cl.CL_program, numInputPrograms)
	for i := cl.CL_uint(0); i < numInputPrograms; i++ {
		inputPrograms[i] = input_programs[i].GetID()
	}

	if program_id := cl.CLLinkProgram(this.context_id, numDevices, deviceIds, options, numInputPrograms, inputPrograms, pfn_notify, user_data, &errCode); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("LinkProgram failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &program{program_id}, nil
	}
}
