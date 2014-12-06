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
	platforms, err = ocl.GetPlatforms()
	if err != nil {
		t.Errorf(err.Error())
	}

	/* Find extensions of all platforms */
	platform_index := -1
	for i := 0; i < len(platforms); i++ {
		/* Get extension data */
		param_value, err = platforms[i].GetInfo(cl.CL_PLATFORM_EXTENSIONS)
		if err != nil {
			t.Errorf(err.Error())
		} else {
			t.Logf("Platform %d supports extensions: %s\n", i, param_value)
		}

		/* Look for ICD extension */
		if strings.Contains(param_value, icd_ext) {
			platform_index = i
			break
		}
	}

	/* Display whether ICD extension is supported */
	if platform_index > -1 {
		t.Logf("Platform %d supports the %s extension.\n", platform_index, icd_ext)
	} else {
		t.Logf("No platforms support the %s extension.\n", icd_ext)
	}
}
