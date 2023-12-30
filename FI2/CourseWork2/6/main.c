#include <stdio.h>

int main()
{
    printf("Struct of the table:\n");
    printf("Surname + 2 initials + number of school + availability of medal (0 or 1) + points + writing test offset (0 or 1)\n");
    printf("Variant 22: find all enrollees with medal and score less than p\n");
    printf("Enter p:\n");
    FILE *f;
    f = fopen("input.txt", "r");
    int r;
    int ch = 0;
    scanf("%d", &r);
    char str[50];
    char in1, in2;
    int sc, me, po, so;
    for (int i = 0; i < 50; i++)
        str[i] = '!';
    while (fscanf(f, "%s %c %c", &str, &in1, &in2) != EOF)
    {
        fscanf(f, "%d %d %d %d", &sc, &me, &po, &so);
        if ((me == 1)&&(po < r))
        {
            printf("%s %c %c\n", str, in1, in2);
            ch++;
        }    
    }
    if (ch == 0)
        printf("No results found\n");
    fclose(f);
    return 0;
}