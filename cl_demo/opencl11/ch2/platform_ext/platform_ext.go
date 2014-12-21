package main

import (
	"fmt"
	"gocl/cl"
	"strings"
)

func main() {

	/* Host data structures */
	var platforms []cl.CL_platform_id
	var num_platforms cl.CL_uint
	var err, i, platform_index cl.CL_int

	platform_index = -1

	/* Extension data */
	var ext_data interface{}
	var ext_size cl.CL_size_t
	const icd_ext string = "cl_khr_icd"

	/* Find number of platforms */
	err = cl.CLGetPlatformIDs(1, nil, &num_platforms)
	if err != cl.CL_SUCCESS {
		println("Couldn't find any platforms.")
		return
	}

	/* Access all installed platforms */
	platforms = make([]cl.CL_platform_id, num_platforms)

	err = cl.CLGetPlatformIDs(num_platforms, platforms, nil)
	if err != cl.CL_SUCCESS {
		println("Couldn't get any platforms.")
		return
	}

	/* Find extensions of all platforms */
	for i = 0; i < cl.CL_int(num_platforms); i++ {

		/* Find size of extension data */
		err = cl.CLGetPlatformInfo(platforms[i],
			cl.CL_PLATFORM_EXTENSIONS, 0, nil, &ext_size)
		if err != cl.CL_SUCCESS {
			println("Couldn't read extension data.")
			return
		}

		/* Access extension data */		
		err = cl.CLGetPlatformInfo(platforms[i], cl.CL_PLATFORM_EXTENSIONS,
			ext_size, &ext_data, nil)
		if err == cl.CL_SUCCESS {
			fmt.Printf("Platform %d supports extensions: %s\n", i, ext_data)
		}

		/* Look for ICD extension */
		if strings.Contains(ext_data.(string), icd_ext) {
			platform_index = i
			break
		}
	}

	/* Display whether ICD extension is supported */
	if platform_index > -1 {
		fmt.Printf("Platform %d supports the %s extension.\n",
			platform_index, icd_ext)
	} else {
		fmt.Printf("No platforms support the %s extension.\n", icd_ext)
	}
}
