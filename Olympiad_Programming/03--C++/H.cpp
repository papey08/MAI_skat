#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main ()
{
	int n;
	cin >> n;
	vector<int> v (n);
	for (int i = 0; i < n; i++)
	{
		cin >> v[i];
	}
	sort (v.begin(), v.end());
	int c;
	c = n;
	for (int g = 0; g < n; g++)
	{
		if (v[g] == v[g+1])
			c -=1;
	}
	if (c == n)
	{
		cout << n;
	}
	else
	{
		cout << c;
	}
	return 0;
}
