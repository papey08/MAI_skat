#include <iostream>
#include <vector>
#include <limits>

std::string to_binary_string(unsigned long long n)
{
    std::string buffer;
    buffer.reserve(std::numeric_limits<unsigned long long>::digits);
    do
    {
        buffer += char('0' + n % 2);
        n = n / 2;
    } while (n > 0);
    return std::string(buffer.crbegin(), buffer.crend());
}

int main() {
    unsigned n;
    std::cin >> n;
    std::vector<std::vector<unsigned>> v(n + 1, std::vector<unsigned>(31, 0));
    for (unsigned i = 1; i <= n; ++i) {
        unsigned deal;
        std::cin >> deal;
        std::copy(v[i - 1].begin(), v[i - 1].end(), v[i].begin());
        ++v[i][deal];
    }
    /* std::copy(v[n - 1].begin(), v[n - 1].end(), v[n].begin());
    for (unsigned i = 0; i < 31; ++i) {
        v[n][i] += v[0][i];
    } */
    
    unsigned q;
    std::cin >> q;
    for (int i = 0; i < q; ++i) {
        unsigned l, r;
        unsigned res = 0;
        unsigned long long m;
        std::cin >> l >> r >> m;
        std::string m_str = to_binary_string(m);
        /* if (r == n) {
            for (unsigned j = 0; j < m_str.length(); ++j) {
                if (m_str[m_str.length() - 1 - j] == '1') {
                    res += v[n - 1][j] * j;
                }
            }
        } else { */
            for (unsigned j = 0; j < m_str.length(); ++j) {
                if (m_str[m_str.length() - 1 - j] == '1') {
                    res += (v[r][j] - v[l][j]) * j;
                }
            }
        // }
        
        std::cout << res << std::endl;
    }

    return 0;
}
