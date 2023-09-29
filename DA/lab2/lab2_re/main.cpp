#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
#include <fstream>

#include "TBTree.hpp"

void StringToChars(char buf[MAX_KEY_LENGTH], std::string& str) {
    for (unsigned i = 0; i < str.length(); ++i) {
        buf[i] = str[i];
    }
    for (unsigned i = str.length(); i < MAX_KEY_LENGTH; ++i) {
        buf[i] = 0;
    }
}

void ToLower(char buf[MAX_KEY_LENGTH], TPair& curPair) {
    bool toClean = false;
    for (unsigned i = 0; i < MAX_KEY_LENGTH; ++i) {
        if (toClean == true) {
            curPair.key[i] = 0;
            buf[i] = 0;
        }
        if (('A' <= buf[i]) && (buf[i] <= 'Z') && toClean == false) {
            buf[i] = buf[i] - 'A' + 'a';
        }
        if ((buf[i] == 0) || (buf[i] == ' ')) {
            toClean = true;
            buf[i] = 0;
        }
        curPair.key[i] = buf[i];
    }
}

int main() {
    std::ios_base::sync_with_stdio(false);
    std::cin.tie(nullptr);
    TBTree btree;
    char buf[MAX_KEY_LENGTH];
    Clear(buf);
    std::string command;
    unsigned long long curValue;
    TPair curPair;
    TBTreeNode* curSearchNode;
    unsigned curSearchPosition;
    while (std::cin >> command) {
        if (command == "+") {
            std::cin >> buf;
            std::cin >> curValue;
            ToLower(buf, curPair);
            curPair.value = curValue;
            curSearchNode = nullptr;
            curSearchPosition = 0;
            btree.Find(curSearchNode, curSearchPosition, curPair);
            if (curSearchNode != nullptr) {
                std::cout << "Exist" << std::endl;
            }
            else {
                btree.Push(curPair);
                std::cout << "OK" << std::endl;
            }
        }
        else if (command == "-") {
            std::cin >> buf;
            ToLower(buf, curPair);
            curSearchNode = nullptr;
            curSearchPosition = 0;
            btree.Find(curSearchNode, curSearchPosition, curPair);
            if (curSearchNode != nullptr) {
                btree.Pop(curPair);
                std::cout << "OK" << std::endl;
            }
            else {
                std::cout << "NoSuchWord" << std::endl;
            }
        }
        else if (command == "!") {
            std::cin >> command;
            if (command == "Load") {
                std::string filePath;
                std::cin >> filePath;
                std::ifstream file(filePath, std::ios::binary);
                btree.Load3(file);
                file.close();
                std::cout << "OK\n";
            }
            else {
                std::string filePath;
                std::cin >> filePath;
                std::ofstream file(filePath, std::ios::trunc | std::ios::binary);
                unsigned size = btree.GetSize();
                file.write(reinterpret_cast<char *>(&size), sizeof(unsigned));
                if (size > 0) {
                    btree.Save3(file);
                }
                file.close();
                std::cout << "OK\n";
            }
        }
        else {
            StringToChars(buf, command);
            ToLower(buf, curPair);
            curSearchNode = nullptr;
            curSearchPosition = 0;
            btree.Find(curSearchNode, curSearchPosition, curPair);
            if (curSearchNode != nullptr) {
                std::cout << "OK: " 
                          << curSearchNode->pairs[curSearchPosition].value
                          << "" << std::endl;
            }
            else {
                std::cout << "NoSuchWord" << std::endl;
            }
        }
        Clear(curPair.key);
    }
    return 0;
}
