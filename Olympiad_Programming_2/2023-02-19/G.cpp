//TaskG
#include <iostream>
#include <vector>


void dfs (int v, std::vector<std::vector<int>>& g, std::vector<bool>& used) {
	used[v] = true;
	//comp.push_back (v);
	for (size_t i=0; i<g[v].size(); ++i) {
		int to = g[v][i];
		if (! used[to])
			dfs(to, g, used);
	}
}

int find_comps(int n, std::vector<std::vector<int>>& g, std::vector<bool>& used) {
    int ans = 0;
    for (int i = 0; i < n; ++i)
        if (!used[i]) {
            //comp.clear();
            //std::cout << i << std::endl;
            dfs(i, g, used);
            ++ans;
        }
    return ans;
}
int main() {

    int n, m;
    std::cin >> n >> m;
    std::vector<std::vector<int>> g(n, std::vector<int>(n));
    std::vector<bool> used(n, false);
    for(int i = 0; i < m; ++i){
        int from, to;
        std::cin >> from >> to;
        g[from].push_back(to);
        g[to].push_back(from);
        used[from] = true;
        used[to] = true;
    }
    for(int i = 0; i < n; ++i){
        used[i] = !used[i];
    }

    //std::vector<int> comp;
    int stars = find_comps(n, g, used);
    std::cout << stars;
    return 0;
}