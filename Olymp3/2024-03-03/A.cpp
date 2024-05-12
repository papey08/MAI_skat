#include <iostream>
#include <fstream>


int main() {
	int n, m;
	std::ifstream in;
	in.open("alter.in");
	in >> n >> m;
	//std::cin >> n >> m;
	in.close();
	std::ofstream out;
	out.open("alter.out");
	int cellAmount = n * m;
	int rowSwaps = n / 2;
	int colSwaps = m / 2;
	int swaps = rowSwaps + colSwaps;
	out << swaps << std::endl;
	//std::cout << swaps << std::endl;
	if (n > 1 && m > 1 && cellAmount > 1) {
		for (int i = 2; i <= n; i += 2) {
			out << i << ' ' << 1 << ' ' << i << ' ' << m << std::endl;
			//std::cout << i << ' ' << 1 << ' ' << i << ' ' << m << std::endl;
		}
		for (int j = 2; j <= m; j += 2) {
			out << 1 << ' ' << j << ' ' << n << ' ' << j << std::endl;
			//std::cout << 1 << ' ' << j << ' ' << n << ' ' << j << std::endl;
		}
	}
	else if (cellAmount > 1) {
		if (n == 1) {
			for (int i = 1; i <= n; ++i) {
				for (int j = (i % 2 == 1 ? 2 : 1); j <= m; j += 2) {
					out << i << ' ' << j << ' ' << i << ' ' << j << std::endl;
					//std::cout << i << ' ' << j << ' ' << i << ' ' << j << std::endl;
				}
			}
		}
		else if (m == 1) {
			for (int j = 1; j <= m; ++j) {
				for (int i = (j % 2 == 1 ? 2 : 1); i <= n; i += 2) {
					out << i << ' ' << j << ' ' << i << ' ' << j << std::endl;
					//std::cout << i << ' ' << j << ' ' << i << ' ' << j << std::endl;
				}
			}
		}
		
	}
	
	out.close();
	return 0;
}
