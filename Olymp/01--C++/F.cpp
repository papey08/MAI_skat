#include <iostream> //Made by Matvey Popov Ì80-108Á-20
using namespace std;
int main ()
{
    int h1, m1, h2, m2, h3, m3, n, M, H;
    cin >> h1 >> m1 >> h2 >> m2 >> h3 >> m3 >> n;
    H = h1 + h2*n + h3*(n-1);
    M = m1 + m2*n + m3*(n-1);
    while (M>=60) {
        M=M-60;
        H=H+1;
    }
    cout << H << " " << M;
    return 0;
}
