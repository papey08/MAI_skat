#include <stdio.h> //������ ����� �8�-108�-20 �й11 ������� �7
#include <cctype>

enum State : unsigned { Start = 0, Finish, Symbol};

State count(const State currentState, const char currentSymbol)
{
    State newState;
    switch (currentState)
    {
        case Start:
        {
            if (currentSymbol != '\0')
                newState = Symbol;
            break;
        }
        case Symbol:
        {
            if (currentSymbol != '\0')
                newState = Symbol;
            if (currentSymbol == '\0')
                newState = Finish;
            break;
        }
    }
    return newState;
}

int check(char d)
{
    if ((d >= 'A')&&(d <= 'Z'))
        return 2;
    if ((d >= 'a')&&(d <= 'z'))
        return 1;
    else
        return 0;
}

int main ()
{
    char a[32768];
    int y = 1, f = 0, e = 32768;
    scanf("%[^\n]", a);
    for (int i = 0; i < e; i++)
    {
        if (check(a[i]) == 0)
        {
            y = 1;
        }
        if (check(a[i]) == 1)
        {
            a[i] += 3+y;
            y++;
            while (y > 26)
                y -= 26;
            while (check(a[i]) != 1)
            {
                while (check(a[i]) != 1)
                {
                    a[i] = a[i]-1;
                    f++;
                }
                a[i] = 'a'+f-1;
                f = 0;
            }
        }
        if (check(a[i]) == 2)
        {
            a[i] += 3+y;
            y++;
            while (y > 26)
                y -= 26;
            while (check(a[i]) != 2)
            {
                while (check(a[i]) != 2)
                {
                    a[i] = a[i]-1;
                    f++;
                }
                a[i] = 'A'+f-1;
                f = 0;
            }
        }
    }
    printf("%s\n", a);
    int bs = 32768;
    char c[bs];
    int n = 1, t = 0, r = 1;
    scanf("%[^\n]", c);
    State currentState = Start;
    while (fgets(c, bs, stdin) != NULL)
    {
       for (int g = 1; g < 32768; g++)
        {
            currentState = count(currentState, c[g]);
            if (currentState == Symbol)
                r++;
        }
        for (int i = 0; i < r; i++)
        {
            if (check(c[i]) == 0)
            {
                n = 1;
            }
            if (check(c[i]) == 1)
            {
                c[i] += 3+n;
                n++;
                while (n > 26)
                    n -= 26;
                while (check(c[i]) != 1)
                {
                    while (check(c[i]) != 1)
                    {
                        c[i] = c[i]-1;
                        t++;
                    }
                    c[i] = 'a'+t-1;
                    t = 0;
                }
            }
            if (check(c[i]) == 2)
            {
                c[i] += 3+n;
                n++;
                while (n > 26)
                    n -= 26;
                while (check(c[i]) != 2)
                {
                    while (check(c[i]) != 2)
                    {
                        c[i] = c[i]-1;
                        t++;
                    }
                    c[i] = 'A'+t-1;
                    t = 0;
                }
            }
        }
        printf("%s\n", c);
    }

    return 0;
}
