#include <iostream>
#include <vector>
#include <string>
#include <algorithm>

#include "TSuffixTree.hpp"

int main() {
    TSuffixTree tree;
    std::string word;
    std::cin >> word;
    tree.NaiveBuild(word);
    unsigned numOfSubstr = 0;
    std::string substr;
    while (std::cin >> substr) {
        ++numOfSubstr;
        std::vector<unsigned> ans = tree.Find(substr);
        if (!ans.empty()) {
            std::sort(ans.begin(), ans.end());
            std::cout << numOfSubstr << ": ";
            for (unsigned i = 0; i < ans.size() - 1; ++i) {
                std::cout << ans[i] << ", ";
            }
            std::cout << ans[ans.size() - 1] << std::endl;
        }
    }
    return 0;
}
