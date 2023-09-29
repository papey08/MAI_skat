#include <iostream>
#include <vector>
#include <string>

using namespace std;

int main()
{
    string a, b;
    cin >> a >> b;
    long long as = a.size(), bs = b.size();
    vector<vector<long long>> v(as + 1, vector<long long> (bs + 1));
    for (int i = 1; i <= as; i++)
    {
        for (int g = 1; g <= bs; g++)
        {
            if (a[i - 1] == b[g - 1])
                v[i][g] = v[i - 1][g - 1] + 1;
            else
                v[i][g] = max(v[i - 1][g], v[i][g - 1]);
        }
    }
    cout << v[as][bs];
    return 0;
}
