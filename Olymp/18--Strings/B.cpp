#include <iostream>
#include <string>
#include <vector>

using namespace std;

vector<int> prefix(string s) 
{
	int n = s.length();
	vector<int> p(n);
	for (int i = 1; i < n; i++) 
    {
		int j = p[i - 1];
		while ((j > 0)&&(s[i] != s[j]))
			j = p[j - 1];
		if (s[i] == s[j])  
            j++;
		p[i] = j;
	}
	return p;
}

int main()
{
    string s;
    cin >> s;
    vector<int> r = prefix(s);
    for (int a: r)
        cout << a << ' ';
    return 0;
}