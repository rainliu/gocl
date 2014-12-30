package main

import (
   "fmt"
   "gocl/cl"
   "unsafe"
   "gocl/cl_demo/utils"
)

const PROGRAM_FILE =       "profile_read.cl"
var    KERNEL_FUNC =[]byte("profile_read")
const NUM_BYTES = 131072
const NUM_ITERATIONS = 2000
const PROFILE_READ = 0

func main() {

   /* OpenCL data structures */
   var device []cl.CL_device_id;
   var context cl.CL_context;
   var queue cl.CL_command_queue;
   var program *cl.CL_program;
   var kernel cl.CL_kernel;
   var err cl.CL_int;

   /* Data and events */
   var num_vectors cl.CL_int
   var data [NUM_BYTES]byte;
   var data_buffer cl.CL_mem;
   var prof_event cl.CL_event;
   var total_time cl.CL_ulong;
   var time_start, time_end interface{}
   var mapped_memory unsafe.Pointer;

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

   /* Create a buffer to hold data */
   data_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_WRITE_ONLY, 
         cl.CL_size_t(unsafe.Sizeof(data[0]))*NUM_BYTES, nil, &err);
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

   /* Tell kernel number of char16 vectors */
   num_vectors = NUM_BYTES/16;
   cl.CLSetKernelArg(kernel, 1, cl.CL_size_t(unsafe.Sizeof(num_vectors)), unsafe.Pointer(&num_vectors));

   /* Create a command queue */
   queue = cl.CLCreateCommandQueue(context, device[0], 
         cl.CL_QUEUE_PROFILING_ENABLE, &err);
   if err < 0 {
      println("Couldn't create a command queue");
      return   
   };

   total_time = 0.0;
   for i:=0; i<NUM_ITERATIONS; i++ {

      /* Enqueue kernel */
      err = cl.CLEnqueueTask(queue, kernel, 0, nil, nil);
      if err < 0 {
         println("Couldn't enqueue the kernel");
         return   
      }

      if PROFILE_READ==1 {
         /* Read the buffer */
         err = cl.CLEnqueueReadBuffer(queue, data_buffer, cl.CL_TRUE, 0, 
               cl.CL_size_t(unsafe.Sizeof(data[0]))*NUM_BYTES, unsafe.Pointer(&data[0]), 0, nil, &prof_event);
         if err < 0 {
            println("Couldn't read the buffer");
            return
         }
      }else{
         /* Create memory map */
         mapped_memory = cl.CLEnqueueMapBuffer(queue, data_buffer, cl.CL_TRUE,
               cl.CL_MAP_READ, 0, cl.CL_size_t(unsafe.Sizeof(data[0]))*NUM_BYTES, 0, nil, &prof_event, &err);
         if err < 0 {
            println("Couldn't map the buffer to host memory");
            return   
         }
      }

      /* Get profiling information */
      cl.CLGetEventProfilingInfo(prof_event, cl.CL_PROFILING_COMMAND_START,
            cl.CL_size_t(unsafe.Sizeof(total_time)), &time_start, nil);
      cl.CLGetEventProfilingInfo(prof_event, cl.CL_PROFILING_COMMAND_END,
            cl.CL_size_t(unsafe.Sizeof(total_time)), &time_end, nil);
      total_time += time_end.(cl.CL_ulong) - time_start.(cl.CL_ulong);

      if PROFILE_READ==0{
         /* Unmap the buffer */
         err = cl.CLEnqueueUnmapMemObject(queue, data_buffer, mapped_memory,
               0, nil, nil);
         if err < 0 {
            println("Couldn't unmap the buffer");
            return   
         }
      }
   }

if PROFILE_READ==1 {
   fmt.Printf("Average read time: %v\n", total_time/NUM_ITERATIONS);
}else{
   fmt.Printf("Average map time: %v\n", total_time/NUM_ITERATIONS);
}

   /* Deallocate resources */
   cl.CLReleaseEvent(prof_event);
   cl.CLReleaseMemObject(data_buffer);
   cl.CLReleaseKernel(kernel);
   cl.CLReleaseCommandQueue(queue);
   cl.CLReleaseProgram(*program);
   cl.CLReleaseContext(context);
}
