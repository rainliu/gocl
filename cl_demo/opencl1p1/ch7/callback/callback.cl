__kernel void callback(__global float *buffer) {
   float4 five_vector = (float4)(5.0);

   for(int i=0; i<1024; i++) {
      vstore4(five_vector, i, buffer);
   }
}
