#include <iostream>
#include <cmath>
#include <iomanip>
using namespace std;
int main ()
{
    double x1, y1, x2, y2, x3, y3, a, b, c, p, S, x12, y12, x13, y13, x23, y23;
    cin >> x1 >> y1 >> x2 >> y2 >> x3 >> y3;
    x12 = x1-x2;
    x13 = x1-x3;
    x23 = x2-x3;
    y12 = y1-y2;
    y13 = y1-y3;
    y23 = y2-y3;
    a = sqrt(x12*x12+y12*y12);
    b = sqrt(x13*x13+y13*y13);
    c = sqrt(x23*x23+y23*y23);
    p = (a+b+c)/2;
    a = p-a;
    b = p-b;
    c = p-c;
    S = sqrt(p*a*b*c);
    cout  << fixed << setprecision(9) << S;
  return 0;

    return 0;
}
