// +build cl11 cl12

package ocl

import (	
	"gocl/cl"
)

type queue1x interface {
	GetID() cl.CL_command_queue
	GetInfo(param_name cl.CL_command_queue_info) (interface{}, error)
	Retain() error
	Release() error
}