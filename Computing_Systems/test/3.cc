#include <stdio.h> //Матвей Попов М8О-108Б-20 Вариант 6 №3

int main()
{
    int n, c = 1, m = 0, e, d = 1;
    scanf("%d \n", &n);
    int a[1001];
    for (int i = 0; i < 1001; i++)
    {
        a[i] = 2147483647;
        a[i] *= -1;
        a[i] -= 1;
    }
    for (int i = 0; i < n; i++)
        scanf("%d ", &a[i]);
    while (d < n-1)
    {
        while (a[d] > a[d - 1])
        {
            c++;
            d++;
        }
        if (c > m)
        {
            m = c;
            e = d;
        }
        c = 1;
        d++;
    }
    for (int i = e - m; i < e; i++)
        printf("%d ", a[i]);
    return 0;
}
