package cl

/*
#cgo CFLAGS: -I CL
#cgo LDFLAGS: -lOpenCL

#include "CL/opencl.h"
*/
import "C"

import (
    "unsafe"
)

func CLGetDeviceIDs(platform CL_platform_id,
    device_type CL_device_type,
    num_entries CL_uint,
    devices []CL_device_id,
    num_devices *CL_uint) CL_int {

    var ret C.cl_int

    if (num_entries == 0 || devices == nil) && num_devices == nil {
        ret = C.clGetDeviceIDs(platform.cl_platform_id,
            C.cl_device_type(device_type),
            0,
            (*C.cl_device_id)(nil),
            (*C.cl_uint)(nil))
    } else {
        var num C.cl_uint

        if num_entries == 0 || devices == nil {
            ret = C.clGetDeviceIDs(platform.cl_platform_id,
                C.cl_device_type(device_type),
                0,
                (*C.cl_device_id)(nil),
                &num)
        } else {
            devices_id := make([]C.cl_device_id, len(devices))
            ret = C.clGetDeviceIDs(platform.cl_platform_id,
                C.cl_device_type(device_type),
                C.cl_uint(num_entries),
                &devices_id[0],
                &num)
            if ret == C.CL_SUCCESS {
                for i := 0; i < len(devices); i++ {
                    devices[i].cl_device_id = devices_id[i]
                }
            }
        }

        if num_devices != nil {
            *num_devices = CL_uint(num)
        }
    }

    return CL_int(ret)
}

