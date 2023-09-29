#include <iostream>
#include <vector>
#include <algorithm>

struct micro_chel {
    unsigned long long w;
    unsigned long long h;
};

bool cmp(micro_chel a, micro_chel b) {
    return a.h < b.h;
}

int main() {
    std::ios_base::sync_with_stdio(false);
    std::cin.tie(nullptr);
    std::cout.tie(nullptr);
    int n;
    std::cin >> n;
    std::vector<micro_chel> v(n);
    for (int i = 0; i < n; ++i) {
        std::cin >> v[i].w;
    }
    for (int i = 0; i < n; ++i) {
        std::cin >> v[i].h;
    }
    std::sort(v.begin(), v.end(), cmp);
    for (int i = 0; i < n-1; ++i) {
        std::cout << v[i].w << " ";
    }
    std::cout << v[n-1].w << std::endl;
    return 0;
}
