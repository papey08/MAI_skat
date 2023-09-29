#ifndef FIGURE_H
#define FIGURE_H

#include <cstddef>
#include "point.h"

using namespace std;

class Figure
{
public:
    virtual ~Figure()
    {};
    virtual double Area() = 0;
    virtual void Print(ostream& os) = 0;
    virtual size_t VertexesNumber() = 0;
};

#endif