func CLGetDeviceInfo(device CL_device_id,
    param_name CL_device_info,
    param_value_size CL_size_t,
    param_value *interface{},
    param_value_size_ret *CL_size_t) CL_int {

    var ret C.cl_int

    if (param_value_size == 0 || param_value == nil) && param_value_size_ret == nil {
        ret = C.clGetDeviceInfo(device.cl_device_id,
            C.cl_device_info(param_name),
            0,
            nil,
            (*C.size_t)(nil))
    } else {
        var size_ret C.size_t

        if param_value_size == 0 || param_value == nil {
            ret = C.clGetDeviceInfo(device.cl_device_id,
                C.cl_device_info(param_name),
                0,
                nil,
                &size_ret)
        } else {
            switch param_name {

            case CL_DEVICE_AVAILABLE,
                CL_DEVICE_COMPILER_AVAILABLE,
                CL_DEVICE_ENDIAN_LITTLE,
                CL_DEVICE_ERROR_CORRECTION_SUPPORT,
                CL_DEVICE_HOST_UNIFIED_MEMORY,
                CL_DEVICE_IMAGE_SUPPORT:
                //DEVICE_LINKER_AVAILABLE, //2.0?
                //DEVICE_PREFERRED_INTEROP_USER_SYNC://2.0?

                var value C.cl_bool
                ret = C.clGetDeviceInfo(device.cl_device_id,
                    C.cl_device_info(param_name),
                    C.size_t(unsafe.Sizeof(value)),
                    unsafe.Pointer(&value),
                    &size_ret)
                *param_value = value == C.CL_TRUE

            case CL_DEVICE_ADDRESS_BITS,
                CL_DEVICE_MAX_CLOCK_FREQUENCY,
                CL_DEVICE_MAX_COMPUTE_UNITS,
                CL_DEVICE_MAX_CONSTANT_ARGS,
                CL_DEVICE_MAX_READ_IMAGE_ARGS,
                CL_DEVICE_MAX_SAMPLERS,
                CL_DEVICE_MAX_WORK_ITEM_DIMENSIONS,
                CL_DEVICE_MAX_WRITE_IMAGE_ARGS,
                CL_DEVICE_MEM_BASE_ADDR_ALIGN,
                CL_DEVICE_MIN_DATA_TYPE_ALIGN_SIZE,
                CL_DEVICE_NATIVE_VECTOR_WIDTH_CHAR,
                CL_DEVICE_NATIVE_VECTOR_WIDTH_SHORT,
                CL_DEVICE_NATIVE_VECTOR_WIDTH_INT,
                CL_DEVICE_NATIVE_VECTOR_WIDTH_LONG,
                CL_DEVICE_NATIVE_VECTOR_WIDTH_FLOAT,
                CL_DEVICE_NATIVE_VECTOR_WIDTH_DOUBLE,
                CL_DEVICE_NATIVE_VECTOR_WIDTH_HALF,
                CL_DEVICE_PREFERRED_VECTOR_WIDTH_CHAR,
                CL_DEVICE_PREFERRED_VECTOR_WIDTH_SHORT,
                CL_DEVICE_PREFERRED_VECTOR_WIDTH_INT,
                CL_DEVICE_PREFERRED_VECTOR_WIDTH_LONG,
                CL_DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT,
                CL_DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE,
                CL_DEVICE_PREFERRED_VECTOR_WIDTH_HALF,
                CL_DEVICE_VENDOR_ID:
                //CL_DEVICE_PARTITION_MAX_SUB_DEVICES,//2.0
                //CL_DEVICE_REFERENCE_COUNT://2.0

                var value C.cl_uint
                ret = C.clGetDeviceInfo(device.cl_device_id,
                    C.cl_device_info(param_name),
                    C.size_t(unsafe.Sizeof(value)),
                    unsafe.Pointer(&value),
                    &size_ret)

                *param_value = CL_uint(value)

            case CL_DEVICE_IMAGE2D_MAX_HEIGHT,
                CL_DEVICE_IMAGE2D_MAX_WIDTH,
                CL_DEVICE_IMAGE3D_MAX_DEPTH,
                CL_DEVICE_IMAGE3D_MAX_HEIGHT,
                CL_DEVICE_IMAGE3D_MAX_WIDTH,
                CL_DEVICE_MAX_PARAMETER_SIZE,
                CL_DEVICE_MAX_WORK_GROUP_SIZE,
                CL_DEVICE_PROFILING_TIMER_RESOLUTION:
                //CL_DEVICE_IMAGE_MAX_BUFFER_SIZE, //2.0
                //CL_DEVICE_IMAGE_MAX_ARRAY_SIZE, //2.0
                //CL_DEVICE_PRINTF_BUFFER_SIZE: //2.0

                var value C.size_t
                ret = C.clGetDeviceInfo(device.cl_device_id,
                    C.cl_device_info(param_name),
                    C.size_t(unsafe.Sizeof(value)),
                    unsafe.Pointer(&value),
                    &size_ret)

                *param_value = CL_size_t(value)

            case CL_DEVICE_GLOBAL_MEM_CACHE_SIZE,
                CL_DEVICE_GLOBAL_MEM_SIZE,
                CL_DEVICE_LOCAL_MEM_SIZE,
                CL_DEVICE_MAX_CONSTANT_BUFFER_SIZE,
                CL_DEVICE_MAX_MEM_ALLOC_SIZE:

                var value C.cl_ulong
                ret = C.clGetDeviceInfo(device.cl_device_id,
                    C.cl_device_info(param_name),
                    C.size_t(unsafe.Sizeof(value)),
                    unsafe.Pointer(&value),
                    &size_ret)

                *param_value = CL_ulong(value)

            case CL_DEVICE_PLATFORM:

                var value C.cl_platform_id
                ret = C.clGetDeviceInfo(device.cl_device_id,
                    C.cl_device_info(param_name),
                    C.size_t(unsafe.Sizeof(value)),
                    unsafe.Pointer(&value),
                    &size_ret)

                *param_value = CL_platform_id{value}

            case CL_DEVICE_PARENT_DEVICE:
                var value C.cl_device_id

                ret = C.clGetDeviceInfo(device.cl_device_id,
                    C.cl_device_info(param_name),
                    C.size_t(unsafe.Sizeof(value)),
                    unsafe.Pointer(&value),
                    &size_ret)

                *param_value = CL_device_id{value}

            case CL_DEVICE_TYPE:
                var value C.cl_device_type
                ret = C.clGetDeviceInfo(device.cl_device_id,
                    C.cl_device_info(param_name),
                    C.size_t(unsafe.Sizeof(value)),
                    unsafe.Pointer(&value),
                    &size_ret)

                *param_value = CL_device_type(value)

            case CL_DEVICE_EXTENSIONS,
                CL_DEVICE_NAME,
                CL_DEVICE_OPENCL_C_VERSION,
                CL_DEVICE_PROFILE,
                CL_DEVICE_VENDOR,
                CL_DEVICE_VERSION,
                CL_DRIVER_VERSION:
                //CL_DEVICE_BUILT_IN_KERNELS://2.0?

                value := make([]C.char, param_value_size)
                ret = C.clGetDeviceInfo(device.cl_device_id,
                    C.cl_device_info(param_name),
                    C.size_t(param_value_size),
                    unsafe.Pointer(&value[0]),
                    &size_ret)

                *param_value = C.GoStringN(&value[0], C.int(size_ret-1))
            default:
                return CL_INVALID_VALUE
            }
        }

        if param_value_size_ret != nil {
            *param_value_size_ret = CL_size_t(size_ret)
        }

    }

    return CL_int(ret)
}

///////////////////////////////////////////////
//OpenCL 1.2 TODO
///////////////////////////////////////////////
/*
func CLCreateSubDevices(in_device CL_device_id,
	properties []CL_device_partition_property,
	num_devices CL_uint,
	out_devices []CL_device_id,
	num_devices_ret *CL_uint) CL_int {

	println("NOT SUPPORT YET!")
	return -1
}

func CLRetainDevice(device CL_device_id) CL_int {
	return CL_int(C.clRetainDevice(device.cl_device_id))
}

func CLReleaseDevice(device CL_device_id) CL_int {
	return CL_int(C.clReleaseDevice(device.cl_device_id))
}
*/
