#include <stdio.h>
#include <stdlib.h>
#include "list.h"

struct list *last(struct list *L, struct list *H)
{
    if (L == NULL)
        return L;
    if (L->next == H)
        return L;
    last(L->next, H);
}

struct list *add(struct list *L, struct list *H, struct list *p, int n)
{
    if (L == NULL)
    {
        struct list *L = malloc(sizeof(struct list));
        L->k = n;
        L->next = L;
        L->pr = L;
        return L;
    }
    if (L == p)
    {
        struct list *L = malloc(sizeof(struct list));
        L->k = n;
        L->next = H;
        L->pr = p;
        p->next = L;
        H->pr = L;
        return H;
    }
    add(L->next, H, p, n);
}

void output(struct list *L, int p, int l)
{
    if (p == l)
        return;
    printf("%d ", L->k);
    p++;
    output(L->next, p, l);
}

int length(struct list *L, struct list *H, int r)
{
    if (L == NULL)
        return 0;
    if ((L == H)&&(r > 0))
        return r;
    r++;
    length(L->next, H, r);
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

struct list *del(struct list *L, struct list *H, int n, int a, int l)
{
    if (a == l)
        return L;
    a++;
    if ((L->k == n)&&(l == 1))
        return NULL;
    if (L->k == n)
    {
        (L->pr)->next = L->next;
        (L->next)->pr = L->pr;
    }
    if (L == NULL)
        return L;
    del(L->next, H, n, a, l);
}

int find(struct list *L, int a, int l, int n, int r)
{
    if (a == l)
        return r;
    a++;
    if (L->k == n)
        r++;
    find(L->next, a, l, n, r);
}

struct list* repl(struct list *L, struct list *H, int n, int r)
{
    if (n == r)
    {
        int pl = (L->next)->k;
        (L->next)->k = (L->pr)->k;
        (L->pr)->k = pl;
        return H;
    }
    r++;
    repl(L->next, H, n, r);
}