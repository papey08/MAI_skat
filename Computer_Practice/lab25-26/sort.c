#include <stdio.h>
#include <stdlib.h>
#include "list.h"
#include "sort.h"

int length(struct list *l, int r)
{
    if (l == NULL)
        return r;
    r++;
    length(l->next, r);
}

int inf(struct list *l, int min)
{
    if (l == NULL)
        return min;
    if (l->k < min)
        min = l->k;
    inf(l->next, min);
}

struct list *replace(struct list *l, int min, int d)
{
    if (l == NULL)
        return l;
    if (l->k == min)
    {
        l->k = d;
        return l;
    }
    replace(l->next, min, d);
    return l;
}

struct list *sort(struct list *l, int n)
{
    if (n == 0)
        return l;
    int min = inf(l, l->k);
    l = replace(l, min, l->k);
    l->k = min;
    n--;
    sort(l->next, n);
    return l;
}