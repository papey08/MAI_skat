#ifndef LZ77_HPP
#define LZ77_HPP

#include <string>
#include <vector>
#include <fstream>
#include <iostream>

namespace NLZ77 {

    const int EPS = 128;
    const int PERCENTS = 100;
    const int BYTE_SIZE = 8;
    const int SUFFIX_LENGTH = 5;

    struct TTriplet {
        int offset;
        int length;
        int nextSymbol;
    };

    std::vector<TTriplet> CompressVector(std::vector<char>& str) {
        std::vector<TTriplet> res;
        int ptr = 0;
        while (ptr < str.size()) {
            int maxOffset = 0;
            int maxSize = 0;
            int nextSymbol = str[ptr];
            for (int i = 0; i < ptr; ++i) {
                if (str[i] == str[ptr]) {
                    int tempOffset = ptr - i;
                    int tempSize = 1;
                    char tempNext = str[ptr + 1];
                    for (int j = 1;
                         j + i < str.size() &&
                         str[j + i] == str[ptr + j]; ++j)
                    {
                        ++tempSize;
                        tempNext = str[ptr + tempSize];
                    }
                    if (tempSize >= maxSize) {
                        maxOffset = tempOffset;
                        maxSize = tempSize;
                        if (ptr + maxSize == str.size()) {
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

    void Compress(std::string& filename, bool flagC, bool flagL) {
        std::ifstream fileToRead(filename);
        std::vector<char> origText = std::vector<char>
                (std::istreambuf_iterator<char>(fileToRead),
                 std::istreambuf_iterator<char>());
        fileToRead.close();
        std::vector<TTriplet> compressed = CompressVector(origText);

        if (!flagC) {
            std::ofstream fileToWrite(filename + ".lz77");
            for (auto x: compressed) {

                fileToWrite << x.offset << " " << x.length << " "
                            << x.nextSymbol
                            << std::endl;
            }
            fileToWrite.close();
        } else {
            for (auto x: compressed) {
                std::cout << x.offset << " " << x.length << " "
                            << x.nextSymbol
                            << std::endl;
            }
        }

        if (flagL) {
            std::cout << filename + ":\n";
            std::cout << "LZ77\n";
            std::cout << "compressed size: " << compressed.size() * BYTE_SIZE
                        << " bytes\n";
            std::cout << "uncompressed size: " << origText.size() << " bytes\n";
        }
    }

    std::string DecompressVector(std::vector<TTriplet>& triples) {
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

    void Decompress(std::string& filename, bool flagC, bool flagL) {
        std::ifstream fileToRead(filename);
        std::vector<TTriplet> triplets;
        TTriplet temp{};
        while (fileToRead >> temp.offset >> temp.length >> temp.nextSymbol) {
            triplets.push_back(temp);
        }
        fileToRead.close();
        std::string res = DecompressVector(triplets);

        if (!flagC) {
            std::string newFilename = "decompressed_" +
                                      filename.substr(0,
                                          filename.length() - SUFFIX_LENGTH);
            std::ofstream fileToWrite(newFilename);
            fileToWrite << res;
            fileToWrite.close();
        } else {
            std::cout << res;
        }

        if (flagL) {
            std::cout << filename + ":\n";
            std::cout << "LZ77\n";
            std::cout << "compressed size: " << triplets.size() * BYTE_SIZE
                        << " bytes\n";
            std::cout << "uncompressed size: " << res.size() << " bytes\n";
        }
    }
}

#endif
