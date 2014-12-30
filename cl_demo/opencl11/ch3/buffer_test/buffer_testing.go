package main

import (
   "fmt"
   "gocl/cl"
   "unsafe"
   "gocl/cl_demo/utils"
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
   var full_matrix, zero_matrix [80]float32;
   var sizeoffloat32 = cl.CL_size_t(unsafe.Sizeof(full_matrix[0])) 
   var buffer_origin =[3]cl.CL_size_t{5*sizeoffloat32, 3, 0};
   var host_origin =[3]cl.CL_size_t {1*sizeoffloat32, 1, 0};
   var region =[3]cl.CL_size_t{4*sizeoffloat32, 4, 1};
   var matrix_buffer cl.CL_mem;

   /* Initialize data */
   for i:=0; i<80; i++ {
      full_matrix[i] = float32(i)*1.0;
      zero_matrix[i] = 0.0;
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
   if program==nil{
      println("Couldn't build program")
      return
   }

   kernel = cl.CLCreateKernel(*program, []byte(KERNEL_FUNC), &err);
   if err < 0 {
      println("Couldn't create a kernel");
      return;   
   }

   /* Create a buffer to hold 80 floats */
   matrix_buffer = cl.CLCreateBuffer(context, cl.CL_MEM_READ_WRITE | 
      cl.CL_MEM_COPY_HOST_PTR, cl.CL_size_t(unsafe.Sizeof(full_matrix)), unsafe.Pointer(&full_matrix[0]), &err);
   if err < 0 {
      println("Couldn't create a buffer object");
      return;   
   }

   /* Set buffer as argument to the kernel */
   err = cl.CLSetKernelArg(kernel, 0, cl.CL_size_t(unsafe.Sizeof(matrix_buffer)), unsafe.Pointer(&matrix_buffer));
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

   /* Enqueue command to write to buffer */
   err = cl.CLEnqueueWriteBuffer(queue, matrix_buffer, cl.CL_TRUE, 0,
         cl.CL_size_t(unsafe.Sizeof(full_matrix)), unsafe.Pointer(&full_matrix[0]), 0, nil, nil); 
   if err < 0 {
      println("Couldn't write to the buffer object");
      return;   
   }

   /* Enqueue command to read rectangle of data */
   err = cl.CLEnqueueReadBufferRect(queue, matrix_buffer, cl.CL_TRUE, 
         buffer_origin, host_origin, region, 10*sizeoffloat32, 0, 
         10*sizeoffloat32, 0, unsafe.Pointer(&zero_matrix[0]), 0, nil, nil);
   if err < 0 {
      println("Couldn't read the rectangle from the buffer object");
      return;   
   }

   /* Display updated buffer */
   for i:=0; i<8; i++ {
      for j:=0; j<10; j++ {
         fmt.Printf("%6.1f", zero_matrix[j+i*10]);
      }
      fmt.Printf("\n");
   }

   /* Deallocate resources */
   cl.CLReleaseMemObject(matrix_buffer);
   cl.CLReleaseKernel(kernel);
   cl.CLReleaseCommandQueue(queue);
   cl.CLReleaseProgram(*program);
   cl.CLReleaseContext(context);
}