#include "ISort.hpp"

int main() {
    std::ios_base::sync_with_stdio(false);
    std::cin.tie(nullptr);
    std::cout.tie(nullptr);
    std::vector<TPair> raw;
    unsigned short inputKey;
    std::string inputString;
    while(std::cin >> inputKey) {
        std::cin >> inputString;
        raw.push_back({inputKey, inputString});
    }
    std::vector<int> count(MAX_KEY, 0);
    std::vector<int> ans(raw.size());
    Count(raw, count);
    Sort(raw, count, ans);
    for (int i = 0; i < ans.size(); ++i) {
        raw[ans[i]].Print();
        std::cout << "\n";
    }
    return 0;
}
