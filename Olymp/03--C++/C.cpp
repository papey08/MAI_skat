#include <iostream>
#include <algorithm>

using namespace std;

int main()
{
    int n, a, b;
    cin >> n >> a >> b;
    n--;
    int t = -1, r = 10000000000;
    int m = 0;
    while (r - t > 1)
    {
        m = (r + t) / 2;
        if (n > m / a + m / b)
            t = m;
         else
            r = m;
    }
    r += min(a, b);
    cout << r;
    return 0;
}
