#include <iostream>
using namespace std;
int main ()
{
	int a, b, c, d;
	long long S = 0;
	cin >> a >> b;
	c = a;
	d = b;
	if (d>c)
	{
	while (b>=a)
	{
		S += a;
		a +=1;
	}
}
	if (d<c)
	{
		while (a>=b)
	{
		S += b;
		b +=1;
	}
}
	if (c==d)
	S = a;

	cout << S << endl;
	return 0;
}
