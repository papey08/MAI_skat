#include <iostream>
#include <math.h> 

int main() {
	int n, T;
	std::cin >> n >> T;
	T = T * 60; // in seconds
	float m;
	std::cin >> m;
	int x, y;
	std::cin >> x >> y;
	float L = 0.0;
	float roadTime = (m / x + (n - m) / y);
	if (roadTime > T) {
		L = roadTime - T;
		//std::cout << L << std::endl;
		L /= 60.0;
		//std::cout << L << std::endl;
		L = ceil(L);
	}

	std::cout << (int)L;

}
