#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main()
{
    int n, k, ans = 0;
    cin >> n >> k;
    vector<int> v(n);
    for (int i = 0; i < n; i++)
    {
        cin >> v[i];
    }
    sort(v.rbegin(), v.rend());
    for(int i = k - 1; i < n; i += k)
    {
        v[i] = 0;
    }
    for (int i = 0; i < n; i++)
    {
        ans += v[i];
    }
    cout << ans << endl;
    return 0;
}
