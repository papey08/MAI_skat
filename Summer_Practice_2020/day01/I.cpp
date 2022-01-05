#include <iostream>
#include <string>
#include <algorithm>
#include <vector>
 
using namespace std;
struct cord 
{
	int a;
	int b;
};

int main() 
{
	ios::sync_with_stdio(false);
	long long n;
	cin >> n;
	vector<long long> v(n);
	for (long long i = 0; i < n; i++)
		cin >> v[i];
	sort(v.begin(), v.end());
	int c;
	cin >> c;
	vector<cord> k(c);
	for (int d = 0; d < c; d++)
		cin >> k[d].a >> k[d].b;
	int z = 0;
	do 
    {
		cout << abs(lower_bound(v.begin(), v.end(), k[z].a) - lower_bound(v.begin(), v.end(), k[z].b)) << '\n';
		z += 1;
	} while (z < c);
    return 0;
}