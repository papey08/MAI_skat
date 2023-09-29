#include <iostream>
#include <vector>

using namespace std;

const int C = 1000000000;

struct edge
{
    int f;
    int t;
    int w;
};

int main()
{
    int n, m, s;
    cin >> n >> m >> s;
    s--;
    vector<edge> e;
    for (int i = 0; i < m; i++)
    {
        int u, v, w;
        cin >> u >> v >> w;
        u--;
        v--;
        e.push_back({u, v, w});
        e.push_back({v, u, w});
    }
    vector<long long> d(n, C);
    d[s] = 0;
    bool ch = 1;
    for (int i = 0; (i < n)&&(ch == 1); i++)
    {
        ch = 0;
        for (size_t j = 0; j < e.size(); j++)
        {
            edge cur_edge = e[j];
            int u = cur_edge.f;
            int v = cur_edge.t;
            int w = cur_edge.w;
            if (d[u] + w < d[v])
            {
                ch = 1;
                d[v] = d[u] + w;
            }
        }
    }
    for (int i = 0; i < n; i++)
    {
        if (d[i] == C)
            cout << "-1" << " ";
        else
            cout << d[i] << " ";
    }
    return 0;
}