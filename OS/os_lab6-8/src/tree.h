#pragma once
#include <vector>

struct Node
{
    int id;
    Node* left;
    Node* right; 
};

class Tree
{
public:
    void push(int);
    void kill(int);
    std::vector<int> get_nodes();
    ~Tree();
private:
    Node* root = NULL;
    Node* push(Node* t, int);
    Node* kill(Node* t, int);
    void get_nodes(Node*, std::vector<int>&);
    void delete_node(Node*);
};
