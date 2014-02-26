gocl
====

Go OpenCL (GOCL) Binding


Library documentation: 

http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/

http://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/

In order to build this, make sure you have the required drivers and SDK installed for your graphics card. You will need at least and opencl.lib:

AMD ATI should include them with their APP SDK: http://developer.amd.com/resources/heterogeneous-computing/opencl-zone/tools-and-libraries/

For NVIDIA, these are included in the CUDA SDK: https://developer.nvidia.com/opencl

The locations of the library and include file can be supplied by way of environment variables, for exampe: 

export CGO_LDFLAGS=-L$AMDAPPSDKROOT/lib/x86_64 (or null for NVIDIA)

export CGO_CFLAGS=-I$AMDAPPSDKROOT/include (or $NVSDKCOMPUTE_ROOT/OpenCL/common/inc for NVIDIA)


To build OpenCL 1.1 compliance: 

go build/install -tags 'opencl1.1' gocl/cl

go test -tags 'opencl1.1' gocl/test

To build OpenCL 1.2 compliance: 

go build/install -tags 'opencl1.2' gocl/cl

go test -tags 'opencl1.2' gocl/test
