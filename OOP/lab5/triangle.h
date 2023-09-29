#ifndef TRIANGLE_H
#define TRIANGLE_H

#include "figure.h"
#include <iostream>

using namespace std;

class Triangle : public Figure
{
private:
    Point a, b, c;
    
public:
    double A;
    Triangle(istream& io);
    Triangle();
    Triangle(Point a, Point b, Point c);
    double GetArea();
    size_t VertexesNumber();
    void Print(ostream &os);
    friend bool operator == (Triangle& t1, Triangle& t2);
    friend ostream& operator << (ostream& os, Triangle& p);
    virtual ~Triangle();
};

#endif
