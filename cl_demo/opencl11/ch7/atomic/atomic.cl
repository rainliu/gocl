__kernel void atomic(__global int* x) {

   __local int a, b;

   a = 0; 
   b = 0;

   /* Increment without atomic add */
   a++;

   /* Increment with atomic add */
   atomic_inc(&b);

   x[0] = a;
   x[1] = b;
}
