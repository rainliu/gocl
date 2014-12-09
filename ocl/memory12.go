// +build cl12

package ocl

import "gocl/cl"

func (this *memory) EnqueueMigrate(queue CommandQueue,
	mem_objects []Memory,
	flags cl.CL_mem_migration_flags,
	event_wait_list []cl.CL_event) (cl.CL_event, error) {
	var event cl.CL_event
	return event, nil
}
