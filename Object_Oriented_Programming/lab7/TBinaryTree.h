#ifndef TBINARYTREE_H
#define TBINARYTREE_H

#include "TBinaryTreeItem.h"

using namespace std;

template <class T>
class TBinaryTree 
{
private:
    shared_ptr <TBinaryTreeItem<T>> node;

public:
    TBinaryTree();
    void Push(const T& tr);
    const T& GetItemNotLess(double area);
    size_t Count(const T& t);
    void Pop(const T& t);
    bool Empty();
    template <class TT>
    friend ostream& operator<<(ostream& os, const TBinaryTree<TT>& tree);
    void Clear();
    virtual ~TBinaryTree();
};

#endif
