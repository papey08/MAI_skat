#include <iostream>
#include <stack>

using namespace std;

int main()
{
	long long n, t = 0, i = 1, a, b;
	stack<long long> s;
	s.push(0);
	s.push(1);
	cin >> n;
	if (n == 0)
	{
		cout << "0";
		return 0;
	}
	if (n == 1)
	{
		cout << "1";
		return 0;
	}
	while (t < n)
	{
		a = s.top();
		s.pop();
		b = s.top();
		s.pop();
		s.push(b);
		s.push(a);
		t = a + b;
		s.push(t);
		i++;
	}
	if (t == n)
		cout << i;
	else
		cout << "-1";
	return 0;
}
