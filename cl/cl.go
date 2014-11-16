// +build cl11 cl12

package cl

/*
#cgo CFLAGS: -I CL
#cgo linux LDFLAGS: -lOpenCL
#cgo darwin LDFLAGS: -framework OpenCL

#include "CL/opencl.h"
*/
import "C"

import (
	"math"
	//"unsafe"
)

/* Primitive Data Type*/
type CL_char int8
type CL_uchar uint8
type CL_short int16
type CL_ushort uint16
type CL_int int32
type CL_uint uint32
type CL_long int64
type CL_ulong uint64
type CL_half uint16
type CL_float float32
type CL_double float64
type CL_size_t int
type CL_intptr_t uintptr

const (
	CL_CHAR_BIT  = 8
	CL_SCHAR_MAX = 127
	CL_SCHAR_MIN = (-127 - 1)
	CL_CHAR_MAX  = CL_SCHAR_MAX
	CL_CHAR_MIN  = CL_SCHAR_MIN
	CL_UCHAR_MAX = 255
	CL_SHRT_MAX  = 32767
	CL_SHRT_MIN  = (-32767 - 1)
	CL_USHRT_MAX = 65535
	CL_INT_MAX   = 2147483647
	CL_INT_MIN   = (-2147483647 - 1)
	CL_UINT_MAX  = math.MaxUint32
	CL_LONG_MAX  = math.MaxInt64
	CL_LONG_MIN  = math.MinInt64
	CL_ULONG_MAX = math.MaxUint64

	CL_FLT_DIG        = 6
	CL_FLT_MANT_DIG   = 24
	CL_FLT_MAX_10_EXP = +38
	CL_FLT_MAX_EXP    = +128
	CL_FLT_MIN_10_EXP = -37
	CL_FLT_MIN_EXP    = -125
	CL_FLT_RADIX      = 2
	CL_FLT_MAX        = math.MaxFloat32
	CL_FLT_MIN        = math.SmallestNonzeroFloat32
	CL_FLT_EPSILON    = CL_float(math.E)

	CL_DBL_DIG        = 15
	CL_DBL_MANT_DIG   = 53
	CL_DBL_MAX_10_EXP = +308
	CL_DBL_MAX_EXP    = +1024
	CL_DBL_MIN_10_EXP = -307
	CL_DBL_MIN_EXP    = -1021
	CL_DBL_RADIX      = 2
	CL_DBL_MAX        = math.MaxFloat64
	CL_DBL_MIN        = math.SmallestNonzeroFloat64
	CL_DBL_EPSILON    = CL_double(math.E)

	CL_M_E        = 2.718281828459045090796
	CL_M_LOG2E    = 1.442695040888963387005
	CL_M_LOG10E   = 0.434294481903251816668
	CL_M_LN2      = 0.693147180559945286227
	CL_M_LN10     = 2.302585092994045901094
	CL_M_PI       = 3.141592653589793115998
	CL_M_PI_2     = 1.570796326794896557999
	CL_M_PI_4     = 0.785398163397448278999
	CL_M_1_PI     = 0.318309886183790691216
	CL_M_2_PI     = 0.636619772367581382433
	CL_M_2_SQRTPI = 1.128379167095512558561
	CL_M_SQRT2    = 1.414213562373095145475
	CL_M_SQRT1_2  = 0.707106781186547572737

	CL_M_E_F        = 2.71828174591064
	CL_M_LOG2E_F    = 1.44269502162933
	CL_M_LOG10E_F   = 0.43429449200630
	CL_M_LN2_F      = 0.69314718246460
	CL_M_LN10_F     = 2.30258512496948
	CL_M_PI_F       = 3.14159274101257
	CL_M_PI_2_F     = 1.57079637050629
	CL_M_PI_4_F     = 0.78539818525314
	CL_M_1_PI_F     = 0.31830987334251
	CL_M_2_PI_F     = 0.63661974668503
	CL_M_2_SQRTPI_F = 1.12837922573090
	CL_M_SQRT2_F    = 1.41421353816986
	CL_M_SQRT1_2_F  = 0.70710676908493

	CL_HUGE_VALF = 1e50
	CL_HUGE_VAL  = 1e500
	CL_MAXFLOAT  = CL_FLT_MAX
	CL_INFINITY  = CL_HUGE_VALF
	CL_NAN       = (CL_INFINITY - CL_INFINITY)
)

////////////////////////////////////////////////////////////////
/* Structure and Type*/
type CL_platform_id struct {
	cl_platform_id C.cl_platform_id
}
type CL_device_id struct {
	cl_device_id C.cl_device_id
}
type CL_context struct {
	cl_context C.cl_context
}
type CL_command_queue struct {
	cl_command_queue C.cl_command_queue
}
type CL_mem struct {
	cl_mem C.cl_mem
}
type CL_program struct {
	cl_program C.cl_program
}
type CL_kernel struct {
	cl_kernel C.cl_kernel
}
type CL_event struct {
	cl_event C.cl_event
}
type CL_sampler struct {
	cl_sampler C.cl_sampler
}

type CL_bool CL_uint /* WARNING!  Unlike cl_ types in cl_platform.h, cl_bool is not guaranteed to be the same size as the bool in kernels. */
type CL_bitfield CL_ulong
type CL_device_type CL_bitfield
type CL_platform_info CL_uint
type CL_device_info CL_uint
type CL_device_fp_config CL_bitfield
type CL_device_mem_cache_type CL_uint
type CL_device_local_mem_type CL_uint
type CL_device_exec_capabilities CL_bitfield
type CL_command_queue_properties CL_bitfield
type CL_device_partition_property CL_intptr_t
type CL_device_affinity_domain CL_bitfield

type CL_context_properties CL_intptr_t
type CL_context_info CL_uint
type CL_command_queue_info CL_uint
type CL_channel_order CL_uint
type CL_channel_type CL_uint
type CL_mem_flags CL_bitfield
type CL_mem_object_type CL_uint
type CL_mem_info CL_uint
type CL_mem_migration_flags CL_bitfield
type CL_image_info CL_uint
type CL_buffer_create_type CL_uint
type CL_addressing_mode CL_uint
type CL_filter_mode CL_uint
type CL_sampler_info CL_uint
type CL_map_flags CL_bitfield
type CL_program_info CL_uint
type CL_program_build_info CL_uint
type CL_program_binary_type CL_uint
type CL_build_status CL_int
type CL_kernel_info CL_uint
type CL_kernel_arg_info CL_uint
type CL_kernel_arg_address_qualifier CL_uint
type CL_kernel_arg_access_qualifier CL_uint
type CL_kernel_arg_type_qualifier CL_bitfield
type CL_kernel_work_group_info CL_uint
type CL_event_info CL_uint
type CL_command_type CL_uint
type CL_profiling_info CL_uint

