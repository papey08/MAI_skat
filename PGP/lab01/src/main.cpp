#include <stdio.h>
#include <stdlib.h>
// #include <chrono>

double get_min(double d1, double d2) {
    if (d1 < d2) {
        return d1;
    } else {
        return d2;
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

    double *res = (double*)malloc(sizeof(double) * n);

    /* std::chrono::steady_clock::time_point start = 
        std::chrono::steady_clock::now(); */

    for (int i = 0; i < n; ++i) {
        res[i] = get_min(arr1[i], arr2[i]);
    }

    /* std::chrono::steady_clock::time_point finish = 
        std::chrono::steady_clock::now();
    unsigned time = 
        std::chrono::duration_cast<std::chrono::nanoseconds>(finish - start).count();

    printf("time: %dns\n", time); */

    for (int i = 0; i < n; ++i) {
        printf("%.10lf ", res[i]);
    }
    printf("\n");

    return 0;
}
