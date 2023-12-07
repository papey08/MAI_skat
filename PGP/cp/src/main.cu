#include <iostream>
#include <string>
#include <math.h>
// #include <chrono>

struct point {
    double x;
    double y;
    double z;

    __host__ __device__ point() {}
    __host__ __device__ point(double x, double y, double z) : x(x), y(y), z(z) {}
};

std::istream& operator>>(std::istream& in, point& p) {
    in >> p.x >> p.y >> p.z;
    return in;
}

std::ostream& operator<<(std::ostream& out, point& p) {
    out << p.x << " " <<  p.y << " " << p.z;
    return out;
}

__host__ __device__ point operator+(point a, point b) {
    return point(
        a.x + b.x,
        a.y + b.y,
        a.z + b.z
    );
}

__host__ __device__ point operator-(point a, point b) {
    return point(
        a.x - b.x,
        a.y - b.y,
        a.z - b.z
    );
}

__host__ __device__ point operator*(point a, double b) {
    return point(
        a.x * b,
        a.y * b,
        a.z * b
    );
}

__host__ __device__ double dot(point a, point b) {
    return a.x * b.x + a.y * b.y + a.z * b.z;
}

__host__ __device__ point prod(point a, point b) {
    return point(
        a.y * b.z - a.z * b.y,
        a.z * b.x - a.x * b.z,
        a.x * b.y - a.y * b.x
    );
}

__host__ __device__ point normalize(point v) {
    double l = sqrt(dot(v, v));
    return point(
        v.x / l,
        v.y / l,
        v.z / l
    );
}

__host__ __device__ point multiply(point a, point b, point c, point v) {
    return point(
        a.x * v.x + b.x * v.y + c.x * v.z,
        a.y * v.x + b.y * v.y + c.y * v.z,
        a.z * v.x + b.z * v.y + c.z * v.z
    );
}

struct polygon {
    point a;
    point b;
    point c;
    uchar4 color;

    __host__ __device__ polygon() {}
    __host__ __device__ polygon(point a, point b, point c, uchar4 color) : a(a), b(b), c(c), color(color) {}
};

__host__ __device__ uchar4 make_ray(
    point position, point direction, 
    point light_position, uchar4 light_color, 
    polygon* polygons, int polygons_amount
) {
    int i_min = -1;
    double ts_min;

    for (int i = 0; i < polygons_amount; ++i) {
        point e1 = polygons[i].b - polygons[i].a;
        point e2 = polygons[i].c - polygons[i].a;
        point p = prod(direction, e2);
        double div = dot(p, e1);
        if (fabs(div) < 1e-10) {
            continue;
        }
        point t = position - polygons[i].a;
        double u = dot(p, t) / div;
        if (u < 0.0 || u > 1.0) {
            continue;
        }
        point q = prod(t, e1);
        double v = dot(q, direction) / div;
        if (v < 0.0 || v + u > 1.0) {
            continue;
        }
        double ts = dot(q, e2) / div; 
        if (ts < 0.0) {
            continue;
        }
        if (i_min == -1 || ts < ts_min) {
            i_min = i;
            ts_min = ts;
        }
    }

    if (i_min == -1) {
        return make_uchar4(0, 0, 0, 255);
    }

    point new_position = direction * ts_min + position;
    point new_direction = light_position - new_position;
    double length = sqrt(dot(new_direction, new_direction));
    new_direction = normalize(new_direction);

    for (int i = 0; i < polygons_amount; ++i) {
        point e1 = polygons[i].b - polygons[i].a;
        point e2 = polygons[i].c - polygons[i].a;
        point p = prod(new_direction, e2);
        double div = dot(p, e1);
        if (fabs(div) < 1e-10)
            continue;
        point t = new_position - polygons[i].a;
        double u = dot(p, t) / div;
        if (u < 0.0 || u > 1.0)
            continue;
        point q = prod(t, e1);
        double v = dot(q, new_direction) / div;
        if (v < 0.0 || v + u > 1.0)
            continue;
        double ts = dot(q, e2) / div; 
        if (ts > 0.0 && ts < length && i != i_min) {
            return make_uchar4(0, 0, 0, 255);
        }
    }

    return make_uchar4(
        polygons[i_min].color.x * light_color.x,
        polygons[i_min].color.y * light_color.y,
        polygons[i_min].color.z * light_color.z,
        255
    );
}

