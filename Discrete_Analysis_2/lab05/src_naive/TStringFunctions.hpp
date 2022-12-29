#ifndef TSTRINGFUNCTIONS_HPP
#define TSTRINGFUNCTIONS_HPP

#include <string>

unsigned BestMatch(std::string& str1, std::string& str2) {
    unsigned res = 0;
    while ((str1[res] == str2[res]) && (res < str1.size()) &&
           (res < str2.size()))
    {
        ++res;
    }
    return res;
}

std::string GetSuffix(std::string& str, unsigned pos) {
    std::string res;
    for (unsigned i = pos; i < str.size(); ++i) {
        res += str[i];
    }
    return res;
}

std::string GetPrefix(std::string& str, unsigned pos) {
    std::string res;
    for (unsigned i = 0; i < pos; ++i) {
        res += str[i];
    }
    return res;
}

#endif
