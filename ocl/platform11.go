// +build cl11

package ocl

import (
	"errors"
	"gocl/cl"
)

func (this *platform) UnloadCompiler() error {
	if errCode := cl.CLUnloadCompiler(); errCode != cl.CL_SUCCESS {
		return errors.New("UnloadCompiler failure with errcode_ret " + string(errCode))
	}
	return nil
}
