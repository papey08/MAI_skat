#include <iostream>
#include <vector>
#include <stack>

using namespace std;

int main()
{
    int n, min = 1000000000, max = -1000000000, lmins, lmaxs, maxp, minp;
    cin >> n;
    vector<int> v(n + 1);
    stack<int> mins;
    stack<int> maxs;
    for (int i = 1; i <= n; i++)
        cin >> v[i];
    if (n == 1)
    {
        cout << 1 << " " << 1 << endl << 1 << " " << 1 << endl << 1 << " " << 1;
        return 0;
    }
    if (v[1] < v[2])
        mins.push(1);
    if (v[1] > v[2])
        maxs.push(1);
    for (int i = 2; i <= n - 1; i++)
    {
        if ((v[i] < v[i - 1])&&(v[i] < v[i + 1]))
            mins.push(i);
        if ((v[i] > v[i - 1])&&(v[i] > v[i + 1]))
            maxs.push(i);
    }
    if (v[n] < v[n - 1])
        mins.push(n);
    if (v[n] > v[n - 1])
        maxs.push(n);
    for (int i = 1; i <= n; i++)
    {
        if (v[i] > max)
        {
            max = v[i];
            maxp = i;
        }
        if (v[i] < min)
        {
            min = v[i];
            minp = i;
        }

    }
    lmins = mins.size();
    lmaxs = maxs.size();
    cout << lmins << " ";
    if (lmins != 0)
    {
        stack<int> rmins;
        for (int i = 0; i < lmins; i++)
        {
            int d = mins.top();
            rmins.push(d);
            mins.pop();
        }
        for (int i = 0; i < lmins; i++)
        {
            int a = rmins.top();
            cout << a << " ";
            rmins.pop();
        }
    }
    cout << endl;
    cout << lmaxs << " ";
    if (lmaxs != 0)
    {
        stack<int> rmaxs;
        for (int i = 0; i < lmaxs; i++)
        {
            int c = maxs.top();
            rmaxs.push(c);
            maxs.pop();
        }
        for (int i = 0; i < lmaxs; i++)
        {
            int b = rmaxs.top();
            cout << b << " ";
            rmaxs.pop();
        }
    }
    cout << endl;
    cout << minp << " " << maxp;
    return 0;
}
