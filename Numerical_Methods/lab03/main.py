import json
import math
import sys

import app.interpolation as ip
import app.spline as spl
from app.format import format_polynom
import app.least_squares as ls
import app.diff as diff
import app.integrate as integrate


def lab03_01():

    def f(x):
        return math.acos(x) + x

    print('-' * 10, '03-01', '-' * 10, sep='', end='\n\n')

    file = open('tests/03-01.json')
    data = json.loads(file.read())
    x_a = data['x_a']
    x_b = data['x_b']
    x_ = data['x*']
    file.close()

    print('Lagrange interpolation:', end='\n\n')

    polynom, error = ip.lagrange_with_check(f, x_a, (x_, f(x_)))
    print('Points A polynom:', polynom)
    print('Error in X*:', error, end='\n\n')

    polynom, error = ip.lagrange_with_check(f, x_b, (x_, f(x_)))
    print('Points B polynom:', polynom)
    print('Error in X*:', error)
    print()

    print('Newton interpolation:', end='\n\n')

    polynom, error = ip.newton(f, x_a, (x_, f(x_)))
    print('Points A polynom:', polynom)
    print('Error in X*:', error, end='\n\n')

    polynom, error = ip.newton(f, x_b, (x_, f(x_)))
    print('Points B polynom:', polynom)
    print('Error in X*:', error)
    print()

    print('-' * 25, end="\n\n\n")


def lab03_02():
    print('-' * 10, '03-02', '-' * 10, sep='', end='\n\n')

    file = open('tests/03-02.json')
    data = json.loads(file.read())
    x_i = data['x_i']
    f_i = data['f_i']
    x_ = data['x*']
    file.close()

    a, b, c, d, y = spl.spline_interpolation(x_i, f_i, x_)
    for i in range(len(x_i) - 1):
        print(f'[{x_i[i]}; {x_i[i+1]})')
        polynom = f's(x) = {a[i]:}+{b[i]:.4f}(x-{x_i[i]:.4f})+' \
                  f'{c[i]:.4f}(x-{x_i[i]:.4f})^2+{d[i]:.4f}(x-{x_i[i]:.4f})^3'
        print(format_polynom(polynom))
    print(f's(x*) = s({x_:.4f}) = {y:.4f}', end='\n\n')
    spl.draw_plot(x_i, f_i, a, b, c, d)

    print('-' * 25, end="\n\n\n")


def lab03_03():
    print('-' * 10, '03-03', '-' * 10, sep='', end='\n\n')

    file = open('tests/03-03.json')
    data = json.loads(file.read())
    x_i = data['x_i']
    y_i = data['y_i']
    file.close()

    print('Least squares method, 1st degree')
    ls1 = ls.least_squares(x_i, y_i, 1)
    print('P(x) =', format_polynom(f'{ls1[0]:.4f}+{ls1[1]:.4f}x'))
    print('Sum of squared errors =', ls.sum_squared_errors(x_i, y_i, ls1),
          end='\n\n')

    print('Least squares method, 2nd degree')
    ls2 = ls.least_squares(x_i, y_i, 2)
    print(f'P(x) =',
          format_polynom(f'{ls2[0]:.4f}+{ls2[1]:.4f}x+{ls2[2]:.4f}x^2'))
    print('Sum of squared errors =', ls.sum_squared_errors(x_i, y_i, ls1),
          end='\n\n')

    ls.draw_plot(x_i, y_i, ls1, ls2)

    print('-' * 25, end="\n\n\n")


def lab03_04():
    print('-' * 10, '03-04', '-' * 10, sep='', end='\n\n')

    file = open('tests/03-04.json')
    data = json.loads(file.read())
    x = data['x']
    y = data['y']
    x_ = data['x*']
    file.close()

    print(f"f'({x_}) = {diff.df(x, y, x_):.4f}")
    print(f"f''({x_}) = {diff.d2f(x, y, x_):.4f}", end='\n\n')

    print('-' * 25, end="\n\n\n")


def lab03_05():

    def f(x):
        return math.sqrt(x) / (4 + 3*x)

    print('-' * 10, '03-05', '-' * 10, sep='', end='\n\n')

    file = open('tests/03-05.json')
    data = json.loads(file.read())
    x0 = data['x0']
    xk = data['xk']
    h1 = data['h1']
    h2 = data['h2']
    p = data['p']
    file.close()

    rectangle_h1 = integrate.rectangle_trapeze(f, x0, xk, h1)
    rectangle_h2 = integrate.rectangle_trapeze(f, x0, xk, h2)
    trapeze_h1 = integrate.rectangle_trapeze(f, x0, xk, h1, False)
    trapeze_h2 = integrate.rectangle_trapeze(f, x0, xk, h2, False)
    simpson_h1 = integrate.simpson(f, x0, xk, h1)
    simpson_h2 = integrate.simpson(f, x0, xk, h2)
    rectangle_runge_rombert = integrate.runge_rombert(h1, h2, rectangle_h1,
                                                      rectangle_h2, p)
    trapeze_runge_rombert = integrate.runge_rombert(h1, h2, trapeze_h1,
                                                    trapeze_h2, p)
    simpson_runge_rombert = integrate.runge_rombert(h1, h2, simpson_h1,
                                                    simpson_h2, p)

    print('Rectangle method')
    print(f'Step {h1}: {rectangle_h1}')
    print(f'Step {h2}: {rectangle_h2}', end='\n\n')

    print('Trapeze method')
    print(f'Step {h1}: {trapeze_h1}')
    print(f'Step {h2}: {trapeze_h2}', end='\n\n')

    print('Simpson method')
    print(f'Step {h1}: {simpson_h1}')
    print(f'Step {h2}: {simpson_h2}', end='\n\n')

    print('Runge Robert method')
    print('Rectangle:', rectangle_runge_rombert)
    print('Trapeze:', trapeze_runge_rombert)
    print('Simpson:', simpson_runge_rombert, end='\n\n')

    print('-' * 25, end="\n\n\n")


if __name__ == '__main__':
    if '1' in sys.argv:
        lab03_01()
    if '2' in sys.argv:
        lab03_02()
    if '3' in sys.argv:
        lab03_03()
    if '4' in sys.argv:
        lab03_04()
    if '5' in sys.argv:
        lab03_05()
