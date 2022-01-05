#include <iostream>
#include <vector>
#include <set>

using namespace std;

struct Edge
{
    long long to;
    long long w;
};

bool operator <(const Edge a,const Edge& b)
{
    if (a.w!=b.w)
        return a.w < b.w;
    return a.to < b.to;
}

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(0);
    cout.tie(0);
    long long n, m;
    cin >> n >> m;
    vector<vector<Edge>> graph(n);
    for (long long i = 0; i < m; i++)
    {
        long long a, b, c;
        cin >> a >> b >> c;
        a--;
        b--;
        graph[a].push_back(Edge{b,c});
        graph[b].push_back(Edge{a,c});
    }
    long long r = 0;
    vector<long long> used(n);
    used[0] = 1;
    set<Edge> edges(graph[0].begin(), graph[0].end());
    while (!edges.empty())
    {
        Edge v = *edges.begin();
        edges.erase (edges.begin());
        if (used[v.to])
            continue;
        used[v.to] = 1;
        r += v.w;
        for (const Edge& next: graph[v.to])
        {
            if (!used[next.to])
                edges.insert(next);
        }
    }
    cout << r;
    return 0;
}