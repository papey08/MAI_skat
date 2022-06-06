#include <iostream>
#include <string>
#include <vector>
#include <chrono>
#include <algorithm>
#include <fstream>

const unsigned MAX_KEY_LENGTH = 257;
const unsigned TREE_DEGREE = 500;
const unsigned MAX_AMOUNT = 2 * TREE_DEGREE - 1;
// const unsigned IOS_BUF_MAX_SIZE = 100;

void Clear(char key[MAX_KEY_LENGTH]) {
    for (unsigned i = 0; i < MAX_KEY_LENGTH; ++i) {
        key[i] = 0;
    }
}

namespace IPair {
    
    struct TPair {
        char key[MAX_KEY_LENGTH];
        unsigned long long value;
        TPair();
        TPair(const TPair& it);
        TPair& operator = (const TPair& it);
    };

    TPair::TPair () {
        Clear(key);
        value = 0;
    }

    TPair::TPair(const TPair& it) {
        value = it.value;
        for (unsigned i = 0; i < MAX_KEY_LENGTH; ++i) {
            key[i] = it.key[i];
        }
    }

    TPair& TPair::operator = (const TPair& it) {
        value = it.value;
        for (unsigned i = 0; i < MAX_KEY_LENGTH; ++i) {
            key[i] = it.key[i];
        }
        return *this;
    }

    bool operator < (TPair& a, TPair& b) {
        for (unsigned i = 0; i < MAX_KEY_LENGTH; ++i) {
            if (a.key[i] != b.key[i]) {
                return a.key[i] < b.key[i];
            }
        }
        return false;
    }

    bool operator == (TPair& a, TPair& b) {
        for (unsigned short i = 0; i < MAX_KEY_LENGTH; ++i) {
            if (a.key[i] != b.key[i]) {
                return false;
            }
        }
        return true;
    }

    bool operator > (TPair& a, TPair& b) {
        return (!(a < b) && !(a == b));
    }

    bool operator >= (TPair& a, TPair& b) {
        return !(a < b);
    }

    bool operator <= (TPair& a, TPair& b) {
        return !(a > b);
    }

    bool operator != (TPair& a, TPair& b) {
        return !(a == b);
    }

};

namespace IBTree {

    using namespace IPair;

    struct TBTreeNode {
        std::vector<TPair> pairs;
        std::vector<TBTreeNode*> children;
        TBTreeNode() : pairs(1), children(2) {}
        ~TBTreeNode();
    };

    unsigned BinarySearch(std::vector<TPair>& v, TPair& element) {
        unsigned left = -1;
        unsigned right = v.size();
        while (left + 1 < right) {
            unsigned mid = (left + right) / 2;
            if (v[mid] < element) {
                left = mid;
            }
            else {
                right = mid;
            }
        }
        return right;
    }

    void FindNode(TBTreeNode* curNode, TBTreeNode*& res, unsigned& position,
                  TPair& element)
    {
        TBTreeNode* node = curNode;
        while (node != nullptr) {
            unsigned index = BinarySearch(node->pairs, element);
            if ((index < node->pairs.size()) && (element == node->pairs[index])) {
                res = node;
                position = index;
                return;
            }
            else {
                node = node->children[index];
            }
        }
        res = nullptr;
    }

    void SplitNode(TBTreeNode*& node, TBTreeNode*& toPushNode) {
        TBTreeNode* leftNode = new TBTreeNode;
        leftNode->pairs.resize(TREE_DEGREE - 1);
        leftNode->children.resize(TREE_DEGREE);
        TBTreeNode* rightNode = new TBTreeNode;
        rightNode->pairs.resize(TREE_DEGREE - 1);
        rightNode->children.resize(TREE_DEGREE);
        for (unsigned i = 0; i < TREE_DEGREE - 1; ++i) {
            leftNode->pairs[i] = node->pairs[i];
            leftNode->children[i] = node->children[i];
            rightNode->pairs[i] = node->pairs[TREE_DEGREE + i];
            rightNode->children[i] = node->children[TREE_DEGREE + i];
        }
        leftNode->children[TREE_DEGREE - 1] = node->children[TREE_DEGREE - 1];
        rightNode->children[TREE_DEGREE - 1] = node->children[MAX_AMOUNT];
        TBTreeNode* midNode = new TBTreeNode;
        midNode->pairs[0] = node->pairs[TREE_DEGREE - 1];
        midNode->children[0] = leftNode;
        midNode->children[1] = rightNode;
        toPushNode = midNode;
        for (unsigned i = 0; i < 2 * TREE_DEGREE; ++i) {
            node->children[i] = nullptr;
        }
        delete node;
    }

