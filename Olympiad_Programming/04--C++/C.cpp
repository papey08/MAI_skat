#include <iostream>
#include <vector>

using namespace std;

int main()
{
    vector<long long> v(100333);
    long long n, a, m, k, mod;
    cin >> n >> a >> m >> k >> mod;
    for (int i = 0; i < n; i++)
    {
        v[a]++;
        a = (a * m + k) % mod;
    }
    long long sum = 0, i = 0, j = 0, q = 1000000007;
    while ((i <= 100333)&&(j < n))
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
