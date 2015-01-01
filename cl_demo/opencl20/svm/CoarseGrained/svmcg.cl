// Array of the structures defined below is built and populated
// with the random values on the host.
// Then it is traversed in the OpenCL kernel on the device.
typedef struct _Element
{
    global float* internal; //points to the "value" of another Element from the same array
    global float* external; //points to the entry in a separate array of floating-point values
    float value;
} Element;

kernel void svmbasic (global Element* elements, global float *dst)
{
    int id = (int)get_global_id(0);

    float internalElement = *(elements[id].internal);
    float externalElement = *(elements[id].external);
    dst[id] = internalElement + externalElement;
}