__host__ __device__ void cpu_render(
    uchar4* data, 
    point camera_position, point camera_view, 
    int width, int height, double view_angle, 
    point light_position, uchar4 light_color, 
    polygon* polygons, int polygons_amount
) {
    double dw = 2. / (width - 1.);
    double dh = 2. / (height - 1.);
    double z = 1. / tan(view_angle * M_PI / 360.);

    point bz = normalize(camera_view - camera_position);
    point bx = normalize(prod(bz, {0., 0., 1.}));
    point by = normalize(prod(bx, bz));

    for (int i = 0; i < width; ++i) {
        for (int j = 0; j < height; ++j) {
            point v = point(-1. + dw * i, (-1. + dh * j) * height / width, z);
            point dir = multiply(bx, by, bz, v);
            data[(height - 1 - j) * width + i] = make_ray(camera_position, normalize(dir), light_position, light_color, polygons, polygons_amount);
        }
    }
}

__global__ void gpu_render(
    uchar4* data,
    point camera_pos, point camera_view, 
    int width, int height, double view_angle,  
    point light_position, uchar4 light_color, 
    polygon* polygons, int polygons_amount
) {
    int i_x = blockDim.x * blockIdx.x + threadIdx.x;
    int i_y = blockDim.y * blockIdx.y + threadIdx.y;
    int offset_x = blockDim.x * gridDim.x;
    int offset_y = blockDim.y * gridDim.y;

    double dw = 2. / (width - 1.);
    double dh = 2. / (height - 1.);
    double z = 1. / tan(view_angle * M_PI / 360.);

    point bz = normalize(camera_view - camera_pos);
    point bx = normalize(prod(bz, {0., 0., 1.}));
    point by = normalize(prod(bx, bz));

    for (int i = i_x; i < width; i += offset_x) {
        for (int j = i_y; j < height; j += offset_y) {
            point v = point(-1. + dw * i, (-1. + dh * j) * height / width, z);
            point dir = multiply(bx, by, bz, v);
            data[(height - 1 - j) * width + i] = make_ray(camera_pos, normalize(dir), light_position, light_color, polygons, polygons_amount);
        }
    }
}

__host__ __device__ void cpu_smoothing(uchar4* data, uchar4* smoothing_data, int width, int height, int sqrt_rpp) {
    for (int x = 0; x < width; ++x) {
        for (int y = 0; y < height; ++y) {
            uint4 temp = make_uint4(0, 0, 0, 0);
            for (int i = 0; i < sqrt_rpp; ++i) {
                for (int j = 0; j < sqrt_rpp; ++j) {
                    uchar4 cur_pixel = data[width * sqrt_rpp * (y * sqrt_rpp + j) + (x * sqrt_rpp + i)];
                    temp.x += cur_pixel.x;
                    temp.y += cur_pixel.y;
                    temp.z += cur_pixel.z;
                }
            }
            int rpp = sqrt_rpp * sqrt_rpp;
            smoothing_data[y * width + x] = make_uchar4(temp.x / rpp, temp.y / rpp, temp.z / rpp, 255);
        }
    }
}

__global__ void gpu_smoothing(uchar4* data, uchar4* smoothing_data, int width, int height, int sqrt_rpp) {
    int i_x = blockDim.x * blockIdx.x + threadIdx.x;
    int i_y = blockDim.y * blockIdx.y + threadIdx.y;
    int offset_x = blockDim.x * gridDim.x;
    int offset_y = blockDim.y * gridDim.y;

    for (int x = i_x; x < width; x += offset_x) {
        for (int y = i_y; y < height; y += offset_y) {
            uint4 temp = make_uint4(0, 0, 0, 0);
            for (int i = 0; i < sqrt_rpp; ++i) {
                for (int j = 0; j < sqrt_rpp; ++j) {
                    uchar4 cur_pixel = data[width * sqrt_rpp * (y * sqrt_rpp + j) + (x * sqrt_rpp + i)];
                    temp.x += cur_pixel.x;
                    temp.y += cur_pixel.y;
                    temp.z += cur_pixel.z;
                }
            }
            int rpp = sqrt_rpp * sqrt_rpp;
            smoothing_data[y * width + x] = make_uchar4(temp.x / rpp, temp.y / rpp, temp.z / rpp, 255);
        }
    }
}


