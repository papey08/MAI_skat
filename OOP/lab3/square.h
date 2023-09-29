#ifndef SQUARE_H
#define SQUARE_H

#include <iostream>
#include "figure.h"

using namespace std;

class Square : public Figure
{
private:
    Point p1, p2, p3, p4;
public:    
    Square();
    Square(istream& is);
    double Area();
    void Print(ostream& os);
    size_t VertexesNumber();
    virtual ~Square();
};

#endif