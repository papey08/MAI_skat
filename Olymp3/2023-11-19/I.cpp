#include <iostream>

int main(){
	int N, M, K;
	std::cin >> N >> M >> K;
	if (N / M >= K) {
		std::cout << "Iron fist Ketil";
	}
	else {
		std::cout << "King Canute";
	}

	return 0;
}
