// +build cl11 cl12 cl20

package ocl

import (
	"gocl/cl"
	"unsafe"
)

type Event interface {
	GetID() cl.CL_event
	GetInfo(param_name cl.CL_event_info) (interface{}, error)
	Retain() error
	Release() error

	SetStatus(execution_status cl.CL_int) error
	SetCallback(command_exec_callback_type cl.CL_int, pfn_notify cl.CL_evt_notify, user_data unsafe.Pointer) error
	GetProfilingInfo(param_name cl.CL_profiling_info) (interface{}, error)
}

type Memory interface {
	GetID() cl.CL_mem
	GetInfo(param_name cl.CL_mem_info) (interface{}, error)
	Retain() error
	Release() error

	SetCallback(pfn_notify cl.CL_mem_notify,
		user_data unsafe.Pointer) error
	EnqueueUnmap(queue CommandQueue,
		mapped_ptr unsafe.Pointer,
		event_wait_list []Event) (Event, error)
}

type Platform interface {
	GetID() cl.CL_platform_id
	GetInfo(param_name cl.CL_platform_info) (interface{}, error)
	GetDevices(deviceType cl.CL_device_type) ([]Device, error)

	UnloadCompiler() error
}

type Sampler interface {
	GetID() cl.CL_sampler
	GetInfo(param_name cl.CL_sampler_info) (interface{}, error)
	Retain() error
	Release() error
}

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
