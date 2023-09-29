#include <iostream>
#include <vector>

const int MAX_YEAR = 2000000;

int main() {
    std::vector<int> v(MAX_YEAR, 0);
    for (int i = 1; i < v.size(); ++i) {
        v[i] = v[i-1];
        if (i % 4 == 0 && (i % 100 != 0 || i % 400 == 0)) {
            ++v[i];
        }
    }

    int t;
    std::cin >> t;
    for (int i = 0; i < t; ++i) {
        int d, m, y, ye;
        std::cin >> d >> m >> y >> ye;
        if (!(d == 29 && m == 2)) {
            std::cout << ye - y << std::endl;
        } else {
            // std::cout << v[ye] - v[y-1] - 1 << std::endl;
            std::cout << ye/4 - y/4 + ye/400 - y/400 + 1 - (ye/100 - y/100 + 1) << std::endl;
        }
    }
    return 0;
}
