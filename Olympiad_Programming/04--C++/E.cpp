#include <iostream>
#include <queue>

using namespace std;

int main()
{
	long long n, s,t;
	cin >> n;
	priority_queue<long long> x;
	for (int i = 0; i < n; i++)
	{
		cin >> t;
		x.push(-t);
	}
    while (x.size() >= 2)
    {
        cout << -x.top() << ' ';
        s = -x.top();
        x.pop();
        cout << -x.top() << endl;
        s += -x.top();
        s *= -1;
        x.pop();
        x.push(s);
    }
    return 0;
}
