#include <iostream>

long long calc(long long a, char OP, long long b) {
    switch (OP) {
    case '+':
        return a + b;
    case '*':
        return a * b;
    case '-':
        return a - b;
    }
}

int main() {
    unsigned n;
    std::cin >> n;
    for (unsigned i = 0; i < n; ++i) {
        long long a, b, c;
        char OP, eq;
        std:: cin >> a >> OP >> b >> eq >> c;
        long long d = calc(a, OP, b);
        if (c == d) {
            std::cout << "correct" << std::endl;
        } else {
            std::cout << "incorrect" << std::endl << a << " " << OP << " " << b << " = " << d << std::endl;
        }
    }
    return 0;
}
