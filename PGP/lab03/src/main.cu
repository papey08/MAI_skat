#include "cuda_runtime.h"
#include "device_launch_parameters.h"
#include <stdio.h>
#include <stdlib.h>
#include <iostream>
#include <vector>
#include <math.h>

const int X_BLOCKS = 8;
const int X_THREADS = 8;
const int Y_BLOCKS = 8;
const int Y_THREADS = 8;

__constant__ double average_cache[32][3];
__constant__ double normalized_cache[32][3];
double average_buffer[32][3];
double normalized_buffer[32][3];

class image {
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

void fill_average_buffer(std::vector<std::vector<std::pair<int, int>>>& coords, uchar4* pixels, int width, int height, int classes_amount) {
    std::vector<double> avgs(32*3);
    for (int i = 0; i < classes_amount; i++) {
        avgs[i*3+0] = 0;
        avgs[i*3+1] = 0;
        avgs[i*3+2] = 0;
        for (int j = 0; j < coords[i].size(); j++) {
            uchar4 pixel = pixels[coords[i][j].second * width + coords[i][j].first];
            double rgb[3];
			rgb[0] = pixel.x;
    		rgb[1] = pixel.y;
    		rgb[2] = pixel.z;
            avgs[i*3+0] += rgb[0];
            avgs[i*3+1] += rgb[1];
            avgs[i*3+2] += rgb[2];
        }
        avgs[i*3+0] /= coords[i].size();
        avgs[i*3+1] /= coords[i].size();
        avgs[i*3+2] /= coords[i].size();
    }
    for (int i = 0; i < classes_amount; i++) {
        average_buffer[i][0] = avgs[i*3+0];
        average_buffer[i][1] = avgs[i*3+1];
        average_buffer[i][2] = avgs[i*3+2];
    }
}

__device__ double spectral_angle_method(uchar4 pixel, int avg) {
    double rgb_pixel[3];
    rgb_pixel[0] = pixel.x;
    rgb_pixel[1] = pixel.y;
    rgb_pixel[2] = pixel.z;
    double t_rgb[3];
    double t_normalized[3];
    t_rgb[0] = rgb_pixel[0];
	t_rgb[1] = rgb_pixel[1];
	t_rgb[2] = rgb_pixel[2];
    t_normalized[0] = normalized_cache[avg][0];
	t_normalized[1] = normalized_cache[avg][1];
	t_normalized[2] = normalized_cache[avg][2];
    return t_rgb[0] * t_normalized[0] + t_rgb[1] * t_normalized[1] + t_rgb[2] * t_normalized[2];
}

__global__ void kernel(uchar4* pixels, int width, int height, int classes_amount) {
    int i_x = blockDim.x * blockIdx.x + threadIdx.x;
    int i_y = blockDim.y * blockIdx.y + threadIdx.y;
    int offset_x = blockDim.x * gridDim.x;
    int offset_y = blockDim.y * gridDim.y;

    for (int y = i_y; y < height; y += offset_y) {
        for (int x = i_x; x < width; x += offset_x) {
            uchar4 pixel = pixels[y * width + x];
            double coeff1 = spectral_angle_method(pixel, 0);
            int idx = 0;
            for (int i = 1; i < classes_amount; i++) {
                double argmax = spectral_angle_method(pixel, i);
                if (coeff1 < argmax) {
                    coeff1 = argmax;
                    idx = i;
                }
            }
            pixels[y*width + x].w = (unsigned char)idx;
        }
	}
}
int main() {
    std::string in_path, out_path;
    std::cin >> in_path >> out_path;
    image img(in_path);
    int classes_amount;
    std::cin >> classes_amount;
    std::vector<std::vector<std::pair<int, int>>> coords(classes_amount);
    for (int i = 0; i < classes_amount; i++) {
		int pairs_amount;
        std::cin >> pairs_amount;
        coords[i].resize(pairs_amount);
        for (int j = 0; j < pairs_amount; j++) {
            std::cin >> coords[i][j].first >> coords[i][j].second;
        }
    }
    fill_average_buffer(coords, img.pixels, img.width, img.height, classes_amount);
    for (int i = 0; i < classes_amount; i++) {
        normalized_buffer[i][0] = (double)average_buffer[i][0] / sqrt(pow(average_buffer[i][0], 2) + pow(average_buffer[i][1], 2) + pow(average_buffer[i][2], 2));
        normalized_buffer[i][1] = (double)average_buffer[i][1] / sqrt(pow(average_buffer[i][0], 2) + pow(average_buffer[i][1], 2) + pow(average_buffer[i][2], 2));
        normalized_buffer[i][2] = (double)average_buffer[i][2] / sqrt(pow(average_buffer[i][0], 2) + pow(average_buffer[i][1], 2) + pow(average_buffer[i][2], 2));
    }
    cudaMemcpyToSymbol(average_cache, average_buffer, 32 * sizeof(double[3]));
    cudaMemcpyToSymbol(normalized_cache, normalized_buffer, 32 * sizeof(double[3]));
    uchar4* dev_out;
    cudaMalloc(&dev_out, sizeof(uchar4) * img.width * img.height);
    cudaMemcpy(dev_out, img.pixels, sizeof(uchar4) * img.width * img.height, cudaMemcpyHostToDevice);
    kernel<<<dim3(X_BLOCKS, X_THREADS), dim3(Y_BLOCKS, Y_THREADS)>>>(dev_out, img.width, img.height, classes_amount);
    cudaGetLastError();
    cudaMemcpy(img.pixels, dev_out, sizeof(uchar4) * img.width * img.height, cudaMemcpyDeviceToHost);
    cudaFree(dev_out);
    img.save_to_file(out_path);
    return 0;
}
		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 
                                                                                                                                                 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 		 				 