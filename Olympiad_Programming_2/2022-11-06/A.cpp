#include <iostream>
#include <fstream>

using namespace std;

int main() {
ifstream file_in;
file_in.open("a.in");
int k;
file_in >> k;
int sum = 1;
for (int i = 0; i < k; i++) {
    int a;
    file_in >> a;
    sum += a;
}
ofstream fout;
fout.open("a.out");
fout << sum;
}