// +build cl11 cl12

package ocl

import (
	"gocl/cl"
	"unsafe"
)

type program1x interface {
	GetID() cl.CL_program
	GetInfo(param_name cl.CL_program_info) (interface{}, error)
	Retain() error
	Release() error

	Build(devices []Device,
		options []byte,
		pfn_notify cl.CL_prg_notify,
		user_data unsafe.Pointer) error
	GetBuildInfo(device Device,
		param_name cl.CL_program_build_info) (interface{}, error)
}
