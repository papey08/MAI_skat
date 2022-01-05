#ifndef TBINARYTREE_H
#define TBINARYTREE_H

#include "TBinaryTreeItem.h"

using namespace std;

class TBinaryTree 
{
private:
    TBinaryTreeItem *node;

public:
    TBinaryTree();
    void Push(const Triangle& tr);
    const Triangle& GetItemNotLess(double area);
    size_t Count(const Triangle& t);
    void Pop(const Triangle& t);
    bool Empty();
    friend ostream& operator<<(ostream& os, const TBinaryTree& tree);
    void Clear();
    virtual ~TBinaryTree();
};

#endif
