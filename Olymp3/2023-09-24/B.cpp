#include <iostream>
#include <vector>

int main() {
	int n;
	std::cin >> n;
	std::vector<char> curRow(n);
	int amountBlackInARow = 0;
	int maxTilesAmount = 0;
	for (int i = 0; i < n; ++i) {
		for (int j = 0; j < n; ++j) {
			std::cin >> curRow[j];
			if (curRow[j] == 'N') {
				++amountBlackInARow;
			}
			else {
				amountBlackInARow = 0;
			}
			if (amountBlackInARow == 2) {
				++maxTilesAmount;
				amountBlackInARow = 0;
			}

		}
		amountBlackInARow = 0;
	}
	std::cout << maxTilesAmount;
	return 0;
}
