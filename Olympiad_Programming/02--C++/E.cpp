#include <iostream>

using namespace std;

int main()
{
    unsigned long long n, r = 1;
    cin >> n;
    if (n == 0)
    {
        cout << 1;
        return 0;
    }
    for (int i = 1; i <= n; i++)
    {
        r = r * i;
        if (r >= 1000000007)
            r = r % 1000000007;
    }
    cout << r;
    return 0;
}
