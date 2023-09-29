#include <iostream>
#include <string>
#include <algorithm>
#include <map>

using namespace std;

struct comp
{
    bool operator()(string l, string r) const
    {
        return l < r;
    }
};

int main()
{
    ios::sync_with_stdio(0);
    cin.tie(0);
    cout.tie(0);
    string input, login, password;
    map<string, string, comp> k;
    map<string, bool> li;
    while(cin >> input)
    {
        if (input == "register") 
        {
            cin >> login >> password;
            if(k.count(login) == 0) 
            {
                k[login] = password;
                cout <<  "account created\n";
                li[login] = 0;
                continue;
            }
            else 
            {
                cout << "login already in use\n";
                continue;
            }
        }
        if (input == "login")
        {
            cin >> login >> password;
            if(k.count(login) == 0 || k[login] != password) 
            {
                cout << "wrong account info\n";
                continue;
            }
            if (!li[login] && k.count(login) > 0) 
            {
                cout << "logged in\n";
                li[login] = 1;
                continue;
            }
            else if (li[login] && k.count(login) > 0) 
            {
                cout << "already logged in\n";
                continue;
            }
        }
        if (input == "logout") 
        {
            cin >> login;
            if(li[login] == 0 || k.count(login) == 0) 
            {
                cout << "incorrect operation\n";
                continue;
            }
            else 
            {
                li[login] = 0;
                cout << "logged out\n";
                continue;
            }
        }
    }
    return 0;
}