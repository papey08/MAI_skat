#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

struct team
{
    string name;
    int s;
    int p;
};

bool cmp(team &a, team &b)
{
    if (a.s > b.s)
        return 1;
    if (a.s < b.s)
        return 0;
    if (a.p < b.p)
        return 1;
    if (a.p > b.p)
        return 0;
    if (a.name < b.name)
        return 1;
    if (a.name >= b.name)
        return 0;
}

int main()
{
    int n;
    cin >> n;
    vector<team> v(n);
    for (int i = 0; i < n; i++)
        cin >> v[i].name >> v[i].s >> v[i].p;
    sort(v.begin(), v.end(), cmp);
    for (int i = 0; i < n; i++)
        cout << v[i].name << endl;
    return 0;
}