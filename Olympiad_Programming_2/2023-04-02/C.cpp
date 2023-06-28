#include <iostream>
#include <vector>

int main() {
    int start_pos = 0;
    int temp;
    for (int i = 0; i < 6; ++i) {
        std::cin >> temp;
        if (temp == 1) {
            start_pos = i;
            ++start_pos;
            break;
        }
    }
    switch (start_pos)
    {
    case 1:
        std::cout << "1\nL1\n";
        break;
    case 2:
        std::cout << "0\n";
        break;
    case 3:
        std::cout << "1\nR1\n";
        break;
    case 4:
        std::cout << "1\nR2\n";
        break;
    case 5:
        std::cout << "1\nD1\n";
        break;
    case 6:
        std::cout << "1\nU1\n";
        break;
    }

    return 0;
}
