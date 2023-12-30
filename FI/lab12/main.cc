#include <stdio.h> // Попов Матвей ЛР №12 Вариант 4

int main()
{
    int n = 0, i = 0, c = 0;
    int a[256];
    scanf("%d", &n);
    while (n > 0)
    {
        a[i] = n%10;
        n /= 10;
        i++;
    }
    struct Nums
    {
        int a1, a2, a3;
    };
    struct Nums N = {0, 0, 0};
    for (int g = 0; g < i-2; g++)
    {
        N.a1 = a[g];
        N.a2 = a[g+1];
        N.a3 = a[g+2];
        if (N.a1 == N.a2 + N.a3)
        {
            printf("%d ", N.a1);
            c++;
        }
    }
    if (c == 0)
        printf("No numbers");
    return 0;
}
