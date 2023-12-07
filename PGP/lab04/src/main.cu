#include <iostream>
#include <vector>
#include <algorithm>

#include <thrust/extrema.h>
#include <thrust/device_vector.h>

const int X_BLOCKS = 32;
const int X_THREADS = 32;
const int Y_BLOCKS = 32;
const int Y_THREADS = 32;
const int BLOCKS = 1024;
const int THREADS = 1024;

const double EPS = 10e-7;

__global__ void div_row(double* dev_data, int shape, double delim, int r) {
    int offset = blockDim.x * gridDim.x;
    for (int i = r + blockDim.x*blockIdx.x + threadIdx.x; i < shape; i += offset) {
        dev_data[r + i*shape] /= delim;
    }
}

__global__ void swap_rows(double* dev_data, int shape, int r1, int r2) {
    int offset = blockDim.x * gridDim.x;
    for (int i = r1 + blockDim.x * blockIdx.x + threadIdx.x; i < shape; i += offset) {
        double temp = dev_data[i*shape + r1];
        dev_data[i*shape + r1] = dev_data[i*shape + r2];
        dev_data[i*shape + r2] = temp;
    }
}

__global__ void kernel(double* dev_data, int shape, int r) {
    int x_offset = blockDim.x * gridDim.x;
    int y_offset = blockDim.y * gridDim.y;
    for (int i = r + blockDim.x*blockIdx.x + threadIdx.x + 1; i < shape; i += x_offset) {
        for (int j = r + blockDim.y*blockIdx.y + threadIdx.y + 1; j < shape; j += y_offset) {
            dev_data[j*shape + i] -= dev_data[r*shape + i] * dev_data[r + j*shape];
        }
    }
}

struct comparator {
    __host__ __device__ bool operator()(double a, double b) {
        return std::fabs(a) < std::fabs(b);
    }
};

class Matrix {
public:
    Matrix(int n) : shape(n) {
        data = (double*)malloc(sizeof(double) * n * n);
    }

    double determinant() {
        comparator comp;                                                                                                                                                    
        double* dev_matrix;                                                                                                                                                    
        cudaMalloc(&dev_matrix, sizeof(double) * shape * shape);                                                                                                                                                    
        cudaMemcpy(dev_matrix, data, sizeof(double) * shape * shape, cudaMemcpyHostToDevice);                                                                                                                                                    
        std::vector<double> delims(shape, 0.);                                                                                                                                                    
        thrust::device_ptr<double> left_ptr;                                                                                                                                                    
        thrust::device_ptr<double> max_ptr;                                                                                                                                                    
        for (int i = 0; i < shape; ++i) {                                                                                                                                                    
            left_ptr = thrust::device_pointer_cast(dev_matrix + i + i * shape);                                                                                                                                                    
            double delim = *left_ptr;                                                                                                                                                    
            if (std::abs(delim) <= EPS) {                                                                                                                                                    
                max_ptr = thrust::max_element(left_ptr, left_ptr + (shape - i), comp);                                                                                                                                                    
                double max_elem = *max_ptr;                                                                                                                                                    
                int max_index = max_ptr - left_ptr;                                                                                                                                                    
                if (std::abs(max_elem) <= EPS) {                                                                                                                                                    
                    cudaFree(dev_matrix);                                                                                                                                                    
                    return 0.;                                                                                                                                                    
                }                                                                                                                                                    
                swap_rows<<<BLOCKS, THREADS>>>(dev_matrix, shape, i, max_index+i);                                                                                                                                                    
                delims[i] -= max_elem;                                                                                                                                                    
                delim = max_elem;                                                                                                                                                    
            }                                                                                                                                                    
            else {                                                                                                                                                    
                delims[i] = delim;                                                                                                                                                    
            }                                                                                                                                                    
            div_row<<<BLOCKS, THREADS>>>(dev_matrix, shape, delim, i);                                                                                                                                                    
            kernel<<<dim3(X_BLOCKS, X_THREADS), dim3(Y_BLOCKS, Y_THREADS)>>>(dev_matrix, shape, i);                                                                                                                                                    
        }                                                                                                                                                    
        cudaFree(dev_matrix);                                                                                                                                                    
        std::sort(delims.begin(), delims.end(), comp);                                                                                                                                                    
        double det = 1.0;                                                                                                                                                    
        int l = 0;                                                                                                                                                    
        int r = delims.size() - 1;                                                                                                                                                    
        while (l <= r) {                                                                                                                                                    
            if (std::abs(det) < EPS) {                                                                                                                                                    
                det *= delims[r];                                                                                                                                                    
                --r;                                                                                                                                                    
            } else {                                                                                                                                                    
                det *= delims[l];                                                                                                                                                    
                ++l;                                                                                                                                                    
            }                                                                                                                                                    
        }                                                                                                                                                    
        return det;                                                                                                                                                    
    }

    friend std::istream& operator>>(std::istream& in, Matrix& matrix) {
        for (int i = 0; i < matrix.shape; ++i) {
            for (int j = 0; j < matrix.shape; ++j) {
                in >> matrix.data[i + j*matrix.shape];
            }
        }
        return in;
    }

    ~Matrix() {
        free(data);
    }

private:
    double* data;
    int shape;
};

int main() {
    std::ios::sync_with_stdio(false);
    std::cin.tie(nullptr);
    std::cout.tie(nullptr);
    int n;
    std::cin >> n;
    Matrix m = Matrix(n);
    std::cin >> m;
    std::cout << m.determinant() << std::endl;
    return 0;
}
