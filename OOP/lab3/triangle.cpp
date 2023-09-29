#include <cmath>
#include "triangle.h"

using namespace std;

Triangle::Triangle(istream& is)
{
    is >> p1 >> p2 >> p3;
}

void Triangle::Print(ostream& os)
{
    os << "Triangle: " << p1 << " " << p2 << " " << p3 << endl;
}

double Triangle::Area()
{
    double a = p1.dist(p2);
    double b = p2.dist(p3);
    double c = p3.dist(p1);
    double p = (a + b + c)/2;
    double s = sqrt(p * (p - a) * (p - b) * (p - c));
    return s;
}

size_t Triangle::VertexesNumber()
{
    return 3;
}

Triangle::~Triangle()
{
    cout << "Done\n";
}