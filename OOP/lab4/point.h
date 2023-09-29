#ifndef POINT_H
#define POINT_H

#include <iostream>

using namespace std;

class Point
{
private:
    double x, y;

public:
    Point();
    Point(istream& is);
    Point(double x, double y);
    double dist(Point& other);
    friend bool operator == (Point& p1, Point& p2);
    double getX();
    double getY();
    friend istream& operator >> (istream& is, Point& p);
    friend ostream& operator << (ostream& os, Point& p);
    friend class Triangle;
};

#endif
