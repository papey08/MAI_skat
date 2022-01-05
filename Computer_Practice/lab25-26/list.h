#ifndef list_h
#define list_h

struct list
{
    int k;
    struct list *next;
};

struct list *add(struct list *l, int n);

int find(struct list *l, int n, int r);

struct list *delete(struct list *l, int n);

short int empty(struct list *l);

void output(struct list *l);

struct list *rem(struct list *l);

#endif