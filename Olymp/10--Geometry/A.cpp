#include <iostream>
#include <cmath>

using namespace std;
const double pi = 3.141592653589;

int main()
{
    double x, y, r;
    int n;
    cin >> n;
    for(int i = 0; i < n; i++)
    {
        cin >> x >> y;
        r = atan2 (y, x);
        if (r < 0)
            r += 2 * pi;
        cout.precision(7);
        cout << r << endl;
    }
    return 0;
}
