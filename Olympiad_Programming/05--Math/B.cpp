#include <iostream>
#include <string>

using namespace std;

int main()
{
	string s;
	cin >> s;
	int a = s[0] - 48;
	for(int i = 1; i < s.size(); i++)
    {
		a = a*10 + s[i] - 48;
		a = a%97;
	}
	if (a == 0)
		cout << "YES";
	else
		cout << "NO";
    return 0;
}
