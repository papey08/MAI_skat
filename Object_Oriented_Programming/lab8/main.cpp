#include "triangle.h" //g++ main.cpp point.cpp triangle.cpp tallocation_block.cpp -Wall -Wextra -o main
#include <iostream>
#include <string>

int main()
{
    Point x1(0, 0);
    Point x2(1, 0);
    Point x3(0, 1);
    Triangle *t1 = new Triangle(x1, x2, x3);
    Triangle *t2 = new Triangle(x2, x1, x3);
    Triangle *t3 = new Triangle(x3, x1, x2);
    std::cout << "Three triangles have been initialized\n";
    delete t1;
    delete t2;
    delete t3;
    std::cout << "Three triangles have been deleted" << std::endl;
    return 0;
}
