#include <iostream>

int main() {
	long long n, a, b;
	std::cin >> n >> a >> b;
	unsigned long long totalPrice = 0;
	if (2 * a <= b) { // low prices makes no sense
		totalPrice = n * a;
	}
	else {
		if (n % 2 == 0) {
			totalPrice = (n / 2) * b;
		}
		else {
			totalPrice = (n / 2) * b + a;
		}
	}
	std::cout << totalPrice;
	return 0;
}
