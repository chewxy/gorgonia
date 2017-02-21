#include <math.h>

#ifdef __cplusplus
extern "C" {
#endif


__global__ void sqrt64(double* A, int size)
{
	int blockId = blockIdx.x + blockIdx.y * gridDim.x + gridDim.x * gridDim.y * blockIdx.z;
	int idx = blockId * (blockDim.x * blockDim.y * blockDim.z) + (threadIdx.z * (blockDim.x * blockDim.y)) + (threadIdx.y * blockDim.x) + threadIdx.x;
	if (idx >= size) {
		return;
	}
	A[idx] = sqrt(A[idx]);
}
	
#ifdef __cplusplus
}
#endif