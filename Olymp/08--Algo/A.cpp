#include <iostream>
#include <vector>
#include <algorithm>
#include <cmath>

using namespace std;
using ll = long long;

int main ()
{
    ios::sync_with_stdio(false);
    int n, q, l, r;
    cin >> n;
    vector<ll> v(n);
    for (int i = 0; i < n; i++)
        cin >> v[i];
    vector<ll> w(n+1);
    w[0] = 0;
    for (int i = 1; i < n+1; i++)
        w[i] = w[i-1] + v[i-1];
    cin >> q;
    for (int i = 0; i < q; i++)
    {
        cin >> l >> r;
        cout << w[r] - w[l-1] << endl;
    }
    return 0;
}
