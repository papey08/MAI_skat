#include <stdio.h>
#include <math.h> //13

double fun0(double x)
{
    return (x * tan(x) - 1 / 3);
}

double fun1(double x)
{
    return (x / pow(cos(x), 2.0) + tan(x));
}

double fun2(double x)
{
    return (4 * (x * (tan(x) + 1)) / (1 + cos(2 * x)));
}

int main()
{
    double a = 0.2, b = 1.0, eps = 1.0, x = (a + b) / 2, r = 0.5472, c = x;
    int k = 0;
    while (1.0 + eps > 1.0)
        eps /= 2.0;
    if (fabs(fun0(x) * fun2(x)) <= fun1(x))
    {
        printf("Newton method does not work");
        return 0;
    }
    while (fabs(fun0(x) * fun2(x)) < fun1(x))
    {
        k++;
        c = x;
        x = c - fun0(c) / fun1(c);
        printf("Step %d: x = %f; answer = %f; difference = %f\n", k, x, r, fabs(x - r));
    }
    return 0;
}
