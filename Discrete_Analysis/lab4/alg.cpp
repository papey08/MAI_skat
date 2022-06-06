#include <iostream>
#include <string>
#include <vector>
// #include <chrono>

std::vector<unsigned> Find(std::string& pattern, std::string& text) {
    std::vector<unsigned> result;
    for (unsigned i = 0; i < text.length() - pattern.length() + 1; ++i) {
        unsigned pos = 0;
        while (text[i + pos] == pattern[pos]) {
            ++pos;
            if (pos == pattern.length()) {
                result.push_back(i);
            }
        }
    }
    return result;
}

int main() {
    /* std::chrono::steady_clock::time_point start = 
        std::chrono::steady_clock::now(); */
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