struct frames_params {
    int amount;
    std::string path_to_save_frames;
    int width, height;
    double view_angle;
};

std::istream& operator>>(std::istream& in, frames_params& f) {
    in >> f.amount >> f.path_to_save_frames >> f.width >> f.height >> f.view_angle;
    return in;
}

struct camera_params {
    double r0c, z0c, phi0c, arc, azc, wrc, wzc, wphic, prc, pzc;
    double r0n, z0n, phi0n, arn, azn, wrn, wzn, wphin, prn, pzn;
};

std::istream& operator>>(std::istream& in, camera_params& c) {
    in >> c.r0c >> c.z0c >> c.phi0c >> c.arc >> c.azc >> c.wrc >> c.wzc >> c.wphic >> c.prc >> c.pzc;
    in >> c.r0n >> c.z0n >> c.phi0n >> c.arn >> c.azn >> c.wrn >> c.wzn >> c.wphin >> c.prn >> c.pzn;
    return in;
}

struct figure_params {
    point center;
    uchar4 color;
    double radius;
};

std::istream& operator>>(std::istream& in, figure_params& f) {
    in >> f.center;

    double r, g, b;
    in >> r >> g >> b;
    f.color = make_uchar4(r * 255, g * 255, b * 255, 255);

    in >> f.radius;
    return in;
}

struct floor_params {
    point p1, p2, p3, p4;
    uchar4 color;
};

std::istream& operator>>(std::istream& in, floor_params& f) {
    in >> f.p1 >> f.p2 >> f.p3 >> f.p4;

    double r, g, b;
    in >> r >> g >> b;
    f.color = make_uchar4(r * 255, g * 255, b * 255, 255);

    return in;
}

struct light_params {
    point position;
    uchar4 color;
    double sqrt_rpp;
};

std::istream& operator>>(std::istream& in, light_params& l) {
    in >> l.position;

    double r, g, b;
    in >> r >> g >> b;
    l.color = make_uchar4(r * 255, g * 255, b * 255, 255);

    in >> l.sqrt_rpp;
    return in;
}


class app {
    bool use_gpu;
    int polygons_amount = 62;
    int x_blocks_amount = 8, x_threads_amount = 8;
    int y_blocks_amount = 8, y_threads_amount = 8;

    frames_params frames;
    camera_params camera;
    figure_params tetrahedron;
    figure_params dodecahedron;
    figure_params icosahedron;
    floor_params floor;
    light_params light;

    void init_floor(polygon* polygons) {
        polygons[0] = polygon(floor.p1, floor.p2, floor.p3, floor.color);
        polygons[1] = polygon(floor.p1, floor.p3, floor.p4, floor.color);
    }

    void init_tetrahedron(polygon* polygons) {
        double a = 4. / sqrt(6) * tetrahedron.radius;

        point v1 = point(tetrahedron.center.x, tetrahedron.center.y + tetrahedron.radius, tetrahedron.center.z);
        point v2 = point(tetrahedron.center.x + 0.578 * a, tetrahedron.center.y - 1/3 * tetrahedron.radius, tetrahedron.center.z);
        point v3 = point(tetrahedron.center.x - 0.289 * a, tetrahedron.center.y - 1/3 * tetrahedron.radius, tetrahedron.center.z + 0.5 * a);
        point v4 = point(tetrahedron.center.x - 0.289 * a, tetrahedron.center.y - 1/3 * tetrahedron.radius, tetrahedron.center.z - 0.5 * a);
    
        polygons[2] = polygon(v1, v2, v3, tetrahedron.color);
        polygons[3] = polygon(v1, v3, v4, tetrahedron.color);
        polygons[4] = polygon(v1, v2, v4, tetrahedron.color);
        polygons[5] = polygon(v2, v3, v4, tetrahedron.color);
    }

