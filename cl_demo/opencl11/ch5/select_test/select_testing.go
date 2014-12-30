package main

import (
   "fmt"
   "gocl/cl"
   "unsafe"
   "gocl/cl_demo/utils"
)

const PROGRAM_FILE =       "select_test.cl"
var    KERNEL_FUNC =[]byte("select_test")

func main() {

   /* OpenCL data structures */
   var device []cl.CL_device_id;
   var context cl.CL_context;
   var queue cl.CL_command_queue;
   var program *cl.CL_program;
   var kernel cl.CL_kernel;
   var err cl.CL_int;

   /* Data and buffers */
   var select1 [4]float32;
   var select2 [2]cl.CL_uchar;
   var select1_buffer, select2_buffer cl.CL_mem;

   /* Create a context */
   device = ch0.Create_device();
   context = cl.CLCreateContext(nil, 1, device[:], nil, nil, &err);
   if err < 0  {
      println ("Couldn't create a context");
      return   
   }

   /* Create a kernel */
   program = ch0.Build_program(context, device[:], PROGRAM_FILE, nil);
   kernel = cl.CLCreateKernel(*program, KERNEL_FUNC, &err);
   if err < 0  {
      println ("Couldn't create a kernel");
      return   
   };

   /* Create a write-only buffer to hold the output data */
   select1_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_WRITE_ONLY, 
         cl.CL_size_t(unsafe.Sizeof(select1)), nil, &err);
   if err < 0  {
      println ("Couldn't create a buffer");
      return   
   };
   select2_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_WRITE_ONLY, 
         cl.CL_size_t(unsafe.Sizeof(select2)), nil, &err);
         
   /* Create kernel argument */
   err = cl.CLSetKernelArg(kernel, 0, cl.CL_size_t(unsafe.Sizeof(select1_buffer)), unsafe.Pointer(&select1_buffer));
   if err < 0  {
      println ("Couldn't set a kernel argument");
      return   
   };
   cl.CLSetKernelArg(kernel, 1, cl.CL_size_t(unsafe.Sizeof(select2_buffer)), unsafe.Pointer(&select2_buffer));
   
   /* Create a command queue */
   queue = cl.CLCreateCommandQueue(context, device[0], 0, &err);
   if err < 0  {
      println ("Couldn't create a command queue");
      return   
   };

   /* Enqueue kernel */
   err = cl.CLEnqueueTask(queue, kernel, 0, nil, nil);
   if err < 0  {
      println ("Couldn't enqueue the kernel");
      return   
   }

   /* Read and print the result */
   err = cl.CLEnqueueReadBuffer(queue, select1_buffer, cl.CL_TRUE, 0, 
      cl.CL_size_t(unsafe.Sizeof(select1)), unsafe.Pointer(&select1), 0, nil, nil);
   if err < 0  {
      println ("Couldn't read the buffer");
      return   
   }
   cl.CLEnqueueReadBuffer(queue, select2_buffer, cl.CL_TRUE, 0, 
      cl.CL_size_t(unsafe.Sizeof(select2)), unsafe.Pointer(&select2), 0, nil, nil);   
   
   fmt.Printf("select: ");
   for i:=0; i<3; i++ {
      fmt.Printf("%.2f, ", select1[i]);
   }
   fmt.Printf("%.2f\n", select1[3]);
   
   fmt.Printf("bitselect: %X, %X\n", select2[0], select2[1]);

   /* Deallocate resources */
   cl.CLReleaseMemObject(select1_buffer);
   cl.CLReleaseMemObject(select2_buffer);   
   cl.CLReleaseKernel(kernel);
   cl.CLReleaseCommandQueue(queue);
   cl.CLReleaseProgram(*program);
   cl.CLReleaseContext(context);
}
