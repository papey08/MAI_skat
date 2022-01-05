cmake_minimum_required(VERSION 3.20)
project(lab8)

set(CMAKE_CXX_STANDARD 11)


add_executable(lab8
  main.cpp figure.h
  triangle.h triangle.cpp
  point.h point.cpp
  tqueue_item.hpp tqueue.hpp
  tallocation_block.h tallocation_block.cpp
  TIterator.h
  TBinaryTree.h TBinaryTree.cpp
  TBinaryTreeItem.h TBinaryTreeItem.cpp
)
