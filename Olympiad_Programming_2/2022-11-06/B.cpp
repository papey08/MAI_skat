// Полонский
#include <iostream>
#include <vector>
#include <set>
#include <string>
#include <algorithm>
#include <fstream>

class Word{
public:
    std::multiset<char> mSet;
    char firstCh, lastCh;
    Word(std::string _word){
        firstCh = _word[0];
        lastCh = _word[_word.size() - 1];
        for(char ch : _word){
            mSet.insert(ch);
        }
    }
};

int main() {
    std::ifstream file("a.in");
    int n;
    file >> n;
    std::vector<std::string> strings;
    for(int i = 0; i < n; ++i){
        std::string curStr;
        file >> curStr;
        strings.push_back(curStr);
    }
    std::sort(strings.begin(), strings.end());

    std::string firstLetters;
    for(std::string& str : strings){
        //std::cout << str << ' ';
        firstLetters.push_back(str[0]);
    }
    //std::cout << "\n" << firstLetters << "\n";
    std::vector<Word> dict;
    for(int i = 0; i < n; ++i){
        Word curWord = Word(strings[i]);
        dict.push_back(curWord);
    }

    int m;
    file >> m;
    int answer = m;
    std::ofstream ansFile("a.out");
    for(int i = 0; i < m; ++i){
        std::string curStr;
        file >> curStr;
        std::multiset<char> curMSet;
        for(char ch : curStr){
            curMSet.insert(ch);
        }
        int index = firstLetters.find_first_of(curStr[0], 0);
        //std::cout << index << "\n";
        if(index < firstLetters.size()){ // в словаре есть слово с такой 1-й буквой
            while(index != firstLetters.size() && firstLetters[index] == curStr[0]){
                if(curStr[curStr.size() - 1] == dict[index].lastCh && curMSet == dict[index].mSet){
                    --answer;
                    break;
                }
                ++index;
            }
        }
    }

    ansFile << answer;
    file.close();
    ansFile.close();
    return 0;
}