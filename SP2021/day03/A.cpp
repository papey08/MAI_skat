#include <iostream>
#include <vector>

using namespace std;

const int a = 1000000007;

int main()
{
    int n, k;
    unsigned long long s;
    cin >> n >> k;
    vector<unsigned long long> dp(n + 1);
    dp[0] = 1;
    dp[1] = 1;
    for (int i = 2; i <= n; i++)
    {
        for (int j = 1; j <= min(k, i); j++)
        {
            dp[i] += dp[i - j];
            if (dp[i] >= a)
                dp[i] %= a;
        }
    }
    s = dp[n];
    cout << s;
    return 0;
}