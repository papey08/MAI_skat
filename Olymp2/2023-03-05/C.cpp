#include <iostream>
#include <vector>

int main() {
	long long n;
	long long k;
	std::cin >> n;
	std::cin >> k;
	long long leftVisitors = k;
	long long curLevel = 1;
	long long M = std::ceil(std::sqrt(n * n + k + 1));
	std::vector<long long> places; // places[i] contains amount of places with distance = i;
	/*places.push_back(1);
	for (long long i = 1; i <= M; ++i) {
		places.push_back(4 * n + 4 * (i - 1));
	}*/
	long long curDist = 1;
	long long sumDist = 0;
	while (true) {
		long long curDPlaces = 4 * n + 4 * (curDist - 1);
		if (leftVisitors - curDPlaces >= 0) {
			sumDist += curDPlaces * curDist;
			leftVisitors -= curDPlaces;
		}
		else {
			sumDist += leftVisitors * curDist;
			break;
		}
		++curDist;
	}
	std::cout << sumDist;
	return 0;
}