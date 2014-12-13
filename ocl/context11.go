// +build cl11

package ocl

import (
	"fmt"
	"gocl/cl"
	"unsafe"
)

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

func (this *context) CreateImage2D(flags cl.CL_mem_flags,
	image_format *cl.CL_image_format,
	image_width cl.CL_size_t,
	image_height cl.CL_size_t,
	image_row_pitch cl.CL_size_t,
	host_ptr unsafe.Pointer) (Image, error) {
	var errCode cl.CL_int

	if memory_id := cl.CLCreateImage2D(this.context_id,
		flags,
		image_format,
		image_width,
		image_height,
		image_row_pitch,
		host_ptr,
		&errCode); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("CreateImage2D failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &image{memory{memory_id}}, nil
	}
}

func (this *context) CreateImage3D(flags cl.CL_mem_flags,
	image_format *cl.CL_image_format,
	image_width cl.CL_size_t,
	image_height cl.CL_size_t,
	image_depth cl.CL_size_t,
	image_row_pitch cl.CL_size_t,
	image_slice_pitch cl.CL_size_t,
	host_ptr unsafe.Pointer) (Image, error) {
	var errCode cl.CL_int

	if memory_id := cl.CLCreateImage3D(this.context_id,
		flags,
		image_format,
		image_width,
		image_height,
		image_depth,
		image_row_pitch,
		image_slice_pitch,
		host_ptr,
		&errCode); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("CreateImage3D failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &image{memory{memory_id}}, nil
	}
}
