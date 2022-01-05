#include <stdio.h>
#include <math.h> //19

double fun(double x)
{
    return (1 / (3 + sin(3.6 * x)));
}

int main ()
{
    double a = 0.0, b = 0.85, r = 0.2624, eps = 1.0, x = (a + b) / 2, c = fun(x);
    int k = 0;
    while (1.0 + eps > 1.0)
        eps /= 2.0;
    eps *= 100;
    while (fabs(x - c) >= eps)
    {
        k++;
        x = c;
        c = fun(x);
        printf("%d) x = %f; answer = %f; difference = %f\n", k, x, r, fabs(x - r));
    }
    return 0;
}
