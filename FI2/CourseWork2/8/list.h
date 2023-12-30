#ifndef list_h
#define list_h

struct list
{
    int k;
    struct list *next;
    struct list *pr;
};

struct list *last(struct list *L, struct list *H);

struct list *add(struct list *L, struct list *H, struct list *p, int n);

void output(struct list *L, int p, int l);

int length(struct list *L, struct list *H, int r);

struct list *rem(struct list *l);

struct list *del(struct list *L, struct list *H, int n, int a, int l);

int find(struct list *L, int a, int l, int n, int r);

struct list* repl(struct list *L, struct list *H, int n, int r);

#endif