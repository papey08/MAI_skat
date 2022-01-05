#include <iostream>
#include <algorithm>
#include <vector>
#include <cmath>

using namespace std;

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(0);
    cout.tie(0);
    int n;
    bool tt;
    cin >> n;
    vector<int> t(n);
    vector<int> num;
    for(int i = 0; i < n; i++)
        cin >> t[i];
    vector<int> c = t;
    sort(c.begin(), c.end());
    for (int i = 0; i < n; i++)
    {
        if (c[i] <= n && (c[i] != c[i + 1]))
            tt = 1;
        else if (c[i] != 0) 
        {
            tt = 0;
            break;
        }
    }
    if (!tt) 
    {
        cout << -1;
        return 0;
    }
    if(tt)
    {
        for(int j = 1; j <= n; j++)
        {
            if (!binary_search(c.begin(), c.end(), j)) 
                num.push_back(j);
            else 
                continue;
        }
    }
    reverse(num.begin(), num.end());
    for(int i = 0; i < n; i++)
    {
        if (t[i] == 0) 
        {
            cout << num[num.size() - 1] << ' ';
            num.pop_back();
        }
        else 
            cout << t[i] << ' ';
    }
    return 0;
}