    void init_dodecahedron(polygon* polygons) {
        double a = (1. + sqrt(5.)) / 2.;
        double b = 1. / a;

        point v[] = {
            point(-b, 0., a),
            point(b, 0., a),
            point(-1., 1., 1.),
            point(1., 1., 1.),
            point(1., -1., 1.),
            point(-1., -1., 1.),
            point(0., -a, b),
            point(0., a, b),
            point(-a, -b, 0.),
            point(-a, b, 0.),
            point(a, b, 0.),
            point(a, -b, 0.),
            point(0., -a, -b),
            point(0., a, -b),
            point(1., 1., -1.),
            point(1., -1., -1.),
            point(-1., -1., -1.),
            point(-1., 1., -1.),
            point(b, 0., -a),
            point(-b, 0., -a)
        };

        for (int i = 0; i < 20; ++i) {
            v[i].x = v[i].x * dodecahedron.radius / sqrt(3.) + dodecahedron.center.x;
            v[i].y = v[i].y * dodecahedron.radius / sqrt(3.) + dodecahedron.center.y;
            v[i].z = v[i].z * dodecahedron.radius / sqrt(3.) + dodecahedron.center.z;
        }

        polygons[6] = polygon(v[4], v[0], v[6], dodecahedron.color);
        polygons[7] = polygon(v[0], v[5], v[6], dodecahedron.color);
        polygons[8] = polygon(v[0], v[4], v[1], dodecahedron.color);
        polygons[9] = polygon(v[0], v[3], v[7], dodecahedron.color);
        polygons[10] = polygon(v[2], v[0], v[7], dodecahedron.color);
        polygons[11] = polygon(v[0], v[1], v[3], dodecahedron.color);
        polygons[12] = polygon(v[10], v[1], v[11], dodecahedron.color);
        polygons[13] = polygon(v[3], v[1], v[10], dodecahedron.color);
        polygons[14] = polygon(v[1], v[4], v[11], dodecahedron.color);
        polygons[15] = polygon(v[5], v[0], v[8], dodecahedron.color);
        polygons[16] = polygon(v[0], v[2], v[9], dodecahedron.color);
        polygons[17] = polygon(v[8], v[0], v[9], dodecahedron.color);
        polygons[18] = polygon(v[5], v[8], v[16], dodecahedron.color);
        polygons[19] = polygon(v[6], v[5], v[12], dodecahedron.color);
        polygons[20] = polygon(v[12], v[5], v[16], dodecahedron.color);
        polygons[21] = polygon(v[4], v[12], v[15], dodecahedron.color);
        polygons[22] = polygon(v[4], v[6], v[12], dodecahedron.color);
        polygons[23] = polygon(v[11], v[4], v[15], dodecahedron.color);
        polygons[24] = polygon(v[2], v[13], v[17], dodecahedron.color);
        polygons[25] = polygon(v[2], v[7], v[13], dodecahedron.color);
        polygons[26] = polygon(v[9], v[2], v[17], dodecahedron.color);
        polygons[27] = polygon(v[13], v[3], v[14], dodecahedron.color);
        polygons[28] = polygon(v[7], v[3], v[13], dodecahedron.color);
        polygons[29] = polygon(v[3], v[10], v[14], dodecahedron.color);
        polygons[30] = polygon(v[8], v[17], v[19], dodecahedron.color);
        polygons[31] = polygon(v[16], v[8], v[19], dodecahedron.color);
        polygons[32] = polygon(v[8], v[9], v[17], dodecahedron.color);
        polygons[33] = polygon(v[14], v[11], v[18], dodecahedron.color);
        polygons[34] = polygon(v[11], v[15], v[18], dodecahedron.color);
        polygons[35] = polygon(v[10], v[11], v[14], dodecahedron.color);
        polygons[36] = polygon(v[12], v[19], v[18], dodecahedron.color);
        polygons[37] = polygon(v[15], v[12], v[18], dodecahedron.color);
        polygons[38] = polygon(v[12], v[16], v[19], dodecahedron.color);
        polygons[39] = polygon(v[19], v[13], v[18], dodecahedron.color);
        polygons[40] = polygon(v[17], v[13], v[19], dodecahedron.color);
        polygons[41] = polygon(v[13], v[14], v[18], dodecahedron.color);
    }

