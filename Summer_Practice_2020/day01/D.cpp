#include <iostream>
#include <algorithm>
#include <cmath>

using namespace std;

int main()
{
    long long l = 0;
    long long r = pow(10, 9) + 1;
    long long m;
    bool f = 0;
    while (f == 0)
    {
        m = (l + r) / 2;
        cout << m << endl;
        string s;
        cin >> s;
        if(s == "<") 
        {
            l = m;
            continue;
        }
        if(s == ">") 
        {
            r = m;
            continue;
        }
        else 
        {
            f = 1;
            break;
        }
    }
    return 0;
}