    void PushNode(TBTreeNode*& node, TBTreeNode*& toPushNode, TPair& element) {
        unsigned index = BinarySearch(node->pairs, element);
        if (node->children[index] == nullptr) {
            node->pairs.insert(node->pairs.begin() + index, element);
            node->children.insert(node->children.begin() + index, nullptr);
        }
        else {
            TBTreeNode* newToPushNode = nullptr;
            PushNode(node->children[index], newToPushNode, element);
            if (newToPushNode != nullptr) {
                node->children[index] = newToPushNode->children[1];
                node->pairs.insert(node->pairs.begin() + index,
                                   newToPushNode->pairs[0]);
                node->children.insert(node->children.begin() + index,
                                      newToPushNode->children[0]);
                newToPushNode->children[0] = nullptr;
                newToPushNode->children[1] = nullptr;
                delete newToPushNode;
            }
        }
        if (node->pairs.size() == MAX_AMOUNT) {
            SplitNode(node, toPushNode);
        }
    }
    
    void PopNode(TBTreeNode*& node, TPair& element) {
        TBTreeNode* curNode = node;
        std::vector< TBTreeNode* > path;
        std::vector<unsigned> indexes;
        unsigned index = BinarySearch(curNode->pairs, element);
        while (1) {
            path.push_back(curNode);
            indexes.push_back(index);
            if ((index < curNode->pairs.size()) && 
                (curNode->pairs[index] == element))
            {
                break;
            }
            curNode = curNode->children[index];
            index = BinarySearch(curNode->pairs, element);
        }
        if (curNode->children[index] != nullptr) {
            TBTreeNode* curPathNode = curNode->children[index];
            while (curPathNode->children[curPathNode->children.size() - 1] != 
                   nullptr)
            {
                curPathNode = curPathNode->children[curPathNode->children.size()
                 - 1];
            }
            TPair value = curPathNode->pairs[curPathNode->pairs.size() - 1];
            PopNode(node, value);
            TBTreeNode* curSearchNode = nullptr;
            unsigned position = 0;
            FindNode(node, curSearchNode, position, element);
            curSearchNode->pairs[position] = value;
            return;
        }
        while (path.size() > 1) {
            TBTreeNode* curPathNode = path[path.size() - 1];
            path.pop_back();
            unsigned curPathIndex = indexes[indexes.size() - 1];
            indexes.pop_back();
            if (curPathNode->pairs.size() == TREE_DEGREE - 1) {
                TBTreeNode* curPathNodeParent = path[path.size() - 1];
                unsigned curPathIndexParent = indexes[indexes.size() - 1];
                TBTreeNode* leftBrother = nullptr;
                if (curPathIndexParent > 0) {
                    leftBrother = curPathNodeParent->children[curPathIndexParent - 1];
                }
                if ((leftBrother != nullptr) &&
                    (leftBrother->pairs.size() > TREE_DEGREE - 1))
                {
                    curPathNode->pairs.erase(curPathNode->pairs.begin() + 
                                             curPathIndex);
                    curPathNode->children.erase(curPathNode->children.begin() + 
                                                curPathIndex);
                    curPathNode->pairs.insert(curPathNode->pairs.begin(), 
                                              curPathNodeParent->pairs[curPathIndexParent - 1]);
                    curPathNode->children.insert(curPathNode->children.begin(),
                                                 leftBrother->children[leftBrother->children.size() - 1]);
                    curPathNodeParent->pairs[curPathIndexParent - 1] = 
                    leftBrother->pairs[leftBrother->pairs.size() - 1];
                    leftBrother->pairs.pop_back();
                    leftBrother->children.pop_back();
                    return;
                    // PopNodeCase1(curPathNode, curPathIndex, curPathNodeParent, 
                    //              curPathIndexParent, leftBrother);
                }
                TBTreeNode* rightBrother = nullptr;
                if (curPathIndexParent < curPathNodeParent->children.size() - 1) {
                    rightBrother = curPathNodeParent->children[curPathIndexParent + 1];	
                }
                if ((rightBrother != nullptr) && 
                    (rightBrother->pairs.size() > TREE_DEGREE - 1))
                {
                    curPathNode->pairs.erase(curPathNode->pairs.begin() + 
                                             curPathIndex);
                    curPathNode->children.erase(curPathNode->children.begin() + 
                                                curPathIndex);
                    curPathNode->pairs.push_back(curPathNodeParent->pairs[curPathIndexParent]);
                    curPathNode->children.push_back(rightBrother->children[0]);
                    curPathNodeParent->pairs[curPathIndexParent] = 
                    rightBrother->pairs[0];
                    rightBrother->pairs.erase(rightBrother->pairs.begin());
                    rightBrother->children.erase(rightBrother->children.begin());
                    return;
                    // PopNodeCase2(curPathNode, curPathIndex, curPathNodeParent, 
                    //              curPathIndexParent, rightBrother);
                }
                curPathNode->pairs.erase(curPathNode->pairs.begin() + curPathIndex);
                curPathNode->children.erase(curPathNode->children.begin() + 
                                            curPathIndex);
                TBTreeNode* newNode = new TBTreeNode;
                newNode->pairs.resize(2 * TREE_DEGREE - 2);
                newNode->children.resize(MAX_AMOUNT);
                if (leftBrother != nullptr) {
                    indexes.pop_back();
                    indexes.push_back(curPathIndexParent - 1);
                    for (unsigned i = 0; i < TREE_DEGREE - 1; ++i) {
                        newNode->pairs[i] = leftBrother->pairs[i];
                        newNode->children[i] = leftBrother->children[i];
                    }
                    newNode->pairs[TREE_DEGREE - 1] = 
                    curPathNodeParent->pairs[curPathIndexParent - 1];
                    newNode->children[TREE_DEGREE - 1] = 
                    leftBrother->children[TREE_DEGREE - 1];
                    for (unsigned i = 0; i < TREE_DEGREE - 2; ++i) {
                        newNode->pairs[TREE_DEGREE + i] = curPathNode->pairs[i];
                        newNode->children[TREE_DEGREE + i] = 
                        curPathNode->children[i];
                    }
                    newNode->children[2 * TREE_DEGREE - 2] = 
                    curPathNode->children[TREE_DEGREE - 2];
                    for (unsigned i = 0; i < leftBrother->children.size(); ++i) {
                        leftBrother->children[i] = nullptr;
                    }
                    delete leftBrother;
                    for (unsigned i = 0; i < curPathNode->children.size(); ++i) {
                        curPathNode->children[i] = nullptr;
                    }
                    delete curPathNode;
                    curPathNodeParent->children[curPathIndexParent] = newNode;
                    curPathNodeParent->children[curPathIndexParent - 1] = nullptr;
                    // PopNodeCase3(indexes, newNode, curPathNode, curPathNodeParent, 
                    //              curPathIndexParent, leftBrother);
                }
                else if (rightBrother != nullptr) {
                    for (unsigned i = 0; i < TREE_DEGREE - 2; ++i) {
                        newNode->pairs[i] = curPathNode->pairs[i];
                        newNode->children[i] = curPathNode->children[i];
                    }
                    newNode->pairs[TREE_DEGREE - 2] = 
                    curPathNodeParent->pairs[curPathIndexParent];
                    newNode->children[TREE_DEGREE - 2] = 
                    curPathNode->children[TREE_DEGREE - 2];
                    for (unsigned i = 0; i < TREE_DEGREE - 1; ++i) {
                        newNode->pairs[TREE_DEGREE - 1 + i] = 
                        rightBrother->pairs[i];
                        newNode->children[TREE_DEGREE - 1 + i] = 
                        rightBrother->children[i];
                    }
                    newNode->children[2 * TREE_DEGREE - 2] = 
                    rightBrother->children[TREE_DEGREE - 1];
                    for (unsigned i = 0; i < curPathNode->children.size(); ++i) {
                        curPathNode->children[i] = nullptr;
                    }
                    delete curPathNode;
                    for (unsigned i = 0; i < rightBrother->children.size(); ++i) {
                        rightBrother->children[i] = nullptr;
                    }
                    delete rightBrother;
                    curPathNodeParent->children[curPathIndexParent] = nullptr;
                    curPathNodeParent->children[curPathIndexParent + 1] = newNode;
                    // PopNodeCase4(indexes, newNode, curPathNode, curPathNodeParent, 
                    //              curPathIndexParent, rightBrother);
                }
            }
            else {
                curPathNode->pairs.erase(curPathNode->pairs.begin() + curPathIndex);
                curPathNode->children.erase(curPathNode->children.begin() + 
                                            curPathIndex);
                return;
            }
        }
        TBTreeNode* curPathNode = path[path.size() - 1];
        path.pop_back();
        unsigned curPathIndex = indexes[indexes.size() - 1];
        indexes.pop_back();
        if (curPathNode->pairs.size() > 1) {
            curPathNode->pairs.erase(curPathNode->pairs.begin() + curPathIndex);
            curPathNode->children.erase(curPathNode->children.begin() + 
                                        curPathIndex);
        }
        else {
            if (curPathNode->children[0] == nullptr) {
                node = curPathNode->children[1];
                curPathNode->children[1] = nullptr;
            }
            else if (curPathNode->children[1] == nullptr) {
                node = curPathNode->children[0];
                curPathNode->children[0] = nullptr;
            }
            delete curPathNode;
        }
    }

