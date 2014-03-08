package main

/*
#include <string.h>
*/
import "C"

import (
   "fmt"
   "gocl/cl"
   "unsafe"
   "gocl/demo/ch0"
)

const PROGRAM_FILE = "blank.cl"
const KERNEL_FUNC = "blank"

func main() {

   /* OpenCL data structures */
   var device []cl.CL_device_id;
   var context cl.CL_context;
   var queue cl.CL_command_queue;
   var program *cl.CL_program;
   var kernel cl.CL_kernel;
   var err cl.CL_int ;

   /* Data and buffers */
   var data_one, data_two, result_array [100]float32;
   var buffer_one, buffer_two cl.CL_mem;
   var mapped_memory unsafe.Pointer;

   /* Initialize arrays */
   for i:=0; i<100; i++ {
      data_one[i] = 1.0*float32(i);
      data_two[i] = -1.0*float32(i);
      result_array[i] = 0.0;
   }

   /* Create a device and context */
   device = ch0.Create_device();
   context = cl.CLCreateContext(nil, 1, device[:], nil, nil, &err);
   if err < 0 {
      println("Couldn't create a context");
      return;   
   }

   /* Build the program and create the kernel */
   program = ch0.Build_program(context, device[:], PROGRAM_FILE, nil);
   kernel = cl.CLCreateKernel(*program, []byte(KERNEL_FUNC), &err);
   if err < 0 {
      println("Couldn't create a kernel");
      return;   
   };

   /* Create buffers */
   buffer_one = cl.CLCreateBuffer(context, cl.CL_MEM_READ_WRITE | 
         cl.CL_MEM_COPY_HOST_PTR, cl.CL_size_t(unsafe.Sizeof(data_one)), unsafe.Pointer(&data_one[0]), &err);
   if err < 0 {
      println("Couldn't create buffer object 1");
      return;   
   }
   buffer_two = cl.CLCreateBuffer(context, cl.CL_MEM_READ_WRITE | 
         cl.CL_MEM_COPY_HOST_PTR, cl.CL_size_t(unsafe.Sizeof(data_two)), unsafe.Pointer(&data_two), &err);
   if err < 0 {
      println("Couldn't create buffer object 2");
      return;   
   }
   /* Set buffers as arguments to the kernel */
   err  = cl.CLSetKernelArg(kernel, 0, cl.CL_size_t(unsafe.Sizeof(buffer_one)), unsafe.Pointer(&buffer_one));
   err |= cl.CLSetKernelArg(kernel, 1, cl.CL_size_t(unsafe.Sizeof(buffer_two)), unsafe.Pointer(&buffer_two));
   if err < 0 {
      println("Couldn't set the buffer as the kernel argument");
      return;   
   }

   /* Create a command queue */
   queue = cl.CLCreateCommandQueue(context, device[0], 0, &err);
   if err < 0 {
      println("Couldn't create a command queue");
      return;   
   };

   /* Enqueue kernel */
   err = cl.CLEnqueueTask(queue, kernel, 0, nil, nil);
   if err < 0 {
      println("Couldn't enqueue the kernel");
      return;   
   }

   /* Enqueue command to copy buffer one to buffer two */
   err = cl.CLEnqueueCopyBuffer(queue, buffer_one, buffer_two, 0, 0,
         cl.CL_size_t(unsafe.Sizeof(data_one)), 0, nil, nil); 
   if err < 0 {
      println("Couldn't perform the buffer copy");
      return;   
   }

   /* Enqueue command to map buffer two to host memory */
   mapped_memory = cl.CLEnqueueMapBuffer(queue, buffer_two, cl.CL_TRUE,
         cl.CL_MAP_READ, 0, cl.CL_size_t(unsafe.Sizeof(data_two)), 0, nil, nil, &err);
   if err < 0 {
      println("Couldn't map the buffer to host memory");
      return;   
   }

   /* Transfer memory and unmap the buffer */
   C.memcpy(unsafe.Pointer(&result_array[0]), mapped_memory, C.size_t(unsafe.Sizeof(data_two)));
   err = cl.CLEnqueueUnmapMemObject(queue, buffer_two, mapped_memory,
         0, nil, nil);
   if err < 0 {
      println("Couldn't unmap the buffer");
      return;   
   }

   /* Display updated buffer */
   for i:=0; i<10; i++ {
      for j:=0; j<10; j++ {
         fmt.Printf("%6.1f", result_array[j+i*10]);
      }
      fmt.Printf("\n");
   }

   /* Deallocate resources */
   cl.CLReleaseMemObject(buffer_one);
   cl.CLReleaseMemObject(buffer_two);
   cl.CLReleaseKernel(kernel);
   cl.CLReleaseCommandQueue(queue);
   cl.CLReleaseProgram(*program);
   cl.CLReleaseContext(context);
}
