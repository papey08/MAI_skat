#include "mlisp.h"

double root(double a, double b);
double half__interval(double a, double b, double fa, double fb);
double __MAB__try(double neg__point, double pos__point);
bool close__enough_Q(double x, double y);
double average(double x, double y);
double f(double z);
extern double total__iterations;
extern double tolerance;
double root(double a, double b)
{
    double temp (0.);
    temp = half__interval(a, b, f(a), f(b));
    display("Total number of iterations=");
    display(total__iterations);newline();
    display("[");
    display(a);
    display(" , ");
    display(b);
    display("]");
    return temp;
}
double half__interval(double a, double b, double fa, double fb)
{
    double root (0.);
    total__iterations = 0.;
    root = (fa < 0. && fb > 0.
            ? __MAB__try(a, b)
            : fa > 0. && fb < 0.
            ? __MAB__try(b, a)
            : (b + 1.));
    newline();
    return root;
}
double __MAB__try(double neg__point, double pos__point)
{
    double midpoint (0.);
    double test__value (0.);
    midpoint = average(neg__point, pos__point);
    return (close__enough_Q(neg__point, pos__point) ? midpoint 
            : (test__value = f(midpoint), 
            display("+"), 
            total__iterations = total__iterations + 1., 
            test__value > 0. ? __MAB__try(neg__point, midpoint) 
            : (test__value < 0. ? __MAB__try(midpoint, pos__point) 
            : midpoint)));
}
bool close__enough_Q(double x, double y)
{
    return (abs(x - y) < tolerance);
}
double average(double x, double y)
{
    return ((x + y) * (1. / 2.0));
}

double tolerance = 0.001;
double total__iterations = 0.;
double f(double z)
{
    return expt(cos(z), 2.) - expt(sin(z), 2.);
}
int main()
{
    display("Variant 208-20\n");
    display(root(1.57, 3.0));
    newline();
    display("(c) Popov Matvey 2022\n");
    newline();

    std::cin.get();
    return 0;
}