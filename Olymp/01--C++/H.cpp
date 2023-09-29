#include <iostream>

using namespace std;

int min(int a, int b)
{
    if (a < b)
        return a;
    else
        return b;
}

int max(int a, int b)
{
    if (a > b)
        return a;
    else
        return b;
}

int main()
{
    short int a[250][250];
    short int r = 0;
    for (int i = 0; i < 250; i++)
    {
        for (int j = 0; j < 250; j++)
            a[i][j] = 0;
    }
    short int x1, y1, x2, y2, x3, y3, x4, y4;
    cin >> x1 >> y1 >> x2 >> y2 >> x3 >> y3 >> x4 >> y4;
    for (int i = min(x1 + 100, x2 + 100); i < max(x1 + 100, x2 + 100); i++)
    {
        for (int j = min(y1 + 100, y2 + 100); j < max(y1 + 100, y2 + 100); j++)
            a[i][j]++;
    }
    for (int i = min(x3 + 100, x4 + 100); i < max(x3 + 100, x4 + 100); i++)
    {
        for (int j = min(y3 + 100, y4 + 100); j < max(y3 + 100, y4 + 100); j++)
            a[i][j]++;
    }
    for (int i = 0; i < 250; i++)
    {
        for (int j = 0; j < 250; j++)
            if (a[i][j] == 2)
                r++;
    }
    cout << r;
    return 0;
}
