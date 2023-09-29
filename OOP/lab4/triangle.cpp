#include "triangle.h"
#include <cmath>

using namespace std;

Triangle::Triangle()
{}

Triangle::Triangle(istream& io)
{
    io >> a >> b >> c;
}

void Triangle::Print(ostream& os)
{
    os << a << " " << b << " " << c << endl;
}

size_t Triangle::VertexesNumber()
{
    return 3;
}

double Triangle::GetArea()
{
    double r1 = a.dist(b);
    double r2 = b.dist(c);
    double r3 = c.dist(a);
    double r = (r1 + r2 + r3)/2;
    A = sqrt(r * (r - r1) * (r - r2) * (r - r3));
    return sqrt(r * (r - r1) * (r - r2) * (r - r3));
}

bool operator == (Triangle& t1, Triangle& t2)
{
    if ((t1.a == t2.a)&&(t1.b == t2.b)&&(t1.c == t2.c))
    {
        return true;
    }
    else
    {
        return false;
    }
}

ostream& operator << (ostream& os, Triangle& t)
{
    os << t.a << " " << t.b << " " << t.c;
    return os;
}

Triangle::Triangle(Point _a, Point _b, Point _c) : a(_a), b(_b), c(_c)
{}

Triangle::~Triangle()
{}
