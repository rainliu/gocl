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
   var flag interface{}//cl.CL_device_fp_config;
   var err cl.CL_int;

   /* Identify a platform */
   err = cl.CLGetPlatformIDs(1, platform[:], nil );
   if err < 0  {
      println ("Couldn't identify a platform");
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

   /* Check float-processing features */
   err = cl.CLGetDeviceInfo(device[0], cl.CL_DEVICE_SINGLE_FP_CONFIG, 
         cl.CL_size_t(unsafe.Sizeof(flag)), &flag, nil );
   if err < 0  {
      println ("Couldn't read floating-point properties");
      return 
   }
   fmt.Printf ("Float Processing Features:\n");
   if (flag.(cl.CL_device_fp_config)  & cl.CL_FP_INF_NAN) > 0{
      fmt.Printf ("INF and NaN values supported.\n");
   }
   if (flag.(cl.CL_device_fp_config)  & cl.CL_FP_DENORM) > 0{
      fmt.Printf ("Denormalized numbers supported.\n");
   }
   if (flag.(cl.CL_device_fp_config)  & cl.CL_FP_ROUND_TO_NEAREST) > 0{ 
      fmt.Printf ("Round To Nearest Even mode supported.\n");
   }
   if (flag.(cl.CL_device_fp_config)  & cl.CL_FP_ROUND_TO_INF) > 0{
      fmt.Printf ("Round To Infinity mode supported.\n");
   }
   if (flag.(cl.CL_device_fp_config)  & cl.CL_FP_ROUND_TO_ZERO) > 0{
      fmt.Printf ("Round To Zero mode supported.\n");
   }
   if (flag.(cl.CL_device_fp_config)  & cl.CL_FP_FMA) > 0{
      fmt.Printf ("Floating-point multiply-and-add operation supported.\n");
   }
   if (flag.(cl.CL_device_fp_config)  & cl.CL_FP_SOFT_FLOAT) > 0{
      fmt.Printf ("Basic floating-point processing performed in software.\n");
   }
}
