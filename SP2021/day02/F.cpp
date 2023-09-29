#include <iostream>

using namespace std;

int main()
{
	long long x, a, n;
	cin >> n >> a;
	long long m = 0, mx = 0, t = 1;
	for (int i = 0; i < n; i++)
    {
		t = (t * 5) % 1000000007;
		x = (a * t) % 1000000007;
		if (x > mx)
        {
			m = mx;
			mx = x;
		}
		else if (x > m)
			m = x;
	}
	if (m == 0)
    {
		t = 1;
		for (int i = 0; i < n; i++)
		{
			t = (t * 5) % 1000000007;
			x = (a * t) % 1000000007;
			if ((x > m) && (x != mx))
				m = x;
		}
	}
	cout << m << ' ' << mx;
	return 0;
}