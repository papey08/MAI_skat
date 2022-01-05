#include <stdio.h> //#18 
#include <stdlib.h>

struct node
{
    int k;
    char op;
    struct node *l;
    struct node *r;
};

struct node *add(struct node *t, int n, char op)
{
    if (t == NULL) 
    {
        struct node *t = malloc(sizeof(struct node));
        t->k = n;
        t->op = op;
        t->r = NULL;
        t->l = NULL;
        return t;
    }
    
    if (n < t->k) 
        t->l = add(t->l, n, op);
    if (n >= t->k) 
        t->r = add(t->r, n, op);
    return t;
};

void inorder(struct node *t)
{
    if (t == NULL) 
        return;
    inorder(t->l);
    if (t->op == '^')
        printf("*%d", t->k);
    else
        printf("*%c", t->op);
    inorder(t->r);
}

struct node *rem(struct node *t)
{
    if (t == NULL) 
        return t;
    if (t->l == NULL && t->r == NULL)
    {
        free(t);
        return NULL;
    }
    if (t->r != NULL)
        t->r = rem(t->r);
    if (t->l != NULL)
        t->l = rem(t->l);
    return rem(t);
}

int main()
{
    struct node* T = NULL;
    int p = 0;
    char c = 'a';
    do
    {
        int n = 0;
        int rn = 0;
        int rd = 0;
        char f;
        scanf("%c", &c);
        if (c != '(')
        {
            f = c;
        }
        if (c == '(')
        {
            scanf("%c", &c);
            scanf("%c", &c);
            f = c;
            rd++;
        }
        if (c == '-')
        {
            scanf("%c", &c);
            f = c;
            rd++;
        }
        while ((c >= '0')&&(c <= '9'))
        {
            n *= 10;
            n += c - '0';
            scanf("%c", &c);
        }
        int g = 0;
        g = n;
        n = 0; 
        while ((c != '+')&&(c != '/')&&(c != '\n')&&(c != '-'))
        {
            scanf("%c", &c);
            while ((c >= '0')&&(c <= '9'))
            {
                n *= 10;
                n += c - '0';
                scanf("%c", &c);
                rn++;
            }
            if (rn != 0)
            {
                T = add(T, n, '^');
                rn = 0;
                n = 0;
            }
            if ((c >= 'a')&&(c <= 'z'))
                T = add(T, 0, c);
            if (c == '(')
            {
                rd++;
                scanf("%c", &c);
                c = 'a';
            } 
        }
        if ((rd % 2 == 0)&&(g == 0)&&(p == 0))
        {
            printf("%c", f);
            inorder(T);
        }
        if ((rd % 2 == 0)&&(g == 0)&&(p > 0))
        {
            printf("%c", f);
            inorder(T);
        }
        if ((rd % 2 == 0)&&(g != 0)&&(p == 0))
        {
            printf("%d", g);
            inorder(T);
        }
        if ((rd % 2 == 0)&&(g != 0)&&(p > 0))
        {
            printf("%d", g);
            inorder(T);
        }
        if ((rd % 2 == 1)&&(g == 0)&&(p == 0))
        {
            printf("-%c", f);
            inorder(T);            
        }
        if ((rd % 2 == 1)&&(g == 0)&&(p > 0))
        {
            printf("(-%c", f);
            inorder(T);
            printf(")");
        }
        if ((rd % 2 == 1)&&(g != 0)&&(p == 0))
        {
            printf("-%d", g);
            inorder(T);
        }
        if ((rd % 2 == 1)&&(g != 0)&&(p > 0))
        {
            printf("(-%d", g);
            inorder(T);
            printf(")");
        }
        T = rem(T);
        printf("%c", c);
        p++;
    } while (c != '\n');
    T = rem(T);
    if (c != '\n')
        printf("\n");
    return 0;
}