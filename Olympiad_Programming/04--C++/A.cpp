#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main(){

	long double sum, sum1 = 0;
	long long a, b, c, d;
	cin >> a;
	vector<int> m(a);

	for (int i = 0; i < a; i++)
	{
		cin >> m[i];
	}
	sort(m.begin(),m.end());
	if (a == 1)
	{
		cout << "Deck looks good";
		return (0);
	}
	for (int i = 1; i < a ; i++)
	{
		if ( m[i] - m[i-1] != 1)
		{
			cout << "Scammed";
			return(0);
		}

	}
	cout << "Deck looks good";
	return 0;
}
