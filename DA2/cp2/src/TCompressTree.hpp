#ifndef THTREE_HPP
#define THTREE_HPP

#include <map>
#include <string>
#include <fstream>
#include <vector>

#include "TCompressTreeNode.hpp"

namespace NHuffman {

    class TCompressTree {
        TCompressTreeNode *root;
        std::string filename;
        std::vector<char> origText;
        unsigned long long compressedBytes = 0;

        // methods for building the tree
        static void SwapNodes(TCompressTreeNode *a, TCompressTreeNode *b);

        void BalanceNodes(TCompressTreeNode *node);

        TCompressTreeNode *GetArt();

        TCompressTreeNode *CheckSymbol(char c);

        void NextSymbol(char c);

        // methods for saving Huffman codes
        void Dfs(TCompressTreeNode *current, std::string &code,
                 std::map<char, std::string> &codes);

        std::map<char, std::string> Codes();

        void SaveKey();

        // method for saving compressed version of file
        void Compress(bool flagC);

    public:
        explicit TCompressTree(std::string &name, bool flagC);

        std::string GetInfo();

        virtual ~TCompressTree();
    };
}

#endif
