cmake_minimum_required(VERSION 3.16)
project(cp)

set(CMAKE_CXX_STANDARD 17)

add_executable(cp main.cpp TTriplet.hpp Compress.hpp Decompress.hpp)
