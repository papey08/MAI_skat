#include <iostream>
#include <vector>

int main() {
	long long x;
	int n;
	std::cin >> n >> x;
	std::vector<int> a(n);
	for (int i = 0; i < n; ++i) {
		std::cin >> a[i];
	}
	std::vector<int> pref(n + 1);
	
	for (int i = 1; i < n + 1; ++i) {
		pref[i] = pref[i - 1] + a[i - 1];
	}
	//for (int el : pref) std::cout << el << ' '; std::cout << "\n";
	bool reachable = false;
	for (int dist = 0; dist < n; ++dist) {
		for (int start = 1; start + dist <= n; ++start) {
			if (pref[start + dist] - pref[start - 1] >= x) {
				reachable = true;
				std::cout << dist + 1;
				return 0;
			}
		}
	}
	std::cout << -1;
	return 0;
}
