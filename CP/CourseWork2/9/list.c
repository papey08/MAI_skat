#include <stdio.h>
#include <stdlib.h>

struct list
{
    double k;
    char name[50];
    struct list *next;
    struct list *pr;
};

int equal(char *a, char *b)
{
    for (int i = 0; i < 50; i++)
    {
        if (a[i] != b[i])
            return 0;    
    }
    return 1;
}

int compare(char *a, char *b)
{
    for (int i = 0; i < 50; i++)
    {
        if (a[i] > b[i])
            return 1;
        if (a[i] < b[i])
            return 0;
    }
    return 0;
}

struct list *last(struct list *L, struct list *H)
{
    if (L == NULL)
        return L;
    if (L->next == H)
        return L;
    last(L->next, H);
}

struct list *add(struct list *L, struct list *H, struct list *p, double n, char *c)
{
    if (L == NULL)
    {
        struct list *L = malloc(sizeof(struct list));
        L->k = n;
        for (int i = 0; i < 50; i++)
            L->name[i] = c[i];
        L->next = L;
        L->pr = L;
        return L;
    }
    if (L == p)
    {
        struct list *L = malloc(sizeof(struct list));
        L->k = n;
        for (int i = 0; i < 50; i++)
            L->name[i] = c[i];
        L->next = H;
        L->pr = p;
        p->next = L;
        H->pr = L;
        return H;
    }
    add(L->next, H, p, n, c);
}

void output(struct list *L, int p, int l)
{
    if (p == l)
        return;
    printf("Key: %lf | Name: %s\n", L->k, L->name);
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

struct list *delkey(struct list *L, struct list *H, double n, int a, int l)
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
    delkey(L->next, H, n, a, l);
}

struct list *delname(struct list *L, struct list *H, char *c, int a, int l)
{
    if (a == l)
        return L;
    a++;
    if ((equal(c, L->name) == 1)&&(l == 1))
        return NULL;
    if (equal(c, L->name) == 1)
    {
        (L->pr)->next = L->next;
        (L->next)->pr = L->pr;
    }
    if (L == NULL)
        return L;
    delname(L->next, H, c, a, l);
}

int findkey(struct list *L, int a, int l, double n, int r)
{
    if (a == l)
        return r;
    a++;
    if (L->k == n)
        r++;
    findkey(L->next, a, l, n, r);
}

int findname(struct list *L, int a, int l, char *c, int r)
{
    if (a == l)
        return r;
    a++;
    if (equal(c, L->name) == 1)
        r++;
    findname(L->next, a, l, c, r);
}

struct list *replace(struct list *a, struct list *b)
{
    double d = a->k;
    a->k = b->k;
    b->k = d;
    char c[50];
    for (int i = 0; i < 50; i++)
    {
        c[i] = a->name[i];
        a->name[i] = b->name[i];
        b->name[i] = c[i];
    }
    return a;
}

struct list *sortkey(struct list *l, struct list *h)
{
    if (l->next == h)
    {
        return h;
    }
    if (l->k > (l->next)->k)
        l = replace(l, l->next);
    sortkey(l->next, h);
}

struct list *sortword(struct list *l, struct list *h)
{
    if (l->next == h)
    {
        return h;
    }
    if (compare(l->name, (l->next)->name) == 1)
        l = replace(l, l->next);
    sortword(l->next, h);
}