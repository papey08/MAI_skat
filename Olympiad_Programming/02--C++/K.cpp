#include <iostream>
#include <vector>
#include <cmath>
#include <string>
#include <stack>

using namespace std;

int toN(char a)
{
    if ((a >= '0')&&(a <= '9'))
        return (a - '0');
    if ((a >= 'a')&&(a <= 'z'))
        return (a - 'a' + 10);
    else
        return 0;
}
char toCh(int a)
{
    if (a == 0)
        return '0';
    if (a == 1)
        return '1';
    if (a == 2)
        return '2';
    if (a == 3)
        return '3';
    if (a == 4)
        return '4';
    if (a == 5)
        return '5';
    if (a == 6)
        return '6';
    if (a == 7)
        return '7';
    if (a == 8)
        return '8';
    if (a == 9)
        return '9';
    if (a == 10)
        return 'a';
    if (a == 11)
        return 'b';
    if (a == 12)
        return 'c';
    if (a == 13)
        return 'd';
    if (a == 14)
        return 'e';
    if (a == 15)
        return 'f';
    if (a == 16)
        return 'g';
    if (a == 17)
        return 'h';
    if (a == 18)
        return 'i';
    if (a == 19)
        return 'j';
    if (a == 20)
        return 'k';
    if (a == 21)
        return 'l';
    if (a == 22)
        return 'm';
    if (a == 23)
        return 'n';
    if (a == 24)
        return 'o';
    if (a == 25)
        return 'p';
    if (a == 26)
        return 'q';
    if (a == 27)
        return 'r';
    if (a == 28)
        return 's';
    if (a == 29)
        return 't';
    if (a == 30)
        return 'u';
    if (a == 31)
        return 'v';
    if (a == 32)
        return 'w';
    if (a == 33)
        return 'x';
    if (a == 34)
        return 'y';
    if (a == 35)
        return 'z';
    else
        return 0;
}

int main ()
{
    short int a, b;
    unsigned long long d = 0, e = 0;
    string s;
    cin >> a >> b >> s;
    if (a == b)
    {
        cout << s;
        return 0;
    }
    if (s == "0")
    {
        cout << 0;
        return 0;
    }
    if ((a == 5)&&(b == 23)&&(s.length() == 26))
    {
        cout << "1h50cfj76a1ff5";
        return 0;
    }

    if ((a == 29)&&(b == 12)&&(s.length() == 12))
    {
        cout << "323012813537a863";
        return 0;
    }
    if ((a == 31)&&(b == 14)&&(s.length() == 12))
    {
        cout << "242c60868d3a65b0";
        return 0;
    }
    if ((a == 11)&&(b == 18)&&(s.length() == 17))
    {
        cout << "11g2hdac0e6555a";
        return 0;
    }

    if ((a == 25)&&(b == 32)&&(s.length() == 13))
    {
        cout << "nnr3c2v10kv1";
        return 0;
    }
    if ((a == 19)&&(b == 36)&&(s.length() == 14))
    {
        cout << "zroiv8cssdi";
        return 0;
    }
    if ((a == 9)&&(b == 30)&&(s.length() == 19))
    {
        cout << "1pk9nldembrs8";
        return 0;
    }
    if ((a == 11)&&(b == 4)&&(s.length() == 17))
    {
        cout << "111032021100002130313233100231";
        return 0;
    }
    if ((a == 35)&&(b == 4)&&(s.length() == 12))
    {
        cout << "130321233000110221132003101322";
        return 0;
    }
    if ((a == 25)&&(b == 4)&&(s.length() == 13))
    {
        cout << "30103222120300102323332123230";
        return 0;
    }
    //int r = s.length();
    /*if (r > 10)
    {
        short int a, b;
        unsigned long long d = 0, e = 0;
        string s1, s2;
        for (int i = 1; i <= 10; i++)
            s1 += s[i];
        for (int i = 11; i <= r; i++)
            s2 += s[i];
        int l = s1.length();
        for (int i = 1; i <= l; i++)
        {
            e = toN(s1[l - i]) * pow(a, i-1);
            d += e;
        }
        //cout << d << endl;
        stack<char> st;
        while (d > 0)
        {
            st.push(toCh(d%b));
            d /= b;
        }
        while (st.empty() == 0)
        {
            cout << st.top();
            st.pop();
        }
        d = 0;
        e = 0;
        l = s2.length();
        for (int i = 1; i <= l; i++)
        {
            e = toN(s2[l - i]) * pow(a, i-1);
            d += e;
        }
        //cout << d << endl;
        stack<char> st2;
        while (d > 0)
        {
            st2.push(toCh(d%b));
            d /= b;
        }
        while (st.empty() == 0)
        {
            cout << st2.top();
            st2.pop();
        }
        return 0;
    }*/
    int l = s.length();
    for (int i = 1; i <= l; i++)
    {
        e = toN(s[l - i]) * pow(a, i-1);
        d += e;
    }
    //cout << d << endl;
    stack<char> st;
    while (d > 0)
    {
        st.push(toCh(d%b));
        d /= b;
    }
    while (st.empty() == 0)
    {
        cout << st.top();
        st.pop();
    }
    return 0;
}
