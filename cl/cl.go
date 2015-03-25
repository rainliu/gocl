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

import "math"

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
type CL_device_svm_capabilities CL_bitfield
type CL_command_queue_properties CL_bitfield
type CL_device_partition_property CL_intptr_t
type CL_device_affinity_domain CL_bitfield

type CL_context_properties CL_intptr_t
type CL_context_info CL_uint
type CL_command_queue_info CL_uint
type CL_channel_order CL_uint
type CL_channel_type CL_uint
type CL_mem_flags CL_bitfield

//type CL_svm_mem_flags CL_bitfield
type CL_mem_object_type CL_uint
type CL_mem_info CL_uint
type CL_mem_migration_flags CL_bitfield
type CL_image_info CL_uint
type CL_buffer_create_type CL_uint
type CL_addressing_mode CL_uint
type CL_filter_mode CL_uint
type CL_sampler_info CL_uint
type CL_map_flags CL_bitfield
type CL_pipe_properties CL_intptr_t
type CL_pipe_info CL_uint
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
type CL_sampler_properties CL_bitfield
type CL_kernel_exec_info CL_uint

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
	CL_SUCCESS                                   = 0
	CL_DEVICE_NOT_FOUND                          = -1
	CL_DEVICE_NOT_AVAILABLE                      = -2
	CL_COMPILER_NOT_AVAILABLE                    = -3
	CL_MEM_OBJECT_ALLOCATION_FAILURE             = -4
	CL_OUT_OF_RESOURCES                          = -5
	CL_OUT_OF_HOST_MEMORY                        = -6
	CL_PROFILING_INFO_NOT_AVAILABLE              = -7
	CL_MEM_COPY_OVERLAP                          = -8
	CL_IMAGE_FORMAT_MISMATCH                     = -9
	CL_IMAGE_FORMAT_NOT_SUPPORTED                = -10
	CL_BUILD_PROGRAM_FAILURE                     = -11
	CL_MAP_FAILURE                               = -12
	CL_MISALIGNED_SUB_BUFFER_OFFSET              = -13
	CL_EXEC_STATUS_ERROR_FOR_EVENTS_IN_WAIT_LIST = -14
	CL_COMPILE_PROGRAM_FAILURE                   = -15
	CL_LINKER_NOT_AVAILABLE                      = -16
	CL_LINK_PROGRAM_FAILURE                      = -17
	CL_DEVICE_PARTITION_FAILED                   = -18
	CL_KERNEL_ARG_INFO_NOT_AVAILABLE             = -19

	CL_INVALID_VALUE                   = -30
	CL_INVALID_DEVICE_TYPE             = -31
	CL_INVALID_PLATFORM                = -32
	CL_INVALID_DEVICE                  = -33
	CL_INVALID_CONTEXT                 = -34
	CL_INVALID_QUEUE_PROPERTIES        = -35
	CL_INVALID_COMMAND_QUEUE           = -36
	CL_INVALID_HOST_PTR                = -37
	CL_INVALID_MEM_OBJECT              = -38
	CL_INVALID_IMAGE_FORMAT_DESCRIPTOR = -39
	CL_INVALID_IMAGE_SIZE              = -40
	CL_INVALID_SAMPLER                 = -41
	CL_INVALID_BINARY                  = -42
	CL_INVALID_BUILD_OPTIONS           = -43
	CL_INVALID_PROGRAM                 = -44
	CL_INVALID_PROGRAM_EXECUTABLE      = -45
	CL_INVALID_KERNEL_NAME             = -46
	CL_INVALID_KERNEL_DEFINITION       = -47
	CL_INVALID_KERNEL                  = -48
	CL_INVALID_ARG_INDEX               = -49
	CL_INVALID_ARG_VALUE               = -50
	CL_INVALID_ARG_SIZE                = -51
	CL_INVALID_KERNEL_ARGS             = -52
	CL_INVALID_WORK_DIMENSION          = -53
	CL_INVALID_WORK_GROUP_SIZE         = -54
	CL_INVALID_WORK_ITEM_SIZE          = -55
	CL_INVALID_GLOBAL_OFFSET           = -56
	CL_INVALID_EVENT_WAIT_LIST         = -57
	CL_INVALID_EVENT                   = -58
	CL_INVALID_OPERATION               = -59
	CL_INVALID_GL_OBJECT               = -60
	CL_INVALID_BUFFER_SIZE             = -61
	CL_INVALID_MIP_LEVEL               = -62
	CL_INVALID_GLOBAL_WORK_SIZE        = -63
	CL_INVALID_PROPERTY                = -64
	CL_INVALID_IMAGE_DESCRIPTOR        = -65
	CL_INVALID_COMPILER_OPTIONS        = -66
	CL_INVALID_LINKER_OPTIONS          = -67
	CL_INVALID_DEVICE_PARTITION_COUNT  = -68
	CL_INVALID_PIPE_SIZE               = -69
	CL_INVALID_DEVICE_QUEUE            = -70

	/* OpenCL Version */
	CL_VERSION_1_0 = 1
	CL_VERSION_1_1 = 1
	CL_VERSION_1_2 = 1
	CL_VERSION_2_0 = 1

	/* cl_bool */
	CL_FALSE        CL_bool = 0
	CL_TRUE         CL_bool = 1
	CL_BLOCKING     CL_bool = CL_TRUE
	CL_NON_BLOCKING CL_bool = CL_FALSE

	/* cl_platform_info */
	CL_PLATFORM_PROFILE    CL_platform_info = 0x0900
	CL_PLATFORM_VERSION    CL_platform_info = 0x0901
	CL_PLATFORM_NAME       CL_platform_info = 0x0902
	CL_PLATFORM_VENDOR     CL_platform_info = 0x0903
	CL_PLATFORM_EXTENSIONS CL_platform_info = 0x0904

	/* cl_device_type - bitfield */
	CL_DEVICE_TYPE_DEFAULT     CL_device_type = (1 << 0)
	CL_DEVICE_TYPE_CPU         CL_device_type = (1 << 1)
	CL_DEVICE_TYPE_GPU         CL_device_type = (1 << 2)
	CL_DEVICE_TYPE_ACCELERATOR CL_device_type = (1 << 3)
	CL_DEVICE_TYPE_CUSTOM      CL_device_type = (1 << 4)
	CL_DEVICE_TYPE_ALL         CL_device_type = 0xFFFFFFFF

	/* cl_device_info */
	CL_DEVICE_TYPE                          CL_device_info = 0x1000
	CL_DEVICE_VENDOR_ID                     CL_device_info = 0x1001
	CL_DEVICE_MAX_COMPUTE_UNITS             CL_device_info = 0x1002
	CL_DEVICE_MAX_WORK_ITEM_DIMENSIONS      CL_device_info = 0x1003
	CL_DEVICE_MAX_WORK_GROUP_SIZE           CL_device_info = 0x1004
	CL_DEVICE_MAX_WORK_ITEM_SIZES           CL_device_info = 0x1005
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_CHAR   CL_device_info = 0x1006
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_SHORT  CL_device_info = 0x1007
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_INT    CL_device_info = 0x1008
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_LONG   CL_device_info = 0x1009
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT  CL_device_info = 0x100A
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE CL_device_info = 0x100B
	CL_DEVICE_MAX_CLOCK_FREQUENCY           CL_device_info = 0x100C
	CL_DEVICE_ADDRESS_BITS                  CL_device_info = 0x100D
	CL_DEVICE_MAX_READ_IMAGE_ARGS           CL_device_info = 0x100E
	CL_DEVICE_MAX_WRITE_IMAGE_ARGS          CL_device_info = 0x100F
	CL_DEVICE_MAX_MEM_ALLOC_SIZE            CL_device_info = 0x1010
	CL_DEVICE_IMAGE2D_MAX_WIDTH             CL_device_info = 0x1011
	CL_DEVICE_IMAGE2D_MAX_HEIGHT            CL_device_info = 0x1012
	CL_DEVICE_IMAGE3D_MAX_WIDTH             CL_device_info = 0x1013
	CL_DEVICE_IMAGE3D_MAX_HEIGHT            CL_device_info = 0x1014
	CL_DEVICE_IMAGE3D_MAX_DEPTH             CL_device_info = 0x1015
	CL_DEVICE_IMAGE_SUPPORT                 CL_device_info = 0x1016
	CL_DEVICE_MAX_PARAMETER_SIZE            CL_device_info = 0x1017
	CL_DEVICE_MAX_SAMPLERS                  CL_device_info = 0x1018
	CL_DEVICE_MEM_BASE_ADDR_ALIGN           CL_device_info = 0x1019
	CL_DEVICE_MIN_DATA_TYPE_ALIGN_SIZE      CL_device_info = 0x101A
	CL_DEVICE_SINGLE_FP_CONFIG              CL_device_info = 0x101B
	CL_DEVICE_GLOBAL_MEM_CACHE_TYPE         CL_device_info = 0x101C
	CL_DEVICE_GLOBAL_MEM_CACHELINE_SIZE     CL_device_info = 0x101D
	CL_DEVICE_GLOBAL_MEM_CACHE_SIZE         CL_device_info = 0x101E
	CL_DEVICE_GLOBAL_MEM_SIZE               CL_device_info = 0x101F
	CL_DEVICE_MAX_CONSTANT_BUFFER_SIZE      CL_device_info = 0x1020
	CL_DEVICE_MAX_CONSTANT_ARGS             CL_device_info = 0x1021
	CL_DEVICE_LOCAL_MEM_TYPE                CL_device_info = 0x1022
	CL_DEVICE_LOCAL_MEM_SIZE                CL_device_info = 0x1023
	CL_DEVICE_ERROR_CORRECTION_SUPPORT      CL_device_info = 0x1024
	CL_DEVICE_PROFILING_TIMER_RESOLUTION    CL_device_info = 0x1025
	CL_DEVICE_ENDIAN_LITTLE                 CL_device_info = 0x1026
	CL_DEVICE_AVAILABLE                     CL_device_info = 0x1027
	CL_DEVICE_COMPILER_AVAILABLE            CL_device_info = 0x1028
	CL_DEVICE_EXECUTION_CAPABILITIES        CL_device_info = 0x1029
	CL_DEVICE_QUEUE_PROPERTIES              CL_device_info = 0x102A /* deprecated */
	CL_DEVICE_QUEUE_ON_HOST_PROPERTIES      CL_device_info = 0x102A
	CL_DEVICE_NAME                          CL_device_info = 0x102B
	CL_DEVICE_VENDOR                        CL_device_info = 0x102C
	CL_DRIVER_VERSION                       CL_device_info = 0x102D
	CL_DEVICE_PROFILE                       CL_device_info = 0x102E
	CL_DEVICE_VERSION                       CL_device_info = 0x102F
	CL_DEVICE_EXTENSIONS                    CL_device_info = 0x1030
	CL_DEVICE_PLATFORM                      CL_device_info = 0x1031
	CL_DEVICE_DOUBLE_FP_CONFIG              CL_device_info = 0x1032
	/* 0x1033 reserved for CL_DEVICE_HALF_FP_CONFIG */
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_HALF          CL_device_info = 0x1034
	CL_DEVICE_HOST_UNIFIED_MEMORY                  CL_device_info = 0x1035 /* deprecated */
	CL_DEVICE_NATIVE_VECTOR_WIDTH_CHAR             CL_device_info = 0x1036
	CL_DEVICE_NATIVE_VECTOR_WIDTH_SHORT            CL_device_info = 0x1037
	CL_DEVICE_NATIVE_VECTOR_WIDTH_INT              CL_device_info = 0x1038
	CL_DEVICE_NATIVE_VECTOR_WIDTH_LONG             CL_device_info = 0x1039
	CL_DEVICE_NATIVE_VECTOR_WIDTH_FLOAT            CL_device_info = 0x103A
	CL_DEVICE_NATIVE_VECTOR_WIDTH_DOUBLE           CL_device_info = 0x103B
	CL_DEVICE_NATIVE_VECTOR_WIDTH_HALF             CL_device_info = 0x103C
	CL_DEVICE_OPENCL_C_VERSION                     CL_device_info = 0x103D
	CL_DEVICE_LINKER_AVAILABLE                     CL_device_info = 0x103E
	CL_DEVICE_BUILT_IN_KERNELS                     CL_device_info = 0x103F
	CL_DEVICE_IMAGE_MAX_BUFFER_SIZE                CL_device_info = 0x1040
	CL_DEVICE_IMAGE_MAX_ARRAY_SIZE                 CL_device_info = 0x1041
	CL_DEVICE_PARENT_DEVICE                        CL_device_info = 0x1042
	CL_DEVICE_PARTITION_MAX_SUB_DEVICES            CL_device_info = 0x1043
	CL_DEVICE_PARTITION_PROPERTIES                 CL_device_info = 0x1044
	CL_DEVICE_PARTITION_AFFINITY_DOMAIN            CL_device_info = 0x1045
	CL_DEVICE_PARTITION_TYPE                       CL_device_info = 0x1046
	CL_DEVICE_REFERENCE_COUNT                      CL_device_info = 0x1047
	CL_DEVICE_PREFERRED_INTEROP_USER_SYNC          CL_device_info = 0x1048
	CL_DEVICE_PRINTF_BUFFER_SIZE                   CL_device_info = 0x1049
	CL_DEVICE_IMAGE_PITCH_ALIGNMENT                CL_device_info = 0x104A
	CL_DEVICE_IMAGE_BASE_ADDRESS_ALIGNMENT         CL_device_info = 0x104B
	CL_DEVICE_MAX_READ_WRITE_IMAGE_ARGS            CL_device_info = 0x104C
	CL_DEVICE_MAX_GLOBAL_VARIABLE_SIZE             CL_device_info = 0x104D
	CL_DEVICE_QUEUE_ON_DEVICE_PROPERTIES           CL_device_info = 0x104E
	CL_DEVICE_QUEUE_ON_DEVICE_PREFERRED_SIZE       CL_device_info = 0x104F
	CL_DEVICE_QUEUE_ON_DEVICE_MAX_SIZE             CL_device_info = 0x1050
	CL_DEVICE_MAX_ON_DEVICE_QUEUES                 CL_device_info = 0x1051
	CL_DEVICE_MAX_ON_DEVICE_EVENTS                 CL_device_info = 0x1052
	CL_DEVICE_SVM_CAPABILITIES                     CL_device_info = 0x1053
	CL_DEVICE_GLOBAL_VARIABLE_PREFERRED_TOTAL_SIZE CL_device_info = 0x1054
	CL_DEVICE_MAX_PIPE_ARGS                        CL_device_info = 0x1055
	CL_DEVICE_PIPE_MAX_ACTIVE_RESERVATIONS         CL_device_info = 0x1056
	CL_DEVICE_PIPE_MAX_PACKET_SIZE                 CL_device_info = 0x1057
	CL_DEVICE_PREFERRED_PLATFORM_ATOMIC_ALIGNMENT  CL_device_info = 0x1058
	CL_DEVICE_PREFERRED_GLOBAL_ATOMIC_ALIGNMENT    CL_device_info = 0x1059
	CL_DEVICE_PREFERRED_LOCAL_ATOMIC_ALIGNMENT     CL_device_info = 0x105A

	/* cl_device_fp_config - bitfield */
	CL_FP_DENORM                        CL_device_fp_config = (1 << 0)
	CL_FP_INF_NAN                       CL_device_fp_config = (1 << 1)
	CL_FP_ROUND_TO_NEAREST              CL_device_fp_config = (1 << 2)
	CL_FP_ROUND_TO_ZERO                 CL_device_fp_config = (1 << 3)
	CL_FP_ROUND_TO_INF                  CL_device_fp_config = (1 << 4)
	CL_FP_FMA                           CL_device_fp_config = (1 << 5)
	CL_FP_SOFT_FLOAT                    CL_device_fp_config = (1 << 6)
	CL_FP_CORRECTLY_ROUNDED_DIVIDE_SQRT CL_device_fp_config = (1 << 7)

	/* cl_device_mem_cache_type */
	CL_NONE             CL_device_mem_cache_type = 0x0
	CL_READ_ONLY_CACHE  CL_device_mem_cache_type = 0x1
	CL_READ_WRITE_CACHE CL_device_mem_cache_type = 0x2

	/* cl_device_local_mem_type */
	CL_LOCAL  CL_device_local_mem_type = 0x1
	CL_GLOBAL CL_device_local_mem_type = 0x2

	/* cl_device_exec_capabilities - bitfield */
	CL_EXEC_KERNEL        CL_device_exec_capabilities = (1 << 0)
	CL_EXEC_NATIVE_KERNEL CL_device_exec_capabilities = (1 << 1)

	/* cl_command_queue_properties - bitfield */
	CL_QUEUE_OUT_OF_ORDER_EXEC_MODE_ENABLE CL_command_queue_properties = (1 << 0)
	CL_QUEUE_PROFILING_ENABLE              CL_command_queue_properties = (1 << 1)
	CL_QUEUE_ON_DEVICE                     CL_command_queue_properties = (1 << 2)
	CL_QUEUE_ON_DEVICE_DEFAULT             CL_command_queue_properties = (1 << 3)

	/* cl_context_info  */
	CL_CONTEXT_REFERENCE_COUNT CL_context_info = 0x1080
	CL_CONTEXT_DEVICES         CL_context_info = 0x1081
	CL_CONTEXT_PROPERTIES      CL_context_info = 0x1082
	CL_CONTEXT_NUM_DEVICES     CL_context_info = 0x1083

	/* cl_context_properties */
	CL_CONTEXT_PLATFORM          CL_context_properties = 0x1084
	CL_CONTEXT_INTEROP_USER_SYNC CL_context_properties = 0x1085

	/* cl_device_partition_property */
	CL_DEVICE_PARTITION_EQUALLY            CL_device_partition_property = 0x1086
	CL_DEVICE_PARTITION_BY_COUNTS          CL_device_partition_property = 0x1087
	CL_DEVICE_PARTITION_BY_COUNTS_LIST_END CL_device_partition_property = 0x0
	CL_DEVICE_PARTITION_BY_AFFINITY_DOMAIN CL_device_partition_property = 0x1088

	/* cl_device_affinity_domain */
	CL_DEVICE_AFFINITY_DOMAIN_NUMA               CL_device_affinity_domain = (1 << 0)
	CL_DEVICE_AFFINITY_DOMAIN_L4_CACHE           CL_device_affinity_domain = (1 << 1)
	CL_DEVICE_AFFINITY_DOMAIN_L3_CACHE           CL_device_affinity_domain = (1 << 2)
	CL_DEVICE_AFFINITY_DOMAIN_L2_CACHE           CL_device_affinity_domain = (1 << 3)
	CL_DEVICE_AFFINITY_DOMAIN_L1_CACHE           CL_device_affinity_domain = (1 << 4)
	CL_DEVICE_AFFINITY_DOMAIN_NEXT_PARTITIONABLE CL_device_affinity_domain = (1 << 5)

	/* cl_device_svm_capabilities */
	CL_DEVICE_SVM_COARSE_GRAIN_BUFFER CL_device_svm_capabilities = (1 << 0)
	CL_DEVICE_SVM_FINE_GRAIN_BUFFER   CL_device_svm_capabilities = (1 << 1)
	CL_DEVICE_SVM_FINE_GRAIN_SYSTEM   CL_device_svm_capabilities = (1 << 2)
	CL_DEVICE_SVM_ATOMICS             CL_device_svm_capabilities = (1 << 3)

	/* cl_command_queue_info */
	CL_QUEUE_CONTEXT         CL_command_queue_info = 0x1090
	CL_QUEUE_DEVICE          CL_command_queue_info = 0x1091
	CL_QUEUE_REFERENCE_COUNT CL_command_queue_info = 0x1092
	CL_QUEUE_PROPERTIES      CL_command_queue_info = 0x1093
	CL_QUEUE_SIZE            CL_command_queue_info = 0x1094

	/* cl_mem_flags and cl_svm_mem_flags - bitfield */
	CL_MEM_READ_WRITE     CL_mem_flags = (1 << 0)
	CL_MEM_WRITE_ONLY     CL_mem_flags = (1 << 1)
	CL_MEM_READ_ONLY      CL_mem_flags = (1 << 2)
	CL_MEM_USE_HOST_PTR   CL_mem_flags = (1 << 3)
	CL_MEM_ALLOC_HOST_PTR CL_mem_flags = (1 << 4)
	CL_MEM_COPY_HOST_PTR  CL_mem_flags = (1 << 5)
	/* reserved                          (1 << 6)    */
	CL_MEM_HOST_WRITE_ONLY       CL_mem_flags = (1 << 7)
	CL_MEM_HOST_READ_ONLY        CL_mem_flags = (1 << 8)
	CL_MEM_HOST_NO_ACCESS        CL_mem_flags = (1 << 9)
	CL_MEM_SVM_FINE_GRAIN_BUFFER CL_mem_flags = (1 << 10) /* used by cl_svm_mem_flags only */
	CL_MEM_SVM_ATOMICS           CL_mem_flags = (1 << 11) /* used by cl_svm_mem_flags only */
	CL_MEM_KERNEL_READ_AND_WRITE CL_mem_flags = (1 << 12)

	/* cl_mem_migration_flags - bitfield */
	CL_MIGRATE_MEM_OBJECT_HOST              CL_mem_migration_flags = (1 << 0)
	CL_MIGRATE_MEM_OBJECT_CONTENT_UNDEFINED CL_mem_migration_flags = (1 << 1)

	/* cl_channel_order */
	CL_R             CL_channel_order = 0x10B0
	CL_A             CL_channel_order = 0x10B1
	CL_RG            CL_channel_order = 0x10B2
	CL_RA            CL_channel_order = 0x10B3
	CL_RGB           CL_channel_order = 0x10B4
	CL_RGBA          CL_channel_order = 0x10B5
	CL_BGRA          CL_channel_order = 0x10B6
	CL_ARGB          CL_channel_order = 0x10B7
	CL_INTENSITY     CL_channel_order = 0x10B8
	CL_LUMINANCE     CL_channel_order = 0x10B9
	CL_Rx            CL_channel_order = 0x10BA
	CL_RGx           CL_channel_order = 0x10BB
	CL_RGBx          CL_channel_order = 0x10BC
	CL_DEPTH         CL_channel_order = 0x10BD
	CL_DEPTH_STENCIL CL_channel_order = 0x10BE
	CL_sRGB          CL_channel_order = 0x10BF
	CL_sRGBx         CL_channel_order = 0x10C0
	CL_sRGBA         CL_channel_order = 0x10C1
	CL_sBGRA         CL_channel_order = 0x10C2
	CL_ABGR          CL_channel_order = 0x10C3

	/* cl_channel_type */
	CL_SNORM_INT8       CL_channel_type = 0x10D0
	CL_SNORM_INT16      CL_channel_type = 0x10D1
	CL_UNORM_INT8       CL_channel_type = 0x10D2
	CL_UNORM_INT16      CL_channel_type = 0x10D3
	CL_UNORM_SHORT_565  CL_channel_type = 0x10D4
	CL_UNORM_SHORT_555  CL_channel_type = 0x10D5
	CL_UNORM_INT_101010 CL_channel_type = 0x10D6
	CL_SIGNED_INT8      CL_channel_type = 0x10D7
	CL_SIGNED_INT16     CL_channel_type = 0x10D8
	CL_SIGNED_INT32     CL_channel_type = 0x10D9
	CL_UNSIGNED_INT8    CL_channel_type = 0x10DA
	CL_UNSIGNED_INT16   CL_channel_type = 0x10DB
	CL_UNSIGNED_INT32   CL_channel_type = 0x10DC
	CL_HALF_FLOAT       CL_channel_type = 0x10DD
	CL_FLOAT            CL_channel_type = 0x10DE
	CL_UNORM_INT24      CL_channel_type = 0x10DF

	/* cl_mem_object_type */
	CL_MEM_OBJECT_BUFFER         CL_mem_object_type = 0x10F0
	CL_MEM_OBJECT_IMAGE2D        CL_mem_object_type = 0x10F1
	CL_MEM_OBJECT_IMAGE3D        CL_mem_object_type = 0x10F2
	CL_MEM_OBJECT_IMAGE2D_ARRAY  CL_mem_object_type = 0x10F3
	CL_MEM_OBJECT_IMAGE1D        CL_mem_object_type = 0x10F4
	CL_MEM_OBJECT_IMAGE1D_ARRAY  CL_mem_object_type = 0x10F5
	CL_MEM_OBJECT_IMAGE1D_BUFFER CL_mem_object_type = 0x10F6
	CL_MEM_OBJECT_PIPE           CL_mem_object_type = 0x10F7

	/* cl_mem_info */
	CL_MEM_TYPE                 CL_mem_info = 0x1100
	CL_MEM_FLAGS                CL_mem_info = 0x1101
	CL_MEM_SIZE                 CL_mem_info = 0x1102
	CL_MEM_HOST_PTR             CL_mem_info = 0x1103
	CL_MEM_MAP_COUNT            CL_mem_info = 0x1104
	CL_MEM_REFERENCE_COUNT      CL_mem_info = 0x1105
	CL_MEM_CONTEXT              CL_mem_info = 0x1106
	CL_MEM_ASSOCIATED_MEMOBJECT CL_mem_info = 0x1107
	CL_MEM_OFFSET               CL_mem_info = 0x1108
	CL_MEM_USES_SVM_POINTER     CL_mem_info = 0x1109

	/* cl_image_info */
	CL_IMAGE_FORMAT         CL_image_info = 0x1110
	CL_IMAGE_ELEMENT_SIZE   CL_image_info = 0x1111
	CL_IMAGE_ROW_PITCH      CL_image_info = 0x1112
	CL_IMAGE_SLICE_PITCH    CL_image_info = 0x1113
	CL_IMAGE_WIDTH          CL_image_info = 0x1114
	CL_IMAGE_HEIGHT         CL_image_info = 0x1115
	CL_IMAGE_DEPTH          CL_image_info = 0x1116
	CL_IMAGE_ARRAY_SIZE     CL_image_info = 0x1117
	CL_IMAGE_BUFFER         CL_image_info = 0x1118
	CL_IMAGE_NUM_MIP_LEVELS CL_image_info = 0x1119
	CL_IMAGE_NUM_SAMPLES    CL_image_info = 0x111A

	/* cl_pipe_info */
	CL_PIPE_PACKET_SIZE CL_pipe_info = 0x1120
	CL_PIPE_MAX_PACKETS CL_pipe_info = 0x1121

	/* cl_addressing_mode */
	CL_ADDRESS_NONE            CL_addressing_mode = 0x1130
	CL_ADDRESS_CLAMP_TO_EDGE   CL_addressing_mode = 0x1131
	CL_ADDRESS_CLAMP           CL_addressing_mode = 0x1132
	CL_ADDRESS_REPEAT          CL_addressing_mode = 0x1133
	CL_ADDRESS_MIRRORED_REPEAT CL_addressing_mode = 0x1134

	/* cl_filter_mode */
	CL_FILTER_NEAREST CL_filter_mode = 0x1140
	CL_FILTER_LINEAR  CL_filter_mode = 0x1141

	/* cl_sampler_info */
	CL_SAMPLER_REFERENCE_COUNT   CL_sampler_info = 0x1150
	CL_SAMPLER_CONTEXT           CL_sampler_info = 0x1151
	CL_SAMPLER_NORMALIZED_COORDS CL_sampler_info = 0x1152
	CL_SAMPLER_ADDRESSING_MODE   CL_sampler_info = 0x1153
	CL_SAMPLER_FILTER_MODE       CL_sampler_info = 0x1154
	CL_SAMPLER_MIP_FILTER_MODE   CL_sampler_info = 0x1155
	CL_SAMPLER_LOD_MIN           CL_sampler_info = 0x1156
	CL_SAMPLER_LOD_MAX           CL_sampler_info = 0x1157

	/* cl_map_flags - bitfield */
	CL_MAP_READ                    CL_map_flags = (1 << 0)
	CL_MAP_WRITE                   CL_map_flags = (1 << 1)
	CL_MAP_WRITE_INVALIDATE_REGION CL_map_flags = (1 << 2)

	/* cl_program_info */
	CL_PROGRAM_REFERENCE_COUNT CL_program_info = 0x1160
	CL_PROGRAM_CONTEXT         CL_program_info = 0x1161
	CL_PROGRAM_NUM_DEVICES     CL_program_info = 0x1162
	CL_PROGRAM_DEVICES         CL_program_info = 0x1163
	CL_PROGRAM_SOURCE          CL_program_info = 0x1164
	CL_PROGRAM_BINARY_SIZES    CL_program_info = 0x1165
	CL_PROGRAM_BINARIES        CL_program_info = 0x1166
	CL_PROGRAM_NUM_KERNELS     CL_program_info = 0x1167
	CL_PROGRAM_KERNEL_NAMES    CL_program_info = 0x1168

	/* cl_program_build_info */
	CL_PROGRAM_BUILD_STATUS                     CL_program_build_info = 0x1181
	CL_PROGRAM_BUILD_OPTIONS                    CL_program_build_info = 0x1182
	CL_PROGRAM_BUILD_LOG                        CL_program_build_info = 0x1183
	CL_PROGRAM_BINARY_TYPE                      CL_program_build_info = 0x1184
	CL_PROGRAM_BUILD_GLOBAL_VARIABLE_TOTAL_SIZE CL_program_build_info = 0x1185

	/* cl_program_binary_type */
	CL_PROGRAM_BINARY_TYPE_NONE            CL_program_binary_type = 0x0
	CL_PROGRAM_BINARY_TYPE_COMPILED_OBJECT CL_program_binary_type = 0x1
	CL_PROGRAM_BINARY_TYPE_LIBRARY         CL_program_binary_type = 0x2
	CL_PROGRAM_BINARY_TYPE_EXECUTABLE      CL_program_binary_type = 0x4

	/* cl_build_status */
	CL_BUILD_SUCCESS     CL_build_status = 0
	CL_BUILD_NONE        CL_build_status = -1
	CL_BUILD_ERROR       CL_build_status = -2
	CL_BUILD_IN_PROGRESS CL_build_status = -3

	/* cl_kernel_info */
	CL_KERNEL_FUNCTION_NAME   CL_kernel_info = 0x1190
	CL_KERNEL_NUM_ARGS        CL_kernel_info = 0x1191
	CL_KERNEL_REFERENCE_COUNT CL_kernel_info = 0x1192
	CL_KERNEL_CONTEXT         CL_kernel_info = 0x1193
	CL_KERNEL_PROGRAM         CL_kernel_info = 0x1194
	CL_KERNEL_ATTRIBUTES      CL_kernel_info = 0x1195

	/* cl_kernel_arg_info */
	CL_KERNEL_ARG_ADDRESS_QUALIFIER CL_kernel_arg_info = 0x1196
	CL_KERNEL_ARG_ACCESS_QUALIFIER  CL_kernel_arg_info = 0x1197
	CL_KERNEL_ARG_TYPE_NAME         CL_kernel_arg_info = 0x1198
	CL_KERNEL_ARG_TYPE_QUALIFIER    CL_kernel_arg_info = 0x1199
	CL_KERNEL_ARG_NAME              CL_kernel_arg_info = 0x119A

	/* cl_kernel_arg_address_qualifier */
	CL_KERNEL_ARG_ADDRESS_GLOBAL   CL_kernel_arg_address_qualifier = 0x119B
	CL_KERNEL_ARG_ADDRESS_LOCAL    CL_kernel_arg_address_qualifier = 0x119C
	CL_KERNEL_ARG_ADDRESS_CONSTANT CL_kernel_arg_address_qualifier = 0x119D
	CL_KERNEL_ARG_ADDRESS_PRIVATE  CL_kernel_arg_address_qualifier = 0x119E

	/* cl_kernel_arg_access_qualifier */
	CL_KERNEL_ARG_ACCESS_READ_ONLY  CL_kernel_arg_access_qualifier = 0x11A0
	CL_KERNEL_ARG_ACCESS_WRITE_ONLY CL_kernel_arg_access_qualifier = 0x11A1
	CL_KERNEL_ARG_ACCESS_READ_WRITE CL_kernel_arg_access_qualifier = 0x11A2
	CL_KERNEL_ARG_ACCESS_NONE       CL_kernel_arg_access_qualifier = 0x11A3

	/* cl_kernel_arg_type_qualifier */
	CL_KERNEL_ARG_TYPE_NONE     CL_kernel_arg_type_qualifier = 0
	CL_KERNEL_ARG_TYPE_CONST    CL_kernel_arg_type_qualifier = (1 << 0)
	CL_KERNEL_ARG_TYPE_RESTRICT CL_kernel_arg_type_qualifier = (1 << 1)
	CL_KERNEL_ARG_TYPE_VOLATILE CL_kernel_arg_type_qualifier = (1 << 2)
	CL_KERNEL_ARG_TYPE_PIPE     CL_kernel_arg_type_qualifier = (1 << 3)

	/* cl_kernel_work_group_info */
	CL_KERNEL_WORK_GROUP_SIZE                    CL_kernel_work_group_info = 0x11B0
	CL_KERNEL_COMPILE_WORK_GROUP_SIZE            CL_kernel_work_group_info = 0x11B1
	CL_KERNEL_LOCAL_MEM_SIZE                     CL_kernel_work_group_info = 0x11B2
	CL_KERNEL_PREFERRED_WORK_GROUP_SIZE_MULTIPLE CL_kernel_work_group_info = 0x11B3
	CL_KERNEL_PRIVATE_MEM_SIZE                   CL_kernel_work_group_info = 0x11B4
	CL_KERNEL_GLOBAL_WORK_SIZE                   CL_kernel_work_group_info = 0x11B5

	/* cl_kernel_exec_info */
	CL_KERNEL_EXEC_INFO_SVM_PTRS              CL_kernel_exec_info = 0x11B6
	CL_KERNEL_EXEC_INFO_SVM_FINE_GRAIN_SYSTEM CL_kernel_exec_info = 0x11B7

	/* cl_event_info  */
	CL_EVENT_COMMAND_QUEUE            CL_event_info = 0x11D0
	CL_EVENT_COMMAND_TYPE             CL_event_info = 0x11D1
	CL_EVENT_REFERENCE_COUNT          CL_event_info = 0x11D2
	CL_EVENT_COMMAND_EXECUTION_STATUS CL_event_info = 0x11D3
	CL_EVENT_CONTEXT                  CL_event_info = 0x11D4

	/* cl_command_type */
	CL_COMMAND_NDRANGE_KERNEL       CL_command_type = 0x11F0
	CL_COMMAND_TASK                 CL_command_type = 0x11F1
	CL_COMMAND_NATIVE_KERNEL        CL_command_type = 0x11F2
	CL_COMMAND_READ_BUFFER          CL_command_type = 0x11F3
	CL_COMMAND_WRITE_BUFFER         CL_command_type = 0x11F4
	CL_COMMAND_COPY_BUFFER          CL_command_type = 0x11F5
	CL_COMMAND_READ_IMAGE           CL_command_type = 0x11F6
	CL_COMMAND_WRITE_IMAGE          CL_command_type = 0x11F7
	CL_COMMAND_COPY_IMAGE           CL_command_type = 0x11F8
	CL_COMMAND_COPY_IMAGE_TO_BUFFER CL_command_type = 0x11F9
	CL_COMMAND_COPY_BUFFER_TO_IMAGE CL_command_type = 0x11FA
	CL_COMMAND_MAP_BUFFER           CL_command_type = 0x11FB
	CL_COMMAND_MAP_IMAGE            CL_command_type = 0x11FC
	CL_COMMAND_UNMAP_MEM_OBJECT     CL_command_type = 0x11FD
	CL_COMMAND_MARKER               CL_command_type = 0x11FE
	CL_COMMAND_ACQUIRE_GL_OBJECTS   CL_command_type = 0x11FF
	CL_COMMAND_RELEASE_GL_OBJECTS   CL_command_type = 0x1200
	CL_COMMAND_READ_BUFFER_RECT     CL_command_type = 0x1201
	CL_COMMAND_WRITE_BUFFER_RECT    CL_command_type = 0x1202
	CL_COMMAND_COPY_BUFFER_RECT     CL_command_type = 0x1203
	CL_COMMAND_USER                 CL_command_type = 0x1204
	CL_COMMAND_BARRIER              CL_command_type = 0x1205
	CL_COMMAND_MIGRATE_MEM_OBJECTS  CL_command_type = 0x1206
	CL_COMMAND_FILL_BUFFER          CL_command_type = 0x1207
	CL_COMMAND_FILL_IMAGE           CL_command_type = 0x1208
	CL_COMMAND_SVM_FREE             CL_command_type = 0x1209
	CL_COMMAND_SVM_MEMCPY           CL_command_type = 0x120A
	CL_COMMAND_SVM_MEMFILL          CL_command_type = 0x120B
	CL_COMMAND_SVM_MAP              CL_command_type = 0x120C
	CL_COMMAND_SVM_UNMAP            CL_command_type = 0x120D

	/* command execution status */
	CL_COMPLETE  CL_int = 0x0
	CL_RUNNING   CL_int = 0x1
	CL_SUBMITTED CL_int = 0x2
	CL_QUEUED    CL_int = 0x3

	/* cl_buffer_create_type  */
	CL_BUFFER_CREATE_TYPE_REGION CL_buffer_create_type = 0x1220

	/* cl_profiling_info  */
	CL_PROFILING_COMMAND_QUEUED   CL_profiling_info = 0x1280
	CL_PROFILING_COMMAND_SUBMIT   CL_profiling_info = 0x1281
	CL_PROFILING_COMMAND_START    CL_profiling_info = 0x1282
	CL_PROFILING_COMMAND_END      CL_profiling_info = 0x1283
	CL_PROFILING_COMMAND_COMPLETE CL_profiling_info = 0x1284
)

