#include <iostream>

int main() {
    std::ios_base::sync_with_stdio(false);
    std::cin.tie(nullptr);
    std::cout.tie(nullptr);
    unsigned n;
    std::cin >> n;
    unsigned num;
    unsigned s = 0;
    for (unsigned i = 0; i < n; ++i) {
        std::cin >> num;
        s ^= num;
    }
    std::cout << s << std::endl;
    return 0;
}
