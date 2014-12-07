package ocl_test

import (
	"gocl/cl"
	"gocl/ocl"
	"strings"
	"testing"
)

func TestPlatform(t *testing.T) {
	/* Host data structures */
	var platforms []ocl.Platform
	var err error

	/* Param value */
	var param_value string
	const icd_ext string = "cl_khr_icd"

	/* Get all installed platforms */
	if platforms, err = ocl.GetPlatforms(); err != nil {
		t.Errorf(err.Error())
		return
	} else {
		t.Logf("Number of platform: %d\n", len(platforms))
	}

	/* Find extensions of all platforms */
	platform_index := -1
	for i := 0; i < len(platforms); i++ {
		DisplayPlatformInfo(t, &platforms[i], cl.CL_PLATFORM_PROFILE, "CL_PLATFORM_PROFILE")
		DisplayPlatformInfo(t, &platforms[i], cl.CL_PLATFORM_VERSION, "CL_PLATFORM_VERSION")
		DisplayPlatformInfo(t, &platforms[i], cl.CL_PLATFORM_NAME, "CL_PLATFORM_NAME")
		DisplayPlatformInfo(t, &platforms[i], cl.CL_PLATFORM_VENDOR, "CL_PLATFORM_VENDOR")
		DisplayPlatformInfo(t, &platforms[i], cl.CL_PLATFORM_EXTENSIONS, "CL_PLATFORM_EXTENSIONS")

		/* Get extension data */
		if param_value, err = platforms[i].GetInfo(cl.CL_PLATFORM_EXTENSIONS); err != nil {
			t.Errorf(err.Error())
			return
		} else {
			/* Look for ICD extension */
			if strings.Contains(param_value, icd_ext) {
				platform_index = i
				break
			}
		}
	}

	/* Display whether ICD extension is supported */
	if platform_index > -1 {
		t.Logf("Platform %d supports the %s extension.\n", platform_index, icd_ext)
	} else {
		t.Logf("No platforms support the %s extension.\n", icd_ext)
	}
}

func DisplayPlatformInfo(t *testing.T, platform *ocl.Platform, param_name cl.CL_platform_info, param_name_str string) {
	if param_value, err := platform.GetInfo(param_name); err != nil {
		t.Errorf(err.Error())
	} else {
		t.Logf("\t %s:\t %s\n", param_name_str, param_value)
	}
}
