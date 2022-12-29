#include <queue>
#include <string>
#include <algorithm>
#include <map>
#include <vector>
#include <iostream>

#include "TCompressTree.hpp"

using namespace NHuffman;

const int BYTE_SIZE = 8;

void TCompressTree::SwapNodes(TCompressTreeNode* a, TCompressTreeNode* b) {
    TCompressTreeNode* parentA = a->parent;
    TCompressTreeNode* parentB = b->parent;
    b->parent = parentA;
    if (parentA->left == a) {
        parentA->left = b;
    } else {
        parentA->right = b;
    }
    a->parent = parentB;
    if (parentB->left == b) {
        parentB->left = a;
    } else {
        parentB->right = a;
    }
}

// swap 2 nodes to keep non-increasing order of the tree
void TCompressTree::BalanceNodes(TCompressTreeNode* node) {
    std::queue<TCompressTreeNode*> bfs;
    bfs.push(root);
    while (!bfs.empty()) {
        TCompressTreeNode* current = bfs.front();
        bfs.pop();
        if (current->left != nullptr) {
            bfs.push(current->left);
        }
        if (current->right != nullptr) {
            bfs.push(current->right);
        }
        if (current == node) {
            return;
        }
        if (current->frequency < node->frequency) {
            SwapNodes(current, node);
            return;
        }
    }
}

// finding the artifical node by dfs
TCompressTreeNode* TCompressTree::GetArt() {
    std::queue<TCompressTreeNode*> bfs;
    bfs.push(root);
    while (!bfs.empty()) {
        TCompressTreeNode* current = bfs.front();
        bfs.pop();
        if (current->left != nullptr) {
            bfs.push(current->left);
        }
        if (current->right != nullptr) {
            bfs.push(current->right);
        }
        if (current->isArt) {
            return current;
        }
    }
    return nullptr;
}

// searching symbol in the tree by bfs
TCompressTreeNode* TCompressTree::CheckSymbol(char c) {
    std::queue<TCompressTreeNode*> bfs;
    bfs.push(root);
    while (!bfs.empty()) {
        TCompressTreeNode* current = bfs.front();
        bfs.pop();
        if (current->left != nullptr) {
            bfs.push(current->left);
        }
        if (current->right != nullptr) {
            bfs.push(current->right);
        }
        if (current->symbol == c && current->isLeaf && !current->isArt) {
            return current;
        }
    }
    // if symbol not in the tree
    return nullptr;
}

// inserting new symbol into the tree
// or incrementing frequency of already existing symbol
void TCompressTree::NextSymbol(char c) {
    TCompressTreeNode* current = CheckSymbol(c);
    // checking if symbol is already in the tree
    if (current != nullptr) {
        // getting into the root and updating frequency of the nodes
        while (current != root) {
            ++current->frequency;
            BalanceNodes(current);
            current = current->parent;
        }
        ++current->frequency;

        // case if symbol not in the tree
    } else {
        TCompressTreeNode* artNode = GetArt();
        TCompressTreeNode* newInsideNode;
        TCompressTreeNode* newSymbolNode;
        // initializing relationships between new nodes
        if (artNode != root) {
            newSymbolNode = new TCompressTreeNode(c, true, false, 1,
                                                  nullptr, nullptr, nullptr);
            newInsideNode = new TCompressTreeNode(0, false, false, 1,
                                                  newSymbolNode, artNode,
                                                  artNode->parent);
            artNode->parent->right = newInsideNode;
            artNode->parent = newInsideNode;
            newSymbolNode->parent = newInsideNode;
        } else {
            newSymbolNode = new TCompressTreeNode(c, true, false, 1,
                                                  nullptr, nullptr, root);
            auto* newArtNode = new TCompressTreeNode(0, true, true, 1, nullptr,
                                                     nullptr, root);
            root->left = newSymbolNode;
            root->right = newArtNode;
            root->isLeaf = false;
            root->isArt = false;
            root->frequency = 2;
            return;
        }
        // getting into the root and updating frequency of the nodes
        current = newInsideNode;
        while (current != root) {
            ++current->frequency;
            BalanceNodes(current);
            current = current->parent;
        }
        ++current->frequency;
    }
}

