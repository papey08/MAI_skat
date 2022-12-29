#ifndef TSUFFIXTREE_HPP
#define TSUFFIXTREE_HPP

#include <map>
#include <set>
#include <vector>
#include <algorithm>
#include <queue>

const char SENTINEL = '$';

class TSuffixTree {

    // Node of TSuffixTree
    // contains number of suffix in suffixNum,
    // positions of substring in leftBound & rightBound
    // map of children and suffixLink
    struct TNode {
        int suffixNum;
        int leftBound, rightBound;
        TNode* suffixLink;
        std::map<char, TNode*> children;

        TNode(int num, int left, int right) {
            suffixNum = num;
            leftBound = left;
            rightBound = right;
            suffixLink = nullptr;
        }

        ~TNode() {
            for (auto& x: children) {
                delete x.second;
            }
        }
    };

    // root of the tree
    TNode* root;

    // original text for usage of leftBound & rightBound
    std::string text;


    // the following 5 fields are for correct work of Ukkonen algorithm
    TNode* currentNode;
    int currentChild{};
    int position{};
    int suffixCount = 0;
    int childrenCounter = 0;


    //
    bool LeafCheck() {
        TNode* temp = currentNode->children[text[currentChild]];
        unsigned len = temp->rightBound - temp->leftBound + 1;
        if (position >= len) {
            currentChild += (int)len;
            position -= (int)len;
            currentNode = temp;
            return true;
        } else {
            return false;
        }
    }

    // Add the next suffix by its number in text
    void Add(unsigned numOfSuffix) {
        TNode* lastAdded = nullptr;
        ++suffixCount;
        while (suffixCount) {
            if (position == 0) {
                currentChild = (int)numOfSuffix;
            }
            // if currentNode contains the next letter we need
            if (currentNode->children.find(text[currentChild]) !=
                currentNode->children.end())
            {
                if (LeafCheck()) {
                    continue;
                } else {
                    TNode* temp = currentNode->children[text[currentChild]];
                    if (text[temp->leftBound + position] == text[numOfSuffix]) {
                        if (lastAdded != nullptr) {
                            lastAdded->suffixLink = currentNode;
                        }
                        ++position;
                        break;
                    } else {
                        TNode* temp2 =
                                currentNode->children[text[currentChild]];
                        auto* toSplit = new TNode(-1,
                                                  temp2->leftBound,
                                                  temp2->leftBound +
                                                    position - 1);
                        currentNode->children[text[currentChild]] = toSplit;
                        temp2->leftBound += position;
                        toSplit->children[text[temp2->leftBound]] = temp2;
                        ++childrenCounter;
                        auto* temp3 =
                                new TNode(childrenCounter,
                                          (int)numOfSuffix,
                                          (int)text.length() - 1);
                        toSplit->children[text[numOfSuffix]] = temp3;
                        if (lastAdded != nullptr) {
                            lastAdded->suffixLink = toSplit;
                        }
                        lastAdded = toSplit;
                    }
                }
            } else {
                //creating a new node and suffix link
                ++childrenCounter;
                auto* temp =
                        new TNode(childrenCounter,
                                  (int)numOfSuffix,
                                  (int)text.length() - 1);
                currentNode->children[text[numOfSuffix]] = temp;
                if (lastAdded != nullptr) {
                    lastAdded->suffixLink = currentNode;
                    lastAdded = nullptr;
                }
            }
            --suffixCount;
            if (currentNode == root) {
                if (position > 0) {
                    --position;
                    ++currentChild;
                }
            } else {
                if (currentNode->suffixLink != nullptr) {
                    currentNode = currentNode->suffixLink;
                } else {
                    currentNode = root;
                }
            }
        }
    }

    // fills set ans with numbers of leafes growing from node curr
    void FindLeafesByBFS(std::set<unsigned>& ans) {
        std::queue<TNode*> bfs;
        bfs.push(currentNode);
        while (!bfs.empty()) {
            TNode* bfsNode = bfs.front();
            bfs.pop();
            if (bfsNode->suffixNum != -1) {
                ans.insert(bfsNode->suffixNum);
                continue;
            }
            for (auto& x: bfsNode->children) {
                if (x.second->suffixNum != -1) {
                    ans.insert(x.second->suffixNum);
                } else {
                    bfs.push(x.second);
                }
            }
        }
    }

    void FillSetWithEntries(std::string& word, std::set<unsigned>& ans) {
        currentNode = root;
        unsigned wordPointer = 0;

        // searching node which contains the word
        while (wordPointer < word.length()) {
            if (currentNode->children.find(word[wordPointer]) ==
                currentNode->children.end())
            {
                return;
            } else {
                currentNode = currentNode->children[word[wordPointer]];
                for (unsigned textPointer = currentNode->leftBound;
                     textPointer <= currentNode->rightBound;
                     ++wordPointer, ++textPointer)
                {
                    if (wordPointer >= word.length()) {
                        break;
                    }
                    if (text[textPointer] != word[wordPointer]) {
                        return;
                    }
                }
            }
        }

        FindLeafesByBFS(ans);
    }


public:

    explicit TSuffixTree(std::string& textToBuild) {
        text = textToBuild + SENTINEL;
        root = new TNode(-1, -1, -1);
        currentNode = root;
        for (unsigned i = 0; i < text.length(); ++i) {
            Add(i);
        }
    }

    std::set<unsigned> Search(std::string& word) {
        std::set<unsigned> ans;
        FillSetWithEntries(word, ans);
        return ans;
    }

    ~TSuffixTree() {
        delete root;
    }
};

#endif
