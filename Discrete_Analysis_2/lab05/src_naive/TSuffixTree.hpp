#ifndef TSUFFIXTREE_HPP
#define TSUFFIXTREE_HPP

#include "TCompactTrie.hpp"
#include "TStringFunctions.hpp"

class TSuffixTree : public TCompactTrie {
public:

    void NaiveBuild(std::string& toBuild) {
        for (unsigned i = toBuild.length() - 1; i >= 0; --i) {
            std::string suffix = GetSuffix(toBuild, i);
            Add(suffix, i+1);
            if (suffix.empty()) {
                break;
            }
        }
    }

};


#endif