    TBTreeNode::~TBTreeNode() {
        for (unsigned i = 0; i < children.size(); ++i) {
            delete children[i];
        }
    }


    class TBTree {
    private:
        TBTreeNode* root;
        unsigned size;
    public:
        TBTree();
        ~TBTree();
        unsigned GetSize();
        void Pop(TPair& element);
        void Find(TBTreeNode*& res, unsigned& position, TPair& element);
        void Push(TPair& element);
        void SaveNode3(TBTreeNode* curNode, std::ofstream& file);
        void Save3(std::ofstream& file);
        void Load3(std::ifstream& file);
        void Free();
    };

    TBTree::TBTree() {
        root = nullptr;
        size = 0;
    }

    TBTree::~TBTree() {
        delete root;
    }

    unsigned TBTree::GetSize() {
        return size;
    }

    void TBTree::Pop(TPair& element) {
        PopNode(root, element);
        --size;
    }

    void TBTree::Find(TBTreeNode*& res, unsigned& position, TPair& element) {
        FindNode(root, res, position, element);
    }

    void TBTree::Push(TPair& element) {
        ++size;
        if (root == nullptr) {
            root = new TBTreeNode;
            root->pairs[0] = element;
            root->children[0] = nullptr;
            root->children[1] = nullptr;
        }
        else {
            TBTreeNode* toPushNode = nullptr;
            PushNode(root, toPushNode, element);
            if (toPushNode != nullptr) {
                root = toPushNode;
            }
        }
    }

