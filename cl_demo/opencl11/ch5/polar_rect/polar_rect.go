package main

import (
   "fmt"
   "gocl/cl"
   "unsafe"
   "gocl/cl_demo/utils"
)

const M_PI = 3.14159265358979323846

const PROGRAM_FILE =       "polar_rect.cl"
var    KERNEL_FUNC =[]byte("polar_rect")

func main() {

   /* OpenCL data structures */
   var device []cl.CL_device_id;
   var context cl.CL_context;
   var queue cl.CL_command_queue;
   var program *cl.CL_program;
   var kernel cl.CL_kernel;
   var err cl.CL_int;

   /* Data and buffers */
   var r_coords = [4]float32{2, 1, 3, 4};
   var angles   = [4]float32 {3*M_PI/8, 3*M_PI/4, 4*M_PI/3, 11*M_PI/6};
   var x_coords, y_coords [4]float32;
   var r_coords_buffer, angles_buffer,
         x_coords_buffer, y_coords_buffer cl.CL_mem;

   /* Create a device and context */
   device = ch0.Create_device();
   context = cl.CLCreateContext(nil, 1, device[:], nil, nil, &err);
   if err < 0  {
      println ("Couldn't create a context");
      return;   
   }

   /* Create a kernel */
   program = ch0.Build_program(context, device[:], PROGRAM_FILE, nil);
   kernel = cl.CLCreateKernel(*program, KERNEL_FUNC, &err);
   if err < 0  {
      println ("Couldn't create a kernel");
      return;   
   };

   /* Create a write-only buffer to hold the output data */
   r_coords_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_READ_ONLY |  cl.CL_MEM_COPY_HOST_PTR,
         cl.CL_size_t(unsafe.Sizeof(r_coords)), unsafe.Pointer(&r_coords[0]), &err);
   if err < 0  {
      println ("Couldn't create a buffer");
      return;   
   };
   angles_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_READ_ONLY |  cl.CL_MEM_COPY_HOST_PTR,
         cl.CL_size_t(unsafe.Sizeof(angles)), unsafe.Pointer(&angles[0]), &err);   
   x_coords_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_READ_WRITE, 
         cl.CL_size_t(unsafe.Sizeof(x_coords)), nil, &err);
   y_coords_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_READ_WRITE, 
         cl.CL_size_t(unsafe.Sizeof(y_coords)), nil, &err);         
         
   /* Create kernel argument */
   err = cl.CLSetKernelArg(kernel, 0, cl.CL_size_t(unsafe.Sizeof(r_coords_buffer)), unsafe.Pointer(&r_coords_buffer));
   if err < 0  {
      println ("Couldn't set a kernel argument");
      return;   
   };
   cl.CLSetKernelArg(kernel, 1, cl.CL_size_t(unsafe.Sizeof(angles_buffer)), unsafe.Pointer(&angles_buffer));
   cl.CLSetKernelArg(kernel, 2, cl.CL_size_t(unsafe.Sizeof(x_coords_buffer)), unsafe.Pointer(&x_coords_buffer));
   cl.CLSetKernelArg(kernel, 3, cl.CL_size_t(unsafe.Sizeof(y_coords_buffer)), unsafe.Pointer(&y_coords_buffer));
   
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
   err = cl.CLEnqueueReadBuffer(queue, x_coords_buffer, cl.CL_TRUE, 0, 
      cl.CL_size_t(unsafe.Sizeof(x_coords)), unsafe.Pointer(&x_coords), 0, nil, nil);
   if err < 0  {
      println ("Couldn't read the buffer");
      return;   
   }
   cl.CLEnqueueReadBuffer(queue, y_coords_buffer, cl.CL_TRUE, 0, 
      cl.CL_size_t(unsafe.Sizeof(y_coords)), unsafe.Pointer(&y_coords), 0, nil, nil);   

   /* Display the results */
   for i:=0; i<4; i++ {
      fmt.Printf("(%6.3f, %6.3f)\n", x_coords[i], y_coords[i]);
   }   
      
   /* Deallocate resources */
   cl.CLReleaseMemObject(r_coords_buffer);
   cl.CLReleaseMemObject(angles_buffer);  
   cl.CLReleaseMemObject(x_coords_buffer);
   cl.CLReleaseMemObject(y_coords_buffer);   
   cl.CLReleaseKernel(kernel);
   cl.CLReleaseCommandQueue(queue);
   cl.CLReleaseProgram(*program);
   cl.CLReleaseContext(context);
}
