#include <iostream>
#include <vector>
#include <string>

// russiaopenhighschoolteamprogrammingcontest = 42, ����� = 46
int main() {
	std::string s;
	std::cin >> s;
	int n;
	std::cin >> n;
	std::vector<int> X(26);
	for (int i = 0; i < n; ++i) {
		char c;
		int x;
		std::cin >> c;
		std::cin >> x;
		X[((int)c) - 97] = x;
	}
	int minClicks = 0;
	std::vector<int> sCount(26);
	for (char c : s) {
		++sCount[((int)c) - 97];
	}
	for (int i = 0; i < 26; ++i) {
		if (X[i] == 0) { // key isn't broken
			minClicks += sCount[i];
		}
		else { // key is broken
			int integerCount = sCount[i] / (X[i] - 1);
			int modCount = sCount[i] % (X[i] - 1);
			if (modCount != 0) {
				++integerCount;
			}
			minClicks += sCount[i] + integerCount;
		}
	}
	std::cout << minClicks;
	return 0;
}
