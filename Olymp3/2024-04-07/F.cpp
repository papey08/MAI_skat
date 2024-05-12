#include <iostream>
#include <map>
#include <set>
#include <string>
#include <vector>
#include <algorithm>

int main() {
	int n;
	std::cin >> n;
	std::vector<std::vector<std::string>> teamsTable(101);
	for (int i = 0; i < n; ++i) {
		std::string code;
		int p;
		std::vector<int> E(6);
		std::cin >> code >> p;
		for (int j = 0; j < 6; ++j) {
			std::cin >> E[j];
		}
		std::sort(E.begin(), E.end());
		int totalScore = 10 * p + E[1] + E[2] + E[3] + E[4];
		teamsTable[totalScore].push_back(code);
	}
	int teamsWBiggerScore = 0, totalMedalTeams = 0;
	for (int i = 100; i > -1 && totalMedalTeams < 1000; --i) {
		if (!teamsTable[i].empty()) {
			if (teamsWBiggerScore > 2) {
				break;
			}
			for (std::string& code : teamsTable[i]) {
				std::cout << code << " " << i << std::endl;
				++teamsWBiggerScore;
				++totalMedalTeams;
			}

			
		}
	}
	//for (auto el : teamsVect) std::cout << el.second << " " << el.first << std::endl;
	//std::cout << "---\n";

	return 0;
}
