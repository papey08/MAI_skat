#include <stdio.h> //14
#include <math.h>

double eps()
{

    double e = 1.0;
    while (1.0 + e > 1.0)
        e /= 2.0;
    return e;
}

double fun(double x)
{
    return (tan(x / 2) - 1 / (tan(x / 2)) + x);
}

int main()
{
    double a = 1.0, b = 2.0, r = 1.0769, av = (a + b) / 2, x = 0.0;
    int k = 0;
    while(fabs(a - b) >= eps())
    {
        if(fun(a) * fun(av) > 0)
        {
            a = (a + b) / 2;
            av = (a + b) / 2;
            k++;
            if (av == x)
                break;
            printf("Step %d: x = %f; answer = %f; difference = %f\n", k , av, r, fabs(av - r));
            x = av;
        }
        if(fun(b) * fun(av) > 0)
        {
            b = (a + b) / 2;
            av = (a + b) / 2;
            k++;
            if (av == x)
                break;
            printf("Step %d: x = %f; answer = %f; difference = %f\n", k , av, r, fabs(av - r));
            x = av;
        }
    }
    return 0;
}
