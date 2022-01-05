#include <stdio.h>
#include <stdlib.h>
#include "matrix.c"

int main()
{
    int m, n, d;
    printf("Variant #7\n");
    printf("Find strings with the biggest amount of non-zero elements and their sum\n");
    printf("Enter size (m * n):\n");
    scanf("%d %d", &m, &n);
    struct matr *M = NULL;
    printf("Enter the matrix:\n");
    for (int i = 0; i < m * n; i++)
    {
        scanf("%d", &d);
        if (d != 0)
            M = add(M, d, i);
    }
    printf("\nThe entered matrix:");
    modern_output(M, n, m);
    printf("\n");
    int r[100];
    for (int i = 0; i < 20; i++)
        r[i] = 0;
    string(M, r, n);
    printf("\n");
    int max = 0;
    for (int i = 0; i < m; i++)
    {
        if (r[i] > max)
            max = r[i];
    }
    if (max == 0)
    {
        printf("No elements instead of 0 was found\n");
        return 0;
    }
    printf("Strings with the biggest amount of non-zero elements: ");
    for (int i = 0; i < m; i++)
    {
        if (r[i] == max)
            printf("%d ", i + 1);
    }
    printf("\n");
    for (int i = 0; i < m; i++)
    {
        if (r[i] == max)
            printf("Sum of string %d is %d\n", i + 1, sum(M, i, n, 0));
    }
    return 0;
}