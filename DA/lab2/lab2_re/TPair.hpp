#include <iostream>
#include <string>
#include <vector>
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
