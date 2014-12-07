// +build cl11

package ocl

import (
	"gocl/cl"
	"unsafe"
)

type Device interface {
	GetID() cl.CL_device_id
	GetInfo(param_name cl.CL_device_info) (interface{}, error)
	CreateContext(properties []cl.CL_context_properties,
		pfn_notify cl.CL_ctx_notify,
		user_data unsafe.Pointer) (Context, error)
}
