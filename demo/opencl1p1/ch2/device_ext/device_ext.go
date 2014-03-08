package main

import (
	"fmt"
	"gocl/cl"
)

func main() {

	/* Host/device data structures */
	var platform [1]cl.CL_platform_id
	var devices []cl.CL_device_id
	var num_devices cl.CL_uint
	var i, err cl.CL_int

	/* Extension data */
	var paramValueSize cl.CL_size_t
	var name_data interface{}
	var ext_data interface{}
	var addr_data interface{}

	/* Identify a platform */
	err = cl.CLGetPlatformIDs(1, platform[:], nil)
	if err != cl.CL_SUCCESS {
		println("Couldn't find any platforms")
		return
	}

	/* Determine number of connected devices */
	err = cl.CLGetDeviceIDs(platform[0], cl.CL_DEVICE_TYPE_ALL, 0, nil, &num_devices)
	if err != cl.CL_SUCCESS {
		println("Couldn't find any devices")
		return
	}

	/* Access connected devices */
	devices = make([]cl.CL_device_id, num_devices)

	err = cl.CLGetDeviceIDs(platform[0], cl.CL_DEVICE_TYPE_ALL,
		num_devices, devices, nil)
	if err != cl.CL_SUCCESS {
		println("Couldn't get any devices.")
		return
	}

	/* Obtain data for each connected device */
	for i = 0; i < cl.CL_int(num_devices); i++ {	

		err = cl.CLGetDeviceInfo(devices[i],
			cl.CL_DEVICE_NAME,
			0,
			nil,
			&paramValueSize)

		if err != cl.CL_SUCCESS {
			fmt.Printf("Failed to find OpenCL device info %s.\n", "NAME")
			return
		}
		
		err = cl.CLGetDeviceInfo(devices[i],
			cl.CL_DEVICE_NAME,
			paramValueSize,
			&name_data,
			nil)
		if err != cl.CL_SUCCESS {
			fmt.Printf("Failed to find OpenCL device info %s.\n", "NAME")
			return
		}

		err = cl.CLGetDeviceInfo(devices[i],
			cl.CL_DEVICE_ADDRESS_BITS,
			0,
			nil,
			&paramValueSize)

		if err != cl.CL_SUCCESS {
			fmt.Printf("Failed to find OpenCL device info %s.\n", "NAME")
			return
		}
		
		err = cl.CLGetDeviceInfo(devices[i],
			cl.CL_DEVICE_ADDRESS_BITS,
			paramValueSize,
			&addr_data,
			nil)
		if err != cl.CL_SUCCESS {
			fmt.Printf("Failed to find OpenCL device info %s.\n", "NAME")
			return
		}

		err = cl.CLGetDeviceInfo(devices[i],
			cl.CL_DEVICE_EXTENSIONS,
			0,
			nil,
			&paramValueSize)

		if err != cl.CL_SUCCESS {
			fmt.Printf("Failed to find OpenCL device info %s.\n", "NAME")
			return
		}
		
		err = cl.CLGetDeviceInfo(devices[i],
			cl.CL_DEVICE_EXTENSIONS,
			paramValueSize,
			&ext_data,
			nil)
		if err != cl.CL_SUCCESS {
			fmt.Printf("Failed to find OpenCL device info %s.\n", "NAME")
			return
		}

		fmt.Printf("NAME: %s\nADDRESS_WIDTH: %d\nEXTENSIONS: %s\n\n",
			name_data.(string), addr_data.(cl.CL_uint), ext_data.(string))
	}
}
