#include <stdio.h>

int main()
{
    int n;
    printf("Size:\n");
    scanf("%d \n", &n);
    int a[n + 1][n + 1];
    short int b[n + 1][n + 1];
    for (int i = 1; i <= n; i++)
    {
        for (int g = 1; g <= n; g++)
        {
            scanf("%d ", &a[i][g]);
            b[i][g] = 1;
        }

    }
    printf("%d ", a[n][n]);
    b[n][n] = 0;
    int c = 1, t = 0, x;
    for (int i = 0; i < n; i++)
    {
        do
        {
            printf("%d ", a[n - 1 - i * 2][n - t]);
            c++;
            if (c == n * n)
                return 0;
            b[n - 1 - i * 2][n - t] = 0;
            t++;
        } while (b[n - i * 2][n - t] == 0);
        x = n - 1 - 2 * i;
        while (x <= n)
        {
            printf("%d ", a[x][n - 1 - 2 * i]);
            c++;
            if (c == n * n)
                return 0;
            b[x][n - 1 - 2 * i] = 0;
            x++;
        }
        t = 0;

        do
        {
            printf("%d ", a[n - t][n - 2 - i * 2]);
            c++;
            if (c == n * n)
                return 0;
            b[n - t][n - 2 - i * 2] = 0;
            t++;
        } while (b[n - t][n - i * 2 - 1] == 0);
        x = n - 2 - 2 * i;
        while (x <= n)
        {
            printf("%d ", a[n - 2 - 2 * i][x]);
            c++;
            if (c == n * n)
                return 0;
            b[n - 2 - 2 * i][x] = 0;
            x++;
        }
        t = 0;
    }
    return 0;
}
