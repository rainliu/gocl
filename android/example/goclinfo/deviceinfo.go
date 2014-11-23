package main

import (
	"fmt"
	"golang.org/x/mobile/cl"
)

func DisplayDeviceInfo(id cl.CL_device_id,
	name cl.CL_device_info,
	str string) {

	var errNum cl.CL_int
	var paramValueSize cl.CL_size_t

	errNum = cl.CLGetDeviceInfo(id,
		name,
		0,
		nil,
		&paramValueSize)

	if errNum != cl.CL_SUCCESS {
		fmt.Printf("Failed to find OpenCL device info %s.\n", str)
		return
	}

	var info interface{}
	errNum = cl.CLGetDeviceInfo(id,
		name,
		paramValueSize,
		&info,
		nil)
	if errNum != cl.CL_SUCCESS {
		fmt.Printf("Failed to find OpenCL device info %s.\n", str)
		return
	}

	// Handle a few special cases

	switch name {

	case cl.CL_DEVICE_TYPE:
		var deviceTypeStr string

		appendBitfield(cl.CL_bitfield(info.(cl.CL_device_type)),
			cl.CL_bitfield(cl.CL_DEVICE_TYPE_CPU),
			"CL_DEVICE_TYPE_CPU",
			&deviceTypeStr)

		appendBitfield(cl.CL_bitfield(info.(cl.CL_device_type)),
			cl.CL_bitfield(cl.CL_DEVICE_TYPE_GPU),
			"CL_DEVICE_TYPE_GPU",
			&deviceTypeStr)

		appendBitfield(cl.CL_bitfield(info.(cl.CL_device_type)),
			cl.CL_bitfield(cl.CL_DEVICE_TYPE_ACCELERATOR),
			"CL_DEVICE_TYPE_ACCELERATOR",
			&deviceTypeStr)

		appendBitfield(cl.CL_bitfield(info.(cl.CL_device_type)),
			cl.CL_bitfield(cl.CL_DEVICE_TYPE_DEFAULT),
			"CL_DEVICE_TYPE_DEFAULT",
			&deviceTypeStr)

		info = deviceTypeStr

		/*
		   case CL_DEVICE_SINGLE_FP_CONFIG:
		   {
		   	std::string fpType;

		   	appendBitfield<cl_device_fp_config>(
		   		*(reinterpret_cast<cl_device_fp_config*>(info)),
		   		CL_FP_DENORM,
		   		"CL_FP_DENORM",
		   		fpType);

		   	appendBitfield<cl_device_fp_config>(
		   		*(reinterpret_cast<cl_device_fp_config*>(info)),
		   		CL_FP_INF_NAN,
		   		"CL_FP_INF_NAN",
		   		fpType);

		   	appendBitfield<cl_device_fp_config>(
		   		*(reinterpret_cast<cl_device_fp_config*>(info)),
		   		CL_FP_ROUND_TO_NEAREST,
		   		"CL_FP_ROUND_TO_NEAREST",
		   		fpType);

		   	appendBitfield<cl_device_fp_config>(
		   		*(reinterpret_cast<cl_device_fp_config*>(info)),
		   		CL_FP_ROUND_TO_ZERO,
		   		"CL_FP_ROUND_TO_ZERO",
		   		fpType);

		   	appendBitfield<cl_device_fp_config>(
		   		*(reinterpret_cast<cl_device_fp_config*>(info)),
		   		CL_FP_ROUND_TO_INF,
		   		"CL_FP_ROUND_TO_INF",
		   		fpType);

		   	appendBitfield<cl_device_fp_config>(
		   		*(reinterpret_cast<cl_device_fp_config*>(info)),
		   		CL_FP_FMA,
		   		"CL_FP_FMA",
		   		fpType);

		   #ifdef CL_FP_SOFT_FLOAT
		   	appendBitfield<cl_device_fp_config>(
		   		*(reinterpret_cast<cl_device_fp_config*>(info)),
		   		CL_FP_SOFT_FLOAT,
		   		"CL_FP_SOFT_FLOAT",
		   		fpType);
		   #endif

		   	std::cout << "\t\t" << str << ":\t" << fpType << std::endl;
		   }
		   case CL_DEVICE_GLOBAL_MEM_CACHE_TYPE:
		   {
		   	std::string memType;

		   	appendBitfield<cl_device_mem_cache_type>(
		   		*(reinterpret_cast<cl_device_mem_cache_type*>(info)),
		   		CL_NONE,
		   		"CL_NONE",
		   		memType);
		   	appendBitfield<cl_device_mem_cache_type>(
		   		*(reinterpret_cast<cl_device_mem_cache_type*>(info)),
		   		CL_READ_ONLY_CACHE,
		   		"CL_READ_ONLY_CACHE",
		   		memType);

		   	appendBitfield<cl_device_mem_cache_type>(
		   		*(reinterpret_cast<cl_device_mem_cache_type*>(info)),
		   		CL_READ_WRITE_CACHE,
		   		"CL_READ_WRITE_CACHE",
		   		memType);

		   	std::cout << "\t\t" << str << ":\t" << memType << std::endl;
		   }
		   break;
		   case CL_DEVICE_LOCAL_MEM_TYPE:
		   {
		   	std::string memType;

		   	appendBitfield<cl_device_local_mem_type>(
		   		*(reinterpret_cast<cl_device_local_mem_type*>(info)),
		   		CL_GLOBAL,
		   		"CL_LOCAL",
		   		memType);

		   	appendBitfield<cl_device_local_mem_type>(
		   		*(reinterpret_cast<cl_device_local_mem_type*>(info)),
		   		CL_GLOBAL,
		   		"CL_GLOBAL",
		   		memType);

		   	std::cout << "\t\t" << str << ":\t" << memType << std::endl;
		   }
		   break;
		   case CL_DEVICE_EXECUTION_CAPABILITIES:
		   {
		   	std::string memType;

		   	appendBitfield<cl_device_exec_capabilities>(
		   		*(reinterpret_cast<cl_device_exec_capabilities*>(info)),
		   		CL_EXEC_KERNEL,
		   		"CL_EXEC_KERNEL",
		   		memType);

		   	appendBitfield<cl_device_exec_capabilities>(
		   		*(reinterpret_cast<cl_device_exec_capabilities*>(info)),
		   		CL_EXEC_NATIVE_KERNEL,
		   		"CL_EXEC_NATIVE_KERNEL",
		   		memType);

		   	std::cout << "\t\t" << str << ":\t" << memType << std::endl;
		   }
		   break;
		   case CL_DEVICE_QUEUE_PROPERTIES:
		   {
		   	std::string memType;

		   	appendBitfield<cl_device_exec_capabilities>(
		   		*(reinterpret_cast<cl_device_exec_capabilities*>(info)),
		   		CL_QUEUE_OUT_OF_ORDER_EXEC_MODE_ENABLE,
		   		"CL_QUEUE_OUT_OF_ORDER_EXEC_MODE_ENABLE",
		   		memType);

		   	appendBitfield<cl_device_exec_capabilities>(
		   		*(reinterpret_cast<cl_device_exec_capabilities*>(info)),
		   		CL_QUEUE_PROFILING_ENABLE,
		   		"CL_QUEUE_PROFILING_ENABLE",
		   		memType);

		   	std::cout << "\t\t" << str << ":\t" << memType << std::endl;
		   }
		   break;
		*/
	default:
	}

	fmt.Printf("\t\t%-20s: %v\n", str, info)
}

func appendBitfield(info, value cl.CL_bitfield, name string, str *string) {
	if (info & value) != 0 {
		if len(*str) > 0 {
			*str += " | "
		}
		*str += name
	}
}
