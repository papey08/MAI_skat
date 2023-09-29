#include <iostream>
#include <cmath>

using namespace std;

struct point
{
    double x, y;
};

int main ()
{
    int n;
    cin >> n;
    for (int i = 0; i < n; i++)
    {
        point o, a, b;
        cin >> o.x >> o.y >> a.x >> a.y >> b.x >> b.y;
        double A, B, r;
        A = (a.x - o.x)*(b.x - o.x) + (a.y - o.y)*(b.y - o.y);
        B = (sqrt((a.x - o.x)*(a.x - o.x) + (a.y - o.y)*(a.y - o.y)) * sqrt((b.x - o.x)*(b.x - o.x) + (b.y - o.y)*(b.y - o.y)));
        r = A/B;
        cout.precision(12);
        cout << acos(r) << endl;
    }
    return 0;
}
