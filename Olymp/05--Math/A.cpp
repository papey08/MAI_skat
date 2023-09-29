#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

int main()
{
	vector<int> z;
	ios::sync_with_stdio(0);
	cin.tie(0);
	vector<int> v;
	unsigned long long n,i;
	cin >> n ;
	for (i = 1; i * i <= n; i++)
	{
		if(n % i == 0)
		{
			if(i * i != n)
				v.push_back(i);
			v.push_back(n / i);
		}
	}
cout << v.size();
return 0;
}
