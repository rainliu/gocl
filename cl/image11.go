// +build cl11

package cl

/*
#cgo CFLAGS: -I CL
#cgo !darwin LDFLAGS: -lOpenCL
#cgo darwin LDFLAGS: -framework OpenCL

#ifdef __APPLE__
#include "OpenCL/opencl.h"
#else
#include "CL/opencl.h"
#endif
 */
import "C"
import "unsafe"

///////////////////////////////////////////////
//OpenCL 1.1 deprecated
///////////////////////////////////////////////

func CLCreateImage2D(context CL_context,
	flags CL_mem_flags,
	image_format *CL_image_format,
	image_width CL_size_t,
	image_height CL_size_t,
	image_row_pitch CL_size_t,
	host_ptr unsafe.Pointer,
	errcode_ret *CL_int) CL_mem {

	var c_image_format C.cl_image_format
	var c_errcode_ret C.cl_int
	var c_image C.cl_mem

	c_image_format.image_channel_order = C.cl_channel_order(image_format.Image_channel_order)
	c_image_format.image_channel_data_type = C.cl_channel_type(image_format.Image_channel_data_type)

	c_image = C.clCreateImage2D(context.cl_context,
		C.cl_mem_flags(flags),
		&c_image_format,
		C.size_t(image_width),
		C.size_t(image_height),
		C.size_t(image_row_pitch),
		host_ptr,
		&c_errcode_ret)

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_mem{c_image}
}

func CLCreateImage3D(context CL_context,
	flags CL_mem_flags,
	image_format *CL_image_format,
	image_width CL_size_t,
	image_height CL_size_t,
	image_depth CL_size_t,
	image_row_pitch CL_size_t,
	image_slice_pitch CL_size_t,
	host_ptr unsafe.Pointer,
	errcode_ret *CL_int) CL_mem {

	var c_image_format C.cl_image_format
	var c_errcode_ret C.cl_int
	var c_image C.cl_mem

	c_image_format.image_channel_order = C.cl_channel_order(image_format.Image_channel_order)
	c_image_format.image_channel_data_type = C.cl_channel_type(image_format.Image_channel_data_type)

	c_image = C.clCreateImage3D(context.cl_context,
		C.cl_mem_flags(flags),
		&c_image_format,
		C.size_t(image_width),
		C.size_t(image_height),
		C.size_t(image_depth),
		C.size_t(image_row_pitch),
		C.size_t(image_slice_pitch),
		host_ptr,
		&c_errcode_ret)

	if errcode_ret != nil {
		*errcode_ret = CL_int(c_errcode_ret)
	}

	return CL_mem{c_image}
}
