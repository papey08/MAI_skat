#ifndef HUFFMAN_HPP
#define HUFFMAN_HPP

#include <string>
#include <iostream>

#include "TCompressTree.hpp"
#include "TDecompressTree.hpp"

namespace NHuffman {

    const int MIN_CHAR = -128;
    const int MAX_CHAR = 127;

    void Compress(std::string& filename, bool flagC, bool flagL) {
        auto tree = TCompressTree(filename, flagC);
        if (flagL) {
            std::string info = tree.GetInfo();
            std::cout << info << std::endl;
        }
    }

    void Decompress(std::string& filename, bool flagC, bool flagL) {
        auto tree = TDecompressTree(filename, flagC);
        if (flagL) {
            std::string info = tree.GetInfo();
            std::cout << info << std::endl;
        }
    }

    // function for -t flag
    bool CheckKey(std::string& filename) {
        std::ifstream file(filename);
        int symbol;
        std::string code;
        while (file >> symbol >> code) {
            if ((symbol < MIN_CHAR) || (symbol > MAX_CHAR)) {
                file.close();
                return false;
            }
            for (char c : code) {
                if ((c != '0') && (c != '1')) {
                    file.close();
                    return false;
                }
            }
        }
        file.close();
        return true;
    }
}

#endif
