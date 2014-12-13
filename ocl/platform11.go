// +build cl11

package ocl

import (
	"fmt"
	"gocl/cl"
)

func (this *platform) UnloadCompiler() error {
	if errCode := cl.CLUnloadCompiler(); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("UnloadCompiler failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}
	return nil
}
