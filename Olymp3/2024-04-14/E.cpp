#include <iostream>
#include <vector>

int main() {
	long long x;
	long long n;
	std::cin >> n >> x;
	std::vector<long long> a(n);
	for (int i = 0; i < n; ++i) {
		std::cin >> a[i];
	}
	std::vector<long long> pref(n + 1);
	
	for (long long i = 1; i < n + 1; ++i) {
		pref[i] = pref[i - 1] + a[i - 1];
	}
	//for (int el : pref) std::cout << el << ' '; std::cout << "\n";
	bool reachable = false;
	for (long long dist = 0; dist < n; ++dist) {
		for (long long start = 1; start + dist <= n; ++start) {
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
