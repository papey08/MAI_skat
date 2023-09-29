#include <iostream>
#include <string>
#include <cmath>

using namespace std;

unsigned long long chull(char a)
{
    return (a - '0');
}

int main ()
{
    int n;
    cin >> n;
    for (int g = 0; g < n; g++)
    {
        unsigned long long b = 0, e = 0;
    string s;
    cin >> s;
    short int l = s.length();
    while (l > 1)
    {
        for (int i = 0; i < l/2; i++)
            b = b * 10 + chull(s[i]);
        for (int i = l/2; i < l; i++)
            e = e * 10 + chull(s[i]);
        s = to_string(b + e);
        b = 0;
        e = 0;
        l = s.length();
    }
    cout << s << endl;
    }
    return 0;
}
