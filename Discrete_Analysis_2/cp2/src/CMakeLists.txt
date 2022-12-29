cmake_minimum_required(VERSION 3.16)
project(cp2)

set(CMAKE_CXX_STANDARD 20)

add_executable(cp2 main.cpp TCompressTreeNode.hpp TCompressTree.hpp TCompressTree.cpp TCompressTree.cpp TCompressTreeNode.cpp TDecompressTree.cpp TDecompressTree.hpp TDecompressTreeNode.cpp TDecompressTreeNode.hpp Huffman.hpp LZ77.hpp)
target_link_libraries(${PROJECT_NAME} stdc++fs)
