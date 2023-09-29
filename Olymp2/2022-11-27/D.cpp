#include <iostream>
#include <algorithm>

int main() {
    unsigned n;
    std::cin >> n;
    double maximum = 0;
    unsigned k, r;
    double temp;
    for (unsigned i = 0; i < n; ++i) {
        std::cin >> k >> r;
        temp = (r - 0.5) / k;
        maximum = std::max(maximum, temp);
    }
    std::cout << maximum << std::endl;
    return 0;
}
