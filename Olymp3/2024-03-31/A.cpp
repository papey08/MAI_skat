#include <iostream>

int main() {
	int p, m;
	std::cin >> p >> m;
	if (p > m) {
		std::swap(p, m);
	}
	int trollTime = 0;
	int trollPiece = 0;
	for (int i = 24; i >= 12; --i) {
		if (i > m) {
			if (i - m >= trollPiece) {
				trollTime = i;
				trollPiece = i - m;
			}
		}
		if (i == m) {
			if ((double)(i - p) / 2.0 >= trollPiece) {
				trollTime = i;
				trollPiece = (double)(i - p) / 2.0;
			}
		}
		if (i > p && i < m) {
			if (i - p >= trollPiece) {
				trollTime = i;
				trollPiece = i - p;
			}
		}
		if (i == p) {
			if ((double)(i - 12) / 2.0 >= trollPiece) {
				trollTime = i;
				trollPiece = (double)(i - 12) / 2.0;
			}
		}
		if (i < p) {
			if (i - 12 >= trollPiece) {
				trollTime = i;
				trollPiece = i - 12;
			}
		}
	}


	std::cout << trollTime;


	return 0;
}
