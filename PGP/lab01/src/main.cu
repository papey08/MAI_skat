#include <stdio.h>
// #include <chrono>

const int THREADS_AMOUNT = 32;
const int BLOCKS_AMOUNT = 32;

__device__ double get_min(double d1, double d2) {
    if (d1 < d2) {
        return d1;
    } else {
        return d2;
    }
}

__global__ void kernel(double *arr1, double *arr2, double *res, int size) {
    int offset = gridDim.x * blockDim.x;
    for (int i = blockIdx.x * blockDim.x + threadIdx.x; i < size; i += offset) {
        res[i] = get_min(arr1[i], arr2[i]);
    }
}

int main() {
    int n;
    scanf("%d", &n);

    double *arr1 = (double*)malloc(sizeof(double) * n);
    for (int i = 0; i < n; ++i) {
        scanf("%lf", &arr1[i]);
    }

    double *arr2 = (double*)malloc(sizeof(double) * n);
    for (int i = 0; i < n; ++i) {
        scanf("%lf", &arr2[i]);
    }

    double *dev_arr1;
    cudaMalloc(&dev_arr1, sizeof(double) * n);
    cudaMemcpy(dev_arr1, arr1, sizeof(double) * n, cudaMemcpyHostToDevice);

    double *dev_arr2;
    cudaMalloc(&dev_arr2, sizeof(double) * n);
    cudaMemcpy(dev_arr2, arr2, sizeof(double) * n, cudaMemcpyHostToDevice);

    double *dev_res;
    cudaMalloc(&dev_res, sizeof(double) * n);

    /* std::chrono::steady_clock::time_point start = 
        std::chrono::steady_clock::now(); */

    kernel<<<BLOCKS_AMOUNT, THREADS_AMOUNT>>>(dev_arr1, dev_arr2, dev_res, n);

    /* std::chrono::steady_clock::time_point finish = 
        std::chrono::steady_clock::now();
    unsigned time = 
        std::chrono::duration_cast<std::chrono::nanoseconds>(finish - start).count(); */

    double *res = (double*)malloc(sizeof(double) * n);
    cudaMemcpy(res, dev_res, sizeof(double) * n, cudaMemcpyDeviceToHost);

    for (int i = 0; i < n; ++i) {
        printf("%.10lf ", res[i]);
    }
    printf("\n");
    // printf("time: %dns\n", time);

    free(arr1);
    free(arr2);
    free(res);
    cudaFree(dev_arr1);
    cudaFree(dev_arr2);
    cudaFree(dev_res);

    return 0;
}
