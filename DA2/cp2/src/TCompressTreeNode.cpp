#include "TCompressTreeNode.hpp"

using namespace NHuffman;

TCompressTreeNode::TCompressTreeNode(char symbol, bool isLeaf, bool isArt,
                       unsigned int frequency, TCompressTreeNode *left,
                       TCompressTreeNode *right, TCompressTreeNode *parent) :

                       symbol(symbol),
                       isLeaf(isLeaf),
                       isArt(isArt),
                       frequency(frequency),
                       left(left),
                       right(right),
                       parent(parent) {}

TCompressTreeNode::~TCompressTreeNode() {
    delete left;
    delete right;
}
