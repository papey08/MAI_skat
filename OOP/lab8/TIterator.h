#ifndef TITERATOR_H
#define TITERATOR_H

#include <memory>
#include "TBinaryTreeItem.h"
#include "TBinaryTree.h"

template <class Node, class T>
class TIterator
{
private:
    std::shared_ptr<Node> node;

public:
    TIterator(std::shared_ptr<Node> n)
    {
        node = n;
    }

    T& operator*()
    {
        return node->tri;
    }

    void Left()
    {
        if (node == NULL)
        {
            return;
        }
        node = node->left;
    }

    void Right()
    {
        if (node == NULL)
        {
            return;
        }
        node = node->right;
    }

    bool operator== (TIterator &i)
    {
        return node == i.node;
    }

    bool operator!= (TIterator &i)
    {
        return !(node == i.node);
    }

    
};

#endif
