#include <iostream>
#include <string>
#include <vector>

using namespace std;

int main()
{
	ios::sync_with_stdio(false);
	cin.tie(0);
	int n;
	cin >> n;
	n *= 2;
	vector<string> v(n);
	for(int i = 0; i < n; i++)
		cin >> v[i];
	string s1, s2;
	long long r = 0;
	cin >> s1 >> s2;
	for(int i = 0; i < max(s1.size(),s2.size()); i++)
    {
		if ((i + 1 <= s1.size()) && (i + 1 <= s2.size()))
			r = r + s2[i] - s1[i];
		if(i + 1 > s1.size())
			r = r + s2[i] - 96;
		if(i + 1 > s2.size())
			r = r - s1[i] + 96;
	}
	cout << r;
	return 0;
}
