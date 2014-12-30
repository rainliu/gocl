// +build cl11

package ocl

type Buffer interface {
	buffer1x
}

type Context interface {
	context1x

	//cl11
	CreateImage2D(flags cl.CL_mem_flags,
		image_format *cl.CL_image_format,
		image_width cl.CL_size_t,
		image_height cl.CL_size_t,
		image_row_pitch cl.CL_size_t,
		host_ptr unsafe.Pointer) (Image, error)
	CreateImage3D(flags cl.CL_mem_flags,
		image_format *cl.CL_image_format,
		image_width cl.CL_size_t,
		image_height cl.CL_size_t,
		image_depth cl.CL_size_t,
		image_row_pitch cl.CL_size_t,
		image_slice_pitch cl.CL_size_t,
		host_ptr unsafe.Pointer) (Image, error)
}

type Device interface {
	device1x
}

type Image interface {
	image1x
}

type Kernel interface {
	kernel1x
}

type Program interface {
	program1x
}

type CommandQueue interface {
	queue1x

	//cl11
	EnqueueMarker() (Event, error)
	EnqueueBarrier() error
	EnqueueWaitForEvents(event_wait_list []Event) error
}