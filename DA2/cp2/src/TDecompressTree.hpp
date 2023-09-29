#ifndef TDECOMPRESSTREE_HPP
#define TDECOMPRESSTREE_HPP

#include <map>
#include <string>

#include "TDecompressTreeNode.hpp"

namespace NHuffman {

    class TDecompressTree {

        TDecompressTreeNode *root;
        unsigned long long compressedBytes = 0;
        unsigned long long decompressedBytes = 0;
        std::string name;

        // loading map of pairs char-code from the file
        static std::map<char, std::string> LoadKey(std::string &filename);

    public:
        explicit TDecompressTree(std::string &filename, bool flagC);

        // getting info for -l flag
        std::string GetInfo() const;

        virtual ~TDecompressTree();
    };
}

#endif
