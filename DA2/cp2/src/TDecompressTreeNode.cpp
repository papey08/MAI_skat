#include "TDecompressTreeNode.hpp"

using namespace NHuffman;

TDecompressTreeNode::TDecompressTreeNode(char symbol, bool isSymbol,
                                         TDecompressTreeNode *left,
                                         TDecompressTreeNode *right) :
                                         symbol(symbol),
                                         isSymbol(isSymbol),
                                         left(left),
                                         right(right) {}

TDecompressTreeNode::~TDecompressTreeNode() {
    delete left;
    delete right;
}
