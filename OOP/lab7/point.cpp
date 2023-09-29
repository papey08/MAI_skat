#include "point.h"
#include <cmath>

using namespace std;

Point::Point() : x(0.0), y(0.0)
{}

Point::Point(double x, double y) : x(x), y(y)
{}

Point::Point(istream& is)
{
    is >> x >> y;
}

double Point::getX()
{
    return x;
}

double Point::getY()
{
    return y;
}

double Point::dist(Point& other)
{
    double dx = abs(x - other.x);
    double dy = abs(y - other.y);
    return sqrt(dx * dx + dy * dy);
}

istream& operator >> (istream& is, Point& p)
{
    is >> p.x >> p.y;
    return is;
}

ostream& operator << (ostream& os, Point& p)
{
    os << "(" << p.getX() << ", " << p.getY() << ")";
    return os;
}

bool operator == (Point& p1, Point& p2)
{
    return((p1.x == p1.y)&&(p2.x == p2.y));
}
