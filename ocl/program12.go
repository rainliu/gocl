// +build cl12

package ocl

import (
	"fmt"
	"gocl/cl"
	"unsafe"
)

func (this *program) Compile(devices []Device,
	options []byte,
	input_headers []Program,
	header_include_names [][]byte,
	pfn_notify cl.CL_prg_notify,
	user_data unsafe.Pointer) error {

	numDevices := cl.CL_uint(len(devices))
	deviceIds := make([]cl.CL_device_id, numDevices)
	for i := cl.CL_uint(0); i < numDevices; i++ {
		deviceIds[i] = devices[i].GetID()
	}

	numInputHeaders := cl.CL_uint(len(input_headers))
	inputHeaders := make([]cl.CL_program, numInputHeaders)
	for i := cl.CL_uint(0); i < numInputHeaders; i++ {
		inputHeaders[i] = input_headers[i].GetID()
	}

	if errCode := cl.CLCompileProgram(this.program_id, numDevices, deviceIds, options, numInputHeaders, inputHeaders, header_include_names, pfn_notify, user_data); errCode != cl.CL_SUCCESS {
		return fmt.Errorf("Compile failure with errcode_ret %d: %s", errCode, ERROR_CODES_STRINGS[-errCode])
	}
	return nil
}
