#include <iostream>

int main() {
    unsigned T;
    std::cin >> T;
    for (unsigned i = 0; i < T; ++i) {
        unsigned r;
        std::cout << "h8\n";
        std::cin >> r;
        if (!r) {
            std::cout << "a2\n";
        }
        else {
            continue;
        }
        std::cin >> r;
        if (!r) {
            std::cout << "c5\n";
        } else {
            continue;
        }
        std::cin >> r;
        if (!r) {
            break;
        }
    }
    return 0;
}
