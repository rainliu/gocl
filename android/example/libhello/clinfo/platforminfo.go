package clinfo

import (
	"fmt"
	"golang.org/x/mobile/cl"
)

///
// Display information for a particular platform.
// Assumes that all calls to clGetPlatformInfo returns
// a value of type char[], which is valid for OpenCL 1.1.
//
func DisplayPlatformInfo(id cl.CL_platform_id,
	name cl.CL_platform_info,
	str string) string{
	var errNum cl.CL_int
	var paramValueSize cl.CL_size_t

	errNum = cl.CLGetPlatformInfo(id,
		name,
		0,
		nil,
		&paramValueSize)
	if errNum != cl.CL_SUCCESS {
		return fmt.Sprintf("Failed to find OpenCL platform %s.\n", str)
	}

	var info interface{}
	errNum = cl.CLGetPlatformInfo(id,
		name,
		paramValueSize,
		&info,
		nil)
	if errNum != cl.CL_SUCCESS {
		return fmt.Sprintf("Failed to find OpenCL platform %s.\n", str)
	}

	return fmt.Sprintf("\t%-24s: %v\n", str, info)
}
