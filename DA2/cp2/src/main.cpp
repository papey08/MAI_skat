#include <iostream>
#include <string>
#include <experimental/filesystem>
#include <chrono>

#include "Huffman.hpp"
#include "LZ77.hpp"

void Remove(std::string& filename) {
    char* toRemove = new char[filename.length() + 1];
    for (unsigned i = 0; i < filename.length(); ++i) {
        toRemove[i] = filename[i];
    }
    toRemove[filename.length()] = '\0';
    std::remove(toRemove);
    delete[] toRemove;
}

void Solve(std::vector<std::string>& filenames, bool flagC, bool flagD,
           bool flagK, bool flagL, bool flag1)
{
    if (flag1) {
        if (flagD) { // decompressing .huffman file
            for (std::string &name: filenames) {
                NHuffman::Decompress(name, flagC, flagL);
                if (!flagK) {
                    std::string keyName = name.substr(0, name.length() - 8) +
                                          ".key";
                    Remove(keyName);
                    Remove(name);
                }
            }
        } else { // compressing using Huffman only
            for (std::string &name: filenames) {
                NHuffman::Compress(name, flagC, flagL);
            }
        }
    } else {
        if (flagD) { // decompressing .lz77.huffman file
            for (std::string &name: filenames) {
                NHuffman::Decompress(name, false, flagL);
                std::string newName = "decompressed_" +
                                      name.substr(0, name.length() - 8);
                NLZ77::Decompress(newName, flagC, flagL);
                Remove(newName);
                if (!flagK) {
                    std::string keyName = name.substr(0, name.length() - 8) +
                                          ".key";
                    Remove(keyName);
                }
                std::string decompressedName = "decompressed_" +
                                               name.substr(0, name.length() - 8);
                Remove(decompressedName);
            }
        } else { // compressing using both LZ77 and Huffman
            for (std::string &name: filenames) {
                NLZ77::Compress(name, false, flagL);
                std::string newName = name + ".lz77";
                NHuffman::Compress(newName, flagC, flagL);
                Remove(newName);
            }
        }
    }
    if (!flagK) {
        for (std::string name : filenames) {
            Remove(name);
        }
    }
}

int main(int args, char** argv) {
    std::chrono::steady_clock::time_point startTime =
         std::chrono::steady_clock::now();

    bool flagC = false; // output to stdout
    bool flagD = false; // decompress
    bool flagK = false; // keep input files
    bool flagL = false; // information
    bool flagR = false; // recursive search of files in folders
    bool flagT = false; // testing key
    bool flag1 = false; // only Huffman compression
    bool flag9 = false; // LZ77 + Huffman compression
    std::vector<std::string> filenames;
    for (unsigned i = 1; i < args; ++i) {
        std::string arg = argv[i];
        if (arg == "-c") {
            flagC = true;
        } else if (arg == "-d") {
            flagD = true;
        } else if (arg == "-k") {
            flagK = true;
        } else if (arg == "-l") {
            flagL = true;
        } else if (arg == "-r") {
            flagR = true;
        } else if (arg == "-t") {
            flagT = true;
        } else if (arg == "-1") {
            flag1 = true;
        } else if (arg == "-9") {
            flag9 = true;
        } else {
            filenames.push_back(arg);
        }
    }
    if ((!flag1) && (!flag9)) {
        flag1 = true;
    }
    if (flagT) { // testing .key file
        if (filenames[0].substr(filenames[0].length() - 3, 3) != "key") {
            std::cerr << "ERROR!! Wrong file extension" << std::endl;
            return 0;
        } else if (NHuffman::CheckKey(filenames[0])) {
            std::cout << filenames[0] << " is correct" << std::endl;
            return 0;
        } else {
            std::cout << filenames[0] << "is incorrect" << std::endl;
            return 0;
        }
    }
    if (filenames.empty()) {
        std::cerr << "ERROR!! No input files" << std::endl;
        return 0;
    }
    if (flag1 && flag9) {
        std::cerr << "ERROR!! Both of -1 and -9 activated" << std::endl;
        return 0;
    }

    if (!flagR) {
        Solve(filenames, flagC, flagD, flagK, flagL, flag1);
    } else {
        std::string path = filenames[0];
        std::experimental::filesystem::recursive_directory_iterator dir(path);
        std::experimental::filesystem::recursive_directory_iterator end;
        std::vector<std::string> newFilenames;
        while (dir != end) {
            if (std::experimental::filesystem::is_regular_file(dir->path())) {
                newFilenames.push_back(dir->path());
            }
            ++dir;
        }
        Solve(newFilenames, flagC, flagD, flagK, flagL, flag1);
    }
    if (flagL) {
        std::chrono::steady_clock::time_point finishTime =
                std::chrono::steady_clock::now();
        unsigned time =
                std::chrono::duration_cast<std::chrono::milliseconds>(finishTime - startTime).count();
        std::cout << "Execution time: " << time << " ms\n";
    }
    return 0;
}
