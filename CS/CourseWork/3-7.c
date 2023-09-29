#include <stdio.h> //7
#include <math.h>

double res1(double x, int g)
{
    return g * (g + 2) * pow(x, g);
}

double res2(double x)
{
    return (x * (3 - x)) / pow(1 - x, 3);
}

void fun(double a, double step, int t, int n, double (*res1)(double, int), double(*res2) (double x))
{
    for (int i = 0; i < t; i++)
    {
        double x = a + step * i, r1 = 0.0;
        for (int g = 1; g <= n; g++)
            r1 += res1(x, g);
        double r2 = res2(x);
        printf ("%d| %lf %lf %lf\n", i+1, x, r1, r2);
    }
}

int main()
{
    double eps = 1.0;
    int n = 0;
    while (1.0 + eps / 2.0 > 1.0)
    {
        eps /= 2.0;
        n += 2;
    }
    double a = 0.0;
    int t;
    scanf("%d", &t);
    double step = 0.5 / t;
    fun(a, step, t, n, res1, res2);
    return 0;
}
