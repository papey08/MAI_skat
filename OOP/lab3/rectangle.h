#ifndef RECTANGLE_H
#define RECTANGLE_H

#include <iostream>
#include "figure.h"

using namespace std;

class Rectangle : public Figure
{
private:
    Point p1, p2, p3, p4;
public:    
    Rectangle();
    Rectangle(istream& is);
    double Area();
    void Print(ostream& os);
    size_t VertexesNumber();
    virtual ~Rectangle();
};

#endif