#include <iostream>
#include <vector>
#include <algorithm>
// #include <chrono>

const int MAX_NUM = 50;
int m, n;

int FindRow(std::vector<std::vector<double>>& v, int t) {
    int minPrice = MAX_NUM + 1;
    int index = -1;
    for (int i = t; i < m; ++i) {
        if ((v[i][t] != 0.0) && (v[i][n] < minPrice)) {
            index = i;
            minPrice = v[i][n];
        }
    }
    return index;
}

void SubtractRows(std::vector<std::vector<double>>& v, int t) {
    for (int i = t + 1; i < m; ++i) {
        double coeff = v[i][t] / v[t][t];
        for (int j = t; j < n; ++j) {
            v[i][j] -= v[t][j] * coeff;
        }
    }
}

int main() {
    // std::chrono::steady_clock::time_point start = 
    //     std::chrono::steady_clock::now();
    std::cin >> m >> n;
    std::vector<int> res;
    std::vector<std::vector<double>> additions
                            (m, std::vector<double> (n + 2));
    for (int i = 0; i < m; ++i) {
        for (int j = 0; j < n + 1; ++j) {
            std::cin >> additions[i][j];
        }
        additions[i][n + 1] = i;
    }

    for (int i = 0; i < n; ++i) {
        int index = FindRow(additions, i);
        if (index == -1) {
            std::cout << "-1" << std::endl;
            return 0;
        }

        std::swap(additions[i], additions[index]);
        res.push_back(additions[i][n + 1]);
        SubtractRows(additions, i);
    }

    std::sort(res.begin(), res.end());
    for (int i = 0; i < res.size(); ++i) {
        std::cout << res[i] + 1;
        if (i == res.size() - 1) {
            std::cout << std::endl;
        } else {
            std::cout << " ";
        }
    }

    // std::chrono::steady_clock::time_point finish = 
    //     std::chrono::steady_clock::now();
    // unsigned time = 
    //     std::chrono::duration_cast<std::chrono::milliseconds>(finish - start).count();
    // std::cout << "!!! " << time << " !!!\n";
    return 0;
}
