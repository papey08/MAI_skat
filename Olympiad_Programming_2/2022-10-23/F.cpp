#include <iostream>

unsigned long long bin_pow(unsigned long long a, int n) {
    unsigned long long res = 1;
    while (n > 0) {
        if (n % 2 == 1)
            res = res * a;
        a = a * a;
        n /= 2;
    }
    return res;
}
int main() {

    int N;
    std::cin >> N;
    std::cout << bin_pow(2, N - 1);
    return 0;
}
