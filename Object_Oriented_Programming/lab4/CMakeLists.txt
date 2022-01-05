cmake_minimum_required(VERSION 3.20)
project(lab4)

set(CMAKE_CXX_STANDARD 11)


add_executable(lab4
  main.cpp figure.cpp
  point.h point.cpp
  triangle.h triangle.cpp
  TBinaryTreeItem.h TBinaryTreeItem.cpp
  TBinaryTree.h TBinaryTree.cpp
)