Pipe Producer-Consumer Kernels

This sample demonstrates the Pipe as a data-sharing FIFO for a producer kernel and a consumer kernel

Prerequisite:

	Graphics Driver : 
		- Install AMD OpenCL 2.0 Driver located here http://support.amd.com/en-us/kb-articles/Pages/OpenCL2-Driver.aspx

	Compatible Operating Systems:
		- Microsoft Windows:
			Windows 8.1 (64-bit version)
		- Linux Distributions:
			Red Hat Enterprise 6.5 (64-bit version)
			Ubuntu 14.04 (64-bit version)

	Compatible Hardware: 
		- Check the AMD Product compatibility from the driver download page here http://support.amd.com/en-us/kb-articles/Pages/OpenCL2-Driver.aspx
		
	Software
		- Windows : Visual Studio 2013
		- Linux : CMake

How to Compile

	Windows:
		- The zip file contains Visual Studio 2013 project file of the sample. Open the project file and compile it in 64-bit configuration
		- The zip file also contains CMakeLists.txt. This can be used to generate project files of any other versions of Visual Studio available on your machine
	Linux:
		- The zip file contains CMakeLists.txt. Use this to generate make files.
			$> cmake 
		- Compile the sample using the make file
			$> make

Run the sample

	Execute the sample typing the following command
		Windows : $> PipeProducerConsumerKernels.exe 
		Linux :	  $> ./PipeProducerConsumerKernels
		
	Command Line Options:
	-h,   	--help                                            Display this information
	        --device        [cpu|gpu]                         Execute the openCL kernel on a device
	-q,   	--quiet                                           Quiet mode. Suppress all text output.
	-e,   	--verify                                          Verify results against reference implementation.
	-t,   	--timing                                          Print timing.
	        --dump          [filename]                        Dump binary image for all devices
	        --load          [filename]                        Load binary image and execute on device
	        --flags         [filename]                        Specify filename containing the compiler flags to build kernel
	-p,   	--platformId    [value]                           Select platformId to be used[0 to N-1 where N is number platforms available].
	-v,   	--version                                         AMD APP SDK version string.
	-d,   	--deviceId      [value]                           Select deviceId to be used[0 to N-1 where N is number devices available].
	-i,   	--iterations                                      Number of iterations to execute kernel
	-s,   	--pipeSize                                        Pipe size
	-b,   	--seed                                            Seed for random number generator
