#include <iostream>
#include <vector>

using namespace std;

void gachi(long long m, vector<int>& v)
{
    v[0] = 1;
    v[1] = 1;
    for (int i = 2; i <= m; i++)
    {
        if (v[i] != 0)
            continue;
        for (int j = i*2; j <= m; j += i)
            v[j] = 1;
    }
 }

int main()
{
    long long n, m, k=0;
    cin >> n >> m;
    int n1 = n;
    long long d = 15485863;
    vector<int> v(d + 1, 0);
    gachi (d, v);
    if (m == n)
    {
        for (int i=0;m==m;i++)
        {
            if(v[i]==0)
                m--;
            if (m==0)
            {
            cout << i;
                break;
            }
        }
    }
    else
    {
        for (long long i = 0; m == m; i++)
        {
            if (v[i] == 0)
            n--;
            if (n == 0)
            {
                for (long long j = i;m >= n1; j++)
                {
                    if (v[j] == 0)
                    {
                        k += j;
                        m--;
                    }
                }
                break;
            }
        }
        cout << k;
    }
    return 0;
}
