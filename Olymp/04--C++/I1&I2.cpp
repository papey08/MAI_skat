#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main()
{
    long long n, r = 0;
    cin >> n;
    vector <int> v(n, 0);
    for (int i=0; i< n; i++)
        cin >> v[i];
    sort(v.begin(), v.end());
    for (int i = 0; i < n; i++)
    {
        for (int g = i + 1; g < n - 1; g++)
        {
            if (v[i] * v[g] != 0)
            {
                auto s = lower_bound(v.begin(), v.end(), v[i] + v[g]) - v.begin();
                r += s - g - 1;
            }
        }
    }
    cout << r << endl;
    return 0;
}
