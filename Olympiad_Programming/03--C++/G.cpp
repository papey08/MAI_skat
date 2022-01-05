#include <iostream>
#include <vector>
#include <string>

using namespace std;

int main()
{
	string s, t;
	int n, a, b;
	cin >> a;
	cin >> s;
	cin >> n;
	vector<short int> v(n);
	for (int i = 0; i < n; i++)
    {
		cin >> b;
		cin >> t;
		int e = 0;
		v[i] = 1;
		for (int j = 0; j < a; j++)
        {
			if (s[j] == t[e])
				e++;
			if (e == b)
			{
				v[i] = 0;
				break;
			}
		}
	}
	for (int i = 0; i < n; i++)
		cout << v[i];
    return 0;
}
