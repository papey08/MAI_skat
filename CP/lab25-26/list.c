#include <stdio.h>
#include <stdlib.h>
#include "list.h"

/*struct list
{
    int k;
    struct list *next;
}; */

struct list *add(struct list *l, int n)
{
    if (l == NULL) 
    {
        struct list *l = malloc(sizeof(struct list));
        l->k = n;
        l->next = NULL;
        return l;
    }
    l->next = add(l->next, n);
}

int find(struct list *l, int n, int r)
{
    if (l == NULL)
        return r;
    if (l->k == n)
        r++;
    find(l->next, n, r);
}

struct list *delete(struct list *l, int n)
{
    if (l->k == n)
        return l->next;
    if (l == NULL)
        return l;
    if ((l->next)->k == n)
    {
        l->next = (l->next)->next;
        return l;
    }
    delete(l->next, n);
    return l;
}

short int empty(struct list *l)
{
    if (l == NULL)
        return 0;
    else
        return 1;
}

void output(struct list *l)
{
    if (l == NULL)
        return;
    printf("%d ", l->k);
    output(l->next);
}

struct list *rem(struct list *l)
{
    if (l == NULL) 
        return l;
    if (l->next == NULL)
    {
        free(l);
        return NULL;
    }
    if (l->next != NULL)
        l->next = rem(l->next);
}