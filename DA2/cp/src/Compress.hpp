#ifndef CODE_HPP
#define CODE_HPP

#include "TTriplet.hpp"

#include <string>
#include <vector>
#include <algorithm>

std::vector<TTriplet> Compress(std::string& str) {
    std::vector<TTriplet> res;
    int ptr = 0;
    while (ptr < str.length()) {
        int maxOffset = 0;
        int maxSize = 0;
        char nextSymbol = str[ptr];
        for (int i = 0; i < ptr; ++i) {
            if (str[i] == str[ptr]) {
                int tempOffset = ptr - i;
                int tempSize = 1;
                char tempNext = str[ptr + 1];
                for (int j = 1;
                     j + i < str.length() && str[j + i] == str[ptr + j]; ++j)
                {
                    ++tempSize;
                    tempNext = str[ptr + tempSize];
                }
                if (tempSize >= maxSize) {
                    maxOffset = tempOffset;
                    maxSize = tempSize;
                    if (ptr + maxSize == str.length()) {
                        nextSymbol = EPS;
                    } else {
                        nextSymbol = tempNext;
                    }
                }
            }
        }
        res.push_back({maxOffset, maxSize, nextSymbol});
        ptr += maxSize + 1;
    }
    return res;
}

#endif
