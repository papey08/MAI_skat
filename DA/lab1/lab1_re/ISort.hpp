#ifndef ISORT_HPP
#define ISORT_HPP

#include "TPair.hpp"

void Count(std::vector<TPair>& raw, std::vector<int>& count) {
    for (int i = 0; i < raw.size(); ++i) {
        ++count[raw[i].GetKey()];
    }
    for (int i = 1; i < MAX_KEY; ++i) {
        count[i] += count[i - 1];
    }
}

void Sort(std::vector<TPair>& raw, std::vector<int>& count, std::vector<int>& ans) {
    for (int i = raw.size() - 1; i >= 0; --i) {
        --count[raw[i].GetKey()];
        ans[count[raw[i].GetKey()]] = i;
    }
}

#endif
