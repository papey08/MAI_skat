#include <iostream>
#include <vector>

using namespace std;

int min(int a, int b)
{
    if (a < b)
        return a;
    else
        return b;
}

int main()
{
    int n, k;
    int a = 1000000007;
    unsigned long long s;
    cin >> n >> k;
    vector<unsigned long long> v(n+1);
    v[0] = 1;
    v[1] = 1;
    for (int i = 2; i <= n; i++)
    {
        for (int j = 1; j <= min(k, i); j++)
        {
            v[i] += v[i - j];
            while (v[i] >= a)
                v[i] -= a;
        }
    }
    s = v[n];
    cout << s;
    return 0;
}
