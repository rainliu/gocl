// Copyright (c) 2014 Intel Corporation
// All rights reserved.
// 
// WARRANTY DISCLAIMER
// 
// THESE MATERIALS ARE PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL INTEL OR ITS
// CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
// EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
// PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY
// OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY OR TORT (INCLUDING
// NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THESE
// MATERIALS, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
// 
// Intel Corporation is the author of the Materials, and requests that all
// problem reports or change requests be submitted to it directly


#include <cstdlib>
#include <iostream>
#include <iomanip>
#include <cstring>
#include <cassert>
#include <exception>

#include <CL/cl.h>

#include "basic.hpp"
#include "cmdparser.hpp"
#include "oclobject.hpp"


// The following piece of code declares the data structure (in file svmbasic.h)
// in a way it is the same on the host and device sides.
// To be used in the OpenCL kernels, the pointers should be defined with 'global' keyword,
// according to OpenCL specification.
// But this keyword is redundant for the host, so we define it as empty.

// Array of the structures defined below is built and populated
// with the random values on the host.
// Then it is traversed in the OpenCL kernel on the device.
typedef struct _Element
{
    float* internal; //points to the "value" of another Element from the same array
    float* external; //points to the entry in a separate array of floating-point values
    float value;
} Element;

using namespace std;


void svmbasic (
    size_t size,
    cl_context context,
    cl_command_queue queue,
    cl_kernel kernel
)
{
    // Prepare input data as follows.
    // Build two arrays:
    //     - an array that consists of the Element structures
    //       (refer to svmbasic.h for the structure definition)
    //     - an array that consists of the float values
    // 
    // Each structure of the first array has the following pointers:
    //     - 'internal', which points to a 'value' field of another entry
    //       of the same array.
    //     - 'external', which points to a float value from the the
    //       second array.
    //
    // Pointers are set randomly. The structures do not reflect any real usage
    // scenario, but are illustrative for a simple device-side traversal.
    //
    //        Array of Element                        Array of floats
    //           structures
    //
    //    ||====================||
    //    ||    .............   ||                   ||============||
    //    ||    .............   ||<-----+            || .......... ||
    //    ||====================||      |            ||    float   ||
    //    ||   float* internal--||------+            ||    float   ||
    //    ||   float* external--||------------------>||    float   ||
    //    ||   float value <----||------+            || .......... ||
    //    ||====================||      |            || .......... ||
    //    ||    .............   ||      |            ||    float   ||
    //    ||    .............   ||      |            ||    float   ||
    //    ||====================||      |            ||    float   ||
    //    ||====================||      |            ||    float   ||
    //    ||   float* internal--||------+            ||    float   ||
    //    ||   float* external--||------------------>||    float   ||
    //    ||   float value      ||                   ||    float   ||
    //    ||====================||                   ||    float   ||
    //    ||    .............   ||                   || .......... ||
    //    ||    .............   ||                   ||============||
    //    ||====================||
    //
    // The two arrays are created independently and are used to illustrate
    // two new OpenCL 2.0 API functions:
    //    - the array of Element structures is passed to the kernel as a
    //      kernel argument with the clSetKernelArgSVMPointer function
    //    - the array of floats is used by the kernel indirectly, and this
    //      dependency should be also specified with the clSetKernelExecInfo
    //      function prior to the kernel execution

    cl_int err = CL_SUCCESS;
    
    // To enable host & device code to share pointer to the same address space 
    // the arrays should be allocated as SVM memory. Use the clSVMAlloc function 
    // to allocate SVM memory.
    //
    // Optionally, this function allows specifying alignment in bytes as its 
    // last argument. As this basic example doesn't require any _special_ alignment,
    // the following code illustrates requesting default alignment via passing
    // zero value.

    Element* inputElements =
        (Element*)clSVMAlloc(
            context,                // the context where this memory is supposed to be used
            CL_MEM_READ_ONLY | CL_MEM_SVM_FINE_GRAIN_BUFFER,
            size*sizeof(Element),   // amount of memory to allocate (in bytes)
            0                       // alignment in bytes (0 means default)
        );

    float* inputFloats =
        (float*)clSVMAlloc(
            context,                // the context where this memory is supposed to be used
            CL_MEM_READ_ONLY | CL_MEM_SVM_FINE_GRAIN_BUFFER,
            size*sizeof(float),     // amount of memory to allocate (in bytes)
            0                       // alignment in bytes (0 means default)
        );

    // The OpenCL kernel uses the aforementioned input arrays to compute
    // values for the output array.

    float* output =
        (float*)clSVMAlloc(
            context,                // the context where this memory is supposed to be used
            CL_MEM_WRITE_ONLY | CL_MEM_SVM_FINE_GRAIN_BUFFER,
            size*sizeof(float),     // amount of memory to allocate (in bytes)
            0                       // alignment in bytes (0 means default)
    );

    if(!inputElements || !inputFloats || !output)
    {
        throw Error(
            "Cannot allocate SVM memory with clSVMAlloc: "
            "it returns null pointer. "
            "You might be out of memory."
        );
    }

    // Note: in the coarse-grained SVM, mapping of inputElement and inputFloats is
    // needed to do the following initialization. While here, in the fine-grained SVM,
    // it is not necessary.

    // Populate data-structures with initial data.

    for (size_t i = 0;  i < size;  i++)
    {
        inputElements[i].internal = &(inputElements[rand_index(size)].value);
        inputElements[i].external = &(inputFloats[rand_index(size)]);
        inputElements[i].value = float(i);
        inputFloats[i] = float(i + size);
    }

    // Note: in the coarse-grained SVM, unmapping of inputElement and inputFloats is
    // needed before scheduling the kernel for execution. While here, in the fine-grained SVM,
    // it is not necessary.

    // Pass arguments to the kernel.
    // According to the OpenCL 2.0 specification, you need to use a special
    // function to pass a pointer from SVM memory to kernel.

    err = clSetKernelArgSVMPointer(kernel, 0, inputElements);
    SAMPLE_CHECK_ERRORS(err);

    err = clSetKernelArgSVMPointer(kernel, 1, output);
    SAMPLE_CHECK_ERRORS(err);

    // For buffer based SVM (both coarse- and fine-grain) if one SVM buffer 
    // points to memory allocated in another SVM buffer, such allocations
    // should be passed to the kernel via clSetKernelExecInfo.

    err = clSetKernelExecInfo(
        kernel, 
        CL_KERNEL_EXEC_INFO_SVM_PTRS,
        sizeof(inputFloats),
        &inputFloats
    );
    SAMPLE_CHECK_ERRORS(err);

    // Run the kernel.
    cout << "Running kernel..." << flush;
    
    err = clEnqueueNDRangeKernel(
        queue,
        kernel,
        1,
        0, &size, 0,
        0, 0, 0
    );
    SAMPLE_CHECK_ERRORS(err);

    // Note: In the fine-grained SVM, after enqueuing the kernel above, the host application is
    // not blocked from accessing SVM allocations that were passed to the kernel. The host
    // can access the same regions of SVM memory as does the kernel if the kernel and the host
    // read/modify different bytes. If one side (host or device) needs to modify the same bytes
    // that are simultaniously read/modified by another side, atomics operations are usually
    // required to maintain sufficient memory consistency. This sample doesn't use this possibility
    // and the host just waits in clFinish below until the kernel is finished.

    err = clFinish(queue);
    SAMPLE_CHECK_ERRORS(err);

    cout << " DONE.\n";

    // Validate output state for correctness.
    
    // Compare: in the coarse-grained SVM case you need to map the output.
    // Here it is not needed.

    cout << "Checking correctness of the output buffer..." << flush;
    for(size_t i = 0; i < size; i++)
    {
        float expectedValue = *(inputElements[i].internal) + *(inputElements[i].external);
        if(output[i] != expectedValue)
        {
            cout << " FAILED.\n";

            cerr
                << "Mismatch at position " << i
                << ", read " << output[i]
                << ", expected " << expectedValue << "\n";

            throw Error("Validation failed");
        }
    }
    cout << " PASSED.\n";

    // Release all SVM buffers and exit.

    clSVMFree(context, output);
    clSVMFree(context, inputFloats);
    clSVMFree(context, inputElements);
}


