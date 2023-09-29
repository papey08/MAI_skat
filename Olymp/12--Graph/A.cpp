#include <iostream>
#include <vector>

using namespace std;

const int C = 1e9;

struct edge
{
    int from;
    int to;
    int weight;
};

int main()
{
    int n, m, start;
    cin >> n >> m >> start;
    start --;
    vector<edge> edges;
    for (int i = 0; i < m; i++)
    {
        int u, v, w;
        cin >> u >> v >> w;
        u--;
        v--;
        edges.push_back({u, v, w});
        edges.push_back({v, u, w});
    }
    vector<long long> d(n, C);
    d[start] = 0;
    bool changed = 1;
    for (int i = 0; (i < n)&&changed; i++)
    {
        changed = 0;
        for (size_t j = 0; j < edges.size(); j++)
        {
            edge cur_edge = edges[j];
            int u = cur_edge.from;
            int v = cur_edge.to;
            int w = cur_edge.weight;
            if (d[u] + w < d[v])
            {
                changed = 1;
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