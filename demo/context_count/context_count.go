package main

import (
	"fmt"
	"gocl/cl"
)

func main() {
	/* Host/device data structures */
	var platform [1]cl.CL_platform_id
	var device [1]cl.CL_device_id
	var context cl.CL_context
	var err cl.CL_int

	var paramValueSize cl.CL_size_t
	var ref_count interface{}

	/* Access the first installed platform */
	err = cl.CLGetPlatformIDs(1, platform[:], nil)
	if err != cl.CL_SUCCESS {
		println("Couldn't find any platforms")
		return
	}

	/* Access the first available device */
	err = cl.CLGetDeviceIDs(platform[0], cl.CL_DEVICE_TYPE_GPU, 1, device[:], nil)
	if err == cl.CL_DEVICE_NOT_FOUND {
		err = cl.CLGetDeviceIDs(platform[0], cl.CL_DEVICE_TYPE_CPU, 1, device[:], nil)
	}
	if err != cl.CL_SUCCESS {
		println("Couldn't find any devices")
		return
	}

	/* Create the context */
	context = cl.CLCreateContext(nil, 1, device[:], nil, nil, &err)
	if err != cl.CL_SUCCESS {
		println("Couldn't create a context")
		return
	}

	/* Determine the reference count */
	err = cl.CLGetContextInfo(context, cl.CL_CONTEXT_REFERENCE_COUNT,
		0, nil, &paramValueSize)

	if err != cl.CL_SUCCESS {
		fmt.Printf("Failed to find context %s.\n", "CL_CONTEXT_REFERENCE_COUNT")
		return
	}

	err = cl.CLGetContextInfo(context, cl.CL_CONTEXT_REFERENCE_COUNT,
		paramValueSize, &ref_count, nil)
	if err != cl.CL_SUCCESS {
		println("Couldn't read the reference count.")
		return
	}
	fmt.Printf("Initial reference count: %d\n", ref_count.(cl.CL_uint))

	/* Update and display the reference count */
	cl.CLRetainContext(context)
	cl.CLGetContextInfo(context, cl.CL_CONTEXT_REFERENCE_COUNT,
		paramValueSize, &ref_count, nil)
	fmt.Printf("Reference count: %d\n", ref_count.(cl.CL_uint))

	cl.CLReleaseContext(context)
	cl.CLGetContextInfo(context, cl.CL_CONTEXT_REFERENCE_COUNT,
		paramValueSize, &ref_count, nil)
	fmt.Printf("Reference count: %d\n", ref_count.(cl.CL_uint))

	cl.CLReleaseContext(context)

	return
}
