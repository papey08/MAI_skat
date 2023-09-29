#include <fstream>
#include <vector>
#include <map>
#include <iostream>

#include "TDecompressTree.hpp"

using namespace NHuffman;

const int BYTE_SIZE = 8;

// loading map of pairs char-code from the file
std::map<char, std::string> TDecompressTree::LoadKey(std::string& filename) {
    std::ifstream file;
    file.open(filename.substr(0, filename.length() - (BYTE_SIZE - 1)) + "key");
    std::map<char, std::string> codes;
    int symbol;
    std::string code;
    while (file >> symbol >> code) {
        codes[char(symbol)] = code;
    }
    file.close();
    return codes;
}

// converts char to a string containing 8-digit binary number
std::string ToBinaryString(char n) {
    std::string buffer;
    bool negative = n < 0;
    if (negative) {
        n *= -1;
    }
    do
    {
        buffer += char('0' + n % 2);
        n /= 2;
    } while (n > 0);
    buffer = std::string(buffer.rbegin(), buffer.rend());
    while (buffer.length() < (BYTE_SIZE - 1)) {
        buffer = '0' + buffer;
    }
    if (negative) {
        buffer = '1' + buffer;
    } else {
        buffer = '0' + buffer;
    }
    return buffer;
}

TDecompressTree::TDecompressTree(std::string& filename, bool flagC) {
    name = filename;
    root = new TDecompressTreeNode(0, false, nullptr, nullptr);
    std::map<char, std::string> codes = LoadKey(filename);

    // building the tree
    for (const auto& x : codes) {
        TDecompressTreeNode* current = root;
        for (char c : x.second) {
            if (c == '0') {
                if (current->left == nullptr) {
                    current->left = new TDecompressTreeNode(0, false, nullptr,
                                                            nullptr);
                    current = current->left;
                } else {
                    current = current->left;
                }
            } else if (c == '1') {
                if (current->right == nullptr) {
                    current->right = new TDecompressTreeNode(0, false, nullptr,
                                                             nullptr);
                    current = current->right;
                } else {
                    current = current->right;
                }
            }
        }
        current->symbol = x.first;
        current->isSymbol = true;
    }

    // loading compressed file

    std::ofstream decompressedFile;
    if (!flagC) {
        decompressedFile.open("decompressed_" +
                              filename.substr(0,filename.length() - BYTE_SIZE));
    }
    std::ifstream compressedFile;
    compressedFile.open(filename);
    std::vector<char> compressedText =
            std::vector<char>(std::istreambuf_iterator<char>(compressedFile),
                    std::istreambuf_iterator<char>());
    compressedBytes = compressedText.size();
    compressedFile.close();
    std::string binaryText;
    for (char c : compressedText) {
        std::string temp = ToBinaryString(c);
        binaryText += temp;
    }
    TDecompressTreeNode* current = root;

    //finding every original symbol by its code
    for (char c : binaryText) {
       if (current->isSymbol) {
           if (!flagC) {
               decompressedFile << current->symbol;
           } else {
               std::cout << current->symbol;
           }
            ++decompressedBytes;
            current = root;
        }
        if (c == '0') {
            if (current->left == nullptr) {
                break;
            }
            current = current->left;
        } else {
            current = current->right;
        }
    }
    if (current->isSymbol) {
        if (!flagC) {
            decompressedFile << current->symbol;
        } else {
            std::cout << current->symbol;
        }
        ++decompressedBytes;
    }
    if (!flagC) {
        decompressedFile.close();
    }
}

// getting info for -l flag
std::string TDecompressTree::GetInfo() const {
    std::string info;
    info += name + ":\n";
    info += "Huffman:\n";
    info += "compressed size: " + std::to_string(compressedBytes) + " bytes\n";
    info += "uncompressed size: " + std::to_string(decompressedBytes) +
            " bytes\n";
    info += "ratio: " +
            std::to_string(100 -((compressedBytes * 100) / decompressedBytes)) +
            "%";
    return info;
}

TDecompressTree::~TDecompressTree() {
    delete root;
}
