// +build cl20

package ocl

import (
	"gocl/cl"
	"unsafe"
)

type context20 interface {
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

	//cl20
	CreateCommandQueueWithProperties(device Device,
		properties []cl.CL_command_queue_properties) (CommandQueue, error)
	CreateSamplerWithProperties(normalized_coords cl.CL_bool,
		addressing_mode cl.CL_addressing_mode,
		filter_mode cl.CL_filter_mode) (Sampler, error)
}
