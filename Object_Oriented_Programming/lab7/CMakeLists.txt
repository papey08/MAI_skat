cmake_minimum_required(VERSION 3.20)
project(lab7)

set(CMAKE_CXX_STANDARD 11)


add_executable(lab7
  main.cpp figure.cpp
  point.h point.cpp
  triangle.h triangle.cpp
  TBinaryTreeItem.h TBinaryTreeItem.cpp
  TBinaryTree.h TBinaryTree.cpp
  TIterator.h
)