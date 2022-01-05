#include <iostream> //Сделал Матвей Попов группа М8О-108Б-20 Вариант 11 сдал 27.05.2021
#include <vector>
#include <queue>
#include <cstdlib>
#include <fstream>

using namespace std;

int number = -1;

void dfs(int v, const vector<vector<int>>& graph, vector<int>& num)
{
    if (num[v] != -1)
        return;
    num[v] = number++;
    for (int to:graph[v])
        dfs(to, graph, num);
}

int main(int argc, char *argv[])
{
    ifstream in(argv[1]);
    int n, m = 0;
    in >> n;
    vector<int> ans(n, 0);
    vector<vector<int>> v(n, vector<int> (n));
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
        {
            in >> v[i][j];
            if (v[i][j] == 1)
                m++;
        }
    }
    in.close();
    if (n == 1)
    {
        fstream out;
        out.open(argv[1]);
        out.clear();
        out << n << endl;
        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < n; j++)
                out << v[i][j] << " ";
            out << endl;
        }
        out << "Text:\n";
        out << "No hinges\n";
        out.close();
        return 0;        
    }
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j <= i; j++)
            v[i][j] = 0;
    }
    for (int k = 0; k < n; k++)
    {
        vector<vector<int>> v2(n, vector<int> (n));
        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < n; j++)
                v2[i][j] = v[i][j];
        }
        for (int i = 0; i < n; i++)
        {
            v2[k][i] = 0;
            v2[i][k] = 0;
        }
        vector<vector<int>> graph(n);
        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < n; j++)
            {
                if (v2[i][j] == 1)
                {
                    graph[i].push_back(j);
                    graph[j].push_back(i);
                }
            }
        }
        vector <int> num(n, -1);
        if (k == 0)
            dfs(1, graph, num);
        else
            dfs(0, graph, num);
        for (int i = 1; i <= n; i++)
        {
            if ((num[i] == -1)&&(i != k))
            {
                ans[k] = 1;
            }
        }
    }
    int re = 0;
    fstream out;
    out.open(argv[1]);
    out.clear();
    out << n << endl;
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
            out << v[i][j] << " ";
        out << endl;
    }
    out << "Text:\n";
    if ((n == 2)||(n == 1))
    {
        out << "No hinges\n";
        out.close();
        return 0;
    }
    for (int i = 0; i < n; i++)
    {
        if (ans[i] == 1)
        {
            out << i << " is a hinge\n";
            re++;
        }
    }
    if (re == 0)
        out << "No hinges\n";
    out.close();
    return 0;
}