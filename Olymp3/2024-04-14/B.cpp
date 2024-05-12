#include <iostream>

int main() {
	int n;
	std::cin >> n;
	for (int height = 1; height <= n; ++height) {
		for (int width = 1; width <= n; ++width) {
			if (height * width == n) {
				std::cout << height << " " << width;
				return 0;
			}
		}
	}
	return 0;
}
