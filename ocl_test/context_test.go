package ocl_test

import (
	"fmt"
	"gocl/cl"
	"gocl/ocl"
	"testing"
	"unsafe"
)

func TestContext(t *testing.T) {
	/* Host/device data structures */
	var platforms []ocl.Platform
	var devices []ocl.Device
	var context ocl.Context
	var err error

	var ref_count interface{}
	user_data := []byte("Hello, I am callback")

	/* Identify a platform */
	if platforms, err = ocl.GetPlatforms(); err != nil {
		t.Errorf(err.Error())
		return
	}

	/* Determine connected devices */
	if devices, err = platforms[0].GetDevices(cl.CL_DEVICE_TYPE_GPU); err != nil {
		if devices, err = platforms[0].GetDevices(cl.CL_DEVICE_TYPE_CPU); err != nil {
			t.Errorf(err.Error())
			return
		}
	}

	/* Create the context */
	if context, err = devices[0].CreateContext(nil, my_contex_notify, unsafe.Pointer(&user_data)); err != nil {
		t.Errorf(err.Error())
		return
	}
	defer context.Release()

	/* Get the reference count */
	if ref_count, err = context.GetInfo(cl.CL_CONTEXT_REFERENCE_COUNT); err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Logf("Initial reference count: %d\n", ref_count.(cl.CL_uint))

	/* Update and display the reference count */
	context.Retain()
	if ref_count, err = context.GetInfo(cl.CL_CONTEXT_REFERENCE_COUNT); err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Logf("Reference count: %d\n", ref_count.(cl.CL_uint))

	context.Release()
	if ref_count, err = context.GetInfo(cl.CL_CONTEXT_REFERENCE_COUNT); err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Logf("Reference count: %d\n", ref_count.(cl.CL_uint))
}

func my_contex_notify(errinfo string, private_info unsafe.Pointer, cb int, user_data unsafe.Pointer) {
	fmt.Printf("my_contex_notify callback: %s\n", *((*[]byte)(user_data)))
}
