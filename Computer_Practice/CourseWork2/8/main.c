#include <stdio.h>
#include <stdlib.h>

int main()
{
    printf("Type of elements -- int\n");
    printf("Type of the list -- circular bidirectional\n");
    printf("Action -- replace (k-1)th and (k+1)th elements of the list (k is parameter)\n");
    printf("Operations:\n");
    printf("an -- add element n to the list\n");
    printf("dn -- delete element n from the list\n");
    printf("kn -- replace (n-1)th element and (n+1)th element of the list\n");
    printf("p -- print the list\n");
    printf("l -- print length of the list\n");
    printf("# -- finish.\n");
    int n;
    char op = 'a';
    struct list *L = NULL;
    while (op != '#')
    {
        scanf("%c", &op);
        if (op == 'a')
        {
            scanf("%d", &n);
            struct list *p = last(L, L);
            L = add(L, L, p, n);
        }
        if (op == 'p')
        {
            output(L, 0, length(L, L, 0));
            printf("\n");
        }
        if (op == 'l')
            printf("Length is %d\n", length(L, L, 0));
        if (op == 'd')
        {
            scanf("%d", &n);
            int l = length(L, L, 0);
            int r = find(L, 0, l, n, 0);
            for (int i = 0; i < r; i++)
                L = del(L, L, n, 0, l);
        }
        if (op == 'k')
        {
            scanf("%d", &n);
            int l = length(L, L, 0);
            if ((n > l)||(n <= 0)||(l < 3))
                printf(":-(\n");
            else
            {
                L = repl(L, L, n, 1);
            }
        }
    }
    L = rem(L);
    return 0;
}