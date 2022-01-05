#ifndef TRIANGLE_H
#define TRIANGLE_H

#include <iostream>
#include "figure.h"

using namespace std;

class Triangle : public Figure
{
private:
    Point p1, p2, p3;
public:    
    Triangle();
    Triangle(istream& is);
    double Area();
    void Print(ostream& os);
    size_t VertexesNumber();
    virtual ~Triangle();
};

#endif