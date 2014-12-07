// +build cl11

package ocl

import "gocl/cl"

type Device interface {
	GetID() cl.CL_device_id
	GetInfo(param_name cl.CL_device_info) (interface{}, error)
}
