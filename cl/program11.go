// +build cl11

package cl

/*
#cgo CFLAGS: -I CL
#cgo !darwin LDFLAGS: -lOpenCL
#cgo darwin LDFLAGS: -framework OpenCL

#ifdef __APPLE__
#include "OpenCL/opencl.h"
#else
#include "CL/opencl.h"
#endif
*/
import "C"

///////////////////////////////////////////////
//OpenCL 1.1
///////////////////////////////////////////////
func CLUnloadCompiler() CL_int {
	return CL_int(C.clUnloadCompiler())
}
