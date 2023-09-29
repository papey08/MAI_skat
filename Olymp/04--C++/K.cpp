#include <iostream>
#include <stack>
#include <string>

using namespace std;

int main ()
{
	string a;
	string x, y, z;
	long long c, d, e;
	stack <string> s;
	while (cin >> a)
	{
		s.push(a);
		if (a == "+")
		{
			s.pop();
			y = s.top();
			s.pop();
			x = s.top();
			s.pop();
			c = stoll(y);
			d = stoll(x);
			e = d+c;
			z = to_string(e);
			s.push(z);
		}
		if (a == "-")
		{
			s.pop();
			y = s.top();
			s.pop();
			x = s.top();
			s.pop();
			c = stoll(y);
			d = stoll(x);
			e = d-c;
			z = to_string(e);
			s.push(z);
		}
		if (a == "*")
		{
			s.pop();
			y = s.top();
			s.pop();
			x = s.top();
			s.pop();
			c = stoll(y);
			d = stoll(x);
			e = d*c;
			z = to_string(e);
			s.push(z);
		}
	}
	cout << s.top();
	return 0;
}
