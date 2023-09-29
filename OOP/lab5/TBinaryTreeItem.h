#ifndef TBINARYTREE_ITEM_H
#define TBINARYTREE_ITEM_H

#include "triangle.h"

class TBinaryTreeItem 
{
public:
    TBinaryTreeItem(const Triangle& tri);
    TBinaryTreeItem(const TBinaryTreeItem& other);
    virtual ~TBinaryTreeItem();
    Triangle tri;
    shared_ptr<TBinaryTreeItem> left;
    shared_ptr<TBinaryTreeItem> right;
    unsigned counter;
};

#endif
