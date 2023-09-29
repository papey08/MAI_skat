#include <iostream>
#include <vector>
#include <map>
// #include <chrono>

unsigned INF = 1000000001;

int main() {
    // std::chrono::steady_clock::time_point startTime = 
    //     std::chrono::steady_clock::now();
    int n, m, start, finish;
    std::cin >> n >> m >> start >> finish;
    --start;
    --finish;
    std::vector<std::map<int, int>> g(n);
    for (unsigned i = 0; i < m; ++i) {
        int from, to, weight;
        std::cin >> from >> to >> weight;
        --from;
        --to;
        g[from][to] = weight;
        g[to][from] = weight;
    }
    std::map<unsigned long long, unsigned> dijkstra;
    dijkstra[0] = start;
    std::vector<bool> used (n);
    used[start] = true;
    std::map<unsigned, unsigned long long> ans;
    ans[start] = 0;
    while (!dijkstra.empty()) {
        unsigned node = dijkstra.begin()->second;
        dijkstra.erase(dijkstra.begin());
        if (node == finish) {
            break;
        }
        for (auto i: g[node]) {
            if (!used[i.first] || 
                ans[node] + i.second < ans[i.first])
            {
                dijkstra[ans[node] + i.second] = i.first;
                ans[i.first] = ans[node] + i.second;
                used[i.first] = true;
            }
        }
    }
    if (!used[finish]) {
        std::cout << "No solution" << std::endl;
    } else {
        std::cout << ans[finish] << std::endl;
    }
    // std::chrono::steady_clock::time_point finishTime = 
    //     std::chrono::steady_clock::now();
    // unsigned time = 
    //     std::chrono::duration_cast<std::chrono::milliseconds>(finishTime - startTime).count();
    // std::cout << "!!! " << time << " !!!\n";
    return 0;
}
