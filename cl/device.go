// +build cl11 cl12 cl20

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
import "unsafe"

func CLGetDeviceIDs(platform CL_platform_id,
	device_type CL_device_type,
	num_entries CL_uint,
	devices []CL_device_id,
	num_devices *CL_uint) CL_int {

	if (num_entries == 0 && devices != nil) || (num_devices == nil && devices == nil) {
		return CL_INVALID_VALUE
	} else {
		var c_num_devices C.cl_uint
		var c_errcode_ret C.cl_int

		if devices == nil {
			c_errcode_ret = C.clGetDeviceIDs(platform.cl_platform_id,
				C.cl_device_type(device_type),
				C.cl_uint(num_entries),
				nil,
				&c_num_devices)
		} else {
			devices_id := make([]C.cl_device_id, len(devices))
			c_errcode_ret = C.clGetDeviceIDs(platform.cl_platform_id,
				C.cl_device_type(device_type),
				C.cl_uint(num_entries),
				&devices_id[0],
				&c_num_devices)
			if c_errcode_ret == C.CL_SUCCESS {
				for i := 0; i < len(devices); i++ {
					devices[i].cl_device_id = devices_id[i]
				}
			}
		}

		if num_devices != nil {
			*num_devices = CL_uint(c_num_devices)
		}

		return CL_int(c_errcode_ret)
	}
}

