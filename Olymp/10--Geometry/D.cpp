#include <iostream>
#include <cmath>

using namespace std;

struct point
{
    int x, y;
};

int main ()
{
    int n;
    cin >> n;
    for (int i = 0; i < n; ++i)
    {
        point p, l1, l2;
        int a, b, c;
        double r;
        cin >> p.x >> p.y >> l1.x >> l1.y >> l2.x >> l2.y;
        a = l1.y - l2.y;
        b = l2.x - l1.x;
        c = l1.x*l2.y - l2.x*l1.y;
        r = (abs(a*p.x + b*p.y + c) / sqrt(a*a + b*b));
        cout.precision(12);
        cout << r << endl;
    }
    return 0;
}
