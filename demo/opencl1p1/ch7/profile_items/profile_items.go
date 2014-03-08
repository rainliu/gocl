package main

import (
   "fmt"
   "gocl/cl"
   "unsafe"
   "gocl/demo/ch0"
)

const PROGRAM_FILE =       "profile_items.cl"
var    KERNEL_FUNC =[]byte("profile_items")
const NUM_INTS = 4096
const NUM_ITEMS = 512
const NUM_ITERATIONS = 2000


func main() {

   /* OpenCL data structures */
   var device []cl.CL_device_id;
   var context cl.CL_context;
   var queue cl.CL_command_queue;
   var program *cl.CL_program;
   var kernel cl.CL_kernel;
   var err cl.CL_int;

   /* Data and events */
   var num_ints cl.CL_int
   var num_items [1]cl.CL_size_t 
   var data [NUM_INTS]cl.CL_int;
   var data_buffer cl.CL_mem;
   var prof_event cl.CL_event;
   var total_time cl.CL_ulong
   var time_start, time_end interface{};

   /* Initialize data */
   for i:=0; i<NUM_INTS; i++ {
      data[i] = cl.CL_int(i);
   }

   /* Set number of data points and work-items */
   num_ints = NUM_INTS;
   num_items[0] = NUM_ITEMS;

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
   data_buffer = cl.CLCreateBuffer(context, 
         cl.CL_MEM_READ_WRITE | cl.CL_MEM_COPY_HOST_PTR, 
         cl.CL_size_t(unsafe.Sizeof(data[0]))*NUM_INTS, unsafe.Pointer(&data[0]), &err);
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
   cl.CLSetKernelArg(kernel, 1, cl.CL_size_t(unsafe.Sizeof(num_ints)), unsafe.Pointer(&num_ints));

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
      cl.CLEnqueueNDRangeKernel(queue, kernel, 1, nil, num_items[:],
            nil, 0, nil, &prof_event);
      if err < 0 {
         println("Couldn't enqueue the kernel");
         return   
      }

      /* Finish processing the queue and get profiling information */
      cl.CLFinish(queue);
      cl.CLGetEventProfilingInfo(prof_event, cl.CL_PROFILING_COMMAND_START,
            cl.CL_size_t(unsafe.Sizeof(total_time)), &time_start, nil);
      cl.CLGetEventProfilingInfo(prof_event, cl.CL_PROFILING_COMMAND_END,
            cl.CL_size_t(unsafe.Sizeof(total_time)), &time_end, nil);
      total_time += time_end.(cl.CL_ulong) - time_start.(cl.CL_ulong);
   }
   fmt.Printf("Average time = %v\n", total_time/NUM_ITERATIONS);

   /* Deallocate resources */
   cl.CLReleaseEvent(prof_event);
   cl.CLReleaseKernel(kernel);
   cl.CLReleaseMemObject(data_buffer);
   cl.CLReleaseCommandQueue(queue);
   cl.CLReleaseProgram(*program);
   cl.CLReleaseContext(context);
}
