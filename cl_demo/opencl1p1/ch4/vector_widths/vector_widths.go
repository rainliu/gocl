package main

import (
   "fmt"
   "gocl/cl"
   "unsafe"
)

func main() {

   /* Host/device data structures */
   var platform [1]cl.CL_platform_id;
   var device [1]cl.CL_device_id;
   var sizeofuint cl.CL_uint;
   var vector_width interface{}
   var err cl.CL_int;

   /* Identify a platform */
   err = cl.CLGetPlatformIDs(1, platform[:], nil );
   if err < 0  {
      println ("Couldn't find any platforms");
      return
   }

   /* Access a device */
   err = cl.CLGetDeviceIDs(platform[0], cl.CL_DEVICE_TYPE_GPU, 1, device[:], nil );
   if err == cl.CL_DEVICE_NOT_FOUND {
      err = cl.CLGetDeviceIDs(platform[0], cl.CL_DEVICE_TYPE_CPU, 1, device[:], nil );
   }
   if err < 0  {
      println ("Couldn't access any devices");
      return   
   }
   
   /* Obtain the device data */
   err = cl.CLGetDeviceInfo(device[0], cl.CL_DEVICE_PREFERRED_VECTOR_WIDTH_CHAR, 
         cl.CL_size_t(unsafe.Sizeof(sizeofuint)), &vector_width, nil );     
   if err < 0  {
      println ("Couldn't read device properties");
      return
   }
   fmt.Printf("Preferred vector width in chars: %v\n", vector_width.(cl.CL_uint));
   cl.CLGetDeviceInfo(device[0], cl.CL_DEVICE_PREFERRED_VECTOR_WIDTH_SHORT, 
         cl.CL_size_t(unsafe.Sizeof(sizeofuint)), &vector_width, nil );      
   fmt.Printf("Preferred vector width in shorts: %v\n", vector_width.(cl.CL_uint));
   cl.CLGetDeviceInfo(device[0], cl.CL_DEVICE_PREFERRED_VECTOR_WIDTH_INT, 
         cl.CL_size_t(unsafe.Sizeof(sizeofuint)), &vector_width, nil );      
   fmt.Printf("Preferred vector width in ints: %v\n", vector_width.(cl.CL_uint));
   cl.CLGetDeviceInfo(device[0], cl.CL_DEVICE_PREFERRED_VECTOR_WIDTH_LONG, 
         cl.CL_size_t(unsafe.Sizeof(sizeofuint)), &vector_width, nil );      
   fmt.Printf("Preferred vector width in longs: %v\n", vector_width.(cl.CL_uint));
   cl.CLGetDeviceInfo(device[0], cl.CL_DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT, 
         cl.CL_size_t(unsafe.Sizeof(sizeofuint)), &vector_width, nil );      
   fmt.Printf("Preferred vector width in floats: %v\n", vector_width.(cl.CL_uint));
   cl.CLGetDeviceInfo(device[0], cl.CL_DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE, 
         cl.CL_size_t(unsafe.Sizeof(sizeofuint)), &vector_width, nil );      
   fmt.Printf("Preferred vector width in doubles: %v\n", vector_width.(cl.CL_uint));
         
   cl.CLGetDeviceInfo(device[0], cl.CL_DEVICE_PREFERRED_VECTOR_WIDTH_HALF, 
         cl.CL_size_t(unsafe.Sizeof(sizeofuint)), &vector_width, nil );      
   fmt.Printf("Preferred vector width in halfs: %v\n", vector_width.(cl.CL_uint));
}
