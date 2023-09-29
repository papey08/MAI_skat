#ifndef FIGURE_H
#define FIGURE_H

#include <memory>
#include "point.h"

using namespace std;

class Figure
{
public:
    virtual double GetArea() = 0;
    virtual void Print(ostream &os) = 0;
    virtual size_t VertexesNumber() = 0; 
    virtual ~Figure() {};
};

#endif