    void init_icosahedron(polygon* polygons) {
        double a = icosahedron.radius / 0.951;
        double r = a * sin(36. * M_PI / 180.);
        double h = 0.25 * a * sqrt(3.);

        point v[] = {
            point(icosahedron.center.x, icosahedron.center.y, icosahedron.center.z),
            point(icosahedron.center.x, icosahedron.center.y + icosahedron.radius, icosahedron.center.z),
            point(icosahedron.center.x + r, icosahedron.center.y + h, icosahedron.center.z),
            point(icosahedron.center.x + r*sin(18. * M_PI/180.), icosahedron.center.y + h, icosahedron.center.z - r*cos(18. * M_PI/180.)),
            point(icosahedron.center.x + r*sin(18. * M_PI/180.), icosahedron.center.y + h, icosahedron.center.z + r*cos(18. * M_PI/180.)),
            point(icosahedron.center.x - r*sin(54. * M_PI/180.), icosahedron.center.y + h, icosahedron.center.z - r*cos(54. * M_PI/180.)),
            point(icosahedron.center.x - r*sin(54. * M_PI/180.), icosahedron.center.y + h, icosahedron.center.z + r*cos(54. * M_PI/180.)),
            point(icosahedron.center.x, icosahedron.center.y - icosahedron.radius, icosahedron.center.z),
            point(icosahedron.center.x - r, icosahedron.center.y - h, icosahedron.center.z),
            point(icosahedron.center.x - r*sin(18. * M_PI/180.), icosahedron.center.y - h, icosahedron.center.z - r*cos(18. * M_PI/180.)),
            point(icosahedron.center.x - r*sin(18. * M_PI/180.), icosahedron.center.y - h, icosahedron.center.z + r*cos(18. * M_PI/180.)),
            point(icosahedron.center.x + r*sin(54. * M_PI/180.), icosahedron.center.y - h, icosahedron.center.z - r*cos(54. * M_PI/180.)),
            point(icosahedron.center.x + r*sin(54. * M_PI/180.), icosahedron.center.y - h, icosahedron.center.z + r*cos(54. * M_PI/180.)),
        };

        polygons[42] = polygon(v[7], v[9], v[8], icosahedron.color);
        polygons[43] = polygon(v[7], v[8], v[10], icosahedron.color);
        polygons[44] = polygon(v[7], v[10], v[12], icosahedron.color);
        polygons[45] = polygon(v[7], v[12], v[11], icosahedron.color);
        polygons[46] = polygon(v[7], v[11], v[9], icosahedron.color);
        polygons[47] = polygon(v[8], v[5], v[6], icosahedron.color);
        polygons[48] = polygon(v[8], v[10], v[6], icosahedron.color);
        polygons[49] = polygon(v[10], v[12], v[4], icosahedron.color);
        polygons[50] = polygon(v[12], v[11], v[2], icosahedron.color);
        polygons[51] = polygon(v[11], v[9], v[3], icosahedron.color);
        polygons[52] = polygon(v[9], v[8], v[5], icosahedron.color);
        polygons[53] = polygon(v[10], v[6], v[4], icosahedron.color);
        polygons[54] = polygon(v[12], v[2], v[4], icosahedron.color);
        polygons[55] = polygon(v[11], v[3], v[2], icosahedron.color);
        polygons[56] = polygon(v[9], v[3], v[5], icosahedron.color);
        polygons[57] = polygon(v[3], v[2], v[1], icosahedron.color);
        polygons[58] = polygon(v[2], v[4], v[1], icosahedron.color);
        polygons[59] = polygon(v[4], v[6], v[1], icosahedron.color);
        polygons[60] = polygon(v[6], v[5], v[1], icosahedron.color);
        polygons[61] = polygon(v[5], v[3], v[1], icosahedron.color);
    }

public:
    app(bool _use_gpu) {
        use_gpu = _use_gpu;
        frames = frames_params{1, "%d.data", 800, 800, 90.};
        camera = camera_params{7., 3., 0., 2., 1., 2., 6., 1., 0., 0., 2., 0., 0., 0.5, 0.1, 1., 4., 1., 0., 0.};
        tetrahedron = figure_params{point(0., -2., 0.), make_uchar4(255, 0, 0, 255), 1.};
        dodecahedron = figure_params{point(0., 0., 0.), make_uchar4(0, 255, 0, 255), 1.};
        icosahedron = figure_params{point(0., 2., 0.), make_uchar4(0, 0, 255, 255), 1.};
        floor = floor_params{point(-5., -5., -1.), point(-5., 5., -1.), point(5., 5., -1.), point(5., -5., -1), make_uchar4(255, 255, 255, 255)};
        light = light_params{point(10., 0., 15.), make_uchar4(75, 50, 25, 255), 4.};
    }