    void TBTree::SaveNode3(TBTreeNode* curNode, std::ofstream &file) {
        for (unsigned i = 0; i < curNode->pairs.size(); ++i) {
            for (unsigned j = 0; j < MAX_KEY_LENGTH; ++j) {
                file.write(reinterpret_cast<char *>(&curNode->pairs[i].key[j]), sizeof(char));
            }
            file.write(reinterpret_cast<char *>(&curNode->pairs[i].value), sizeof(unsigned long long));
            file.flush();
        }
        if (curNode->children[0] != nullptr) {
            for (unsigned i = 0; i < curNode->children.size(); ++i) {
                SaveNode3(curNode->children[i], file);
            }
        }
    }

    void TBTree::Save3(std::ofstream& file) {
        SaveNode3(root, file);
    }

    void TBTree::Load3(std::ifstream& file) {
        Free();
        unsigned size;
        file.read(reinterpret_cast<char *>(&size), sizeof(unsigned));
        for (unsigned i = 0; i < size; ++i) {
            TPair tempPair;
            for (unsigned j = 0; j < MAX_KEY_LENGTH; ++j) {
                file.read(reinterpret_cast<char *>(&tempPair.key[j]), sizeof(char));
            }
            file.read(reinterpret_cast<char *>(&tempPair.value), sizeof(unsigned long long));
            this->Push(tempPair);
        }
    }

