gocl
====

Go OpenCL (GOCL) Binding (http://www.gocl.org)


Library documentation: 

http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/

http://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/

In order to build this, make sure you have the required drivers and SDK installed for your graphics card. You will need at least opencl.lib:

AMD ATI should include them with their APP SDK: http://developer.amd.com/tools-and-sdks/heterogeneous-computing/amd-accelerated-parallel-processing-app-sdk/

For NVIDIA, these are included in the CUDA SDK: https://developer.nvidia.com/opencl

The locations of the library and include file can be supplied by way of environment variables, for example: 

export CGO_LDFLAGS=-L$AMDAPPSDKROOT/lib/x86_64     			(or null for NVIDIA and Mac OSX)

export CGO_CFLAGS=-I$GOPATH/src/gocl/cl/CL     				(gocl/cl/CL have the latest OpenCL 2.0 include files from https://www.khronos.org/registry/cl/)

===============================================

In gocl/cl/*.go files, please change #cgo LDFLAGS according to your OS:

for Mac OSX:  "#cgo LDFLAGS: -framework OpenCL"

for Linux :   "#cgo LDFLAGS: -lOpenCL"

===============================================

To build OpenCL 1.1 compliance: 

go build/install -tags 'cl11' gocl/cl

go build/install -tags 'cl11' gocl/demo/opencl1p1/ch(x)         (Examples in "OpenCL in Action")

go test -tags 'cl11' gocl/test

===============================================

To build OpenCL 1.2 compliance: 

go build/install -tags 'cl12' gocl/cl

go build/install -tags 'cl12' gocl/demo/opencl1p2/chapter(x)    (Examples in "Heterogeneous Computing with OpenCL, 2nd Edition")

go test -tags 'cl12' gocl/test