    app(
        bool _use_gpu,
        frames_params _frames, 
        camera_params _camera, 
        figure_params _tetrahedron,
        figure_params _dodecahedron,
        figure_params _icosahedron,
        floor_params _floor,
        light_params _light
    ) {
        use_gpu = _use_gpu;
        frames = _frames;
        camera = _camera;
        tetrahedron = _tetrahedron;
        dodecahedron = _dodecahedron;
        icosahedron = _icosahedron;
        floor = _floor;
        light = _light;
    }

    void run() {
        uchar4* data = (uchar4*)malloc(sizeof(uchar4) * frames.width * frames.height * light.sqrt_rpp * light.sqrt_rpp);
        uchar4* smoothing_data = (uchar4*)malloc(sizeof(uchar4) * frames.width * frames.height);
        uchar4* dev_data;
        uchar4* dev_smoothing_data;
        polygon polygons[polygons_amount];
        polygon* dev_polygons;

        init_floor(polygons);
        init_tetrahedron(polygons);
        init_dodecahedron(polygons);
        init_icosahedron(polygons);
        
        if (use_gpu) {
            cudaMalloc(&dev_data, sizeof(uchar4) * frames.width * frames.height * light.sqrt_rpp * light.sqrt_rpp);
            cudaMalloc(&dev_smoothing_data, sizeof(uchar4) * frames.width * frames.height);
            cudaMalloc(&dev_polygons, sizeof(polygon) * polygons_amount);
            cudaMemcpy(dev_polygons, polygons, sizeof(polygon) * polygons_amount, cudaMemcpyHostToDevice);
        }

        for (int i = 0; i < frames.amount; ++i) {
            double t = 2 * M_PI * i / frames.amount;
            point camera_current_position = point(
                (camera.r0c + camera.arc * sin(camera.wrc * t + camera.prc)) * cos(camera.phi0c + camera.wphic * t),
                (camera.r0c + camera.arc * sin(camera.wrc * t + camera.prc)) * sin(camera.phi0c + camera.wphic * t),
                camera.z0c + camera.azc * sin(camera.wzc * t + camera.pzc)
            );
            point camera_current_view = point(
                (camera.r0n + camera.arn * sin(camera.wrn * t + camera.prn)) * cos(camera.phi0n + camera.wphin * t),
                (camera.r0n + camera.arn * sin(camera.wrn * t + camera.prn)) * sin(camera.phi0n + camera.wphin * t),
                camera.z0n + camera.azn * sin(camera.wzn * t + camera.pzn)
            );

            cudaEvent_t start, stop;
            cudaEventCreate(&start);
            cudaEventCreate(&stop);
            cudaEventRecord(start);

            if (use_gpu) {
                gpu_render<<<dim3(x_blocks_amount, x_threads_amount), dim3(y_blocks_amount, y_threads_amount)>>>(
                    dev_data,
                    camera_current_position, 
                    camera_current_view, 
                    frames.width * light.sqrt_rpp, 
                    frames.height * light.sqrt_rpp, 
                    frames.view_angle,
                    light.position, 
                    light.color, 
                    dev_polygons, 
                    polygons_amount
                );

                gpu_smoothing<<<dim3(x_blocks_amount, x_threads_amount), dim3(y_blocks_amount, y_threads_amount)>>>(
                    dev_data, 
                    dev_smoothing_data, 
                    frames.width, 
                    frames.height, 
                    light.sqrt_rpp
                );

                cudaMemcpy(smoothing_data, dev_smoothing_data, sizeof(uchar4) * frames.width * frames.height, cudaMemcpyDeviceToHost);
            } else {
                cpu_render(
                    data,
                    camera_current_position, 
                    camera_current_view, 
                    frames.width * light.sqrt_rpp, 
                    frames.height * light.sqrt_rpp, 
                    frames.view_angle,
                    light.position, 
                    light.color, 
                    polygons, 
                    polygons_amount
                );

                cpu_smoothing(
                    data, 
                    smoothing_data, 
                    frames.width, 
                    frames.height, 
                    light.sqrt_rpp
                );
            }

            cudaEventRecord(stop);
            cudaEventSynchronize(stop);
            cudaEventDestroy(start);
            cudaEventDestroy(stop);

            
            char path_to_save_ith_frame[frames.path_to_save_frames.length() + std::to_string(i).length() - 2];
            sprintf(path_to_save_ith_frame, frames.path_to_save_frames.c_str(), i);
            FILE* output_file = fopen(path_to_save_ith_frame, "w");
            fwrite(&frames.width, sizeof(int), 1, output_file);
            fwrite(&frames.height, sizeof(int), 1, output_file);
            fwrite(smoothing_data, sizeof(uchar4), frames.width * frames.height, output_file);
            fclose(output_file);
        }

        free(data);
        free(smoothing_data);
        if (use_gpu) {
            cudaFree(dev_data);
            cudaFree(dev_smoothing_data);
            cudaFree(dev_polygons);
        }
    }