type CL_image_format struct {
	Image_channel_order     CL_channel_order
	Image_channel_data_type CL_channel_type
}

type CL_image_desc struct {
	Image_type        CL_mem_object_type
	Image_width       CL_size_t
	Image_height      CL_size_t
	Image_depth       CL_size_t
	Image_array_size  CL_size_t
	Image_row_pitch   CL_size_t
	Image_slice_pitch CL_size_t
	Num_mip_levels    CL_uint
	Num_samples       CL_uint
	Buffer            CL_mem
}

type CL_buffer_region struct {
	Origin CL_size_t
	Size   CL_size_t
}

const (
	/* Error Codes */
	CL_SUCCESS                                   = C.CL_SUCCESS
	CL_DEVICE_NOT_FOUND                          = C.CL_DEVICE_NOT_FOUND
	CL_DEVICE_NOT_AVAILABLE                      = C.CL_DEVICE_NOT_AVAILABLE
	CL_COMPILER_NOT_AVAILABLE                    = C.CL_COMPILER_NOT_AVAILABLE
	CL_MEM_OBJECT_ALLOCATION_FAILURE             = C.CL_MEM_OBJECT_ALLOCATION_FAILURE
	CL_OUT_OF_RESOURCES                          = C.CL_OUT_OF_RESOURCES
	CL_OUT_OF_HOST_MEMORY                        = C.CL_OUT_OF_HOST_MEMORY
	CL_PROFILING_INFO_NOT_AVAILABLE              = C.CL_PROFILING_INFO_NOT_AVAILABLE
	CL_MEM_COPY_OVERLAP                          = C.CL_MEM_COPY_OVERLAP
	CL_IMAGE_FORMAT_MISMATCH                     = C.CL_IMAGE_FORMAT_MISMATCH
	CL_IMAGE_FORMAT_NOT_SUPPORTED                = C.CL_IMAGE_FORMAT_NOT_SUPPORTED
	CL_BUILD_PROGRAM_FAILURE                     = C.CL_BUILD_PROGRAM_FAILURE
	CL_MAP_FAILURE                               = C.CL_MAP_FAILURE
	CL_MISALIGNED_SUB_BUFFER_OFFSET              = C.CL_MISALIGNED_SUB_BUFFER_OFFSET
	CL_EXEC_STATUS_ERROR_FOR_EVENTS_IN_WAIT_LIST = C.CL_EXEC_STATUS_ERROR_FOR_EVENTS_IN_WAIT_LIST
	CL_COMPILE_PROGRAM_FAILURE                   = C.CL_COMPILE_PROGRAM_FAILURE
	CL_LINKER_NOT_AVAILABLE                      = C.CL_LINKER_NOT_AVAILABLE
	CL_LINK_PROGRAM_FAILURE                      = C.CL_LINK_PROGRAM_FAILURE
	CL_DEVICE_PARTITION_FAILED                   = C.CL_DEVICE_PARTITION_FAILED
	CL_KERNEL_ARG_INFO_NOT_AVAILABLE             = C.CL_KERNEL_ARG_INFO_NOT_AVAILABLE

	CL_INVALID_VALUE                   = C.CL_INVALID_VALUE
	CL_INVALID_DEVICE_TYPE             = C.CL_INVALID_DEVICE_TYPE
	CL_INVALID_PLATFORM                = C.CL_INVALID_PLATFORM
	CL_INVALID_DEVICE                  = C.CL_INVALID_DEVICE
	CL_INVALID_CONTEXT                 = C.CL_INVALID_CONTEXT
	CL_INVALID_QUEUE_PROPERTIES        = C.CL_INVALID_QUEUE_PROPERTIES
	CL_INVALID_COMMAND_QUEUE           = C.CL_INVALID_COMMAND_QUEUE
	CL_INVALID_HOST_PTR                = C.CL_INVALID_HOST_PTR
	CL_INVALID_MEM_OBJECT              = C.CL_INVALID_MEM_OBJECT
	CL_INVALID_IMAGE_FORMAT_DESCRIPTOR = C.CL_INVALID_IMAGE_FORMAT_DESCRIPTOR
	CL_INVALID_IMAGE_SIZE              = C.CL_INVALID_IMAGE_SIZE
	CL_INVALID_SAMPLER                 = C.CL_INVALID_SAMPLER
	CL_INVALID_BINARY                  = C.CL_INVALID_BINARY
	CL_INVALID_BUILD_OPTIONS           = C.CL_INVALID_BUILD_OPTIONS
	CL_INVALID_PROGRAM                 = C.CL_INVALID_PROGRAM
	CL_INVALID_PROGRAM_EXECUTABLE      = C.CL_INVALID_PROGRAM_EXECUTABLE
	CL_INVALID_KERNEL_NAME             = C.CL_INVALID_KERNEL_NAME
	CL_INVALID_KERNEL_DEFINITION       = C.CL_INVALID_KERNEL_DEFINITION
	CL_INVALID_KERNEL                  = C.CL_INVALID_KERNEL
	CL_INVALID_ARG_INDEX               = C.CL_INVALID_ARG_INDEX
	CL_INVALID_ARG_VALUE               = C.CL_INVALID_ARG_VALUE
	CL_INVALID_ARG_SIZE                = C.CL_INVALID_ARG_SIZE
	CL_INVALID_KERNEL_ARGS             = C.CL_INVALID_KERNEL_ARGS
	CL_INVALID_WORK_DIMENSION          = C.CL_INVALID_WORK_DIMENSION
	CL_INVALID_WORK_GROUP_SIZE         = C.CL_INVALID_WORK_GROUP_SIZE
	CL_INVALID_WORK_ITEM_SIZE          = C.CL_INVALID_WORK_ITEM_SIZE
	CL_INVALID_GLOBAL_OFFSET           = C.CL_INVALID_GLOBAL_OFFSET
	CL_INVALID_EVENT_WAIT_LIST         = C.CL_INVALID_EVENT_WAIT_LIST
	CL_INVALID_EVENT                   = C.CL_INVALID_EVENT
	CL_INVALID_OPERATION               = C.CL_INVALID_OPERATION
	CL_INVALID_GL_OBJECT               = C.CL_INVALID_GL_OBJECT
	CL_INVALID_BUFFER_SIZE             = C.CL_INVALID_BUFFER_SIZE
	CL_INVALID_MIP_LEVEL               = C.CL_INVALID_MIP_LEVEL
	CL_INVALID_GLOBAL_WORK_SIZE        = C.CL_INVALID_GLOBAL_WORK_SIZE
	CL_INVALID_PROPERTY                = C.CL_INVALID_PROPERTY
	CL_INVALID_IMAGE_DESCRIPTOR        = C.CL_INVALID_IMAGE_DESCRIPTOR
	CL_INVALID_COMPILER_OPTIONS        = C.CL_INVALID_COMPILER_OPTIONS
	CL_INVALID_LINKER_OPTIONS          = C.CL_INVALID_LINKER_OPTIONS
	CL_INVALID_DEVICE_PARTITION_COUNT  = C.CL_INVALID_DEVICE_PARTITION_COUNT

	/* OpenCL Version */
	CL_VERSION_1_0 = C.CL_VERSION_1_0
	CL_VERSION_1_1 = C.CL_VERSION_1_1
	CL_VERSION_1_2 = C.CL_VERSION_1_2

	/* cl_bool */
	CL_FALSE        CL_bool = C.CL_FALSE
	CL_TRUE         CL_bool = C.CL_TRUE
	CL_BLOCKING     CL_bool = C.CL_BLOCKING
	CL_NON_BLOCKING CL_bool = C.CL_NON_BLOCKING

	/* cl_platform_info */
	CL_PLATFORM_PROFILE    CL_platform_info = C.CL_PLATFORM_PROFILE
	CL_PLATFORM_VERSION    CL_platform_info = C.CL_PLATFORM_VERSION
	CL_PLATFORM_NAME       CL_platform_info = C.CL_PLATFORM_NAME
	CL_PLATFORM_VENDOR     CL_platform_info = C.CL_PLATFORM_VENDOR
	CL_PLATFORM_EXTENSIONS CL_platform_info = C.CL_PLATFORM_EXTENSIONS

	/* cl_device_type - bitfield */
	CL_DEVICE_TYPE_DEFAULT     CL_device_type = C.CL_DEVICE_TYPE_DEFAULT
	CL_DEVICE_TYPE_CPU         CL_device_type = C.CL_DEVICE_TYPE_CPU
	CL_DEVICE_TYPE_GPU         CL_device_type = C.CL_DEVICE_TYPE_GPU
	CL_DEVICE_TYPE_ACCELERATOR CL_device_type = C.CL_DEVICE_TYPE_ACCELERATOR
	CL_DEVICE_TYPE_CUSTOM      CL_device_type = C.CL_DEVICE_TYPE_CUSTOM
	CL_DEVICE_TYPE_ALL         CL_device_type = C.CL_DEVICE_TYPE_ALL

	/* cl_device_info */
	CL_DEVICE_TYPE                          CL_device_info = C.CL_DEVICE_TYPE
	CL_DEVICE_VENDOR_ID                     CL_device_info = C.CL_DEVICE_VENDOR_ID
	CL_DEVICE_MAX_COMPUTE_UNITS             CL_device_info = C.CL_DEVICE_MAX_COMPUTE_UNITS
	CL_DEVICE_MAX_WORK_ITEM_DIMENSIONS      CL_device_info = C.CL_DEVICE_MAX_WORK_ITEM_DIMENSIONS
	CL_DEVICE_MAX_WORK_GROUP_SIZE           CL_device_info = C.CL_DEVICE_MAX_WORK_GROUP_SIZE
	CL_DEVICE_MAX_WORK_ITEM_SIZES           CL_device_info = C.CL_DEVICE_MAX_WORK_ITEM_SIZES
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_CHAR   CL_device_info = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_CHAR
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_SHORT  CL_device_info = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_SHORT
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_INT    CL_device_info = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_INT
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_LONG   CL_device_info = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_LONG
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT  CL_device_info = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE CL_device_info = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE
	CL_DEVICE_MAX_CLOCK_FREQUENCY           CL_device_info = C.CL_DEVICE_MAX_CLOCK_FREQUENCY
	CL_DEVICE_ADDRESS_BITS                  CL_device_info = C.CL_DEVICE_ADDRESS_BITS
	CL_DEVICE_MAX_READ_IMAGE_ARGS           CL_device_info = C.CL_DEVICE_MAX_READ_IMAGE_ARGS
	CL_DEVICE_MAX_WRITE_IMAGE_ARGS          CL_device_info = C.CL_DEVICE_MAX_WRITE_IMAGE_ARGS
	CL_DEVICE_MAX_MEM_ALLOC_SIZE            CL_device_info = C.CL_DEVICE_MAX_MEM_ALLOC_SIZE
	CL_DEVICE_IMAGE2D_MAX_WIDTH             CL_device_info = C.CL_DEVICE_IMAGE2D_MAX_WIDTH
	CL_DEVICE_IMAGE2D_MAX_HEIGHT            CL_device_info = C.CL_DEVICE_IMAGE2D_MAX_HEIGHT
	CL_DEVICE_IMAGE3D_MAX_WIDTH             CL_device_info = C.CL_DEVICE_IMAGE3D_MAX_WIDTH
	CL_DEVICE_IMAGE3D_MAX_HEIGHT            CL_device_info = C.CL_DEVICE_IMAGE3D_MAX_HEIGHT
	CL_DEVICE_IMAGE3D_MAX_DEPTH             CL_device_info = C.CL_DEVICE_IMAGE3D_MAX_DEPTH
	CL_DEVICE_IMAGE_SUPPORT                 CL_device_info = C.CL_DEVICE_IMAGE_SUPPORT
	CL_DEVICE_MAX_PARAMETER_SIZE            CL_device_info = C.CL_DEVICE_MAX_PARAMETER_SIZE
	CL_DEVICE_MAX_SAMPLERS                  CL_device_info = C.CL_DEVICE_MAX_SAMPLERS
	CL_DEVICE_MEM_BASE_ADDR_ALIGN           CL_device_info = C.CL_DEVICE_MEM_BASE_ADDR_ALIGN
	CL_DEVICE_MIN_DATA_TYPE_ALIGN_SIZE      CL_device_info = C.CL_DEVICE_MIN_DATA_TYPE_ALIGN_SIZE
	CL_DEVICE_SINGLE_FP_CONFIG              CL_device_info = C.CL_DEVICE_SINGLE_FP_CONFIG
	CL_DEVICE_GLOBAL_MEM_CACHE_TYPE         CL_device_info = C.CL_DEVICE_GLOBAL_MEM_CACHE_TYPE
	CL_DEVICE_GLOBAL_MEM_CACHELINE_SIZE     CL_device_info = C.CL_DEVICE_GLOBAL_MEM_CACHELINE_SIZE
	CL_DEVICE_GLOBAL_MEM_CACHE_SIZE         CL_device_info = C.CL_DEVICE_GLOBAL_MEM_CACHE_SIZE
	CL_DEVICE_GLOBAL_MEM_SIZE               CL_device_info = C.CL_DEVICE_GLOBAL_MEM_SIZE
	CL_DEVICE_MAX_CONSTANT_BUFFER_SIZE      CL_device_info = C.CL_DEVICE_MAX_CONSTANT_BUFFER_SIZE
	CL_DEVICE_MAX_CONSTANT_ARGS             CL_device_info = C.CL_DEVICE_MAX_CONSTANT_ARGS
	CL_DEVICE_LOCAL_MEM_TYPE                CL_device_info = C.CL_DEVICE_LOCAL_MEM_TYPE
	CL_DEVICE_LOCAL_MEM_SIZE                CL_device_info = C.CL_DEVICE_LOCAL_MEM_SIZE
	CL_DEVICE_ERROR_CORRECTION_SUPPORT      CL_device_info = C.CL_DEVICE_ERROR_CORRECTION_SUPPORT
	CL_DEVICE_PROFILING_TIMER_RESOLUTION    CL_device_info = C.CL_DEVICE_PROFILING_TIMER_RESOLUTION
	CL_DEVICE_ENDIAN_LITTLE                 CL_device_info = C.CL_DEVICE_ENDIAN_LITTLE
	CL_DEVICE_AVAILABLE                     CL_device_info = C.CL_DEVICE_AVAILABLE
	CL_DEVICE_COMPILER_AVAILABLE            CL_device_info = C.CL_DEVICE_COMPILER_AVAILABLE
	CL_DEVICE_EXECUTION_CAPABILITIES        CL_device_info = C.CL_DEVICE_EXECUTION_CAPABILITIES
	CL_DEVICE_QUEUE_PROPERTIES              CL_device_info = C.CL_DEVICE_QUEUE_PROPERTIES
	CL_DEVICE_NAME                          CL_device_info = C.CL_DEVICE_NAME
	CL_DEVICE_VENDOR                        CL_device_info = C.CL_DEVICE_VENDOR
	CL_DRIVER_VERSION                       CL_device_info = C.CL_DRIVER_VERSION
	CL_DEVICE_PROFILE                       CL_device_info = C.CL_DEVICE_PROFILE
	CL_DEVICE_VERSION                       CL_device_info = C.CL_DEVICE_VERSION
	CL_DEVICE_EXTENSIONS                    CL_device_info = C.CL_DEVICE_EXTENSIONS
	CL_DEVICE_PLATFORM                      CL_device_info = C.CL_DEVICE_PLATFORM
	CL_DEVICE_DOUBLE_FP_CONFIG              CL_device_info = C.CL_DEVICE_DOUBLE_FP_CONFIG

	/* 0x1033 reserved for CL_DEVICE_HALF_FP_CONFIG */
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_HALF  CL_device_info = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_HALF
	CL_DEVICE_HOST_UNIFIED_MEMORY          CL_device_info = C.CL_DEVICE_HOST_UNIFIED_MEMORY
	CL_DEVICE_NATIVE_VECTOR_WIDTH_CHAR     CL_device_info = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_CHAR
	CL_DEVICE_NATIVE_VECTOR_WIDTH_SHORT    CL_device_info = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_SHORT
	CL_DEVICE_NATIVE_VECTOR_WIDTH_INT      CL_device_info = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_INT
	CL_DEVICE_NATIVE_VECTOR_WIDTH_LONG     CL_device_info = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_LONG
	CL_DEVICE_NATIVE_VECTOR_WIDTH_FLOAT    CL_device_info = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_FLOAT
	CL_DEVICE_NATIVE_VECTOR_WIDTH_DOUBLE   CL_device_info = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_DOUBLE
	CL_DEVICE_NATIVE_VECTOR_WIDTH_HALF     CL_device_info = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_HALF
	CL_DEVICE_OPENCL_C_VERSION             CL_device_info = C.CL_DEVICE_OPENCL_C_VERSION
	CL_DEVICE_LINKER_AVAILABLE             CL_device_info = C.CL_DEVICE_LINKER_AVAILABLE
	CL_DEVICE_BUILT_IN_KERNELS             CL_device_info = C.CL_DEVICE_BUILT_IN_KERNELS
	CL_DEVICE_IMAGE_MAX_BUFFER_SIZE        CL_device_info = C.CL_DEVICE_IMAGE_MAX_BUFFER_SIZE
	CL_DEVICE_IMAGE_MAX_ARRAY_SIZE         CL_device_info = C.CL_DEVICE_IMAGE_MAX_ARRAY_SIZE
	CL_DEVICE_PARENT_DEVICE                CL_device_info = C.CL_DEVICE_PARENT_DEVICE
	CL_DEVICE_PARTITION_MAX_SUB_DEVICES    CL_device_info = C.CL_DEVICE_PARTITION_MAX_SUB_DEVICES
	CL_DEVICE_PARTITION_PROPERTIES         CL_device_info = C.CL_DEVICE_PARTITION_PROPERTIES
	CL_DEVICE_PARTITION_AFFINITY_DOMAIN    CL_device_info = C.CL_DEVICE_PARTITION_AFFINITY_DOMAIN
	CL_DEVICE_PARTITION_TYPE               CL_device_info = C.CL_DEVICE_PARTITION_TYPE
	CL_DEVICE_REFERENCE_COUNT              CL_device_info = C.CL_DEVICE_REFERENCE_COUNT
	CL_DEVICE_PREFERRED_INTEROP_USER_SYNC  CL_device_info = C.CL_DEVICE_PREFERRED_INTEROP_USER_SYNC
	CL_DEVICE_PRINTF_BUFFER_SIZE           CL_device_info = C.CL_DEVICE_PRINTF_BUFFER_SIZE
	CL_DEVICE_IMAGE_PITCH_ALIGNMENT        CL_device_info = C.CL_DEVICE_IMAGE_PITCH_ALIGNMENT
	CL_DEVICE_IMAGE_BASE_ADDRESS_ALIGNMENT CL_device_info = C.CL_DEVICE_IMAGE_BASE_ADDRESS_ALIGNMENT

	/* cl_device_fp_config - bitfield */
	CL_FP_DENORM                        CL_device_fp_config = C.CL_FP_DENORM
	CL_FP_INF_NAN                       CL_device_fp_config = C.CL_FP_INF_NAN
	CL_FP_ROUND_TO_NEAREST              CL_device_fp_config = C.CL_FP_ROUND_TO_NEAREST
	CL_FP_ROUND_TO_ZERO                 CL_device_fp_config = C.CL_FP_ROUND_TO_ZERO
	CL_FP_ROUND_TO_INF                  CL_device_fp_config = C.CL_FP_ROUND_TO_INF
	CL_FP_FMA                           CL_device_fp_config = C.CL_FP_FMA
	CL_FP_SOFT_FLOAT                    CL_device_fp_config = C.CL_FP_SOFT_FLOAT
	CL_FP_CORRECTLY_ROUNDED_DIVIDE_SQRT CL_device_fp_config = C.CL_FP_CORRECTLY_ROUNDED_DIVIDE_SQRT

	/* cl_device_mem_cache_type */
	CL_NONE             CL_device_mem_cache_type = C.CL_NONE
	CL_READ_ONLY_CACHE  CL_device_mem_cache_type = C.CL_READ_ONLY_CACHE
	CL_READ_WRITE_CACHE CL_device_mem_cache_type = C.CL_READ_WRITE_CACHE

	/* cl_device_local_mem_type */
	CL_LOCAL  CL_device_local_mem_type = C.CL_LOCAL
	CL_GLOBAL CL_device_local_mem_type = C.CL_GLOBAL

	/* cl_device_exec_capabilities - bitfield */
	CL_EXEC_KERNEL        CL_device_exec_capabilities = C.CL_EXEC_KERNEL
	CL_EXEC_NATIVE_KERNEL CL_device_exec_capabilities = C.CL_EXEC_NATIVE_KERNEL

	/* cl_command_queue_properties - bitfield */
	CL_QUEUE_OUT_OF_ORDER_EXEC_MODE_ENABLE CL_command_queue_properties = C.CL_QUEUE_OUT_OF_ORDER_EXEC_MODE_ENABLE
	CL_QUEUE_PROFILING_ENABLE              CL_command_queue_properties = C.CL_QUEUE_PROFILING_ENABLE

	/* cl_context_info  */
	CL_CONTEXT_REFERENCE_COUNT CL_context_info = C.CL_CONTEXT_REFERENCE_COUNT
	CL_CONTEXT_DEVICES         CL_context_info = C.CL_CONTEXT_DEVICES
	CL_CONTEXT_PROPERTIES      CL_context_info = C.CL_CONTEXT_PROPERTIES
	CL_CONTEXT_NUM_DEVICES     CL_context_info = C.CL_CONTEXT_NUM_DEVICES

	/* cl_context_properties */
	CL_CONTEXT_PLATFORM          CL_context_properties = C.CL_CONTEXT_PLATFORM
	CL_CONTEXT_INTEROP_USER_SYNC CL_context_properties = C.CL_CONTEXT_INTEROP_USER_SYNC

	/* cl_device_partition_property */
	CL_DEVICE_PARTITION_EQUALLY            CL_device_partition_property = C.CL_DEVICE_PARTITION_EQUALLY
	CL_DEVICE_PARTITION_BY_COUNTS          CL_device_partition_property = C.CL_DEVICE_PARTITION_BY_COUNTS
	CL_DEVICE_PARTITION_BY_COUNTS_LIST_END CL_device_partition_property = C.CL_DEVICE_PARTITION_BY_COUNTS_LIST_END
	CL_DEVICE_PARTITION_BY_AFFINITY_DOMAIN CL_device_partition_property = C.CL_DEVICE_PARTITION_BY_AFFINITY_DOMAIN

	/* cl_device_affinity_domain */
	CL_DEVICE_AFFINITY_DOMAIN_NUMA               CL_device_affinity_domain = C.CL_DEVICE_AFFINITY_DOMAIN_NUMA
	CL_DEVICE_AFFINITY_DOMAIN_L4_CACHE           CL_device_affinity_domain = C.CL_DEVICE_AFFINITY_DOMAIN_L4_CACHE
	CL_DEVICE_AFFINITY_DOMAIN_L3_CACHE           CL_device_affinity_domain = C.CL_DEVICE_AFFINITY_DOMAIN_L3_CACHE
	CL_DEVICE_AFFINITY_DOMAIN_L2_CACHE           CL_device_affinity_domain = C.CL_DEVICE_AFFINITY_DOMAIN_L2_CACHE
	CL_DEVICE_AFFINITY_DOMAIN_L1_CACHE           CL_device_affinity_domain = C.CL_DEVICE_AFFINITY_DOMAIN_L1_CACHE
	CL_DEVICE_AFFINITY_DOMAIN_NEXT_PARTITIONABLE CL_device_affinity_domain = C.CL_DEVICE_AFFINITY_DOMAIN_NEXT_PARTITIONABLE

	/* cl_command_queue_info */
	CL_QUEUE_CONTEXT         CL_command_queue_info = C.CL_QUEUE_CONTEXT
	CL_QUEUE_DEVICE          CL_command_queue_info = C.CL_QUEUE_DEVICE
	CL_QUEUE_REFERENCE_COUNT CL_command_queue_info = C.CL_QUEUE_REFERENCE_COUNT
	CL_QUEUE_PROPERTIES      CL_command_queue_info = C.CL_QUEUE_PROPERTIES

	/* cl_mem_flags - bitfield */
	CL_MEM_READ_WRITE     CL_mem_flags = C.CL_MEM_READ_WRITE
	CL_MEM_WRITE_ONLY     CL_mem_flags = C.CL_MEM_WRITE_ONLY
	CL_MEM_READ_ONLY      CL_mem_flags = C.CL_MEM_READ_ONLY
	CL_MEM_USE_HOST_PTR   CL_mem_flags = C.CL_MEM_USE_HOST_PTR
	CL_MEM_ALLOC_HOST_PTR CL_mem_flags = C.CL_MEM_ALLOC_HOST_PTR
	CL_MEM_COPY_HOST_PTR  CL_mem_flags = C.CL_MEM_COPY_HOST_PTR

	// reserved
	CL_MEM_HOST_WRITE_ONLY = C.CL_MEM_HOST_WRITE_ONLY
	CL_MEM_HOST_READ_ONLY  = C.CL_MEM_HOST_READ_ONLY
	CL_MEM_HOST_NO_ACCESS  = C.CL_MEM_HOST_NO_ACCESS

	/* cl_mem_migration_flags - bitfield */
	CL_MIGRATE_MEM_OBJECT_HOST              CL_mem_migration_flags = C.CL_MIGRATE_MEM_OBJECT_HOST
	CL_MIGRATE_MEM_OBJECT_CONTENT_UNDEFINED CL_mem_migration_flags = C.CL_MIGRATE_MEM_OBJECT_CONTENT_UNDEFINED

	/* cl_channel_order */
	CL_R             CL_channel_order = C.CL_R
	CL_A             CL_channel_order = C.CL_A
	CL_RG            CL_channel_order = C.CL_RG
	CL_RA            CL_channel_order = C.CL_RA
	CL_RGB           CL_channel_order = C.CL_RGB
	CL_RGBA          CL_channel_order = C.CL_RGBA
	CL_BGRA          CL_channel_order = C.CL_BGRA
	CL_ARGB          CL_channel_order = C.CL_ARGB
	CL_INTENSITY     CL_channel_order = C.CL_INTENSITY
	CL_LUMINANCE     CL_channel_order = C.CL_LUMINANCE
	CL_Rx            CL_channel_order = C.CL_Rx
	CL_RGx           CL_channel_order = C.CL_RGx
	CL_RGBx          CL_channel_order = C.CL_RGBx
	CL_DEPTH         CL_channel_order = C.CL_DEPTH
	CL_DEPTH_STENCIL CL_channel_order = C.CL_DEPTH_STENCIL

	/* cl_channel_type */
	CL_SNORM_INT8       CL_channel_type = C.CL_SNORM_INT8
	CL_SNORM_INT16      CL_channel_type = C.CL_SNORM_INT16
	CL_UNORM_INT8       CL_channel_type = C.CL_UNORM_INT8
	CL_UNORM_INT16      CL_channel_type = C.CL_UNORM_INT16
	CL_UNORM_SHORT_565  CL_channel_type = C.CL_UNORM_SHORT_565
	CL_UNORM_SHORT_555  CL_channel_type = C.CL_UNORM_SHORT_555
	CL_UNORM_INT_101010 CL_channel_type = C.CL_UNORM_INT_101010
	CL_SIGNED_INT8      CL_channel_type = C.CL_SIGNED_INT8
	CL_SIGNED_INT16     CL_channel_type = C.CL_SIGNED_INT16
	CL_SIGNED_INT32     CL_channel_type = C.CL_SIGNED_INT32
	CL_UNSIGNED_INT8    CL_channel_type = C.CL_UNSIGNED_INT8
	CL_UNSIGNED_INT16   CL_channel_type = C.CL_UNSIGNED_INT16
	CL_UNSIGNED_INT32   CL_channel_type = C.CL_UNSIGNED_INT32
	CL_HALF_FLOAT       CL_channel_type = C.CL_HALF_FLOAT
	CL_FLOAT            CL_channel_type = C.CL_FLOAT
	CL_UNORM_INT24      CL_channel_type = C.CL_UNORM_INT24

	/* cl_mem_object_type */
	CL_MEM_OBJECT_BUFFER         CL_mem_object_type = C.CL_MEM_OBJECT_BUFFER
	CL_MEM_OBJECT_IMAGE2D        CL_mem_object_type = C.CL_MEM_OBJECT_IMAGE2D
	CL_MEM_OBJECT_IMAGE3D        CL_mem_object_type = C.CL_MEM_OBJECT_IMAGE3D
	CL_MEM_OBJECT_IMAGE2D_ARRAY  CL_mem_object_type = C.CL_MEM_OBJECT_IMAGE2D_ARRAY
	CL_MEM_OBJECT_IMAGE1D        CL_mem_object_type = C.CL_MEM_OBJECT_IMAGE1D
	CL_MEM_OBJECT_IMAGE1D_ARRAY  CL_mem_object_type = C.CL_MEM_OBJECT_IMAGE1D_ARRAY
	CL_MEM_OBJECT_IMAGE1D_BUFFER CL_mem_object_type = C.CL_MEM_OBJECT_IMAGE1D_BUFFER

	/* cl_mem_info */
	CL_MEM_TYPE                 CL_mem_info = C.CL_MEM_TYPE
	CL_MEM_FLAGS                CL_mem_info = C.CL_MEM_FLAGS
	CL_MEM_SIZE                 CL_mem_info = C.CL_MEM_SIZE
	CL_MEM_HOST_PTR             CL_mem_info = C.CL_MEM_HOST_PTR
	CL_MEM_MAP_COUNT            CL_mem_info = C.CL_MEM_MAP_COUNT
	CL_MEM_REFERENCE_COUNT      CL_mem_info = C.CL_MEM_REFERENCE_COUNT
	CL_MEM_CONTEXT              CL_mem_info = C.CL_MEM_CONTEXT
	CL_MEM_ASSOCIATED_MEMOBJECT CL_mem_info = C.CL_MEM_ASSOCIATED_MEMOBJECT
	CL_MEM_OFFSET               CL_mem_info = C.CL_MEM_OFFSET

	/* cl_image_info */
	CL_IMAGE_FORMAT         CL_image_info = C.CL_IMAGE_FORMAT
	CL_IMAGE_ELEMENT_SIZE   CL_image_info = C.CL_IMAGE_ELEMENT_SIZE
	CL_IMAGE_ROW_PITCH      CL_image_info = C.CL_IMAGE_ROW_PITCH
	CL_IMAGE_SLICE_PITCH    CL_image_info = C.CL_IMAGE_SLICE_PITCH
	CL_IMAGE_WIDTH          CL_image_info = C.CL_IMAGE_WIDTH
	CL_IMAGE_HEIGHT         CL_image_info = C.CL_IMAGE_HEIGHT
	CL_IMAGE_DEPTH          CL_image_info = C.CL_IMAGE_DEPTH
	CL_IMAGE_ARRAY_SIZE     CL_image_info = C.CL_IMAGE_ARRAY_SIZE
	CL_IMAGE_BUFFER         CL_image_info = C.CL_IMAGE_BUFFER
	CL_IMAGE_NUM_MIP_LEVELS CL_image_info = C.CL_IMAGE_NUM_MIP_LEVELS
	CL_IMAGE_NUM_SAMPLES    CL_image_info = C.CL_IMAGE_NUM_SAMPLES

	/* cl_addressing_mode */
	CL_ADDRESS_NONE            CL_addressing_mode = C.CL_ADDRESS_NONE
	CL_ADDRESS_CLAMP_TO_EDGE   CL_addressing_mode = C.CL_ADDRESS_CLAMP_TO_EDGE
	CL_ADDRESS_CLAMP           CL_addressing_mode = C.CL_ADDRESS_CLAMP
	CL_ADDRESS_REPEAT          CL_addressing_mode = C.CL_ADDRESS_REPEAT
	CL_ADDRESS_MIRRORED_REPEAT CL_addressing_mode = C.CL_ADDRESS_MIRRORED_REPEAT

	/* cl_filter_mode */
	CL_FILTER_NEAREST CL_filter_mode = C.CL_FILTER_NEAREST
	CL_FILTER_LINEAR  CL_filter_mode = C.CL_FILTER_LINEAR

	/* cl_sampler_info */
	CL_SAMPLER_REFERENCE_COUNT   CL_sampler_info = C.CL_SAMPLER_REFERENCE_COUNT
	CL_SAMPLER_CONTEXT           CL_sampler_info = C.CL_SAMPLER_CONTEXT
	CL_SAMPLER_NORMALIZED_COORDS CL_sampler_info = C.CL_SAMPLER_NORMALIZED_COORDS
	CL_SAMPLER_ADDRESSING_MODE   CL_sampler_info = C.CL_SAMPLER_ADDRESSING_MODE
	CL_SAMPLER_FILTER_MODE       CL_sampler_info = C.CL_SAMPLER_FILTER_MODE

	/* cl_map_flags - bitfield */
	CL_MAP_READ                    CL_map_flags = C.CL_MAP_READ
	CL_MAP_WRITE                   CL_map_flags = C.CL_MAP_WRITE
	CL_MAP_WRITE_INVALIDATE_REGION CL_map_flags = C.CL_MAP_WRITE_INVALIDATE_REGION

	/* cl_program_info */
	CL_PROGRAM_REFERENCE_COUNT CL_program_info = C.CL_PROGRAM_REFERENCE_COUNT
	CL_PROGRAM_CONTEXT         CL_program_info = C.CL_PROGRAM_CONTEXT
	CL_PROGRAM_NUM_DEVICES     CL_program_info = C.CL_PROGRAM_NUM_DEVICES
	CL_PROGRAM_DEVICES         CL_program_info = C.CL_PROGRAM_DEVICES
	CL_PROGRAM_SOURCE          CL_program_info = C.CL_PROGRAM_SOURCE
	CL_PROGRAM_BINARY_SIZES    CL_program_info = C.CL_PROGRAM_BINARY_SIZES
	CL_PROGRAM_BINARIES        CL_program_info = C.CL_PROGRAM_BINARIES
	CL_PROGRAM_NUM_KERNELS     CL_program_info = C.CL_PROGRAM_NUM_KERNELS
	CL_PROGRAM_KERNEL_NAMES    CL_program_info = C.CL_PROGRAM_KERNEL_NAMES

	/* cl_program_build_info */
	CL_PROGRAM_BUILD_STATUS  CL_program_build_info = C.CL_PROGRAM_BUILD_STATUS
	CL_PROGRAM_BUILD_OPTIONS CL_program_build_info = C.CL_PROGRAM_BUILD_OPTIONS
	CL_PROGRAM_BUILD_LOG     CL_program_build_info = C.CL_PROGRAM_BUILD_LOG
	CL_PROGRAM_BINARY_TYPE   CL_program_build_info = C.CL_PROGRAM_BINARY_TYPE

	/* cl_program_binary_type */
	CL_PROGRAM_BINARY_TYPE_NONE            CL_program_binary_type = C.CL_PROGRAM_BINARY_TYPE_NONE
	CL_PROGRAM_BINARY_TYPE_COMPILED_OBJECT CL_program_binary_type = C.CL_PROGRAM_BINARY_TYPE_COMPILED_OBJECT
	CL_PROGRAM_BINARY_TYPE_LIBRARY         CL_program_binary_type = C.CL_PROGRAM_BINARY_TYPE_LIBRARY
	CL_PROGRAM_BINARY_TYPE_EXECUTABLE      CL_program_binary_type = C.CL_PROGRAM_BINARY_TYPE_EXECUTABLE

	/* cl_build_status */
	CL_BUILD_SUCCESS     CL_build_status = C.CL_BUILD_SUCCESS
	CL_BUILD_NONE        CL_build_status = C.CL_BUILD_NONE
	CL_BUILD_ERROR       CL_build_status = C.CL_BUILD_ERROR
	CL_BUILD_IN_PROGRESS CL_build_status = C.CL_BUILD_IN_PROGRESS

	/* cl_kernel_info */
	CL_KERNEL_FUNCTION_NAME   CL_kernel_info = C.CL_KERNEL_FUNCTION_NAME
	CL_KERNEL_NUM_ARGS        CL_kernel_info = C.CL_KERNEL_NUM_ARGS
	CL_KERNEL_REFERENCE_COUNT CL_kernel_info = C.CL_KERNEL_REFERENCE_COUNT
	CL_KERNEL_CONTEXT         CL_kernel_info = C.CL_KERNEL_CONTEXT
	CL_KERNEL_PROGRAM         CL_kernel_info = C.CL_KERNEL_PROGRAM
	CL_KERNEL_ATTRIBUTES      CL_kernel_info = C.CL_KERNEL_ATTRIBUTES

	/* cl_kernel_arg_info */
	CL_KERNEL_ARG_ADDRESS_QUALIFIER CL_kernel_arg_info = C.CL_KERNEL_ARG_ADDRESS_QUALIFIER
	CL_KERNEL_ARG_ACCESS_QUALIFIER  CL_kernel_arg_info = C.CL_KERNEL_ARG_ACCESS_QUALIFIER
	CL_KERNEL_ARG_TYPE_NAME         CL_kernel_arg_info = C.CL_KERNEL_ARG_TYPE_NAME
	CL_KERNEL_ARG_TYPE_QUALIFIER    CL_kernel_arg_info = C.CL_KERNEL_ARG_TYPE_QUALIFIER
	CL_KERNEL_ARG_NAME              CL_kernel_arg_info = C.CL_KERNEL_ARG_NAME

	/* cl_kernel_arg_address_qualifier */
	CL_KERNEL_ARG_ADDRESS_GLOBAL   CL_kernel_arg_address_qualifier = C.CL_KERNEL_ARG_ADDRESS_GLOBAL
	CL_KERNEL_ARG_ADDRESS_LOCAL    CL_kernel_arg_address_qualifier = C.CL_KERNEL_ARG_ADDRESS_LOCAL
	CL_KERNEL_ARG_ADDRESS_CONSTANT CL_kernel_arg_address_qualifier = C.CL_KERNEL_ARG_ADDRESS_CONSTANT
	CL_KERNEL_ARG_ADDRESS_PRIVATE  CL_kernel_arg_address_qualifier = C.CL_KERNEL_ARG_ADDRESS_PRIVATE

	/* cl_kernel_arg_access_qualifier */
	CL_KERNEL_ARG_ACCESS_READ_ONLY  CL_kernel_arg_access_qualifier = C.CL_KERNEL_ARG_ACCESS_READ_ONLY
	CL_KERNEL_ARG_ACCESS_WRITE_ONLY CL_kernel_arg_access_qualifier = C.CL_KERNEL_ARG_ACCESS_WRITE_ONLY
	CL_KERNEL_ARG_ACCESS_READ_WRITE CL_kernel_arg_access_qualifier = C.CL_KERNEL_ARG_ACCESS_READ_WRITE
	CL_KERNEL_ARG_ACCESS_NONE       CL_kernel_arg_access_qualifier = C.CL_KERNEL_ARG_ACCESS_NONE

	/* cl_kernel_arg_type_qualifier */
	CL_KERNEL_ARG_TYPE_NONE     CL_kernel_arg_type_qualifier = C.CL_KERNEL_ARG_TYPE_NONE
	CL_KERNEL_ARG_TYPE_CONST    CL_kernel_arg_type_qualifier = C.CL_KERNEL_ARG_TYPE_CONST
	CL_KERNEL_ARG_TYPE_RESTRICT CL_kernel_arg_type_qualifier = C.CL_KERNEL_ARG_TYPE_RESTRICT
	CL_KERNEL_ARG_TYPE_VOLATILE CL_kernel_arg_type_qualifier = C.CL_KERNEL_ARG_TYPE_VOLATILE

	/* cl_kernel_work_group_info */
	CL_KERNEL_WORK_GROUP_SIZE                    CL_kernel_work_group_info = C.CL_KERNEL_WORK_GROUP_SIZE
	CL_KERNEL_COMPILE_WORK_GROUP_SIZE            CL_kernel_work_group_info = C.CL_KERNEL_COMPILE_WORK_GROUP_SIZE
	CL_KERNEL_LOCAL_MEM_SIZE                     CL_kernel_work_group_info = C.CL_KERNEL_LOCAL_MEM_SIZE
	CL_KERNEL_PREFERRED_WORK_GROUP_SIZE_MULTIPLE CL_kernel_work_group_info = C.CL_KERNEL_PREFERRED_WORK_GROUP_SIZE_MULTIPLE
	CL_KERNEL_PRIVATE_MEM_SIZE                   CL_kernel_work_group_info = C.CL_KERNEL_PRIVATE_MEM_SIZE
	CL_KERNEL_GLOBAL_WORK_SIZE                   CL_kernel_work_group_info = C.CL_KERNEL_GLOBAL_WORK_SIZE

	/* cl_event_info  */
	CL_EVENT_COMMAND_QUEUE            CL_event_info = C.CL_EVENT_COMMAND_QUEUE
	CL_EVENT_COMMAND_TYPE             CL_event_info = C.CL_EVENT_COMMAND_TYPE
	CL_EVENT_REFERENCE_COUNT          CL_event_info = C.CL_EVENT_REFERENCE_COUNT
	CL_EVENT_COMMAND_EXECUTION_STATUS CL_event_info = C.CL_EVENT_COMMAND_EXECUTION_STATUS
	CL_EVENT_CONTEXT                  CL_event_info = C.CL_EVENT_CONTEXT

	/* cl_command_type */
	CL_COMMAND_NDRANGE_KERNEL       CL_command_type = C.CL_COMMAND_NDRANGE_KERNEL
	CL_COMMAND_TASK                 CL_command_type = C.CL_COMMAND_TASK
	CL_COMMAND_NATIVE_KERNEL        CL_command_type = C.CL_COMMAND_NATIVE_KERNEL
	CL_COMMAND_READ_BUFFER          CL_command_type = C.CL_COMMAND_READ_BUFFER
	CL_COMMAND_WRITE_BUFFER         CL_command_type = C.CL_COMMAND_WRITE_BUFFER
	CL_COMMAND_COPY_BUFFER          CL_command_type = C.CL_COMMAND_COPY_BUFFER
	CL_COMMAND_READ_IMAGE           CL_command_type = C.CL_COMMAND_READ_IMAGE
	CL_COMMAND_WRITE_IMAGE          CL_command_type = C.CL_COMMAND_WRITE_IMAGE
	CL_COMMAND_COPY_IMAGE           CL_command_type = C.CL_COMMAND_COPY_IMAGE
	CL_COMMAND_COPY_IMAGE_TO_BUFFER CL_command_type = C.CL_COMMAND_COPY_IMAGE_TO_BUFFER
	CL_COMMAND_COPY_BUFFER_TO_IMAGE CL_command_type = C.CL_COMMAND_COPY_BUFFER_TO_IMAGE
	CL_COMMAND_MAP_BUFFER           CL_command_type = C.CL_COMMAND_MAP_BUFFER
	CL_COMMAND_MAP_IMAGE            CL_command_type = C.CL_COMMAND_MAP_IMAGE
	CL_COMMAND_UNMAP_MEM_OBJECT     CL_command_type = C.CL_COMMAND_UNMAP_MEM_OBJECT
	CL_COMMAND_MARKER               CL_command_type = C.CL_COMMAND_MARKER
	CL_COMMAND_ACQUIRE_GL_OBJECTS   CL_command_type = C.CL_COMMAND_ACQUIRE_GL_OBJECTS
	CL_COMMAND_RELEASE_GL_OBJECTS   CL_command_type = C.CL_COMMAND_RELEASE_GL_OBJECTS
	CL_COMMAND_READ_BUFFER_RECT     CL_command_type = C.CL_COMMAND_READ_BUFFER_RECT
	CL_COMMAND_WRITE_BUFFER_RECT    CL_command_type = C.CL_COMMAND_WRITE_BUFFER_RECT
	CL_COMMAND_COPY_BUFFER_RECT     CL_command_type = C.CL_COMMAND_COPY_BUFFER_RECT
	CL_COMMAND_USER                 CL_command_type = C.CL_COMMAND_USER
	CL_COMMAND_BARRIER              CL_command_type = C.CL_COMMAND_BARRIER
	CL_COMMAND_MIGRATE_MEM_OBJECTS  CL_command_type = C.CL_COMMAND_MIGRATE_MEM_OBJECTS
	CL_COMMAND_FILL_BUFFER          CL_command_type = C.CL_COMMAND_FILL_BUFFER
	CL_COMMAND_FILL_IMAGE           CL_command_type = C.CL_COMMAND_FILL_IMAGE

	/* command execution status */
	CL_COMPLETE  = C.CL_COMPLETE
	CL_RUNNING   = C.CL_RUNNING
	CL_SUBMITTED = C.CL_SUBMITTED
	CL_QUEUED    = C.CL_QUEUED

	/* cl_buffer_create_type  */
	CL_BUFFER_CREATE_TYPE_REGION CL_buffer_create_type = C.CL_BUFFER_CREATE_TYPE_REGION

	/* cl_profiling_info  */
	CL_PROFILING_COMMAND_QUEUED CL_profiling_info = C.CL_PROFILING_COMMAND_QUEUED
	CL_PROFILING_COMMAND_SUBMIT CL_profiling_info = C.CL_PROFILING_COMMAND_SUBMIT
	CL_PROFILING_COMMAND_START  CL_profiling_info = C.CL_PROFILING_COMMAND_START
	CL_PROFILING_COMMAND_END    CL_profiling_info = C.CL_PROFILING_COMMAND_END
)

/*func CL_sizeof(any unsafe.ArbitraryType) CL_size_t{
	return CL_size_t(unsafe.Sizeof(any))
}*/
