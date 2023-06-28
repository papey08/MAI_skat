#include <iostream>

int main() {
    int n,k;
    std::cin >> n >> k;
    if (n == 1) {
        std::cout << "1\n";
    } else if (k == 1) {
        std::cout << "1\n";
    } else {
        std::cout << (n-1)*(k-1) + 1 << std::endl;
    }
    return 0;
}
