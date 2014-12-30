// +build cl12

package ocl

import (
	"gocl/cl"
	"unsafe"
)

type Buffer interface {
	buffer1x

	//cl12
	EnqueueFill(queue CommandQueue,
		pattern unsafe.Pointer,
		pattern_size cl.CL_size_t,
		offset cl.CL_size_t,
		cb cl.CL_size_t,
		event_wait_list []Event) (Event, error)
}

type Context interface {
	context1x

	//cl12
	CreateImage(flags cl.CL_mem_flags,
		image_format *cl.CL_image_format,
		image_desc *cl.CL_image_desc,
		host_ptr unsafe.Pointer) (Image, error)
	CreateProgramWithBuiltInKernels(devices []Device,
		kernel_names []byte) (Program, error)
	LinkProgram(devices []Device,
		options []byte,
		input_programs []Program,
		pfn_notify cl.CL_prg_notify,
		user_data unsafe.Pointer) (Program, error)
}

type Device interface {
	device1x

	//cl12
	CreateSubDevices(properties []cl.CL_device_partition_property) ([]Device, error)
	Retain() error
	Release() error
}

type Image interface {
	image1x

	//cl12
	EnqueueFill(queue CommandQueue,
		fill_color unsafe.Pointer,
		origin [3]cl.CL_size_t,
		region [3]cl.CL_size_t,
		event_wait_list []Event) (Event, error)
}

type Kernel interface {
	kernel1x

	//cl12
	GetArgInfo(arg_index cl.CL_uint,
		param_name cl.CL_kernel_arg_info) (interface{}, error)
}

type Program interface {
	program1x

	//cl12
	Compile(devices []Device,
		options []byte,
		input_headers []Program,
		header_include_names [][]byte,
		pfn_notify cl.CL_prg_notify,
		user_data unsafe.Pointer) error
}

type CommandQueue interface {
	queue1x

	//cl12
	EnqueueMarkerWithWaitList(event_wait_list []Event) (Event, error)
	EnqueueBarrierWithWaitList(event_wait_list []Event) (Event, error)
	EnqueueMigrateMemObjects(mem_objects []Memory, flags cl.CL_mem_migration_flags, event_wait_list []Event) (Event, error)
}
