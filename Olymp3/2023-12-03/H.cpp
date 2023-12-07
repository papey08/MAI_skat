#include <iostream>
#include <vector>
#include <algorithm>

int main() {
	int n, m;
	std::cin >> n >> m;
	std::vector<int> army(n * m, 0);
	for (int i = 0; i < n * m; ++i) {
		std::cin >> army[i];
	}
	std::sort(army.begin(), army.end());
	std::reverse(army.begin(), army.end());
	for (int i = 0; i < n - 1; ++i) {
		for (int j = 0; j < m - 1; ++j) {
			std::cout << army[i * m + j] << ' ';
		}
		std::cout << army[i * m + m - 1] << std::endl;
	}
	for (int j = 0; j < m - 1; ++j) {
		std::cout << army[(n - 1) * m + j] << ' ';
	}
	std::cout << army[(n - 1) * m + m - 1];

	return 0;
}
