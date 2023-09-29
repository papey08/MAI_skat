#ifndef THTREENODE_HPP
#define THTREENODE_HPP


namespace NHuffman {

    class TCompressTreeNode {

        friend class TCompressTree;

        char symbol;
        bool isLeaf;
        bool isArt;
        unsigned frequency;

        TCompressTreeNode *left;
        TCompressTreeNode *right;
        TCompressTreeNode *parent;

    public:
        TCompressTreeNode(char symbol, bool isLeaf, bool isArt,
                          unsigned int frequency,
                          TCompressTreeNode *left, TCompressTreeNode *right,
                          TCompressTreeNode *parent);

        ~TCompressTreeNode();
    };
}

#endif