    void TBTree::Free() {
        delete root;
        root = nullptr;
        size = 0;
    }

    void TreeLoad(TBTree& tree, FILE* file) {
        tree.Free();
        TPair cPair;
        unsigned newSize;
        fread(&newSize, sizeof(unsigned), 1, file);
        for (unsigned i = 0; i < newSize; ++i) {
            fread(&cPair.key, sizeof(char), MAX_KEY_LENGTH, file);
            fread(&cPair.value, sizeof(unsigned long long), 1, file);
            tree.Push(cPair);
        }
    }
};



using namespace IPair;
using namespace IBTree;

void StringToChars(char buf[MAX_KEY_LENGTH], std::string& str) {
    for (unsigned i = 0; i < str.length(); ++i) {
        buf[i] = str[i];
    }
    for (unsigned i = str.length(); i < MAX_KEY_LENGTH; ++i) {
        buf[i] = 0;
    }
}

void ToLower(char buf[MAX_KEY_LENGTH], TPair& curPair) {
    bool toClean = false;
    for (unsigned i = 0; i < MAX_KEY_LENGTH; ++i) {
        if (toClean == true) {
            curPair.key[i] = 0;
            buf[i] = 0;
        }
        if (('A' <= buf[i]) && (buf[i] <= 'Z') && toClean == false) {
            buf[i] = buf[i] - 'A' + 'a';
        }
        if ((buf[i] == 0) || (buf[i] == ' ')) {
            toClean = true;
            buf[i] = 0;
        }
        curPair.key[i] = buf[i];
    }
}

int /* sub */main() {
    std::ios_base::sync_with_stdio(false);
    std::cin.tie(nullptr);
    // std::cout.tie(nullptr);
    TBTree btree;
    char buf[MAX_KEY_LENGTH];
    Clear(buf);
    std::string command;
    unsigned long long curValue;
    TPair curPair;
    TBTreeNode* curSearchNode;
    unsigned curSearchPosition;
    while (std::cin >> command) {
        if (command == "+") {
            // scanf("%s", buf);
            std::cin >> buf;
            std::cin >> curValue;
            ToLower(buf, curPair);
            curPair.value = curValue;
            curSearchNode = nullptr;
            curSearchPosition = 0;
            btree.Find(curSearchNode, curSearchPosition, curPair);
            if (curSearchNode != nullptr) {
                std::cout << "Exist" << std::endl;
            }
            else {
                btree.Push(curPair);
                std::cout << "OK" << std::endl;
            }
        }
        else if (command == "-") {
            // scanf("%s", buf);
            std::cin >> buf;
            ToLower(buf, curPair);
            curSearchNode = nullptr;
            curSearchPosition = 0;
            btree.Find(curSearchNode, curSearchPosition, curPair);
            if (curSearchNode != nullptr) {
                btree.Pop(curPair);
                std::cout << "OK" << std::endl;
            }
            else {
                std::cout << "NoSuchWord" << std::endl;
            }
        }
        else if (command == "!") {
            std::cin >> command;
            if (command == "Load") {
                std::string filePath;
                std::cin >> filePath;
                std::ifstream file(filePath, std::ios::binary);
                btree.Load3(file);
                file.close();
                std::cout << "OK\n";
            }
            else {
                std::string filePath;
                std::cin >> filePath;
                std::ofstream file(filePath, std::ios::trunc | std::ios::binary);
                unsigned size = btree.GetSize();
                file.write(reinterpret_cast<char *>(&size), sizeof(unsigned));
                if (size > 0) {
                    btree.Save3(file);
                }
                file.close();
                std::cout << "OK\n";
            }
        }
        else {
            StringToChars(buf, command);
            ToLower(buf, curPair);
            curSearchNode = nullptr;
            curSearchPosition = 0;
            btree.Find(curSearchNode, curSearchPosition, curPair);
            if (curSearchNode != nullptr) {
                std::cout << "OK: " 
                          << curSearchNode->pairs[curSearchPosition].value
                          << "" << std::endl;
            }
            else {
                std::cout << "NoSuchWord" << std::endl;
            }
        }
        Clear(curPair.key);
    }
    return 0;
}

