#ifndef TPAIR_HPP
#define TPAIR_HPP

#include <iostream>
#include <string>
#include <vector>

const int STR_SIZE = 64;
const int MAX_KEY = 65536;

class TPair {
    public:
        TPair() : key(0), value("")
        {}

        TPair(unsigned short n, std::string s) : key(n), value(s)
        {}

        TPair operator=(TPair b) {
            key = b.key;
            value = b.value;
            return *this;
        }

        int GetKey() {
            return key;
        }

        void Print() {
            std::cout << key << '\t' << value;
        }

        ~TPair() {}

    private:
        unsigned short key;
        std::string value;
};

#endif
