#include <stdio.h>

int check(int a)
{
    if (a%2 == 0)
        return 1;
    else
        return 0;
}

int main()
{
    int n, c = 0, d = 0, b;
    scanf("%d \n", &n);
    for (int i = 0; i < n; i++)
    {
        scanf("%d ", &b);
        if (check(b) == 1)
            c += b;
        else
            d += b;
    }
    printf("%d\n", c - d);
    return 0;
}
