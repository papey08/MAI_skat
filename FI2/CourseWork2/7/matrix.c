#include <stdio.h>
#include <stdlib.h>

struct matr
{
    int data;
    int index;
    struct matr *next;
};

struct matr *add(struct matr *M, int d, int i)
{
    if (M == NULL) 
    {
        struct matr *M = malloc(sizeof(struct matr));
        M->data = d;
        M->index = i;
        M->next = NULL;
        return M;
    }
    M->next = add(M->next, d, i);
}

void output(struct matr *M)
{
    if (M == NULL)
        return;
    printf("%d %d\n", M->data, M->index);
    output(M->next);
}

void modern_output(struct matr *M, int m, int n)
{
    struct matr *N = M;
    for (int i = 0; i < n * m; i++)
    {
        if (i % m == 0)
            printf("\n");
        if (N != NULL)
        {
            if (i == N->index)
            {
                printf("%d ", N->data);
                N = N->next;
            }
            else
                printf("0 ");
            continue;
        }
        else
            printf("0 ");
    }
}

void string(struct matr *M, int *r, int n)
{
    if (M == NULL)
        return;
    r[M->index / n]++;
    string(M->next, r, n);
}

int sum(struct matr *M, int i, int n, int ans)
{
    if (M == NULL)
        return ans;
    if (M->index / n == i)
        ans += M->data;
    sum(M->next, i, n, ans);
}