// required for building map of pairs char-code from the tree
void TCompressTree::Dfs(TCompressTreeNode* current, std::string& code,
                        std::map<char, std::string>& codes)
{
    if (current->isLeaf && !current->isArt) {
        codes[current->symbol] = code;
    }
    if (current->left != nullptr) {
        std::string temp = code + '0';
        Dfs(current->left, temp, codes);
    }
    if (current->right != nullptr) {
        std::string temp = code + '1';
        Dfs(current->right, temp, codes);
    }
}

std::map<char, std::string> TCompressTree::Codes() {
    std::map<char, std::string> codes;
    std::string temp;
    Dfs(root, temp, codes);
    return codes;
}

// saving map of pairs char-code to the file
void TCompressTree::SaveKey() {
    std::ofstream keyFile;
    keyFile.open(filename + ".key");
    std::map<char, std::string> codes = Codes();
    for (const auto& x : codes) {
        keyFile << int(x.first) << " " << x.second << '\n';
    }
    keyFile.close();
}

char GetPow(unsigned n) {
    switch (n) {
        case 0:
            return (char)1;
        case 1:
            return (char)2;
        case 2:
            return (char)4;
        case 3:
            return (char)8;
        case 4:
            return (char)16;
        case 5:
            return (char)32;
        case 6:
            return (char)64;
        default:
            return (char)0;
    }
}

// converts string containing 8-digit binary number into the char
char BinaryToChar(std::string& bin) {
    if (bin == "10000000") {
        return -128;
    }
    char res = 0;
    for (unsigned i = bin.length() - 1; i >= 1; --i) {
        if (bin[i] == '1') {
            res += GetPow(bin.length() - 1 - i);
        }
    }
    if (bin[0] == '1') {
        res *= -1;
    }
    return res;
}

// compressing text of original file to file with .huffman extention
void TCompressTree::Compress(bool flagC) {
    if (!flagC) {
        std::ofstream resFile;
        resFile.open(filename + ".huffman");
        std::map<char, std::string> codes = Codes();
        std::string temp;
        for (char c: origText) {
            temp += codes[c];
            while (temp.length() >= BYTE_SIZE) {
                std::string binByte = temp.substr(0, BYTE_SIZE);
                char temp2 = BinaryToChar(binByte);
                resFile << temp2;
                temp = temp.substr(BYTE_SIZE, temp.length() - BYTE_SIZE);
                ++compressedBytes;
            }
        }
        if (!temp.empty()) {
            while (temp.length() != BYTE_SIZE) {
                temp += '0';
            }
            resFile << BinaryToChar(temp);
            ++compressedBytes;
        }
        resFile.close();


    } else {
        std::map<char, std::string> codes = Codes();
        std::string temp;
        for (char c: origText) {
            temp += codes[c];
            while (temp.length() >= BYTE_SIZE) {
                std::string binByte = temp.substr(0, BYTE_SIZE);
                char temp2 = BinaryToChar(binByte);
                std::cout << temp2;
                temp = temp.substr(BYTE_SIZE, temp.length() - BYTE_SIZE);
                ++compressedBytes;
            }
        }
        if (!temp.empty()) {
            while (temp.length() != BYTE_SIZE) {
                temp += '0';
            }
            std::cout << BinaryToChar(temp);
            ++compressedBytes;
        }
    }
}

TCompressTree::TCompressTree(std::string& name, bool flagC) {
    filename = name;
    std::ifstream file(filename);
    origText = std::vector<char>(std::istreambuf_iterator<char>(file),
            std::istreambuf_iterator<char>());
    file.close();
    root = new TCompressTreeNode(0, true, true, 1, nullptr, nullptr, nullptr);
    for (char c : origText) {
        NextSymbol(c);
    }
    SaveKey();
    Compress(flagC);
}

// getting info for -l flag
std::string TCompressTree::GetInfo() {
    std::string info;
    info += filename + ":\n";
    info += "Huffman:\n";
    info += "compressed size: " + std::to_string(compressedBytes) + " bytes\n";
    info += "uncompressed size: " + std::to_string(origText.size()) +
            " bytes\n";
    info += "ratio: " +
            std::to_string(100 - ((compressedBytes * 100) / origText.size())) +
            "%";
    return info;
}

TCompressTree::~TCompressTree() = default;
