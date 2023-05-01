import math
import sys

import app.equation as eq
import app.equations_system as eqs

eps = 0.001


def lab02_01():
    l = 1
    r = 2

    def f(x):
        return math.tan(x) - 5. * x**2 + 1.

    def phi(x):
        return math.atan(5. * x**2 - 1.)

    def df(x):
        return math.cos(x)**(-2) - 10.*x

    print('-' * 10, '02-01', '-' * 10, sep='', end='\n\n')
    print('tg(x) - 5x^2 + 1 = 0', end='\n\n')

    x, i = eq.iterations(f, phi, l, r, eps)
    print('Iterations method:')
    if i != -1:
        print('x =', x)
        print('f(x) =', f(x))
        print('Number of iterations:', i)
    else:
        print('Iterations limit exceeded')
    print()

    x, i = eq.newton(f, df, l, r, eps)
    print('Newton method:')
    if i != -1:
        print('x =', x)
        print('f(x) =', f(x))
        print('Number of iterations:', i)
    else:
        print('Iterations limit exceeded')
    print('-' * 25, end="\n\n\n")


def lab02_02():
    a = 2
    l = 1
    r = 2

    def f1(x):
        return x[0]**2 - 2*math.log10(x[1]) - 1

    def f2(x):
        return x[0]**2 - a*x[0]*x[1] + a

    def phi1(x):
        return math.sqrt(2*math.log10(x[1]) + 1)

    def phi2(x):
        return (x[0]**2 + a) / (a * x[0])

    def dphi1_dx1(x):
        return 0

    def dphi1_dx2(x):
        return (x[1] * math.sqrt(math.log(10)) *
                math.sqrt(2*math.log(x[1]) + math.log(10)))**(-1)

    def dphi2_dx1(x):
        return a**(-1) - x[0]**(-2)

    def dphi2_dx2(x):
        return 0

    def df1_dx1(x):
        return 2 * x[0]

    def df1_dx2(x):
        return -2 / (x[1] * math.log(10))

    def df2_dx1(x):
        return 2*x[0] - a*x[1]

    def df2_dx2(x):
        return -a * x[0]

    print('-' * 10, '02-02', '-' * 10, sep='', end='\n\n')
    print('Equation system:')
    print('x1^2 - 2lg(x2) - 1 = 0')
    print('x1^2 - ax1x2 + a = 0, a = 2', end='\n\n')

    x, i = eqs.iterations(f1, f2, phi1, phi2, dphi1_dx1, dphi1_dx2, dphi2_dx1, dphi2_dx2, l, r, l, r, eps)
    print('Iterations method:')
    if i != -1:
        print('x1 =', x[0])
        print('x2 =', x[1])
        print('f1(x1, x2) =', f1(x))
        print('f2(x1, x2) =', f2(x))
        print('Number of iterations:', i)
    else:
        print('Iterations limit exceeded')
    print()

    x, i = eqs.newton(f1, f2, df1_dx1, df1_dx2, df2_dx1, df2_dx2, l, r, l, r, eps)
    print('Newton method:')
    if i != -1:
        print('x1 =', x[0])
        print('x2 =', x[1])
        print('f1(x1, x2) =', f1(x))
        print('f2(x1, x2) =', f2(x))
        print('Number of iterations:', i)
    else:
        print('Iterations limit exceeded')
    print()

    x, i = eqs.zeidel([l, r], eps, phi1, phi2, dphi1_dx1, dphi1_dx2, dphi2_dx1,
                      dphi2_dx2)
    print('Zeidel method:')
    if i != -1:
        print('x1 =', x[0])
        print('x2 =', x[1])
        print('f1(x1, x2) =', f1(x))
        print('f2(x1, x2) =', f2(x))
        print('Number of iterations:', i)
    else:
        print('Iterations limit exceeded')
    print()
    print('-' * 25, end="\n\n\n")


if __name__ == '__main__':
    if '1' in sys.argv:
        lab02_01()
    if '2' in sys.argv:
        lab02_02()
