#include <iostream>

using namespace std;


long long
gachi (long long a, long long b)
{

return a == 0 ? b : gachi (b % a, a);

}



long long
muchi (long long a, long long b, long long &x, long long &y)
{

if (a == 0)
    {

x = 0;

y = 1;

return b;

}

long long x1, y1;

long long d = muchi (b % a, a, x1, y1);

x = y1 - (b / a) * x1;

y = x1;

return d;

}



void
res (long long a, long long b, long long c)
{

long long g, x, y;

g = muchi (a, b, x, y);

if (c % g != 0)
    {

cout << -1 << '\n';

}

  else
    {

x *= c / g;

y *= c / g;

cout << g << ' ' << x << ' ' << y << '\n';

}

}



int
main ()
{

ios::sync_with_stdio (false);

cin.tie (0);

cout.tie (0);

long long t;

cin >> t;

for (long long i = 0; i < t; i++)
    {

long long a, b, c;

cin >> a >> b >> c;

res (a, b, c);

}
return 0;

}
