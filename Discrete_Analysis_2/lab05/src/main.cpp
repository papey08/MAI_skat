#include <iostream>
#include <string>
#include <set>

#include "TSuffixTree.hpp"

int main() {

    std::string text;
    std::cin >> text;
    TSuffixTree tr(text);
    std::string word;
    unsigned numOfWord = 1;
    while (std::cin >> word) {
        std::set<unsigned> ans;
        ans = tr.Search(word);
        if (!ans.empty()) {
            std::cout << numOfWord << ": ";
            for (auto x: ans) {
                std::cout << x;
                if (x != *(ans.rbegin())) {
                    std::cout << ", ";
                } else {
                    std::cout << std::endl;
                }
            }
        }
        ++numOfWord;
    }
    return 0;
}
