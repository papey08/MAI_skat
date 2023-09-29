#include <iostream>
#include <vector>
#include <queue>

using namespace std;
using graph = vector<vector<int>>;

int main()
{
    int n, m, k;
    cin >> n >> m >> k;
    k--;
    graph g(n);
    for (int i = 0; i < m; i++)
    {
        int u, v;
        cin >> u >> v;
        u--;
        v--;
        g[u].push_back(v);
        g[v].push_back(u);
    }
    queue<int> q;
    vector<int> lvl(n, -1);
    q.push(k);
    lvl[k] = 0;
    while (!q.empty())
    {
        int t = q.front();
        q.pop();
        for (size_t i = 0; i < g[t].size(); i++)
        {
            int h = g[t][i];
            if (lvl[h] == -1)
            {
                lvl[h] = lvl[t] + 1;
                q.push(h);
            }
        }
    }
    for (int i = 0; i < n; i++)
        cout << lvl[i] << " ";
    return 0;
}