#include <stdio.h>
#include "functions.h"

int main ()
{
	struct Chords
	{
		int i, j, l;
	};
	struct Chords d = {22, 10, 10};
	int a, b, c;
	for (int k = 0; k < 50; k++)
	{
		a = d.i;
		b = d.j;
		c = d.l;
		d.i = min((c) % 5, (a*k) % 5) + b + k/3;
		d.j = (max(-3*a, 2*b))/5 - abs(b - c);
		d.l = b + (c) % 7 + (k*sign(a)) % 10;
		if (check(d.i, d.j) == 1)
		{
			printf("Yes\n");
			printf("i = %d", d.i);
			printf(" j = %d", d.j);
			printf(" l = %d", d.l);
			printf(" k = %d", k);
			return 0;
		}
	}
	printf("No\n");
	printf("i = %d", d.i);
	printf(" j = %d", d.j);
	printf(" l = %d", d.l);
	printf(" k = 50");
	return 0;
}
