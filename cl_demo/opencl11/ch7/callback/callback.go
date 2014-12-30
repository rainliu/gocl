package main

import (
   "fmt"
   "gocl/cl"
   "unsafe"
   "gocl/cl_demo/utils"
)

const PROGRAM_FILE =       "callback.cl"
var    KERNEL_FUNC =[]byte("callback")


func kernel_complete(e cl.CL_event, status cl.CL_int, data unsafe.Pointer) {
   fmt.Printf("%s", *((*[]byte)(data)))
}

func read_complete(e cl.CL_event, status cl.CL_int, data unsafe.Pointer) {
   var check cl.CL_bool;
   var buffer_data []float32;

   buffer_data = *(*[]float32)(data);
   check = cl.CL_TRUE;
   for i:=0; i<4096; i++ {
      if buffer_data[i] != 5.0 {
         check = cl.CL_FALSE;
         break;
      }  
   }
   if check==cl.CL_TRUE {
      fmt.Printf("The data has been initialized successfully..\n");
   }else{
      fmt.Printf("The data has not been initialized successfully..\n");
   }
}

func main() {

   /* OpenCL data structures */
   var device []cl.CL_device_id;
   var context cl.CL_context;
   var queue cl.CL_command_queue;
   var program *cl.CL_program;
   var kernel cl.CL_kernel;
   var err cl.CL_int;

   /* Data and events */
   var kernel_msg []byte;
   var data []float32;
   var data_buffer cl.CL_mem;
   var kernel_event, read_event cl.CL_event;   
   
   data = make([]float32, 4096);   
   kernel_msg = []byte("The kernel finished successfully.\n");

   /* Create a device and context */
   device = ch0.Create_device();
   context = cl.CLCreateContext(nil, 1, device[:], nil, nil, &err);
   if err < 0 {
      println("Couldn't create a context");
      return   
   }     

   /* Build the program and create a kernel */
   program = ch0.Build_program(context, device[:], PROGRAM_FILE, nil);
   kernel = cl.CLCreateKernel(*program, KERNEL_FUNC, &err);
   if err < 0 {
      println("Couldn't create a kernel");
      return   
   };

   /* Create a write-only buffer to hold the output data */
   data_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_WRITE_ONLY, 
         cl.CL_size_t(unsafe.Sizeof(data[0]))*4096, nil, &err);
   if err < 0 {
      println("Couldn't create a buffer");
      return   
   };         

   /* Create kernel argument */
   err = cl.CLSetKernelArg(kernel, 0, cl.CL_size_t(unsafe.Sizeof(data_buffer)), unsafe.Pointer(&data_buffer));
   if err < 0 {
      println("Couldn't set a kernel argument");
      return   
   };

   /* Create a command queue */
   queue = cl.CLCreateCommandQueue(context, device[0], 0, &err);
   if err < 0 {
      println("Couldn't create a command queue");
      return   
   };

   /* Enqueue kernel */
   err = cl.CLEnqueueTask(queue, kernel, 0, nil, &kernel_event);
   if err < 0 {
      println("Couldn't enqueue the kernel");
      return   
   }

   err = cl.CLSetEventCallback(kernel_event, cl.CL_COMPLETE, 
         kernel_complete, unsafe.Pointer(&kernel_msg));
   if err < 0 {
      println("Couldn't set callback for event");
      return   
   }

   /* Read the buffer */
   err = cl.CLEnqueueReadBuffer(queue, data_buffer, cl.CL_FALSE, 0, 
      cl.CL_size_t(unsafe.Sizeof(data[0]))*4096, unsafe.Pointer(&data[0]), 0, nil, &read_event);
   if err < 0 {
      println("Couldn't read the buffer");
      return   
   }
   /* Set event handling routines */
   cl.CLSetEventCallback(read_event, cl.CL_COMPLETE, read_complete, unsafe.Pointer(&data));

   /* Deallocate resources */
   cl.CLReleaseMemObject(data_buffer);
   cl.CLReleaseKernel(kernel);
   cl.CLReleaseCommandQueue(queue);
   cl.CLReleaseProgram(*program);
   cl.CLReleaseContext(context);
}
