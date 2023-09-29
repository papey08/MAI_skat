#include <iostream>
#include <string>

using namespace std;

int main()
{
    int n;
    cin >> n;
    string s;
    cin >> s;
    if ((s[0] == ')')||(s[0] == '}')||(s[0] == ']'))
    {
        cout << "Nein";
        return 0;
    }
    if ((s[n - 1] == '(')||(s[n - 1] == '{')||(s[n - 1] == '['))
    {
        cout << "Nein";
        return 0;
    }
    for (int i = 0; i < n; i++)
    {
        if ((s[i] == '(')&&((s[i + 1] == ']')||(s[i + 1] == '}')))
        {
            cout << "Nein";
            return 0;
        }
        if ((s[i] == '[')&&((s[i + 1] == '}')||(s[i + 1] == ')')))
        {
            cout << "Nein";
            return 0;
        }
        if ((s[i] == '{')&&((s[i + 1] == ']')||(s[i + 1] == ')')))
        {
            cout << "Nein";
            return 0;
        }
    }
    int amcl = 0, amsl = 0, amfl = 0, amcr = 0, amsr = 0, amfr = 0;
    for (int i = 0; i < n; i++)
    {
        if (s[i] == '(')
            amcl++;
        if (s[i] == '[')
            amsl++;
        if (s[i] == '{')
            amfl++;
        if (s[i] == ')')
            amcr++;
        if (s[i] == ']')
            amsr++;
        if (s[i] == '}')
            amfr++;
    }
    if ((amcl != amcr)||(amsl != amsr)||(amfl != amfr))
    {
        cout << "Nein";
        return 0;
    }
    cout << "Ja";
    return 0;
}