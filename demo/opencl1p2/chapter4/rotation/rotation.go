package main

import (
  "log"
  "math"
   "os"
   "fmt"
    "unsafe"
    "gocl/cl"
    "gocl/demo/ch0"
)

func chk( status cl.CL_int, cmd string) {

   if status != cl.CL_SUCCESS {
      fmt.Printf("%s failed (%d)\n", cmd, status);
      os.Exit(1);
   }
}

func main() {
   var i cl.CL_size_t
   // Set the image rotation (in degrees)
   theta     := float64(3.14159/6);
   cos_theta := float32(math.Cos(theta));
   sin_theta := float32(math.Sin(theta));
   fmt.Printf("theta = %f (cos theta = %f, sin theta = %f)\n", theta, cos_theta, 
      sin_theta);

   inputFile  := "test.png";
   outputFile := "output.png";

   inputpixels, imageWidth, imageHeight, err1:= ch0.Read_image_data(inputFile);
   if err1!=nil{
      log.Fatal(err1)
      return
   }else{
      fmt.Printf("width=%d, height=%d\n", imageWidth, imageHeight)
   }

   // Output image on the host
   outputpixels:= make([]uint16,  imageHeight*imageWidth);
   inputImage  := make([]float32, imageHeight*imageWidth);
   outputImage := make([]float32, imageHeight*imageWidth);

   for i=0; i<imageHeight*imageWidth; i++{
      inputImage[i] = float32(inputpixels[i])
   }

   // Set up the OpenCL environment
   var status cl.CL_int;

   // Discovery platform
   var platform [1]cl.CL_platform_id;
   status = cl.CLGetPlatformIDs(1, platform[:], nil);
   chk(status, "clGetPlatformIDs");

   // Discover device
   var device [1]cl.CL_device_id;
   cl.CLGetDeviceIDs(platform[0], cl.CL_DEVICE_TYPE_ALL, 1, device[:], nil);
   chk(status, "clGetDeviceIDs");

   // Create context
   var context cl.CL_context;
   context = cl.CLCreateContext(nil, 1, device[:], nil, nil, &status);
   chk(status, "clCreateContext");

   // Create command queue
   var queue cl.CL_command_queue;
   queue = cl.CLCreateCommandQueue(context, device[0], 0, &status);
   chk(status, "clCreateCommandQueue");

   dataSize := imageWidth*imageHeight*cl.CL_size_t(unsafe.Sizeof(inputImage[0]))
   // Create the input and output buffers
   d_input := cl.CLCreateBuffer(context, cl.CL_MEM_READ_ONLY, dataSize, nil,
       &status);
   chk(status, "clCreateBuffer");

   d_output := cl.CLCreateBuffer(context, cl.CL_MEM_WRITE_ONLY, dataSize, nil,
       &status);
   chk(status, "clCreateBuffer");

   // Copy the input image to the device
   status = cl.CLEnqueueWriteBuffer(queue, d_input, cl.CL_TRUE, 0, dataSize, 
         unsafe.Pointer(&inputImage[0]), 0, nil, nil);
   chk(status, "clEnqueueWriteBuffer");

   // Create a program object with source and build it
   program := ch0.Build_program(context, device[:], "rotation.cl", nil);
   kernel := cl.CLCreateKernel(*program, []byte("img_rotate"), &status);
   chk(status, "clCreateKernel")

   // Set the kernel arguments
   w := cl.CL_int(imageWidth)
   h := cl.CL_int(imageHeight)
   status  = cl.CLSetKernelArg(kernel, 0, cl.CL_size_t(unsafe.Sizeof(d_output)),  unsafe.Pointer(&d_output));
   status |= cl.CLSetKernelArg(kernel, 1, cl.CL_size_t(unsafe.Sizeof(d_input)),   unsafe.Pointer(&d_input));
   status |= cl.CLSetKernelArg(kernel, 2, cl.CL_size_t(unsafe.Sizeof(w)),         unsafe.Pointer(&w));
   status |= cl.CLSetKernelArg(kernel, 3, cl.CL_size_t(unsafe.Sizeof(h)),         unsafe.Pointer(&h));
   status |= cl.CLSetKernelArg(kernel, 4, cl.CL_size_t(unsafe.Sizeof(sin_theta)), unsafe.Pointer(&sin_theta));
   status |= cl.CLSetKernelArg(kernel, 5, cl.CL_size_t(unsafe.Sizeof(cos_theta)), unsafe.Pointer(&cos_theta));
   chk(status, "clSetKernelArg");

   // Set the work item dimensions
   var globalSize =[2]cl.CL_size_t{imageWidth, imageHeight};
   status = cl.CLEnqueueNDRangeKernel(queue, kernel, 2, nil, globalSize[:], nil, 0,
      nil, nil);
   chk(status, "clEnqueueNDRange");

   // Read the image back to the host
   status = cl.CLEnqueueReadBuffer(queue, d_output, cl.CL_TRUE, 0, dataSize, 
         unsafe.Pointer(&outputImage[0]), 0, nil, nil); 
   chk(status, "clEnqueueReadBuffer");

   // Write the output image to file
  for i=0; i<imageHeight*imageWidth; i++{
      outputpixels[i]=uint16(outputImage[i])
   }
   ch0.Write_image_data(outputFile, outputpixels, imageWidth, imageHeight);

   // Free OpenCL resources
    cl.CLReleaseKernel(kernel);
    cl.CLReleaseProgram(*program);
    cl.CLReleaseCommandQueue(queue);
    cl.CLReleaseMemObject(d_input);
    cl.CLReleaseMemObject(d_output);
    cl.CLReleaseContext(context);
}
