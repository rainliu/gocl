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

// This function takes a positive integer and rounds it up to
// the nearest multiple of another provided integer
func roundUp( value,  multiple uint32) uint32{

  // Determine how far past the nearest multiple the value is
  remainder := value % multiple;

  // Add the difference to make the value a multiple
  if(remainder != 0) {
      value += (multiple-remainder);
  }

  return value;
}


func chk( status cl.CL_int, cmd string) {

   if status != cl.CL_SUCCESS {
      fmt.Printf("%s failed (%d)\n", cmd, status);
      os.Exit(1);
   }
}

func main() {
   var i, j cl.CL_size_t
   // Rows and columns in the input image
   inputFile  := "test.png";
   outputFile := "output.png";
   refFile    := "ref.png";

   // Homegrown function to read a BMP from file
   inputpixels, imageWidth, imageHeight, err1:= ch0.Read_image_data(inputFile);
   if err1!=nil{
      log.Fatal(err1)
      return
   }else{
      fmt.Printf("width=%d, height=%d (%d)\n", imageWidth, imageHeight, inputpixels[0])
   }

   // Output image on the host
   outputpixels:= make([]uint16,  imageHeight*imageWidth);
   inputImage  := make([]float32, imageHeight*imageWidth);
   outputImage := make([]float32, imageHeight*imageWidth);
   refImage    := make([]float32, imageHeight*imageWidth);

   for i=0; i<imageHeight*imageWidth; i++{
      inputImage[i] = float32(inputpixels[i])
   }

   // 45 degree motion blur
   var filter =[49]float32{0,      0,      0,      0,      0,      0,      0,
       0,      0,      0,      0,      0,      0,      0,
       0,      0,     -1,      0,      1,      0,      0,
       0,      0,     -2,      0,      2,      0,      0,
       0,      0,     -1,      0,      1,      0,      0,
       0,      0,      0,      0,      0,      0,      0,
       0,      0,      0,      0,      0,      0,      0};

   // The convolution filter is 7x7
   filterWidth := cl.CL_size_t(7);  
   filterSize  := cl.CL_size_t(filterWidth*filterWidth);  // Assume a square kernel

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
   //var props =[3]cl.CL_context_properties{cl.CL_CONTEXT_PLATFORM,
   //    (cl.CL_context_properties)(unsafe.Pointer(&platform[0])), 0};

   var context cl.CL_context;
   context = cl.CLCreateContext(nil, 1, device[:], nil, nil, &status);
   chk(status, "clCreateContext");

   // Create command queue
   var queue cl.CL_command_queue;
   queue = cl.CLCreateCommandQueue(context, device[0], 0, &status);
   chk(status, "clCreateCommandQueue");

   // The image format describes how the data will be stored in memory
   var format cl.CL_image_format;
   format.Image_channel_order     = cl.CL_R;     // single channel
   format.Image_channel_data_type = cl.CL_FLOAT; // float data type

   var desc cl.CL_image_desc
  desc.Image_type       = cl.CL_MEM_OBJECT_IMAGE2D
  desc.Image_width      = imageWidth;
  desc.Image_height     = imageHeight;
  desc.Image_depth      = 0;
  desc.Image_array_size = 0;
  desc.Image_row_pitch  = 0;
  desc.Image_slice_pitch= 0;
  desc.Num_mip_levels   = 0;
  desc.Num_samples      = 0;
  desc.Buffer           = cl.CL_mem{};

   // Create space for the source image on the device
   d_inputImage := cl.CLCreateImage(context, cl.CL_MEM_READ_ONLY, &format, &desc, 
      nil, &status);
   chk(status, "clCreateImage");

   // Create space for the output image on the device
   d_outputImage := cl.CLCreateImage(context, cl.CL_MEM_WRITE_ONLY, &format, &desc, 
      nil, &status);
   chk(status, "clCreateImage");

   // Create space for the 7x7 filter on the device
   d_filter := cl.CLCreateBuffer(context, 0, filterSize*cl.CL_size_t(unsafe.Sizeof(filter[0])), 
      nil, &status);
   chk(status, "clCreateBuffer");

   // Copy the source image to the device
   var origin =[3]cl.CL_size_t{0, 0, 0};  // Offset within the image to copy from
   var region =[3]cl.CL_size_t{cl.CL_size_t(imageWidth), cl.CL_size_t(imageHeight), 1}; // Elements to per dimension
   status = cl.CLEnqueueWriteImage(queue, d_inputImage, cl.CL_FALSE, origin, region, 
      0, 0, unsafe.Pointer(&inputImage[0]), 0, nil, nil);
   chk(status, "clEnqueueWriteImage");
    
   // Copy the 7x7 filter to the device
   status = cl.CLEnqueueWriteBuffer(queue, d_filter, cl.CL_FALSE, 0, 
      filterSize*cl.CL_size_t(unsafe.Sizeof(filter[0])), unsafe.Pointer(&filter[0]), 0, nil, nil);
   chk(status, "clEnqueueWriteBuffer");

   // Create the image sampler
   sampler := cl.CLCreateSampler(context, cl.CL_FALSE, 
      cl.CL_ADDRESS_CLAMP_TO_EDGE, cl.CL_FILTER_NEAREST, &status);
   chk(status, "clCreateSampler");

   // Create a program object with source and build it
   program := ch0.Build_program(context, device[:], "convolution.cl", nil);
   kernel := cl.CLCreateKernel(*program, []byte("convolution"), &status);
   chk(status, "clCreateKernel")

   // Set the kernel arguments
   var w,h,f cl.CL_int
   w = cl.CL_int(imageWidth);
   h = cl.CL_int(imageHeight);
   f = cl.CL_int(filterWidth);
   status  = cl.CLSetKernelArg(kernel, 0, cl.CL_size_t(unsafe.Sizeof(d_inputImage)),   unsafe.Pointer(&d_inputImage));
   status |= cl.CLSetKernelArg(kernel, 1, cl.CL_size_t(unsafe.Sizeof(d_outputImage)),  unsafe.Pointer(&d_outputImage));
   status |= cl.CLSetKernelArg(kernel, 2, cl.CL_size_t(unsafe.Sizeof(h)),              unsafe.Pointer(&h));
   status |= cl.CLSetKernelArg(kernel, 3, cl.CL_size_t(unsafe.Sizeof(w)),              unsafe.Pointer(&w));
   status |= cl.CLSetKernelArg(kernel, 4, cl.CL_size_t(unsafe.Sizeof(d_filter)),       unsafe.Pointer(&d_filter));
   status |= cl.CLSetKernelArg(kernel, 5, cl.CL_size_t(unsafe.Sizeof(f)),              unsafe.Pointer(&f));
   status |= cl.CLSetKernelArg(kernel, 6, cl.CL_size_t(unsafe.Sizeof(sampler)),        unsafe.Pointer(&sampler));
   chk(status, "clSetKernelArg");

   // Set the work item dimensions
   var globalSize =[2]cl.CL_size_t{imageWidth, imageHeight};
   status = cl.CLEnqueueNDRangeKernel(queue, kernel, 2, nil, globalSize[:], nil, 0,
      nil, nil);
   chk(status, "clEnqueueNDRange");

   // Read the image back to the host
   status = cl.CLEnqueueReadImage(queue, d_outputImage, cl.CL_TRUE, origin, 
      region, 0, 0, unsafe.Pointer(&outputImage[0]), 0, nil, nil); 
   chk(status, "clEnqueueReadImage");

   // Write the output image to file
  for i=0; i<imageHeight*imageWidth; i++{
      outputpixels[i]=uint16(outputImage[i])
   }
   ch0.Write_image_data(outputFile, outputpixels, imageWidth, imageHeight);

   // Compute the reference image
   for i = 0; i < imageHeight; i++ {
      for j = 0; j < imageWidth; j++ {
         refImage[i*imageWidth+j] = 0;
      }
   }

   // Iterate over the rows of the source image
   halfFilterWidth := filterWidth/2;
   var sum float32;
   for i = 0; i < imageHeight; i++ {
      // Iterate over the columns of the source image
      for j = 0; j < imageWidth; j++ {
         sum = 0; // Reset sum for new source pixel
         // Apply the filter to the neighborhood
         for k := - halfFilterWidth; k <= halfFilterWidth; k++ {
            for l := - halfFilterWidth; l <= halfFilterWidth; l++ {
               if i+k >= 0 && i+k < imageHeight && 
                  j+l >= 0 && j+l < imageWidth {
                  sum += inputImage[(i+k)*imageWidth + j+l] * 
                         filter[(k+halfFilterWidth)*filterWidth + 
                            l+halfFilterWidth];
               }else{
                  i_k := i+k;
                  j_l := j+l;
                  if i+k < 0 {
                    i_k = 0
                  }else if i+k>=imageHeight{
                    i_k = imageHeight-1
                  }
                  if j+l < 0 {
                    j_l = 0
                  }else if j+l >=imageWidth{
                    j_l = imageWidth-1
                  }
                  sum += inputImage[(i_k)*imageWidth + j_l] * 
                         filter[(k+halfFilterWidth)*filterWidth + 
                            l+halfFilterWidth];
               }
            } 
         }
         refImage[i*imageWidth+j] = sum;
      }
   }
 // Write the ref image to file
  for i=0; i<imageHeight*imageWidth; i++{
      outputpixels[i]=uint16(refImage[i])
   }
   ch0.Write_image_data(refFile, outputpixels, imageWidth, imageHeight);

   failed := 0;
   for i = 0; i < imageHeight; i++ {
      for j = 0; j < imageWidth; j++ {
         if math.Abs(float64(outputImage[i*imageWidth+j]-refImage[i*imageWidth+j])) > 0.01 {
            //fmt.Printf("Results are INCORRECT\n");
            //fmt.Printf("Pixel mismatch at <%d,%d> (%f vs. %f) %f\n", i, j,
            //   outputImage[i*imageWidth+j], refImage[i*imageWidth+j], inputImage[i*imageWidth+j]);
            failed++;
         }
      }
   }
   fmt.Printf("Mismatch Pixel number/Total pixel number = %d/%d\n", failed, imageWidth*imageHeight);
}
