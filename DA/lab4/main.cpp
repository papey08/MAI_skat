#include <iostream>
#include <vector>
#include <string>
// #include <chrono>

std::vector<unsigned> ZFunctionSquare(std::string& str) {
    std::vector<unsigned> result(str.length());
    result[0] = str.length();
    for (unsigned i = 1; i < str.length(); ++i) {
        unsigned pos = 0;
        while (str[pos] == str[i + pos]) {
            ++pos;
        }
        result[i] = pos;
    }
    return result;
}

std::vector<unsigned> ZFunctionLin(std::string& str) {
    std::vector<unsigned> result(str.length());
    result[0] = str.length();
    for (unsigned i = 1, l = 0, r = 0; i < str.length(); ++i) {
        if (i <= r) {
            result[i] = std::min(r - i + 1, result[i - l]);
        }
        while ((i + result[i] < str.size()) && 
               (str[result[i]] == str[i + result[i]]))
        {
            ++result[i];
        }
        if (i + result[i] - 1 > r) {
            l = i;
            r = i + result[i] - 1;
        }
    }
    return result;
}

std::vector<unsigned> Find(std::string& pattern, std::string text) {
    std::string str = pattern + "$" + text;
    std::vector<unsigned> result;
    std::vector<unsigned> zFunc = ZFunctionLin(str);
    for (unsigned i = pattern.size(); i < zFunc.size(); ++i) {
        if (zFunc[i] == pattern.size()) {
            result.push_back(i - pattern.size() - 1);
        }
    }
    return result;
}

int main() {
/*     std::chrono::steady_clock::time_point start = 
        std::chrono::steady_clock::now(); */
    std::ios_base::sync_with_stdio(false);
    std::cin.tie(nullptr);
    std::cout.tie(nullptr);
    std::string pattern, text;
    std::cin >> text >> pattern;
    std::vector<unsigned> result = Find(pattern, text);
    for (auto& x: result) {
        std::cout << x << std::endl;
    }
    /* std::chrono::steady_clock::time_point finish = 
        std::chrono::steady_clock::now();
    unsigned time = 
        std::chrono::duration_cast<std::chrono::milliseconds>(finish - start).count();
    std::cout << "!!! " << time << " !!!\n"; */
    return 0;
}
