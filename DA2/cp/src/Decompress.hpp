#ifndef DECODE_HPP
#define DECODE_HPP

#include "TTriplet.hpp"

#include <string>
#include <vector>

std::string Decompress(std::vector<TTriplet>& triples) {
    std::string res;
    for (auto & triple : triples) {
        int ptr = (int)res.length();
        for (int j = 0; j < triple.length; ++j) {
            res.push_back(res[ptr - triple.offset + j]);
        }
        if (triple.nextSymbol == EPS) {
            break;
        }
        res.push_back(triple.nextSymbol);
    }
    return res;
}

#endif
