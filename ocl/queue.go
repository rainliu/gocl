// +build cl11 cl12

package ocl

import (
	//"errors"
	"gocl/cl"
	//"unsafe"
)

type CommandQueue interface {
	//GetInfo(param_name cl.CL_context_info) (interface{}, error)
	Retain() error
	Release() error
}

type command_queue struct {
	command_queue_id cl.CL_command_queue
}