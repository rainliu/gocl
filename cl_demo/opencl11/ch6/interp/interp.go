package main

import (
	"fmt"
	"gocl/cl"
	"gocl/cl_demo/utils"
	"unsafe"
)

const SCALE_FACTOR = 3
const INPUT_FILE = "input.png"
const OUTPUT_FILE = "output.png"
const PROGRAM_FILE = "interp.cl"

var KERNEL_FUNC = []byte("interp")

func main() {

	/* Host/device data structures */
	var device []cl.CL_device_id
	var context cl.CL_context
	var queue cl.CL_command_queue
	var program *cl.CL_program
	var kernel cl.CL_kernel
	var err cl.CL_int

	var err1 error
	var global_size [2]cl.CL_size_t

	/* Image data */
	var input_pixels, output_pixels []uint16
	var png_format cl.CL_image_format
	var input_image, output_image cl.CL_mem
	var origin, region [3]cl.CL_size_t
	var width, height cl.CL_size_t

	/* Open input file and read image data */
	input_pixels, width, height, err1 = utils.Read_image_data(INPUT_FILE)
	if err1 != nil {
		return
	} else {
		fmt.Printf("width=%d, height=%d", width, height)
	}
	output_pixels = make([]uint16, width*height*SCALE_FACTOR*SCALE_FACTOR)

	/* Create a device and context */
	device = utils.Create_device()
	context = cl.CLCreateContext(nil, 1, device[:], nil, nil, &err)
	if err < 0 {
		println("Couldn't create a context")
		return
	}

	/* Build the program and create a kernel */
	program = utils.Build_program(context, device[:], PROGRAM_FILE, nil)
	kernel = cl.CLCreateKernel(*program, KERNEL_FUNC, &err)
	if err < 0 {
		fmt.Printf("Couldn't create a kernel: %d", err)
		return
	}

	/* Create input image object */
	png_format.Image_channel_order = cl.CL_LUMINANCE
	png_format.Image_channel_data_type = cl.CL_UNORM_INT16
	input_image = cl.CLCreateImage2D(context,
		cl.CL_MEM_READ_ONLY|cl.CL_MEM_COPY_HOST_PTR,
		&png_format, width, height, 0, unsafe.Pointer(&input_pixels[0]), &err)
	if err < 0 {
		println("Couldn't create the image object")
		return
	}

	/* Create output image object */
	output_image = cl.CLCreateImage2D(context,
		cl.CL_MEM_WRITE_ONLY, &png_format, SCALE_FACTOR*width,
		SCALE_FACTOR*height, 0, nil, &err)
	if err < 0 {
		println("Couldn't create the image object")
		return
	}

	/* Create kernel arguments */
	err = cl.CLSetKernelArg(kernel, 0, cl.CL_size_t(unsafe.Sizeof(input_image)), unsafe.Pointer(&input_image))
	err |= cl.CLSetKernelArg(kernel, 1, cl.CL_size_t(unsafe.Sizeof(output_image)), unsafe.Pointer(&output_image))
	if err < 0 {
		println("Couldn't set a kernel argument")
		return
	}

	/* Create a command queue */
	queue = cl.CLCreateCommandQueue(context, device[0], 0, &err)
	if err < 0 {
		println("Couldn't create a command queue")
		return
	}

	/* Enqueue kernel */
	global_size[0] = width
	global_size[1] = height
	err = cl.CLEnqueueNDRangeKernel(queue, kernel, 2, nil, global_size[:],
		nil, 0, nil, nil)
	if err < 0 {
		println("Couldn't enqueue the kernel")
		return
	}

	/* Read the image object */
	origin[0] = 0
	origin[1] = 0
	origin[2] = 0
	region[0] = SCALE_FACTOR * width
	region[1] = SCALE_FACTOR * height
	region[2] = 1
	err = cl.CLEnqueueReadImage(queue, output_image, cl.CL_TRUE, origin,
		region, 0, 0, unsafe.Pointer(&output_pixels[0]), 0, nil, nil)
	if err < 0 {
		println("Couldn't read from the image object")
		return
	}

	/* Create output PNG file and write data */
	utils.Write_image_data(OUTPUT_FILE, output_pixels, SCALE_FACTOR*width, SCALE_FACTOR*height)

	/* Deallocate resources */
	cl.CLReleaseMemObject(input_image)
	cl.CLReleaseMemObject(output_image)
	cl.CLReleaseKernel(kernel)
	cl.CLReleaseCommandQueue(queue)
	cl.CLReleaseProgram(*program)
	cl.CLReleaseContext(context)
}
