#include <iostream>
#include <string>

using namespace std;

int main ()
{
	string a;
	cin >> a;
	int l = a.size(), n = 0;
	for (int i =0; i<=l; i++)
	{
		if ((a[l - i] == '0')|(a[l - i] == '4')|(a[l - i] == '6')|(a[l - i] == '9'))
		{
			n += 1;
		}
		if (a[l - i] == '8')
		{
			n += 2;
		}
	}
	cout << n;
	return 0;
}
