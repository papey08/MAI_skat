#include <iostream>

using namespace std;

int main()
{
    int R, r;
        cin >> R >> r;
    if (r == R)
        cout << 1;
    else if (r > R)
        cout << 2;
    else
        cout << 1;
    return 0;
}
