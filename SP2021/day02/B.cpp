#include<iostream>
#include<string>
#include<set>
#include<algorithm>

using namespace std;

int main()
{
    long long n;
    cin >> n;
    string s;
    set<string> q;
    for (int i=0; i < n; i++)
    {
        cin >> s;
        sort (s.begin(), s.end());
        q.insert(s);
    }
    cout << q.size();
    return 0;
}