#include <cmath>
#include "square.h"

using namespace std;

Square::Square(istream& is)
{
    is >> p1 >> p2 >> p3 >> p4;
}

void Square::Print(ostream& os)
{
    os << "Square: " << p1 << " " << p2 << " " << p3 << " " << p4 << endl;
}

double Square::Area()
{
    double a = p1.dist(p2);
    double b = p1.dist(p3);
    double c = p1.dist(p4);
    double d = a;
    if (d > b)
        d = b;
    if (d > c)
        d = c;
    return d * d;
}

size_t Square::VertexesNumber()
{
    return 4;
}

Square::~Square()
{
    cout << "Done\n";
}