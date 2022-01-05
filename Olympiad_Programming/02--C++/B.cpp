#include <iostream>
#include <cmath>
#include <string>
#include <vector>

using namespace std;

int main()
{
    int c;
    vector<short int> v(15);
    long long A, B, sA = 0, sB = 0;
    char t;
    string s;
    while (cin >> A >> t >> B)
    {
        sA += A;
        sB += B;
    }
    sA = sA + sB/1000000000000000;
    sB = sB % 1000000000000000;
    cout << sA << '.';
    s = to_string(sB);
    c = 15 - s.length();
    for(int i = 0; i < c; i++)
        cout << "0";
    cout << sB;
    return 0;
}
