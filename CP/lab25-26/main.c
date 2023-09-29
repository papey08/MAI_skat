#include <stdio.h> //linear list, procedure & method -- 2
#include <stdlib.h>
#include "list.h"
#include "sort.h"

int main()
{
    printf("Operations:\n");
    printf("an to add element n to the list\n");
    printf("fn to find element n in the list\n");
    printf("dn to delete element n from the list\n");
    printf("$n to add element n in the right order\n");
    printf("e to check the list if it is empty\n");
    printf("p to print the list\n");
    printf("s to sort the list\n");
    printf("# to finish.\n");
    char op = 'a';
    int e = 1;
    struct list *L = NULL;
    while (op != '#')
    {
        scanf("%c", &op);
        if (op == 'a')
        {
            scanf("%d", &e);
            L = add(L, e);
        }
        if (op == 'f')
        {
            scanf("%d", &e);
            int r = find(L, e, 0);
            if (r == 0)
                printf("Not found\n");
            if (r == 1)
                printf("Found 1 element\n");
            if (r > 1)
                printf("Found %d elements\n", r);
        }
        if (op == 'd')
        {
            scanf("%d", &e);
            int r = find(L, e, 0);
            for (int i = 0; i < r; i++)
                L = delete(L, e);
        }
        if (op == '$')
        {
            scanf("%d", &e);
            L = add(L, e);
            L = sort(L, length(L, 0) - 1);
        }
        if (op == 'e')
        {
            if (empty(L) == 0)
                printf("List is empty\n");
            else
                printf("List is not empty\n");
        }
        if (op == 'p')
        {
            output(L);
            printf("\n");
        }
        if (op == 's')
        {
            int f = length(L, 0);
            L = sort(L, f-1);
        }
    }
    L = rem(L);
    return 0;
}