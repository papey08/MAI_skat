#include <iostream>
#include <vector>
#include <string>
// #include <chrono>

#include "TTriplet.hpp"
#include "Compress.hpp"
#include "Decompress.hpp"

int main() {
    // std::chrono::steady_clock::time_point startTime =
    //     std::chrono::steady_clock::now();
    std::string command;
    std::cin >> command;
    if (command == "compress") {
        std::string word;
        std::cin >> word;
        std::vector<TTriplet> code = Compress(word);
        for (auto& x: code) {
            std::cout << x.offset << " " << x.length << " " << x.nextSymbol
                << std::endl;
        }
    } else if (command == "decompress") {
        std::vector<TTriplet> triplets;
        TTriplet temp{};
        while (std::cin >> temp.offset >> temp.length >> temp.nextSymbol) {
            triplets.push_back(temp);
        }
        std::string res = Decompress(triplets);
        std::cout << res << std::endl;
    }
    // std::chrono::steady_clock::time_point finishTime = 
    //     std::chrono::steady_clock::now();
    // unsigned time =
    //     std::chrono::duration_cast<std::chrono::milliseconds>(finishTime - startTime).count();
    // std::cout << "!!! " << time << " !!!\n";
    return 0;
}
