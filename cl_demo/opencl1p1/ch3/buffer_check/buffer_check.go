package main

import (
	"fmt"
	"gocl/cl"
	"unsafe"
	"gocl/demo/ch0"
)

func main() {

	/* Host/device data structures */
	var device []cl.CL_device_id
	var context cl.CL_context
	var err cl.CL_int

	/* Data and buffers */
	var main_data [100]float32
	var main_buffer, sub_buffer cl.CL_mem
	var main_buffer_mem, sub_buffer_mem interface{}
	var main_buffer_size, sub_buffer_size interface{}
	var buffer_size cl.CL_size_t
	var buffer_mem cl.CL_ulong
	var region cl.CL_buffer_region

	/* Create device and context */
	device = ch0.Create_device()
	context = cl.CLCreateContext(nil, 1, device[:], nil, nil, &err)
	if err < 0 {
		println("Couldn't create a context")
		return
	}

	/* Create a buffer to hold 100 floating-point values */
	main_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_READ_ONLY|
		cl.CL_MEM_COPY_HOST_PTR, cl.CL_size_t(unsafe.Sizeof(main_data)), unsafe.Pointer(&main_data[0]), &err)
	if err < 0 {
		println("Couldn't create a buffer")
		return
	}

	/* Create a sub-buffer containing values 30-49 */
	region.Origin = 30 * cl.CL_size_t(unsafe.Sizeof(main_data[0]))
	region.Size = 20 * cl.CL_size_t(unsafe.Sizeof(main_data[0]))
	fmt.Printf("origin=%d, size=%d\n", region.Origin, region.Size)

	sub_buffer = cl.CLCreateSubBuffer(main_buffer, cl.CL_MEM_READ_ONLY|
		cl.CL_MEM_COPY_HOST_PTR, cl.CL_BUFFER_CREATE_TYPE_REGION, unsafe.Pointer(&region), &err)
	if err < 0 {
		fmt.Printf("Couldn't create a sub-buffer, errcode=%d\n", err)
		return
	}

	/* Obtain size information about the buffers */
	cl.CLGetMemObjectInfo(main_buffer, cl.CL_MEM_SIZE,
		cl.CL_size_t(unsafe.Sizeof(buffer_size)), &main_buffer_size, nil)
	cl.CLGetMemObjectInfo(sub_buffer, cl.CL_MEM_SIZE,
		cl.CL_size_t(unsafe.Sizeof(buffer_size)), &sub_buffer_size, nil)
	fmt.Printf("Main buffer size: %v\n", main_buffer_size.(cl.CL_size_t))
	fmt.Printf("Sub-buffer size:  %v\n", sub_buffer_size.(cl.CL_size_t))

	/* Obtain the host pointers */
	cl.CLGetMemObjectInfo(main_buffer, cl.CL_MEM_HOST_PTR, cl.CL_size_t(unsafe.Sizeof(buffer_mem)),
		&main_buffer_mem, nil)
	cl.CLGetMemObjectInfo(sub_buffer, cl.CL_MEM_HOST_PTR, cl.CL_size_t(unsafe.Sizeof(buffer_mem)),
		&sub_buffer_mem, nil)
	fmt.Printf("Main buffer memory address: %v\n", main_buffer_mem.(cl.CL_ulong))
	fmt.Printf("Sub-buffer memory address:  %v\n", sub_buffer_mem.(cl.CL_ulong))

	/* Print the address of the main data */
	fmt.Printf("Main array address: %v\n", main_data)

	/* Deallocate resources */
	cl.CLReleaseMemObject(main_buffer)
	cl.CLReleaseMemObject(sub_buffer)
	cl.CLReleaseContext(context)
}
