#include <stdio.h>
#include <stdlib.h>
#include "list.c"

int main()
{
    printf("Type of key -- double\n");
    printf("Type of the list -- circular bidirectional\n");
    printf("Sorting method -- bubble sort\n");
    printf("Operations:\n");
    printf("a n s -- add element with key n and name s to the table\n");
    printf("dk n -- delete element with key n from the table\n");
    printf("dn s -- delete element with name s from the table\n");
    printf("fk n -- find element with key n in the table\n");
    printf("fn s -- find element with name s in the table\n");
    printf("sk -- sort the table by key\n");
    printf("sn -- sort the table by name\n");
    printf("p -- print the table\n");
    printf("# -- finish.\n");
    double n;
    char op = 'a';
    char na[50];
    struct list *L = NULL;
    while (op != '#')
    {
        scanf("%c", &op);
        if (op == 'a')
        {
            for (int i = 0; i < 50; i++)
                na[i] = '!';
            scanf("%lf %s", &n, &na);
            struct list *p = last(L, L);
            L = add(L, L, p, n, na);
        }
        if (op == 'p')
        {
            output(L, 0, length(L, L, 0));
            printf("\n");
        }
        if (op == 'f')
        {
            scanf("%c", &op);
            if (op == 'k')
            {
                scanf("%lf", &n);
                int l = length(L, L, 0);
                int r = findkey(L, 0, l, n, 0);
                printf("Elements with key %lf in the table: %d\n", n, r);
            }
            if (op == 'n')
            {
                char wo[50];
                for (int i = 0; i < 50; i++)
                    wo[i] = '!';
                scanf("%s", &wo);
                int l = length(L, L, 0);
                int r = findname(L, 0, l, wo, 0);
                printf("Elements with name %s in the table: %d\n", wo, r);
            }
        }
        if (op == 'd')
        {
            scanf("%c", &op);
            if (op == 'k')
            {
                scanf("%lf", &n);
                int l = length(L, L, 0);
                int r = findkey(L, 0, l, n, 0);
                for (int i = 0; i < r; i++)
                    L = delkey(L, L, n, 0, l);
            }
            if (op == 'n')
            {
                char wo[50];
                for (int i = 0; i < 50; i++)
                    wo[i] = '!';
                scanf("%s", &wo);
                int l = length(L, L, 0);
                int r = findname(L, 0, l, wo, 0);
                for (int i = 0; i < r; i++)
                    L = delname(L, L, wo, 0, l);
            }
        }
        if (op == 's')
        {
            scanf("%c", &op);
            if (op == 'k')
            {
                int f = length(L, L, 0);
                for (int i = 0; i < f - 1; i++)
                    L = sortkey(L, L);
            }
            if (op == 'n')
            {
                int f = length(L, L, 0);
                for (int i = 0; i < f - 1; i++)
                    L = sortword(L, L);
            }
        }
    }
    L = rem(L);
    return 0;
}