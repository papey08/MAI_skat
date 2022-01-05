#include "TBinaryTreeItem.h"

TBinaryTreeItem::TBinaryTreeItem(const Triangle &t) 
{
    this->tri = t;
    this->left = NULL;
    this->right = NULL;
    this->counter = 1;
}

TBinaryTreeItem::TBinaryTreeItem(const TBinaryTreeItem &other) 
{
    this->tri = other.tri;
    this->left = other.left;
    this->right = other.right;
    this->counter = other.counter;
}

TBinaryTreeItem::~TBinaryTreeItem() 
{}
