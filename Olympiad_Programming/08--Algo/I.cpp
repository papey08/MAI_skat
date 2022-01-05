#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main ()
{
    int n, s, f, c = 0, m = 0, ans = 0, a, b, d = 0;
    cin >> n;
    vector<pair <int, int>> e;
    for (int i = 0; i < n; i++)
    {
        cin >> s >> f;
        e.push_back(make_pair(s, -1));
        e.push_back(make_pair(f, 1));
    }
    sort(e.begin(), e.end());
    for (auto event: e)
    {
        int a = event.first;
        int b = event.second;
        c -= b;
        if (c > m)
        {
            m = c;
            ans = a;
            d++;
        }
    }
    cout << d << " " << ans;
    return 0;
}