/* int main(int args, char** argv) {
    unsigned amount;
    std::cin >> amount;
    TBTree btree;
    TPair curPair;
    TBTreeNode* curSearchNode;
    unsigned curSearchPosition;
    char command;
    char buf[MAX_KEY_LENGTH];
    Clear(buf);
    unsigned long long curValue;

    std::chrono::steady_clock::time_point pushStart = std::chrono::steady_clock::now();
    for (unsigned i = 0; i < amount; ++i) {
        scanf("%s%llu", buf, &curValue);
        ToLower(buf, curPair);
        curPair.value = curValue;
        curSearchNode = nullptr;
        curSearchPosition = 0;
        btree.Find(curSearchNode, curSearchPosition, curPair);
        if (curSearchNode != nullptr) {
            printf("Exist\n");
        }
        else {
            btree.Push(curPair);
            printf("OK\n");
        }
        Clear(curPair.key);
    }
    std::chrono::steady_clock::time_point pushFinish = std::chrono::steady_clock::now();
    unsigned timeOfPush = std::chrono::duration_cast<std::chrono::milliseconds>(pushFinish - pushStart).count();
    
    std::chrono::steady_clock::time_point findStart = std::chrono::steady_clock::now();
    for (unsigned i = 0; i < amount; ++i) {
        scanf("%s", buf);
        ToLower(buf, curPair);
        curSearchNode = nullptr;
        curSearchPosition = 0;
        btree.Find(curSearchNode, curSearchPosition, curPair);
        if (curSearchNode != nullptr) {
            printf("OK: %llu\n", curSearchNode->pairs[curSearchPosition].value);
        }
        else {
            printf("NoSuchWord\n");
        }
        Clear(curPair.key);
    }
    std::chrono::steady_clock::time_point findFinish = std::chrono::steady_clock::now();
    unsigned timeOfFind = std::chrono::duration_cast<std::chrono::milliseconds>(findFinish - findStart).count();
    
    std::chrono::steady_clock::time_point popStart = std::chrono::steady_clock::now();
    for (unsigned i = 0; i < amount; ++i) {
        scanf("%s", buf);
        ToLower(buf, curPair);
        curSearchNode = nullptr;
        curSearchPosition = 0;
        btree.Find(curSearchNode, curSearchPosition, curPair);
        if (curSearchNode != nullptr) {
            btree.Pop(curPair);
            printf("OK\n");
        }
        else {
            printf("NoSuchWord\n");
        }
        Clear(curPair.key);
    }
    std::chrono::steady_clock::time_point popFinish = std::chrono::steady_clock::now();
    unsigned timeOfPop = std::chrono::duration_cast<std::chrono::milliseconds>(popFinish - popStart).count();
    
    std::cout << "!!!!! Time to push " << amount << " elements: " << timeOfPush << std::endl;
    std::cout << "!!!!! Time to find " << amount << " elements: " << timeOfFind << std::endl;
    std::cout << "!!!!! Time to pop " << amount << " elements: " << timeOfPop << std::endl;
    return 0;
} */
