#include <iostream>
#include <string>

int main() {
	int x, y;
	std::cin >> x >> y;
	std::string s;
	std::cin >> s;
	for (char el : s) {
		if (el == 'U') {
			++y;
		}
		else if (el == 'D') {
			--y;
		}
		else if (el == 'L') {
			--x;
		}
		else {
			++x;
		}
	}
	std::cout << x << " " << y;
	return 0;
}
