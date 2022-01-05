#include <cmath>
#include "rectangle.h"

using namespace std;

Rectangle::Rectangle(istream& is)
{
    is >> p1 >> p2 >> p3 >> p4;
}

void Rectangle::Print(ostream& os)
{
    os << "Rectangle: " << p1 << " " << p2 << " " << p3 << " " << p4 << endl;
}

double Rectangle::Area()
{
    double a = p1.dist(p2);
    double b = p1.dist(p3);
    double c = p1.dist(p4);
    double d1 = max(a, b);
    double d2 = max(d1, c);
    if (d2 == a)
        return b * c;
    if (d2 == b)
        return a * c;
    if (d2 == c)
        return a * b;
    return 0.0; //How?
}

size_t Rectangle::VertexesNumber()
{
    return 4;
}

Rectangle::~Rectangle()
{
    cout << "Done\n";
}