// +build cl11 cl12

package ocl

import (
	"fmt"
	"gocl/cl"
	"unsafe"
)

type context1x interface {
	GetID() cl.CL_context
	GetInfo(param_name cl.CL_context_info) (interface{}, error)
	Retain() error
	Release() error

	CreateBuffer(flags cl.CL_mem_flags,
		size cl.CL_size_t,
		host_ptr unsafe.Pointer) (Buffer, error)
	CreateEvent() (Event, error)
	CreateProgramWithSource(count cl.CL_uint,
		strings [][]byte,
		lengths []cl.CL_size_t) (Program, error)
	CreateProgramWithBinary(devices []Device,
		lengths []cl.CL_size_t,
		binaries [][]byte,
		binary_status []cl.CL_int) (Program, error)

	GetSupportedImageFormats(flags cl.CL_mem_flags,
		image_type cl.CL_mem_object_type) ([]cl.CL_image_format, error)

	//cl1x only, not in cl20
	CreateCommandQueue(device Device,
		properties []cl.CL_command_queue_properties) (CommandQueue, error)
	CreateSampler(normalized_coords cl.CL_bool,
		addressing_mode cl.CL_addressing_mode,
		filter_mode cl.CL_filter_mode) (Sampler, error)
}

func (this *context) CreateCommandQueue(device Device,
	properties []cl.CL_command_queue_properties) (CommandQueue, error) {
	var property cl.CL_command_queue_properties
	var errCode cl.CL_int

	if properties == nil {
		property = 0
	} else {
		property = properties[0]
	}

	if command_queue_id := cl.CLCreateCommandQueue(this.context_id, device.GetID(), property, &errCode); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("CreateCommandQueue failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	} else {
		return &command_queue{command_queue_id}, nil
	}
}

func (this *context) CreateSampler(normalized_coords cl.CL_bool,
	addressing_mode cl.CL_addressing_mode,
	filter_mode cl.CL_filter_mode) (Sampler, error) {
	var errCode cl.CL_int

	if sampler_id := cl.CLCreateSampler(this.context_id, normalized_coords, addressing_mode, filter_mode, &errCode); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("CreateSampler failure with errcode_ret %d: %s", errCode, cl.ERROR_CODES_STRINGS[-errCode])
	} else {
		return &sampler{sampler_id}, nil
	}
}
