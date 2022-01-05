#include <iostream>

using namespace std;

int main()
{
    int n, m, x, y;
    cin >> n >> m >> x >> y;
    int S = n * m, s = x * y;
    if (S % s == 0)
        cout << S / s;
    else if ((n % x == 0) && (m % y == 0))
        cout << (n / x) * (m / y);
    else
        cout << (n / x + 1) * (m / y + 1);
    return 0;
}
