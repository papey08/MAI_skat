#include <iostream>

using namespace std;

unsigned long long gcd(unsigned long long a, unsigned long long b)
{
	while (a != 0)
	{
	 b %= a;
	 swap (a, b);
	}
	return b;
}

unsigned long long lcm(unsigned long long a, unsigned long long b)
{
	return a / gcd(a, b) * b;
}

int main()
{
	ios::sync_with_stdio(false);
	cin.tie(0);
	unsigned long long c = 1, d, e, t;
	cin >> e;
	for (unsigned long long g = 0; g < e; g++)
	{
		cin >> t;
		d = lcm(c, t);
		c = d;
	}
	cout << d;
	return 0;
}
