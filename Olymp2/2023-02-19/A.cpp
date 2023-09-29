#include <iostream>
#include <string>

int main() {
    std::string n;
    std::cin >> n;
    int ans = 0;
    for (int i = n.length()-1; i >= 0; i--) {
        if (n[i] == '0') {
            ++ans;
        } else {
            break;
        }
    }
    std::cout << ans;
    return 0;
}
