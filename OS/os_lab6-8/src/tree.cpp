#include <iostream>
#include <vector>
#include <algorithm>
#include "tree.h"

Tree::~Tree()
{
    delete_node(root);
}

void Tree::push(int id)
{
    root = push(root, id);
}

void Tree::kill(int id)
{
    root = kill(root, id);
}

void Tree::delete_node(Node* node)
{
    if(node == NULL)
	{
		return;
	}
    delete_node(node->right);
    delete_node(node->left);
    delete node;
}

std::vector<int> Tree::get_nodes()
{
    std::vector<int> result;
    get_nodes(root, result);
    return result;
}

void Tree::get_nodes(Node* node, std::vector<int>& v)
{
    if (node == NULL)
	{
		return;
	}
    get_nodes(node->left, v);
    v.push_back(node->id);
    get_nodes(node->right, v);
}

Node* Tree::push(Node* root, int val)
{
    if (root == NULL)
	{
        root = new Node;
        root->id = val;
        root->left = NULL;
        root->right = NULL;
        return root;
    }
    else if (val < root->id)
	{
        root->left = push(root->left, val);
    }
    else if (val >= root->id)
	{
        root->right = push(root->right, val);
    }
    return root;
}

Node* Tree::kill(Node* root_node, int val)
{
    Node* node;
    if (root_node == NULL)
	{
        return NULL;
    }
    else if (val < root_node->id) 
	{
        root_node->left = kill(root_node->left, val); 
    }
    else if (val >root_node->id)
	{
        root_node->right = kill(root_node->right, val); 
    }
    else
	{
        node = root_node;
        if (root_node->left == NULL)
		{
            root_node = root_node->right; 
        }
        else if (root_node->right == NULL)
		{
             root_node = root_node->left; 
        }
        delete node;
    }
    if (root_node == NULL)
	{
         return root_node; 
    }
	return root_node;
}
