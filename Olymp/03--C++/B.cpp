#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main ()
{
	int N;
	cin >> N;
	vector<int> v(N);
	for (int& a: v)
	{
		cin >> a;
	}
	sort (v.begin() , v.end());
	for (int a: v)
	{
		cout << a << ' ';
	}
	return 0;
}
