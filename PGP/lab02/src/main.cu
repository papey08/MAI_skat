#include <iostream>
#include <string>
#include <vector>
#include <cuda_runtime.h>
#include "math.h"
// #include <chrono>

const int X_BLOCKS = 8;
const int X_THREADS = 8;
const int Y_BLOCKS = 8;
const int Y_THREADS = 8;

class image{
public:
    int width;
    int height;

    uchar4* pixels;

    image(std::string path) {
        FILE *img = fopen(path.c_str(), "rb");
        fread(&width, sizeof(int), 1, img);
        fread(&height, sizeof(int), 1, img);
        pixels = (uchar4*)malloc(sizeof(uchar4) * width * height);
        fread(pixels, sizeof(uchar4), width * height, img);
        fclose(img);
    }

    ~image() {
        free(pixels);
    }

    void save_to_file(std::string path) {
        FILE *img = fopen(path.c_str(), "wb");
        fwrite(&width, sizeof(int), 1, img);
        fwrite(&height, sizeof(int), 1, img);
        fwrite(pixels, sizeof(uchar4), width * height, img);
        fclose(img);
    }
};

const double R_COEFF = 0.299;
const double G_COEFF = 0.587;
const double B_COEFF = 0.114;

__device__ double to_black_white(uchar4 p) {
    return R_COEFF * p.x + G_COEFF * p.y + B_COEFF * p.z;
}

__device__ uchar4 sobel(double w[3][3]) {
    double gx = w[0][2] + 2 * w[1][2] + w[2][2] - w[0][0] - 2 * w[1][0] - w[2][0];
    double gy = w[2][0] + 2 * w[2][1] + w[2][2] - w[0][0] - 2 * w[0][1] - w[0][2];

    int res = min(255, int(sqrt(gx * gx + gy * gy)));
    return make_uchar4(res, res, res, res);
}

texture<uchar4, 2, cudaReadModeElementType> tex;

__global__ void kernel(uchar4 *out, int width, int height) {
    int i_y = blockDim.y * blockIdx.y + threadIdx.y;
    int i_x = blockDim.x * blockIdx.x + threadIdx.x;
    int offset_y = blockDim.y * gridDim.y;
    int offset_x = blockDim.x * gridDim.x;

    for (int y = i_y; y < height; y += offset_y) {
        for (int x = i_x; x < width; x += offset_x) {
            double w[3][3];

            for (int i = 0; i < 3; i++) {
                for (int j = 0; j < 3; j++) {
                    w[i][j] = to_black_white(tex2D(tex, x-1+i, y-1+j));
                }
            }

            out[y*width + x] = sobel(w);
        }
    }
}

int main() {
    std::string in_path, out_path;
    std::cin >> in_path >> out_path;
    image img(in_path);

    cudaArray *arr;
    cudaChannelFormatDesc ch = cudaCreateChannelDesc<uchar4>();
    cudaMallocArray(&arr, &ch, img.width, img.height);
    cudaMemcpyToArray(arr, 0, 0, img.pixels, sizeof(uchar4) * img.width * img.height, cudaMemcpyHostToDevice);

    tex.addressMode[0] = cudaAddressModeClamp;
    tex.addressMode[1] = cudaAddressModeClamp;
    tex.channelDesc = ch;
    tex.filterMode = cudaFilterModePoint;
    tex.normalized = false;

    cudaBindTextureToArray(tex, arr, ch);
    uchar4* dev_out;
    cudaMalloc(&dev_out, sizeof(uchar4) * img.width * img.height);

    /* std::chrono::steady_clock::time_point start = 
        std::chrono::steady_clock::now(); */

    kernel<<<dim3(X_BLOCKS, X_THREADS), dim3(Y_BLOCKS, Y_THREADS)>>>(dev_out, img.width, img.height);

    /* std::chrono::steady_clock::time_point finish = 
        std::chrono::steady_clock::now();
    unsigned time = 
        std::chrono::duration_cast<std::chrono::nanoseconds>(finish - start).count();

    printf("time: %dns\n", time); */

    cudaMemcpy(img.pixels, dev_out, sizeof(uchar4) * img.width * img.height, cudaMemcpyDeviceToHost);
    
    cudaUnbindTexture(tex);
    cudaFreeArray(arr);
    cudaFree(dev_out);

    img.save_to_file(out_path);

    return 0;
}
