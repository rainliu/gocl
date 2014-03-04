package ch0

import (
   "fmt"
   "gocl/cl"
   "os"
)

/* Find a GPU or CPU associated with the first available platform */
func Create_device() []cl.CL_device_id {

   var platform [1]cl.CL_platform_id
   var dev [1]cl.CL_device_id
   var err cl.CL_int

   /* Identify a platform */
   err = cl.CLGetPlatformIDs(1, platform[:], nil)
   if err < 0 {
      println("Couldn't identify a platform")
      return nil
   }

   /* Access a device */
   err = cl.CLGetDeviceIDs(platform[0], cl.CL_DEVICE_TYPE_GPU, 1, dev[:], nil)
   if err == cl.CL_DEVICE_NOT_FOUND {
      err = cl.CLGetDeviceIDs(platform[0], cl.CL_DEVICE_TYPE_CPU, 1, dev[:], nil)
   }
   if err < 0 {
      println("Couldn't access any devices")
      return nil
   }

   return dev[:]
}

/* Create program from a file and compile it */
func Build_program(context cl.CL_context, device []cl.CL_device_id, 
   filename string, options []byte) *cl.CL_program {
   var program cl.CL_program;
   //var program_handle;
   var program_buffer [1][]byte
   var program_log interface{}
   var program_size [1]cl.CL_size_t
   var log_size cl.CL_size_t
   var err cl.CL_int
   
   /* Read each program file and place content into buffer array */
   program_handle, err1 := os.Open(filename)
   if err1 != nil {
      fmt.Printf("Couldn't find the program file %s\n", filename)
      return nil
   }
   defer program_handle.Close()

   fi, err2 := program_handle.Stat()
   if err2 != nil {
      fmt.Printf("Couldn't find the program stat\n")
      return nil
   }
   program_size[0] = cl.CL_size_t(fi.Size())
   program_buffer[0] = make([]byte, program_size[0])
   read_size, err3 := program_handle.Read(program_buffer[0])
   if err3 != nil || cl.CL_size_t(read_size) != program_size[0] {
      fmt.Printf("read file error or file size wrong\n")
      return nil
   }

   /* Create a program containing all program content */
   program = cl.CLCreateProgramWithSource(context, 1,
      program_buffer[:], program_size[:], &err)
   if err < 0 {
      fmt.Printf("Couldn't create the program\n")
   }

   /* Build program */
   err = cl.CLBuildProgram(program, 1, device[:], options, nil, nil)
   if err < 0 {
      /* Find size of log and print to std output */
      cl.CLGetProgramBuildInfo(program, device[0], cl.CL_PROGRAM_BUILD_LOG,
         0, nil, &log_size)
      cl.CLGetProgramBuildInfo(program, device[0], cl.CL_PROGRAM_BUILD_LOG,
         log_size, &program_log, nil)
      fmt.Printf("%s\n", program_log)
      return nil
   }

   return &program;
}