func CLGetDeviceInfo(device CL_device_id,
	param_name CL_device_info,
	param_value_size CL_size_t,
	param_value *interface{},
	param_value_size_ret *CL_size_t) CL_int {

	if (param_value_size == 0 || param_value == nil) && param_value_size_ret == nil {
		return CL_INVALID_VALUE
	} else {
		var c_param_value_size_ret C.size_t
		var c_errcode_ret C.cl_int

		if param_value_size == 0 || param_value == nil {
			c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
				C.cl_device_info(param_name),
				C.size_t(param_value_size),
				nil,
				&c_param_value_size_ret)
		} else {
			switch param_name {

			case CL_DEVICE_AVAILABLE,
				CL_DEVICE_COMPILER_AVAILABLE,
				CL_DEVICE_ENDIAN_LITTLE,
				CL_DEVICE_ERROR_CORRECTION_SUPPORT,
				CL_DEVICE_HOST_UNIFIED_MEMORY,
				CL_DEVICE_IMAGE_SUPPORT,
				CL_DEVICE_LINKER_AVAILABLE,
				CL_DEVICE_PREFERRED_INTEROP_USER_SYNC:

				var value C.cl_bool
				c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
					C.cl_device_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)
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
				CL_DEVICE_VENDOR_ID,
				CL_DEVICE_GLOBAL_MEM_CACHELINE_SIZE,
				CL_DEVICE_PARTITION_MAX_SUB_DEVICES,
				CL_DEVICE_REFERENCE_COUNT,
				CL_DEVICE_MAX_READ_WRITE_IMAGE_ARGS,
				CL_DEVICE_QUEUE_ON_DEVICE_PREFERRED_SIZE,
				CL_DEVICE_QUEUE_ON_DEVICE_MAX_SIZE,
				CL_DEVICE_MAX_ON_DEVICE_QUEUES,
				CL_DEVICE_MAX_ON_DEVICE_EVENTS,
				CL_DEVICE_MAX_PIPE_ARGS,
				CL_DEVICE_PIPE_MAX_ACTIVE_RESERVATIONS,
				CL_DEVICE_PIPE_MAX_PACKET_SIZE,
				CL_DEVICE_PREFERRED_PLATFORM_ATOMIC_ALIGNMENT,
				CL_DEVICE_PREFERRED_GLOBAL_ATOMIC_ALIGNMENT,
				CL_DEVICE_PREFERRED_LOCAL_ATOMIC_ALIGNMENT:

				var value C.cl_uint
				c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
					C.cl_device_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_uint(value)

			case CL_DEVICE_IMAGE2D_MAX_HEIGHT,
				CL_DEVICE_IMAGE2D_MAX_WIDTH,
				CL_DEVICE_IMAGE3D_MAX_DEPTH,
				CL_DEVICE_IMAGE3D_MAX_HEIGHT,
				CL_DEVICE_IMAGE3D_MAX_WIDTH,
				CL_DEVICE_MAX_PARAMETER_SIZE,
				CL_DEVICE_MAX_WORK_GROUP_SIZE,
				CL_DEVICE_PROFILING_TIMER_RESOLUTION,
				CL_DEVICE_IMAGE_MAX_BUFFER_SIZE,
				CL_DEVICE_IMAGE_MAX_ARRAY_SIZE,
				CL_DEVICE_PRINTF_BUFFER_SIZE,
				CL_DEVICE_MAX_GLOBAL_VARIABLE_SIZE,
				CL_DEVICE_GLOBAL_VARIABLE_PREFERRED_TOTAL_SIZE:

				var value C.size_t
				c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
					C.cl_device_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_size_t(value)

			case CL_DEVICE_GLOBAL_MEM_CACHE_SIZE,
				CL_DEVICE_GLOBAL_MEM_SIZE,
				CL_DEVICE_LOCAL_MEM_SIZE,
				CL_DEVICE_MAX_CONSTANT_BUFFER_SIZE,
				CL_DEVICE_MAX_MEM_ALLOC_SIZE:

				var value C.cl_ulong
				c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
					C.cl_device_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_ulong(value)

			case CL_DEVICE_PLATFORM:

				var value C.cl_platform_id
				c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
					C.cl_device_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_platform_id{value}

			case CL_DEVICE_PARENT_DEVICE:
				var value C.cl_device_id

				c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
					C.cl_device_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_device_id{value}

			case CL_DEVICE_TYPE:
				var value C.cl_device_type
				c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
					C.cl_device_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_device_type(value)

			case CL_DEVICE_EXTENSIONS,
				CL_DEVICE_NAME,
				CL_DEVICE_OPENCL_C_VERSION,
				CL_DEVICE_PROFILE,
				CL_DEVICE_VENDOR,
				CL_DEVICE_VERSION,
				CL_DRIVER_VERSION,
				CL_DEVICE_BUILT_IN_KERNELS:

				value := make([]C.char, param_value_size)
				c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
					C.cl_device_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value[0]),
					&c_param_value_size_ret)

				*param_value = C.GoStringN(&value[0], C.int(c_param_value_size_ret-1))

			case CL_DEVICE_SINGLE_FP_CONFIG,
				CL_DEVICE_DOUBLE_FP_CONFIG:
				var value C.cl_device_fp_config
				c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
					C.cl_device_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_device_fp_config(value)

			case CL_DEVICE_GLOBAL_MEM_CACHE_TYPE:
				var value C.cl_device_mem_cache_type
				c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
					C.cl_device_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_device_mem_cache_type(value)

			case CL_DEVICE_LOCAL_MEM_TYPE:
				var value C.cl_device_local_mem_type
				c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
					C.cl_device_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_device_local_mem_type(value)

			case CL_DEVICE_EXECUTION_CAPABILITIES:
				var value C.cl_device_exec_capabilities
				c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
					C.cl_device_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_device_exec_capabilities(value)

			//case CL_DEVICE_QUEUE_PROPERTIES,//deprecated
			case CL_DEVICE_QUEUE_ON_HOST_PROPERTIES,
				CL_DEVICE_QUEUE_ON_DEVICE_PROPERTIES:
				var value C.cl_command_queue_properties
				c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
					C.cl_device_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_command_queue_properties(value)

			case CL_DEVICE_PARTITION_PROPERTIES,
				CL_DEVICE_PARTITION_TYPE:
				var param C.cl_device_partition_property
				length := int(C.size_t(param_value_size) / C.size_t(unsafe.Sizeof(param)))

				value1 := make([]C.cl_device_partition_property, length)
				value2 := make([]CL_device_partition_property, length)

				c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
					C.cl_device_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value1[0]),
					&c_param_value_size_ret)

				for i := 0; i < length; i++ {
					value2[i] = CL_device_partition_property(value1[i])
				}

				*param_value = value2

			case CL_DEVICE_MAX_WORK_ITEM_SIZES:
				var param C.size_t
				length := int(C.size_t(param_value_size) / C.size_t(unsafe.Sizeof(param)))

				value1 := make([]C.size_t, length)
				value2 := make([]CL_size_t, length)

				c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
					C.cl_device_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value1[0]),
					&c_param_value_size_ret)

				for i := 0; i < length; i++ {
					value2[i] = CL_size_t(value1[i])
				}

				*param_value = value2

			case CL_DEVICE_PARTITION_AFFINITY_DOMAIN:
				var value C.cl_device_affinity_domain
				c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
					C.cl_device_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_device_affinity_domain(value)

			case CL_DEVICE_SVM_CAPABILITIES:
				var value C.cl_bitfield //C.cl_device_svm_capabilities //use cl_bitfield to make darwin pass
				c_errcode_ret = C.clGetDeviceInfo(device.cl_device_id,
					C.cl_device_info(param_name),
					C.size_t(param_value_size),
					unsafe.Pointer(&value),
					&c_param_value_size_ret)

				*param_value = CL_device_svm_capabilities(value)

			default:
				return CL_INVALID_VALUE
			}
		}

		if param_value_size_ret != nil {
			*param_value_size_ret = CL_size_t(c_param_value_size_ret)
		}

		return CL_int(c_errcode_ret)
	}
}
