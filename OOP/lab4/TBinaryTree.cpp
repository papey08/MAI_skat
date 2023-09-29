#include "TBinaryTree.h"

using namespace std;

TBinaryTree::TBinaryTree()
{
    node = NULL;
}

void print_tree(ostream& os, TBinaryTreeItem* node)
{
    if (!node)
    {
        return;
    }
    if (node->left)
    {
        os << node->counter << "*" << node->tri.GetArea() << ": [";
        print_tree(os, node->left);
        if (node->right)
        {
            os << ", ";
            print_tree(os, node->right);
        }
        os << "]";
    } 
    else if (node->right)
    {
       os << node->counter << "*" << node->tri.GetArea() << ": [";
        print_tree(os, node->right);
        if (node->left) 
        {
            os << ", ";
            print_tree(os, node->left);
        }
        os << "]";
    }
    else
    {
        os << node->counter << "*" << node->tri.GetArea();
    }
}

std::ostream& operator << (ostream& os, const TBinaryTree& tree)
{
    print_tree(os, tree.node);
    os;
    return os;
}

void TBinaryTree::Push(const Triangle &tr) 
{
    Triangle t = tr;
    if (node == NULL) 
    {
        node = new TBinaryTreeItem(t);
    }
    else if (node->tri.GetArea() == t.GetArea()) 
    {
        node->counter++;
    }
    else 
    {
        TBinaryTreeItem* prev = node;
        TBinaryTreeItem* cur;
        bool bebra = true;
        if (t.GetArea() < prev->tri.GetArea()) 
        {
            cur = node->left;
        }
        else if (t.GetArea() > prev->tri.GetArea()) 
        {
            cur = node->right;
            bebra = false;
        }
        while (cur != NULL) 
        {
            if (cur->tri == t) 
            {
                cur->counter++;
            }
            else 
            {
                if (t.GetArea() < cur->tri.GetArea()) 
                {
                    prev = cur;
                    cur = prev->left;
                    bebra = true;
                }
                else if (t.GetArea() > cur->tri.GetArea()) 
                {
                    prev = cur;
                    cur = prev->right;
                    bebra = false;
                }
            }
        }
        cur = new TBinaryTreeItem(t);
        if (bebra == true) 
        {
            prev->left = cur;
        }
        else 
        {
            prev->right = cur;
        }
    }
}

TBinaryTreeItem* __Pop(TBinaryTreeItem* node)
{
    if (node->left == NULL) 
    {
        return node;
    }
    return __Pop(node->left);
}

TBinaryTreeItem* _Pop(TBinaryTreeItem* node, Triangle &t)
{
    if (node == NULL) 
    {
        return node;
    }
    else if (t.GetArea() < node->tri.GetArea()) 
    {
        node->left = _Pop(node->left, t);
    }
    else if (t.GetArea() > node->tri.GetArea()) 
    {
        node->right = _Pop(node->right, t);
    }
    else 
    {
        if (node->left == NULL && node->right == NULL) 
        {
            if (node->counter > 1)
            {
                --node->counter;
                return node;
            }
            node = NULL;
            delete node;
            return node;
        }
        else if (node->left == NULL && node->right != NULL) 
        {
            if (node->counter > 1)
            {
                --node->counter;
                return node;
            }
            node = node->right;
            node->right = NULL;
            delete node->right;
            return node;
        }
        else if (node->right == NULL && node->left != NULL) 
        {
            if (node->counter > 1)
            {
                --node->counter;
                return node;
            }
            node = node->left;
            node->left = NULL;
            delete node->left;
            return node;
        }
        else 
        {
            TBinaryTreeItem* bebra = __Pop(node->right);
            node->tri.A = bebra->tri.GetArea();
            node->right = _Pop(node->right, bebra->tri);
        }
    }
    return node;
}

void TBinaryTree::Pop(const Triangle &t)
{
    Triangle tr = t;
    node = _Pop(node, tr);
}

unsigned _Count(TBinaryTreeItem* cur, unsigned res, Triangle& t)
{
    if (cur != NULL) 
    {
        _Count(cur->left, res, t);
        _Count(cur->right, res, t);
        if (cur->tri.GetArea() == t.GetArea()) 
        {
            return cur->counter;
        }
    }
    return 0;
}

size_t TBinaryTree::Count(const Triangle& t)
{
    Triangle tr = t;
    return _Count(node, 0, tr);
}

Triangle bebra;

Triangle& _GetItemNotLess(double area, TBinaryTreeItem* node)
{
    if (node->tri.GetArea() >= area)
    {
        return node->tri;
    }
    else
    {
        _GetItemNotLess(area, node->right);
    }
    return bebra;
}

const Triangle& TBinaryTree::GetItemNotLess(double area)
{
    return _GetItemNotLess(area, node);
}

void _Clear(TBinaryTreeItem* cur)
{
    if (cur!= NULL)
    {
        _Clear(cur->left);
        _Clear(cur->right);
        cur = NULL;
        delete cur;
    }
}

void TBinaryTree::Clear()
{
    _Clear(node);
    delete node;
    node = NULL;
}

bool TBinaryTree::Empty()
{
    return (node == NULL);
}

TBinaryTree::~TBinaryTree()
{
    Clear();
}
