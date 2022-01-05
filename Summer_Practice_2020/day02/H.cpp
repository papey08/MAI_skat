#include <iostream>
#include <vector>

using namespace std;

const int c = 100333;

int main()
{
    vector<long long> v(c);
    long long n, a, m, k, mod;
    cin >> n >> a >> m >> k >> mod;
    for (int i = 0; i < n; i++)
    {
        v[a]++;
        a = (a * m + k) % mod;
    }
    long long sum = 0, i = 0, j = 0, q = 1000000007;
    while ((i <= c)&&(j < n))
    {
        if (v[i] > 0)
        {
            sum += ((j + 1) * i) % q;
            j++;
            v[i]--;
        }
        else
            i++;
    }
    sum %= q;
    cout << sum;
    return 0;
}