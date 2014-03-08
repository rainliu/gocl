package main

import (
   "fmt"
   "gocl/cl"
   "unsafe"
   "gocl/demo/ch0"
   "time"
   "bufio"
   "os"
)

const PROGRAM_FILE =       "user_event.cl"
var    KERNEL_FUNC =[]byte("user_event")


func read_complete(e cl.CL_event, status cl.CL_int, data unsafe.Pointer) {
   var float_data []float32
   float_data = *(*[]float32)(data);
   fmt.Printf("New data: %4.2f, %4.2f, %4.2f, %4.2f\n", 
      float_data[0], float_data[1], float_data[2], float_data[3]);
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
   var data []float32;
   var data_buffer cl.CL_mem;
   var user_event, kernel_event, read_event [1]cl.CL_event;
   
   /* Initialize data */
   data = make([]float32, 4)
   for i:=0; i<4; i++{
      data[i] = float32(i) * 1.0;
   }

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
         cl.CL_size_t(unsafe.Sizeof(data[0]))*4, unsafe.Pointer(&data[0]), &err);
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
   queue = cl.CLCreateCommandQueue(context, device[0], 
         cl.CL_QUEUE_OUT_OF_ORDER_EXEC_MODE_ENABLE, &err);
   if err < 0 {
      println("Couldn't create a command queue");
      return   
   };

   /* Configure events */
   user_event[0] = cl.CLCreateUserEvent(context, &err);
   if err < 0 {
      println("Couldn't enqueue the kernel");
      return   
   }

   /* Enqueue kernel */
   err = cl.CLEnqueueTask(queue, kernel, 1, user_event[:], &kernel_event[0]);
   if err < 0 {
      println("Couldn't enqueue the kernel");
      return   
   }

   /* Read the buffer */
   err = cl.CLEnqueueReadBuffer(queue, data_buffer, cl.CL_FALSE, 0, 
      cl.CL_size_t(unsafe.Sizeof(data[0]))*4, unsafe.Pointer(&data[0]), 1, kernel_event[:], &read_event[0]);
   if err < 0 {
      println("Couldn't read the buffer");
      return
   }

   /* Set callback for event */
   err = cl.CLSetEventCallback(read_event[0], cl.CL_COMPLETE, 
         read_complete, unsafe.Pointer(&data));
   if err < 0 {
      println("Couldn't set callback for event");
      return   
   }

   /* Sleep for a second to demonstrate the that commands haven't
      started executing. Then prompt user */
   time.Sleep(1);
   fmt.Printf("Old data: %4.2f, %4.2f, %4.2f, %4.2f\n", 
      data[0], data[1], data[2], data[3]);
   fmt.Printf("Press ENTER to continue.\n");
   //getchar();
   reader := bufio.NewReader(os.Stdin)
   reader.ReadString('\n')

   /* Set user event to success */
   cl.CLSetUserEventStatus(user_event[0], cl.CL_SUCCESS);

   /* Deallocate resources */
   cl.CLReleaseEvent(read_event[0]);
   cl.CLReleaseEvent(kernel_event[0]);
   cl.CLReleaseEvent(user_event[0]);
   cl.CLReleaseMemObject(data_buffer);
   cl.CLReleaseKernel(kernel);
   cl.CLReleaseCommandQueue(queue);
   cl.CLReleaseProgram(*program);
   cl.CLReleaseContext(context);
}
