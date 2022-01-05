#include "TBinaryTree.h"

using namespace std;

template <class T>
TBinaryTree<T>::TBinaryTree() : node(NULL)
{}

template <class T>
void print_tree(ostream& os, shared_ptr <TBinaryTreeItem<T>> node)
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

template <class TT>
std::ostream& operator << (ostream& os, const TBinaryTree<TT>& tree)
{
    print_tree(os, tree.node);
    os;
    return os;
}

template <class T>
void TBinaryTree<T>::Push(const T &tr) 
{
    T t = tr;
    if (node == NULL) 
    {
        shared_ptr <TBinaryTreeItem<T>> c(new TBinaryTreeItem<T>(t));
        node = c;
    }
    else if (node->tri.GetArea() == t.GetArea()) 
    {
        node->counter++;
    }
    else 
    {
        shared_ptr<TBinaryTreeItem<T>> prev = node;
        shared_ptr<TBinaryTreeItem<T>> cur;
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
        shared_ptr<TBinaryTreeItem<T>> c(new TBinaryTreeItem<T>(t));
        cur = c;
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

template <class T>
shared_ptr<TBinaryTreeItem<T>> __Pop(shared_ptr<TBinaryTreeItem<T>> node)
{
    if (node->left == NULL) 
    {
        return node;
    }
    return __Pop(node->left);
}

template <class T>
shared_ptr<TBinaryTreeItem<T>> _Pop(shared_ptr<TBinaryTreeItem<T>> node, T &t)
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
            return node;
        }
        else 
        {
            shared_ptr<TBinaryTreeItem<T>> bebra = __Pop(node->right);
            node->tri.A = bebra->tri.GetArea();
            node->right = _Pop(node->right, bebra->tri);
        }
    }
    return node;
}

template <class T>
void TBinaryTree<T>::Pop(const T &t)
{
    T tr = t;
    node = _Pop(node, tr);
}

template <class T>
unsigned _Count(shared_ptr<TBinaryTreeItem<T>> cur, unsigned res, T& t)
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

template <class T>
size_t TBinaryTree<T>::Count(const T& t)
{
    T tr = t;
    return _Count(node, 0, tr);
}

template <class T>
T& _GetItemNotLess(double area, shared_ptr<TBinaryTreeItem<T>> node)
{
    if (node->tri.GetArea() >= area)
    {
        return node->tri;
    }
    else
    {
        _GetItemNotLess(area, node->right);
    }
}

template <class T>
const T& TBinaryTree<T>::GetItemNotLess(double area)
{
    return _GetItemNotLess(area, node);
}

template <class T>
void _Clear(shared_ptr<TBinaryTreeItem<T>> cur)
{
    if (cur!= NULL)
    {
        _Clear(cur->left);
        _Clear(cur->right);
        cur = NULL;
    }
}

template <class T>
void TBinaryTree<T>::Clear()
{
    _Clear(node);
    node = NULL;
}

template <class T>
bool TBinaryTree<T>::Empty()
{
    return (node == NULL);
}

template <class T>
TBinaryTree<T>::~TBinaryTree()
{
    Clear();
}

template class TBinaryTree<Triangle>;
template ostream& operator<<(ostream& os, const TBinaryTree<Triangle>& tr);
