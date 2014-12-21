// +build cl11 cl12 cl20

package ocl

import (
	"fmt"
	"gocl/cl"
	"unsafe"
)

type context struct {
	context_id cl.CL_context
}

func CreateContext(properties []cl.CL_context_properties,
	devices []Device,
	pfn_notify cl.CL_ctx_notify,
	user_data unsafe.Pointer) (Context, error) {
	var numDevices cl.CL_uint
	var deviceIds []cl.CL_device_id
	var errCode cl.CL_int

	numDevices = cl.CL_uint(len(devices))
	deviceIds = make([]cl.CL_device_id, numDevices)
	for i := cl.CL_uint(0); i < numDevices; i++ {
		deviceIds[i] = devices[i].GetID()
	}

	/* Create the context */
	if context_id := cl.CLCreateContext(properties, numDevices, deviceIds, pfn_notify, user_data, &errCode); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("CreateContext failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &context{context_id}, nil
	}
}

func CreateContextFromType(properties []cl.CL_context_properties,
	device_type cl.CL_device_type,
	pfn_notify cl.CL_ctx_notify,
	user_data unsafe.Pointer) (Context, error) {
	var errCode cl.CL_int

	/* Create the context */
	if context_id := cl.CLCreateContextFromType(properties, device_type, pfn_notify, user_data, &errCode); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("CreateContext failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &context{context_id}, nil
	}
}

func (this *context) GetID() cl.CL_context {
	return this.context_id
}

func (this *context) GetInfo(param_name cl.CL_context_info) (interface{}, error) {
	/* param data */
	var param_value interface{}
	var param_size cl.CL_size_t
	var errCode cl.CL_int

	/* Find size of param data */
	if errCode = cl.CLGetContextInfo(this.context_id, param_name, 0, nil, &param_size); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetInfo failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}

	/* Access param data */
	if errCode = cl.CLGetContextInfo(this.context_id, param_name, param_size, &param_value, nil); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetInfo failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}

	return param_value, nil
}

func (this *context) Retain() error {
	if errCode := cl.CLRetainContext(this.context_id); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("Retain failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}
	return nil
}

func (this *context) Release() error {
	if errCode := cl.CLReleaseContext(this.context_id); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("Release failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}
	return nil
}

func (this *context) CreateBuffer(flags cl.CL_mem_flags, size cl.CL_size_t, host_ptr unsafe.Pointer) (Buffer, error) {
	var errCode cl.CL_int

	if memory_id := cl.CLCreateBuffer(this.context_id, flags, size, host_ptr, &errCode); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("CreateBuffer failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &buffer{memory{memory_id}}, nil
	}
}

func (this *context) CreateEvent() (Event, error) {
	var errCode cl.CL_int

	if event_id := cl.CLCreateUserEvent(this.context_id, &errCode); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("CreateUserEvent failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &event{event_id}, nil
	}
}

func (this *context) GetSupportedImageFormats(flags cl.CL_mem_flags, image_type cl.CL_mem_object_type) ([]cl.CL_image_format, error) {
	var numImageFormats cl.CL_uint
	var imageFormats []cl.CL_image_format
	var errCode cl.CL_int

	/* Find size of param data */
	if errCode = cl.CLGetSupportedImageFormats(this.context_id,
		flags,
		image_type,
		0,
		nil,
		&numImageFormats); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetSupportedImageFormats failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}

	imageFormats = make([]cl.CL_image_format, numImageFormats)

	/* Access param data */
	if errCode = cl.CLGetSupportedImageFormats(this.context_id,
		flags,
		image_type,
		numImageFormats,
		imageFormats,
		nil); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("GetSupportedImageFormats failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}

	return imageFormats, nil
}

func (this *context) CreateProgramWithSource(count cl.CL_uint,
	strings [][]byte,
	lengths []cl.CL_size_t) (Program, error) {
	var errCode cl.CL_int

	if program_id := cl.CLCreateProgramWithSource(this.context_id, count, strings, lengths, &errCode); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("CreateProgramWithSource failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &program{program_id}, nil
	}
}

func (this *context) CreateProgramWithBinary(devices []Device,
	lengths []cl.CL_size_t,
	binaries [][]byte,
	binary_status []cl.CL_int) (Program, error) {
	var errCode cl.CL_int

	numDevices := cl.CL_uint(len(devices))
	deviceIds := make([]cl.CL_device_id, numDevices)
	for i := cl.CL_uint(0); i < numDevices; i++ {
		deviceIds[i] = devices[i].GetID()
	}

	if program_id := cl.CLCreateProgramWithBinary(this.context_id, numDevices, deviceIds, lengths, binaries, binary_status, &errCode); errCode != cl.CL_SUCCESS {
		return nil, fmt.Errorf("CreateProgramWithBinary failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	} else {
		return &program{program_id}, nil
	}
}
