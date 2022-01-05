#include "functions.h"

int min(int a, int b)
{
	if (a < b)
		return a;
	else
		return b;
}

int max(int a, int b)
{
	if (a > b)
		return a;
	else
		return b;
}

int sign(int a)
{
	if (a < 0)
		return -1;
	if (a > 0)
		return 1;
	if (a == 0)
		return 0;
}

int abs(int a)
{
    if (a < 0)
        return (a*(-1));
    else
        return a;
}

int check(int i, int j)
{
	if ((i >= 5)&&(i <= 15)&&(j <= -5)&&(j >= -15))
		return 1;
	else
		return 0;
}
