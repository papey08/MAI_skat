#include <iostream>
#include <vector>

using namespace std;

int main()
{
    short int n;
    cin >> n;
    vector<int> v(n + 1);
    for (int i = 1; i <= n; i++)
        cin >> v[i];
    int l, r, t;
    cin >> l >> r;
    while ((l != 0)&&(r != 0))
    {
        for (int i = 0; i < (r - l + 1)/2; i++)
        {
            t = v[l + i];
            v[l + i] = v[r - i];
            v[r - i] = t;
        }
        cin >> l >> r;
    }
    for (int i = 1; i <= n; i++)
        cout << v[i] << " ";
    return 0;
}
