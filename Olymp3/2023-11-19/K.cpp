#include <iostream>
#include <vector>

int main() {
	int N, K;
	std::cin >> N >> K;
	std::vector<std::vector<int>> map(N, std::vector<int>(N, 0));
	for (int i = 0; i < N; ++i) {
		for (int j = 0; j < N; ++j) {
			std::cin >> map[i][j];
		}
	}
	int knockoutAreas = 0;
	for (int i = 0; i < N - K + 1; ++i) {
		for (int j = 0; j < N - K + 1; ++j) {
			int cornLU = map[i][j];
			int cornRU = map[i][j + K - 1];
			int cornLD = map[i + K - 1][j];
			int cornRD = map[i + K - 1][j + K - 1];
			//std::cout << cornLU << " " << cornRU << " " << cornLD << " " << cornRD << std::endl;
			if ((cornLU == cornRU) && (cornLU == cornLD) && (cornLU == cornRD)) {
				++knockoutAreas;
			}
		}
	}
	std::cout << knockoutAreas;
	return 0;
}
