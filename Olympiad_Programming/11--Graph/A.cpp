#include <iostream>
#include <vector>
#include <queue>

using namespace std;

int number = -1;

void dfs(int v, const vector<vector<int>>& graph, vector<int>& num)
{
    if (num[v] != -1)
        return;
    num[v] = ++number;
    for (int to:graph[v])
        dfs(to, graph, num);
}

int main() 
{
    ios::sync_with_stdio(false);
    cin.tie(0);
    cout.tie(0);
    int n, m, k;
    cin >> n >> m >> k;
    k--;
    vector<vector<int>> graph(n);
    for (int i = 0; i < m; i++)
    {
        int a, b;
        cin >> a >> b;
        a--;
        b--;
        graph[a].push_back(b);
        graph[b].push_back(a);
    }
    vector <int> num(n, -1);
    dfs(k, graph, num);
    for (int x: num)
        cout << x << ' ';
    cout << endl;
    return 0;
}