package main

import (
   "fmt"
   "gocl/cl"
   "unsafe"
   "gocl/cl_demo/utils"
)

const PROGRAM_FILE =       "op_test.cl"
var    KERNEL_FUNC =[]byte("op_test")

func main() {

   /* OpenCL data structures */
   var device []cl.CL_device_id;
   var context cl.CL_context;
   var queue cl.CL_command_queue;
   var program *cl.CL_program;
   var kernel cl.CL_kernel;
   var err cl.CL_int;

   /* Data and buffers */
   var test [4]cl.CL_int;      
   var test_buffer cl.CL_mem;

   /* Create a device and context */
   device = ch0.Create_device();
   context = cl.CLCreateContext(nil, 1, device[:], nil, nil, &err);
   if err < 0  {
      println ("Couldn't create a context");
      return;   
   }

   /* Build the program and create a kernel */
   program = ch0.Build_program(context, device[:], PROGRAM_FILE, nil);
   kernel = cl.CLCreateKernel(*program, KERNEL_FUNC, &err);
   if err < 0  {
      println ("Couldn't create a kernel");
      return;   
   };

   /* Create a write-only buffer to hold the output data */
   test_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_WRITE_ONLY, 
         cl.CL_size_t(unsafe.Sizeof(test)), nil, &err);
   if err < 0  {
      println ("Couldn't create a buffer");
      return;   
   };

   /* Create kernel argument */
   err = cl.CLSetKernelArg(kernel, 0, cl.CL_size_t(unsafe.Sizeof(test_buffer)), unsafe.Pointer(&test_buffer));
   if err < 0  {
      println ("Couldn't set a kernel argument");
      return;   
   };

   /* Create a command queue */
   queue = cl.CLCreateCommandQueue(context, device[0], 0, &err);
   if err < 0  {
      println ("Couldn't create a command queue");
      return;   
   };

   /* Enqueue kernel */
   err = cl.CLEnqueueTask(queue, kernel, 0, nil, nil);
   if err < 0  {
      println ("Couldn't enqueue the kernel");
      return;   
   }

   /* Read and print the result */
   err = cl.CLEnqueueReadBuffer(queue, test_buffer, cl.CL_TRUE, 0, 
      cl.CL_size_t(unsafe.Sizeof(test)), unsafe.Pointer(&test), 0, nil, nil);
   if err < 0  {
      println ("Couldn't read the buffer");
      return;   
   }
   for i:=0; i<3; i++ {
      fmt.Printf("%d, ", test[i]);
   }
   fmt.Printf("%d\n", test[3]);

   /* Deallocate resources */
   cl.CLReleaseMemObject(test_buffer);
   cl.CLReleaseKernel(kernel);
   cl.CLReleaseCommandQueue(queue);
   cl.CLReleaseProgram(*program);
   cl.CLReleaseContext(context);
}
