#include <iostream>
#include <vector>

using namespace std;

int main ()
{
    ios::sync_with_stdio(0);
    cin.tie(0);
    cout.tie(0);
    int n, q, l, r;
    cin >> n;
    vector<long long> v(n);
    for (int i = 0; i < n; i++)
        cin >> v[i];
    vector<long long> prev(n+1);
    prev[0] = 0;
    for (int i = 1; i < n+1; i++)
        prev[i] = prev[i-1] + v[i-1];
    cin >> q;
    for (int i = 0; i < q; i++)
    {
        cin >> l >> r;
        cout << prev[r] - prev[l-1] << endl;
    }
    return 0;
}