__kernel void polar_rect(__global float4 *r_vals, 
                         __global float4 *angles,
                         __global float4 *x_coords, 
                         __global float4 *y_coords) {

   *y_coords = sincos(*angles, x_coords);
   *x_coords *= *r_vals;
   *y_coords *= *r_vals;
}
