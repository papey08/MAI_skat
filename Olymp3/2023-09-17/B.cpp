#include <iostream>
#include <vector>
#include <string>

int main() {
	int n;

	std::cin >> n;
	std::vector<std::string> words(n);
	for (int i = 0; i < n; ++i) {
		std::cin >> words[i];
	}
	//for (std::string el : words) std::cout << el << std::endl;
	bool contains = false;
	for (int i = 0; i < n; ++i) {
		if (words[i] == "codecup" ||
			words[i] == "odecup" ||
			words[i] == "cdecup" ||
			words[i] == "coecup" ||
			words[i] == "codcup" ||
			words[i] == "codeup" ||
			words[i] == "codecp" ||
			words[i] == "codecu") {
			contains = true;
			break;
		}

	}
	if (contains) {
		std::cout << "Yes";
	}
	else {
		std::cout << "No";
	}
	return 0;
}