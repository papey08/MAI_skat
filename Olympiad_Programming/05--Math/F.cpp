#include <iostream>
#include <algorithm>
#include <vector>

 using namespace std;


int
main ()
{

long long n, p = 0, s = 0;

cin >> n;

vector < long long >v (n);

for (int i = 0; i < n; i++)

cin >> v[i];

sort (v.begin (), v.end ());

for (int i = 1; i < n; i++)

    {

s = s + v[i] * i - (p + v[i - 1]);

p = p + v[i - 1];

}
cout << s;

return 0;

}
