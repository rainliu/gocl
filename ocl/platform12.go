// +build cl12

package ocl

import (
	"errors"
	"gocl/cl"
)

func (this *platform) UnloadCompiler() error {
	if errCode := cl.CLUnloadPlatformCompiler(this.platform_id); errCode != cl.CL_SUCCESS {
		return errors.New("UnloadCompiler failure with errcode_ret " + string(errCode))
	}
	return nil
}