/* Error Codes Strings*/
var ERROR_CODES_STRINGS = []string{
	"CL_SUCCESS",                                   //0
	"CL_DEVICE_NOT_FOUND",                          //-1
	"CL_DEVICE_NOT_AVAILABLE",                      //-2
	"CL_COMPILER_NOT_AVAILABLE",                    //-3
	"CL_MEM_OBJECT_ALLOCATION_FAILURE",             //-4
	"CL_OUT_OF_RESOURCES",                          //-5
	"CL_OUT_OF_HOST_MEMORY",                        //-6
	"CL_PROFILING_INFO_NOT_AVAILABLE",              //-7
	"CL_MEM_COPY_OVERLAP",                          //-8
	"CL_IMAGE_FORMAT_MISMATCH",                     //-9
	"CL_IMAGE_FORMAT_NOT_SUPPORTED",                //-10
	"CL_BUILD_PROGRAM_FAILURE",                     //-11
	"CL_MAP_FAILURE",                               //-12
	"CL_MISALIGNED_SUB_BUFFER_OFFSET",              //-13
	"CL_EXEC_STATUS_ERROR_FOR_EVENTS_IN_WAIT_LIST", //-14
	"CL_COMPILE_PROGRAM_FAILURE",                   //-15
	"CL_LINKER_NOT_AVAILABLE",                      //-16
	"CL_LINK_PROGRAM_FAILURE",                      //-17
	"CL_DEVICE_PARTITION_FAILED",                   //-18
	"CL_KERNEL_ARG_INFO_NOT_AVAILABLE",             //-19
	"CL_RESERVED_20",                               //-20
	"CL_RESERVED_21",                               //-21
	"CL_RESERVED_22",                               //-22
	"CL_RESERVED_23",                               //-23
	"CL_RESERVED_24",                               //-24
	"CL_RESERVED_25",                               //-25
	"CL_RESERVED_26",                               //-26
	"CL_RESERVED_27",                               //-27
	"CL_RESERVED_28",                               //-28
	"CL_RESERVED_29",                               //-29
	"CL_INVALID_VALUE",                             //-30
	"CL_INVALID_DEVICE_TYPE",                       //-31
	"CL_INVALID_PLATFORM",                          //-32
	"CL_INVALID_DEVICE",                            //-33
	"CL_INVALID_CONTEXT",                           //-34
	"CL_INVALID_QUEUE_PROPERTIES",                  //-35
	"CL_INVALID_COMMAND_QUEUE",                     //-36
	"CL_INVALID_HOST_PTR",                          //-37
	"CL_INVALID_MEM_OBJECT",                        //-38
	"CL_INVALID_IMAGE_FORMAT_DESCRIPTOR",           //-39
	"CL_INVALID_IMAGE_SIZE",                        //-40
	"CL_INVALID_SAMPLER",                           //-41
	"CL_INVALID_BINARY",                            //-42
	"CL_INVALID_BUILD_OPTIONS",                     //-43
	"CL_INVALID_PROGRAM",                           //-44
	"CL_INVALID_PROGRAM_EXECUTABLE",                //-45
	"CL_INVALID_KERNEL_NAME",                       //-46
	"CL_INVALID_KERNEL_DEFINITION",                 //-47
	"CL_INVALID_KERNEL",                            //-48
	"CL_INVALID_ARG_INDEX",                         //-49
	"CL_INVALID_ARG_VALUE",                         //-50
	"CL_INVALID_ARG_SIZE",                          //-51
	"CL_INVALID_KERNEL_ARGS",                       //-52
	"CL_INVALID_WORK_DIMENSION",                    //-53
	"CL_INVALID_WORK_GROUP_SIZE",                   //-54
	"CL_INVALID_WORK_ITEM_SIZE",                    //-55
	"CL_INVALID_GLOBAL_OFFSET",                     //-56
	"CL_INVALID_EVENT_WAIT_LIST",                   //-57
	"CL_INVALID_EVENT",                             //-58
	"CL_INVALID_OPERATION",                         //-59
	"CL_INVALID_GL_OBJECT",                         //-60
	"CL_INVALID_BUFFER_SIZE",                       //-61
	"CL_INVALID_MIP_LEVEL",                         //-62
	"CL_INVALID_GLOBAL_WORK_SIZE",                  //-63
	"CL_INVALID_PROPERTY",                          //-64
	"CL_INVALID_IMAGE_DESCRIPTOR",                  //-65
	"CL_INVALID_COMPILER_OPTIONS",                  //-66
	"CL_INVALID_LINKER_OPTIONS",                    //-67
	"CL_INVALID_DEVICE_PARTITION_COUNT",            //-68
	"CL_INVALID_PIPE_SIZE",                         //-69
	"CL_INVALID_DEVICE_QUEUE",                      //-70
	"GOCL_RESERVED_71",                             //-71
	"GOCL_RESERVED_72",                             //-72
	"GOCL_RESERVED_73",                             //-73
	"GOCL_RESERVED_74",                             //-74
	"GOCL_RESERVED_75",                             //-75
	"GOCL_RESERVED_76",                             //-76
	"GOCL_RESERVED_77",                             //-77
	"GOCL_RESERVED_78",                             //-78
	"GOCL_RESERVED_79",                             //-79
	"GOCL_RESERVED_81",                             //-81
	"GOCL_RESERVED_82",                             //-82
	"GOCL_RESERVED_83",                             //-83
	"GOCL_RESERVED_84",                             //-84
	"GOCL_RESERVED_85",                             //-85
	"GOCL_RESERVED_86",                             //-86
	"GOCL_RESERVED_87",                             //-87
	"GOCL_RESERVED_88",                             //-88
	"GOCL_RESERVED_89",                             //-89
	"GOCL_RESERVED_90",                             //-90
	"GOCL_RESERVED_91",                             //-91
	"GOCL_RESERVED_92",                             //-92
	"GOCL_RESERVED_93",                             //-93
	"GOCL_RESERVED_94",                             //-94
	"GOCL_RESERVED_95",                             //-95
	"GOCL_RESERVED_96",                             //-96
	"GOCL_RESERVED_97",                             //-97
	"GOCL_RESERVED_98",                             //-98
	"GOCL_RESERVED_99",                             //-99
}