    ~app() {}

    friend std::ostream& operator<<(std::ostream& out, app& a) {
        out << a.frames.amount << std::endl;
        out << a.frames.path_to_save_frames << std::endl;
        out << a.frames.width << " " << a.frames.height << " " << a.frames.view_angle << std::endl;
        
        out << a.camera.r0c << " " << a.camera.z0c << " " << a.camera.phi0c << " " 
            << a.camera.arc << " " << a.camera.azc << " " << a.camera.wrc << " " 
            << a.camera.wzc << " " << a.camera.wphic << " " << a.camera.prc << " " 
            << a.camera.pzc << std::endl;
        out << a.camera.r0n << " " << a.camera.z0n << " " << a.camera.phi0n << " " 
            << a.camera.arn << " " << a.camera.azn << " " << a.camera.wrn << " " 
            << a.camera.wzn << " " << a.camera.wphin << " " << a.camera.prn << " " 
            << a.camera.pzn << std::endl;

        out << a.tetrahedron.center << " " << int(a.tetrahedron.color.x) / 255. << " " 
            << int(a.tetrahedron.color.y) / 255. << " " << int(a.tetrahedron.color.z) / 255. << " " 
            << a.tetrahedron.radius << std::endl;

        out << a.dodecahedron.center << " " << int(a.dodecahedron.color.x) / 255. << " " 
            << int(a.dodecahedron.color.y) / 255. << " " << int(a.dodecahedron.color.z) / 255. << " " 
            << a.dodecahedron.radius << std::endl;

        out << a.icosahedron.center << " " << int(a.icosahedron.color.x) / 255. << " " 
            << int(a.icosahedron.color.y) / 255. << " " << int(a.icosahedron.color.z) / 255. << " " 
            << a.icosahedron.radius << std::endl;

        out << a.floor.p1 << " " << a.floor.p2 << " " << a.floor.p3 << " " << a.floor.p4 << " " 
            << int(a.floor.color.x) / 255. << " " << int(a.floor.color.y) / 255. << " " << int(a.floor.color.z) / 255. << std::endl;

        out << a.light.position << " " << int(a.light.color.x) / 255. << " " << int(a.light.color.y) / 255. << " "
            << int(a.light.color.z) / 255. << " " << a.light.sqrt_rpp;
        
        return out;
    }
};

int main(int args, char** argv) {
    bool gpu_flag = false, cpu_flag = false, default_flag = false;
    for (int i = 1; i < args; ++i) {
        std::string argument = std::string(argv[i]);
        if (argument == "--cpu") {
            cpu_flag = true;
        } else if (argument == "--gpu") {
            gpu_flag = true;
        } else if (argument == "--default") {
            default_flag = true;
        }
    }
    bool use_gpu = true;
    if (cpu_flag && !gpu_flag) {
        use_gpu = false;
    }

    /* std::chrono::steady_clock::time_point start = 
        std::chrono::steady_clock::now(); */

    if (default_flag) {
        app App = app(use_gpu);
        std::cout << App << std::endl;
        App.run();
    } else {
        frames_params frames;
        camera_params camera;
        figure_params tetrahedron;
        figure_params dodecahedron;
        figure_params icosahedron;
        floor_params floor;
        light_params light;

        std::cin >> frames;
        std::cin >> camera;
        std::cin >> tetrahedron;
        std::cin >> dodecahedron;
        std::cin >> icosahedron;
        std::cin >> floor;
        std::cin >> light;

        app App = app(
            use_gpu,
            frames,
            camera,
            tetrahedron,
            dodecahedron,
            icosahedron,
            floor,
            light
        );
        App.run();
    }

    /* std::chrono::steady_clock::time_point finish = 
        std::chrono::steady_clock::now();
    unsigned time = 
        std::chrono::duration_cast<std::chrono::seconds>(finish - start).count();

    printf("time: %d sec\n", time); */

    return 0;
}
