package main

import (
	"fmt"
	"gocl/cl"
	"testing"
	"unsafe"
)

func TestContext(t *testing.T) {
	/* Host/device data structures */
	var platform [1]cl.CL_platform_id
	var device [1]cl.CL_device_id
	var context cl.CL_context
	var err cl.CL_int

	var paramValueSize cl.CL_size_t
	var ref_count interface{}
	user_data := []byte("Hello, I am callback")

	/* Access the first installed platform */
	err = cl.CLGetPlatformIDs(1, platform[:], nil)
	if err != cl.CL_SUCCESS {
		t.Errorf("Couldn't find any platforms")
	}

	/* Access the first available device */
	err = cl.CLGetDeviceIDs(platform[0], cl.CL_DEVICE_TYPE_GPU, 1, device[:], nil)
	if err == cl.CL_DEVICE_NOT_FOUND {
		err = cl.CLGetDeviceIDs(platform[0], cl.CL_DEVICE_TYPE_CPU, 1, device[:], nil)
	}
	if err != cl.CL_SUCCESS {
		t.Errorf("Couldn't find any devices")
	}

	/* Create the context */
	context = cl.CLCreateContext(nil, 1, device[:], my_contex_notify, unsafe.Pointer(&user_data), &err)
	if err != cl.CL_SUCCESS {
		t.Errorf("Couldn't create a context")
	}

	/* Determine the reference count */
	err = cl.CLGetContextInfo(context, cl.CL_CONTEXT_REFERENCE_COUNT,
		0, nil, &paramValueSize)

	if err != cl.CL_SUCCESS {
		t.Errorf("Failed to find context %s.\n", "CL_CONTEXT_REFERENCE_COUNT")
	}

	err = cl.CLGetContextInfo(context, cl.CL_CONTEXT_REFERENCE_COUNT,
		paramValueSize, &ref_count, nil)
	if err != cl.CL_SUCCESS {
		t.Errorf("Couldn't read the reference count.")
	}
	t.Logf("Initial reference count: %d\n", ref_count.(cl.CL_uint))

	/* Update and display the reference count */
	cl.CLRetainContext(context)
	cl.CLGetContextInfo(context, cl.CL_CONTEXT_REFERENCE_COUNT,
		paramValueSize, &ref_count, nil)
	t.Logf("Reference count: %d\n", ref_count.(cl.CL_uint))

	cl.CLReleaseContext(context)
	cl.CLGetContextInfo(context, cl.CL_CONTEXT_REFERENCE_COUNT,
		paramValueSize, &ref_count, nil)
	t.Logf("Reference count: %d\n", ref_count.(cl.CL_uint))

	cl.CLReleaseContext(context)
}

func my_contex_notify(errinfo string, private_info unsafe.Pointer, cb int, user_data unsafe.Pointer) {
	fmt.Printf("my_contex_notify callback: %s\n", *((*[]byte)(user_data)))
}
