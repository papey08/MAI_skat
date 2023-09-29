#include <iostream>
#include <vector>
#include <bitset>
// #include <chrono>

const int MAX_N = 100;

int main() {
    // std::chrono::steady_clock::time_point start = 
    //     std::chrono::steady_clock::now();
    int n, m;
    std::cin >> n >> m;
    std::vector<int> w(n);
    std::vector<long long> c(n);
    for (int i = 0; i < n; ++i) {
        std::cin >> w[i] >> c[i];
    }

    std::vector<std::vector<long long>> ansPrev
                    (n + 1, std::vector<long long>(m + 1));
    std::vector<std::vector<std::bitset<MAX_N>>> resPrev
                    (n + 1, std::vector<std::bitset<MAX_N>>(m + 1));
    
    long long ans = 0;
    std::bitset<MAX_N> res;

    for (int i = 1; i < n + 1; ++i) {

        // std::copy(ansPrev[i - 1].begin(), ansPrev[i - 1].end(),
        //             ansPrev[i].begin());
        // std::copy(resPrev[i - 1].begin(), resPrev[i - 1].end(),
        //             resPrev[i].begin());

        for (int j = 1; j < m + 1; ++j) {
            ansPrev[i][j] = ansPrev[i - 1][j];
            resPrev[i][j] = resPrev[i - 1][j];
            if ((c[i - 1] > ansPrev[i][j]) && (j - w[i - 1] == 0)) {
                ansPrev[i][j] = c[i - 1];
                resPrev[i][j] = 0;
                resPrev[i][j][i - 1] = 1;
            }
            if (ansPrev[i][j] > ans) {
                ans = ansPrev[i][j];
                res = resPrev[i][j];
            }
        }
    }

    std::vector<std::vector<long long>> ansCur
                    (n + 1, std::vector<long long>(m + 1));
    std::vector<std::vector<std::bitset<MAX_N>>> resCur
                    (n + 1, std::vector<std::bitset<MAX_N>>(m + 1));

    for (long long i = 2; i < n + 1; ++i) {
        for (int j = 1; j < n + 1; ++j) {

            // std::copy(ansCur[j - 1].begin(), ansCur[j - 1].end(),
            //             ansCur[j].begin());
            // std::copy(resCur[i - 1].begin(), resCur[i - 1].end(),
            //             resCur[i].begin());

            for (int k = 1; k < m + 1; ++k) {
                ansCur[j][k] = ansCur[j - 1][k];
                resCur[j][k] = resCur[j - 1][k];
                if ((k - w[j - 1] > 0) && (ansPrev[j - 1][k - w[j - 1]] > 0)) {
                    if (i * (c[j - 1] + 
                        ansPrev[j - 1][k - w[j - 1]] / (i - 1)) > ansCur[j][k]) 
                    {
                        ansCur[j][k] = i * (c[j - 1] + 
                            ansPrev[j - 1][k - w[j - 1]] / (i - 1));
                        resCur[j][k] = resPrev[j - 1][k - w[j - 1]];
                        resCur[j][k][j - 1] = 1;
                    }
                }
                if (ansCur[j][k] > ans) {
                    ans = ansCur[j][k];
                    res = resCur[j][k];
                }
            }
        }

        std::swap(ansCur, ansPrev);
        std::swap(resCur, resPrev);
    }

    std::cout << ans << std::endl;
    for (int i = 0; i < n; ++i) {
        if (res[i]) {
            std::cout << i + 1 << ' ';
        }
    }
    std::cout << std::endl;

    // std::chrono::steady_clock::time_point finish = 
    //     std::chrono::steady_clock::now();
    // unsigned time = 
    //     std::chrono::duration_cast<std::chrono::milliseconds>(finish - start).count();
    // std::cout << "!!! " << time << " !!!\n";

    return 0;
}
