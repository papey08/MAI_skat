#ifndef TDECOMPRESSTREENODE_HPP
#define TDECOMPRESSTREENODE_HPP

namespace NHuffman {

    class TDecompressTreeNode {

        friend class TDecompressTree;

        char symbol;
        bool isSymbol;
        TDecompressTreeNode *left;
        TDecompressTreeNode *right;

    public:
        TDecompressTreeNode(char symbol, bool isSymbol,
                            TDecompressTreeNode *left,
                            TDecompressTreeNode *right);

        virtual ~TDecompressTreeNode();
    };
}

#endif