bool checkSVMAvailability (cl_device_id device)
{
    cl_device_svm_capabilities caps;

    cl_int err = clGetDeviceInfo(
        device,
        CL_DEVICE_SVM_CAPABILITIES,
        sizeof(cl_device_svm_capabilities),
        &caps,
        0
    );

    // Check for fine-grained buffer SVM type availability:
    return err == CL_SUCCESS && (caps & CL_DEVICE_SVM_FINE_GRAIN_BUFFER);
}


int main (int argc, const char** argv)
{
    try
    {
        // Define and parse command-line arguments.
        CmdParserCommon cmdparser(argc, argv);
        cmdparser.device_type.setDefaultValue("gpu");
        // Additional option to set array/global size:
        CmdOption<size_t> size(cmdparser, 's', "size", "<integer>", "Global size.",  1024*1024);

        cmdparser.parse();

        // Immediately exit if user wanted to see the usage information only.
        if(cmdparser.help.isSet())
        {
            return 0;
        }

        // Create the necessary OpenCL objects up to device queue.
        OpenCLBasic oclobjects(
            cmdparser.platform.getValue(),
            cmdparser.device_type.getValue(),
            cmdparser.device.getValue()
        );

        if(!checkSVMAvailability(oclobjects.device))
        {
            throw Error(
                "Cannot detect fine-grained buffer SVM capabilities on the device. "
                "The device seemingly doesn't support fine-grained buffer SVM."
            );
        }

        // Build kernel.
        OpenCLProgramOneKernel executable(
            oclobjects,
            L"svmbasic.cl",
            "",
            "svmbasic",
            "-I."    // directory to search for #include directives
        );

        // Then call the main sample routine - resource allocations, OpenCL kernel
        // execution, and so on.
        svmbasic(size.getValue(), oclobjects.context, oclobjects.queue, executable.kernel);

        // All resource deallocations happen in destructors of helper objects.

        return 0;
    }
    catch(const CmdParser::Error& error)
    {
        cerr
            << "[ ERROR ] In command line: " << error.what() << "\n"
            << "Run " << argv[0] << " -h for usage info.\n";
        return 1;
    }
    catch(const Error& error)
    {
        cerr << "[ ERROR ] Sample application specific error: " << error.what() << "\n";
        return 1;
    }
    catch(const exception& error)
    {
        cerr << "[ ERROR ] " << error.what() << "\n";
        return 1;
    }
    catch(...)
    {
        cerr << "[ ERROR ] Unknown/internal error happened.\n";
        return 1;
    }
}