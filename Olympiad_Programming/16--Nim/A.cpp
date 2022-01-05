#include <stdio.h>
#include <vector>

int main()
{
    int n, a, r = 0;
    scanf("%d", &n);
    for (int i = 0; i < n; i++)
    {
        scanf("%d", &a);
        r = r ^ a;
    }
    if (r == 0)
        printf("2\n");
    else
        printf("1\n");
    return 0;
}