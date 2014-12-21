__kernel void op_test(__global int4 *output) {

   int4 vec = (int4)(1, 2, 3, 4);
   
   /* Adds 4 to every element of vec */
   vec += 4;
   
   /* Sets the third element to 0
      Doesn't change the other elements 
      (-1 in hexadecimal = 0xFFFFFFFF */
   if(vec.s2 == 7)
      vec &= (int4)(-1, -1, 0, -1);
   
   /* Sets the first element to -1, the second to 0 */
   vec.s01 = vec.s23 < 7; 
   
   /* Divides the last element by 2 until it is less than or equal to 7 */
   while(vec.s3 > 7 && (vec.s0 < 16 || vec.s1 < 16))
      vec.s3 >>= 1; 
      
   *output = vec;
}
