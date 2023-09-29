#ifndef TCOMPACTTRIE_HPP
#define TCOMPACTTRIE_HPP

#include "TStringFunctions.hpp"

#include <string>
#include <vector>

const char SENTINEL = '$';

class TCompactTrie {
    struct TCompactTrieItem {
        std::string value;
        std::vector<TCompactTrieItem*> children;
        unsigned num{};
        bool isLeaf{true};

        ~TCompactTrieItem() {
            for (auto& x: children) {
                delete x;
            }
        }
    };

    static void Split(TCompactTrieItem*& toSplit, unsigned pos) {
        auto* itemAfterSplit = new TCompactTrieItem;
        itemAfterSplit->value = GetSuffix(toSplit->value, pos);
        itemAfterSplit->children = std::move(toSplit->children);
        itemAfterSplit->num = toSplit->num;
        itemAfterSplit->isLeaf = itemAfterSplit->children.empty();
        toSplit->value = GetPrefix(toSplit->value, pos);
        toSplit->children.push_back(itemAfterSplit);
    }

    void AddToItem(TCompactTrieItem*& current, std::string& toAdd, unsigned n) {
        unsigned bm = BestMatch(current->value, toAdd);
        if (bm == 0) {
            auto* itemToAdd = new TCompactTrieItem;
            itemToAdd->value = toAdd;
            itemToAdd->num = n;
            itemToAdd->isLeaf = true;
            current->children.push_back(itemToAdd);
            current->isLeaf = false;
        } else if (bm < current->value.length()) {
            Split(current, bm);
            auto* itemToAdd = new TCompactTrieItem;
            itemToAdd->value = GetSuffix(toAdd, bm);
            if (toAdd.empty()) {
                return;
            }
            itemToAdd->num = n;
            itemToAdd->isLeaf = true;
            current->children.push_back(itemToAdd);
            current->isLeaf = false;
        } else if (bm >= current->value.length()) {
            toAdd = GetSuffix(toAdd, bm);
            if (toAdd.empty()) {
                return;
            }
            for (auto& x: current->children) {
                if (x->value[0] == toAdd[0]) {
                    TCompactTrieItem* itemToAdd = x;
                    AddToItem(itemToAdd, toAdd, n);
                    return;
                }
            }
            auto* itemToAdd = new TCompactTrieItem;
            itemToAdd->value = toAdd;
            itemToAdd->num = n;
            itemToAdd->isLeaf = true;
            current->children.push_back(itemToAdd);
        }
    }

    void FillVector(TCompactTrieItem*& current, std::vector<unsigned>& res) {
        for (auto& x: current->children) {
            if (x->isLeaf) {
                res.push_back(x->num);
            } else {
                FillVector(x, res);
            }
        }
    }

    void FindVectorOfItems(TCompactTrieItem*& current,
                           std::vector<unsigned>& res,
                           std::string& toFind)
    {
        if (toFind.empty()) {
            if (current->isLeaf) {
                res.push_back(current->num);
            } else {
                FillVector(current, res);
            }
            return;
        }
        for (auto& x: current->children) {
            if (x->value.length() == BestMatch(toFind, x->value)) {
                TCompactTrieItem* nextCurrent = x;
                std::string nextToFind = GetSuffix(toFind, x->value.length());
                FindVectorOfItems(nextCurrent, res, nextToFind);
            } else if (toFind.length() == BestMatch(toFind, x->value)) {
                if (x->isLeaf) {
                    res.push_back(x->num);
                } else {
                    FillVector(x, res);
                }
                return;
            }
        }
    }


    TCompactTrieItem* root;

public:

    TCompactTrie() {
        root = new TCompactTrieItem;
        root->value = SENTINEL;
        root->isLeaf = true;
    }

    void Add(std::string toAdd, unsigned n) {
        root->isLeaf = false;
        toAdd = SENTINEL + toAdd + SENTINEL;
        AddToItem(root, toAdd, n);
    }

    std::vector<unsigned> Find(std::string& toFind) {
        std::vector<unsigned> res;
        if (toFind.empty()) {
            return res;
        }
        FindVectorOfItems(root, res, toFind);
        return res;
    }

    ~TCompactTrie() {
        delete root;
    }

};


#endif
