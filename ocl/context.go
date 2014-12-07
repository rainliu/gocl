// +build cl11 cl12

package ocl

import (
	"gocl/cl"
)

type Context interface {
}

type context struct {
	context_id cl.CL_context
}

func CreateContext() Context {

	return nil
}
