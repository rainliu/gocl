package main

import (
    "fmt"
    "golang.org/x/mobile/cl"
)

func initCL() {
    var errNum cl.CL_int
    //var numPlatforms cl.CL_uint
    var platformIds []cl.CL_platform_id
    //var context cl.CL_context

    // First, query the total number of platforms
    errNum = cl.CLGetPlatformIDs(0, nil, &numPlatforms)
    if errNum != cl.CL_SUCCESS || numPlatforms <= 0 {
        println("Failed to find any OpenCL platform.")
        return
    }

    // Next, allocate memory for the installed plaforms, and qeury
    // to get the list.
    platformIds = make([]cl.CL_platform_id, numPlatforms)

    // First, query the total number of platforms
    errNum = cl.CLGetPlatformIDs(numPlatforms, platformIds, nil)
    if errNum != cl.CL_SUCCESS {
        println("Failed to find any OpenCL platforms.")
        return
    }

    fmt.Printf("Number of platforms: \t%d\n", numPlatforms)

    // Iterate through the list of platforms displaying associated information
    for i := cl.CL_uint(0); i < numPlatforms; i++ {
        // First we display information associated with the platform
        DisplayPlatformInfo(
            platformIds[i],
            cl.CL_PLATFORM_PROFILE,
            "CL_PLATFORM_PROFILE")
        DisplayPlatformInfo(
            platformIds[i],
            cl.CL_PLATFORM_VERSION,
            "CL_PLATFORM_VERSION")
        DisplayPlatformInfo(
            platformIds[i],
            cl.CL_PLATFORM_VENDOR,
            "CL_PLATFORM_VENDOR")
        DisplayPlatformInfo(
            platformIds[i],
            cl.CL_PLATFORM_EXTENSIONS,
            "CL_PLATFORM_EXTENSIONS")

        // Now query the set of devices associated with the platform
        var numDevices cl.CL_uint
        errNum = cl.CLGetDeviceIDs(platformIds[i],
            cl.CL_DEVICE_TYPE_ALL,
            0,
            nil,
            &numDevices)
        if errNum != cl.CL_SUCCESS {
            println("Failed to find OpenCL devices.")
            return
        }

        devices := make([]cl.CL_device_id, numDevices)
        errNum = cl.CLGetDeviceIDs(platformIds[i],
            cl.CL_DEVICE_TYPE_ALL,
            numDevices,
            devices,
            nil)
        if errNum != cl.CL_SUCCESS {
            println("Failed to find OpenCL devices.")
            return
        }

        fmt.Printf("\n\tNumber of devices: \t%d\n", numDevices)

        // Iterate through each device, displaying associated information
        for j := cl.CL_uint(0); j < numDevices; j++ {
            DisplayDeviceInfo(devices[j],
                cl.CL_DEVICE_TYPE,
                "CL_DEVICE_TYPE")

            DisplayDeviceInfo(devices[j],
                cl.CL_DEVICE_NAME,
                "CL_DEVICE_NAME")

            DisplayDeviceInfo(devices[j],
                cl.CL_DEVICE_VENDOR,
                "CL_DEVICE_VENDOR")

            //DisplayDeviceInfo(devices[j],
            //	cl.CL_DRIVER_VERSION,
            //	"CL_DRIVER_VERSION")

            DisplayDeviceInfo(devices[j],
                cl.CL_DEVICE_PROFILE,
                "CL_DEVICE_PROFILE")
			
			fmt.Printf("\n")
        }
    }
}
