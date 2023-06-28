#include <iostream>

int main() {
    std::ios_base::sync_with_stdio(false);
    std::cin.tie(nullptr);
    std::cout.tie(nullptr);
    int n;
    std::cin >> n;
    if (n % 2 == 0) {
        std::cout << "No" << std::endl;
    } else {
        std::cout << "Yes" << std::endl;
        for (unsigned i = n - 1; i > 0; --i) {
            for (unsigned j = 0; j < i; ++j) {
                if (j % 2 == 0) {
                    std::cout << "1" << " ";
                } else {
                    std::cout << "0" << " ";
                }
            }
            std::cout << std::endl;
        }
    }
    return 0;
}
