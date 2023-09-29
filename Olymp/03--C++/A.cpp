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
        return true;
    if (a.s < b.s)
        return false;
    if (a.p < b.p)
        return true;
    if (a.p > b.p)
        return false;
    if (a.name < b.name)
        return true;
    if (a.name >= b.name)
        return false;
}

int main()
{
    int n;
    cin >> n;
    vector<team> v(n);
    for (int i = 0; i < n; i++)
    {
        cin >> v[i].name >> v[i].s >> v[i].p;
    }
    sort(v.begin(), v.end(), cmp);
    for (int i = 0; i < n; i++)
    {
        cout << v[i].name << endl;
    }
    /*vector<string> name(n);
    vector<int> s(n);
    vector<int> p(n);
    for (int i = 0; i < n; i++)
    {
        cin >> name[i] >> s[i] >> p[i];
    }
    int d = s.size();
    while (d > 0)
    {
        d -= 1;
        for (int i = 0; i < d; i++)
        {
            while (s[i] < s[i + 1])
            {
                int c = s[i];
                s[i] = s[i + 1];
                s[i + 1] = c;
                string e = name[i];
                name[i] = name[i + 1];
                name[i + 1] = e;
                int f = p[i];
                p[i] = p[i + 1];
                p[i + 1] = f;
                i++;
            }
            while (s[i] == s[i + 1])
            {
                while (p[i] > p[i + 1])
                {
                    string e = name[i];
                    name[i] = name[i + 1];
                    name[i + 1] = e;
                    int f = p[i];
                    p[i] = p[i + 1];
                    p[i + 1] = f;
                    i++;
                }
                while (p[i] == p[i + 1])
                {
                    while (name[i] < name[i + 1])
                    {
                        string e = name[i];
                        name[i] = name[i + 1];
                        name[i + 1] = e;
                        i++;
                    }
                }
                if(p[i] < p[i + 1])
                    i++;
            }
        }
    }
    for (int i = 0; i < n; i++)
    {
        cout << name[i] << endl;
    }
    */
    return 0;
}
