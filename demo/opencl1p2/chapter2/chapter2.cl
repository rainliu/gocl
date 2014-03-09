// OpenCL kernel to perform an element-wise 
// add of two arrays     

__kernel 
void vecadd(__global int *A,                        
            __global int *B,                        
            __global int *C)                        
{                                                   
                                                    
   // Get the work-item’s unique ID                 
   int idx = get_global_id(0);                      
                                                    
   // Add the corresponding locations of            
   // 'A' and 'B', and store the result in 'C'.     
   C[idx] = A[idx] + B[idx];                        
}                                                   