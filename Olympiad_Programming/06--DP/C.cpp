#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int min(int a, int b)
{
    if (a < b)
        return a;
    else
        return b;
}

int main ()
{
    int n, k, m, d = 0;
    int a = 1000000007;
    unsigned long long s;
    cin >> n >> k >> m;
    if (m == 0)
    {
        vector<unsigned long long> v(n+1);
        v[0] = 1;
        v[1] = 1;
        for (int i = 2; i <= n; i++)
        {
            for (int j = 1; j <= min(k, i); j++)
            {
                v[i] += v[i-j];
                while (v[i] >= a)
                    v[i] -= a;
            }
        }
        s = v[n];
        cout << s;
        return 0;
    }
    vector<int> e(m);
    for (int t = 0; t < m; t++)
        cin >> e[t];
    sort(e.begin(), e.end());
    vector<unsigned long long> v(n+1);
    v[0] = 1;
    v[1] = 1;
    if (e[0] == 1)
    {
        v[1] = 0;
        d = 1;
    }
    for (int i = 2; i <= n; i++)
    {
        if (i != e[d])
        {
        for (int j = 1; j <= min(k, i); j++)
        {
            v[i] += v[i-j];
            while (v[i] >= a)
                v[i] -= a;
        }
        }
        else
        {
            v[i] = 0;
            d ++;
        }
    }
    s = v[n];
    cout << s;
    return 0;
}
