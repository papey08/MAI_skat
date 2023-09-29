#include "TBinaryTreeItem.h"


template <class T>
TBinaryTreeItem<T>::TBinaryTreeItem(const T &t) 
{
    this->tri = t;
    this->left = NULL;
    this->right = NULL;
    this->counter = 1;
}

template <class T>
TBinaryTreeItem<T>::TBinaryTreeItem(const TBinaryTreeItem<T> &other) 
{
    this->tri = other.tri;
    this->left = other.left;
    this->right = other.right;
    this->counter = other.counter;
}

template <class T>
TBinaryTreeItem<T>::~TBinaryTreeItem() 
{}

template <class TT>
ostream& operator<<(ostream& os, TBinaryTreeItem<TT> tr)
{
    os << tr.tri << " ";
    return os;
}

#include "triangle.h"
template class TBinaryTreeItem<Triangle>;
template ostream& operator<<(ostream& os,  TBinaryTreeItem<Triangle> t);
