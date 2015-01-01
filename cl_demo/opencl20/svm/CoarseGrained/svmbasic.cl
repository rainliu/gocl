#include "svmbasic.h"

kernel void svmbasic (global Element* elements, global float *dst)
{
    int id = (int)get_global_id(0);

    float internalElement = *(elements[id].internal);
    float externalElement = *(elements[id].external);
    dst[id] = internalElement + externalElement;
}
