#include <iostream>
#include <cmath>
#include <vector>

using namespace std;

int main ()
{
    int n;
    double a = 0, s = 0, m = 0;
    cin >> n;
    vector<double> x(n + 1);
    vector<double> y(n + 1);
    for (int i = 1; i <= n; i++)
        cin >> x[i] >> y[i];
    for (int i = 1; i <= n - 1; i++)
        s += x[i] * y[i + 1];
    s += x[n] * y[1];
    for (int i = 1; i <= n - 1; i++)
        m -= y[i] * x[i + 1];
    m -= y[n] * x[1];
    a = 0.5 * abs(s + m);
    cout.precision(12);
    cout << a << endl;
    return 0;
}
