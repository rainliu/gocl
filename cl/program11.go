// +build cl11 cl12 cl20

package cl

/*
#cgo CFLAGS: -I CL
#cgo !darwin LDFLAGS: -lOpenCL
#cgo darwin LDFLAGS: -framework OpenCL

#include "CL/opencl.h"
*/
import "C"

///////////////////////////////////////////////
//OpenCL 1.1
///////////////////////////////////////////////
func CLUnloadCompiler() CL_int {
	return CL_int(C.clUnloadCompiler())
}
