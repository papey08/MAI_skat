#include <iostream>

int main() {
    int n, m, x, y;
    std::cin >> n >> m >> x >> y;
    int ans = (((x - 1) % 7 + 1) + ((y - 1) % 7) - 1) % 7 + 1;
    std::cout << ans << std::endl;
    return 0;
}
