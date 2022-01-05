#ifndef TBINARYTREE_ITEM_H
#define TBINARYTREE_ITEM_H

#include "triangle.h"

template<class T>
class TBinaryTreeItem
{
public:
    TBinaryTreeItem(const T& tri);
    TBinaryTreeItem(const TBinaryTreeItem<T>& other);
    virtual ~TBinaryTreeItem();
    T tri;
    shared_ptr<TBinaryTreeItem<T>> left;
    shared_ptr<TBinaryTreeItem<T>> right;
    unsigned counter;
    
    template<class TT>
    friend ostream &operator<<(ostream &os, const TBinaryTreeItem<TT> &t);
};

#endif
