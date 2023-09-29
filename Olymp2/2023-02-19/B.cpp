#include <iostream>
#include <map>
#include <string>

std::map<char, std::string> init_map() {
    std::map<char, std::string> res{
        {'a', ".-"},
        {'b', "-..."},
        {'c', "-.-."},
        {'d', "-.."},
        {'e', "."},
        {'f', "..-."},
        {'g', "--."},
        {'h', "...."},
        {'i', ".."},
        {'j', ".---"},
        {'k', "-.-"},
        {'l', ".-.."},
        {'m', "--"},
        {'n', "-."},
        {'o', "---"},
        {'p', ".--."},
        {'q', "--.-"},
        {'r', ".-."},
        {'s', "..."},
        {'t', "-"},
        {'u', "..-"},
        {'v', "...-"},
        {'w', ".--"},
        {'x', "-..-"},
        {'y', "-.--"},
        {'z', "--.."},
    };
    return res;
}

std::string get_coded(std::string& str) {
    std::string res = "";
    std::map<char, std::string> keys = init_map();
    for (auto c: str) {
        res += keys[c];
    }
    return res;
}

int main() {
    std::string str;
    std::cin >> str;
    std::string coded_str = get_coded(str);
    bool ans = true;
    for (int i = 0; i < coded_str.length() / 2; i++) {
        if (coded_str[i] != coded_str[coded_str.length() - 1- i]) {
            ans = false;
            break;
        }
    }
    if (ans) {
        std::cout << "yes\n";
    } else {
        std::cout << "no\n";
    }
    return 0;
}
