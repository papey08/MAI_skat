#ifndef sort_h
#define sort_h

int length(struct list *l, int r);

int inf(struct list *l, int min);

struct list *replace(struct list *l, int min, int d);

struct list *sort(struct list *l, int n);

#endif
