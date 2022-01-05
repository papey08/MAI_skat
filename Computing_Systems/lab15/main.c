#include <stdio.h>

int main()
{
    int n, min = 2147483647;
    printf ("Enter size:\n");
    scanf("%d\n", &n);
    int a[n][n];
    int sum[n];
    for (int i = 0; i < n; i++)
    {
        sum[i] = 0;
    }
    for (int i = 0; i < n; i++)
    {
        for (int g = 0; g < n; g++)
        {
            scanf ("%d ", &a[i][g]);
            if (a[i][g] < min)
                min = a[i][g];
        }
    }
    for (int i = 0; i < n; i++)
    {
        for (int g = 0; g < n; g++)
        {
            sum[i] += a[g][i];
        }
    }
    for (int i = 0; i < n; i++)
    {
        for (int g = 0; g < n; g++)
        {
            if (a[i][g] == min)
                a[i][g] = sum[g];
            printf ("%d ", a[i][g]);
        }
        printf ("\n");
    }
    return 0;
}
