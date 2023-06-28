import sys
import json
import numpy as np
import matplotlib.pyplot as plt

import app.cauchy_problem as cp
import app.boundary_value_problem as bvp
from app.tools import *


def lab04_01():

    def f(x, y, z):
        return 12*y / (x**2)

    def g(x, y, z):
        return z

    def exact(x):
        c2 = 1
        c1 = 1
        return c2*(x**4) + c1/(x**3)

    print('-' * 10, '04-01', '-' * 10, sep='', end='\n\n')

    file = open('tests/04-01.json')
    data = json.loads(file.read())
    y0 = data['fx']
    dy0 = data['dfx']
    borders = (data['xl'], data['xr'])
    h = data['h']
    file.close()

    h_euler = h;x_euler, y_euler = cp.euler(f, g, y0, dy0, borders, h_euler)
    plt.plot(x_euler, y_euler, label=f'Euler, h={h}')
    x_euler2, y_euler2 = cp.euler(f, g, y0, dy0, borders, h_euler/2)
    plt.plot(x_euler2, y_euler2, label=f'Euler, h={h/2}')

    x_i_euler, y_i_euler = cp.implicit_euler(f, y0, dy0, borders, h)
    plt.plot(x_i_euler, y_i_euler, label=f'Implicit Euler, h={h}')
    x_i_euler2, y_i_euler2 = cp.implicit_euler(f, y0, dy0, borders, h/2)
    plt.plot(x_i_euler2, y_i_euler2, label=f'Implicit Euler, h={h/2}')

    x_runge, y_runge = cp.runge_kutta(f, g, y0, dy0, borders, h)
    plt.plot(x_runge, y_runge, label=f'Runge Kutta, h={h}')
    x_runge2, y_runge2 = cp.runge_kutta(f, g, y0, dy0, borders, h / 2)
    plt.plot(x_runge2, y_runge2, label=f'Runge Kutta, h={h/2}')

    x_adams, y_adams = cp.adams(f, g, y0, dy0, borders, h)
    plt.plot(x_adams, y_adams, label=f'Adams, h={h}')
    x_adams2, y_adams2 = cp.adams(f, g, y0, dy0, borders, h / 2)
    plt.plot(x_adams2, y_adams2, label=f'Adams, h={h/2}')

    x_exact = [i for i in np.arange(borders[0], borders[1] + h, h)]
    x_exact2 = [i for i in np.arange(borders[0], borders[1] + h / 2, h / 2)]
    y_exact = [exact(x_i) for x_i in x_exact]
    y_exact2 = [exact(x_i) for x_i in x_exact2]
    x_exact_for_euler = [i for i in np.arange(borders[0], borders[1] + h_euler, h_euler)]
    x_exact2_for_euler = [i for i in np.arange(borders[0], borders[1] + h_euler / 2, h_euler / 2)]
    y_exact_for_euler = [exact(x_i) for x_i in x_exact_for_euler]
    y_exact2_for_euler = [exact(x_i) for x_i in x_exact2_for_euler]
    plt.plot(x_exact, y_exact, label='Exact')

    print('Errors')
    print(f'h = {h}')
    print('Euler:', error(y_euler, y_exact_for_euler))
    print('Implicit Euler:', error(y_i_euler, y_exact))
    print('Runge Kutta:', error(y_runge, y_exact))
    print('Adams:', error(y_adams, y_exact), end='\n\n')
    print(f'h = {h/2}')
    print('Euler:', error(y_euler2, y_exact2_for_euler))
    print('Implicit Euler:', error(y_i_euler2, y_exact2))
    print('Runge Kutta:', error(y_runge2, y_exact2))
    print('Adams:', error(y_adams2, y_exact2), end='\n\n')
    print()

    print('Runge Romberg')
    print('Euler:', runge_rombert(h, h / 2, y_euler, y_euler2, 4))
    print('Implicit Euler:', runge_rombert(h, h/2, y_i_euler, y_i_euler2, 4))
    print('Runge Kutta:', runge_rombert(h, h / 2, y_runge, y_runge2, 4))
    print('Adams:', runge_rombert(h, h / 2, y_adams, y_adams2, 4), end='\n\n')

    plt.title('04-01')
    plt.legend()
    plt.show()

    print('-' * 25, end="\n\n\n")


def lab04_02():

    def ddf(x, f, df):
        return ((2*x+1)*df - 2*f) / (x+0.00001)

    def f(x):
        return 2*x + 1 + np.exp(2*x)

    def p(x):
        return -(2*x + 1)

    def q(x):
        return 2

    def right_f(x):
        return 0

    print('-' * 10, '04-02', '-' * 10, sep='', end='\n\n')

    file = open('tests/04-02.json')
    data = json.loads(file.read())
    equation = {'p': p, 'q': q, 'f': right_f}
    bcondition1 = data['bcondition1']
    bcondition2 = data['bcondition2']
    borders = data['borders']
    h = data['h']
    file.close()

    x = np.arange(borders[0], borders[1]+h, h)
    y = f(x)
    y1 = bvp.shooting_method(ddf, borders, bcondition1, bcondition2, h, f)
    y2 = bvp.finite_difference_method(ddf, f, bcondition1, bcondition2, equation, borders, h)

    h2 = h / 2
    y1_2 = bvp.shooting_method(ddf, borders, bcondition1, bcondition2, h2, f)
    y2_2 = bvp.finite_difference_method(ddf, f, bcondition1, bcondition2, equation, borders, h2)

    print("Runge Rombert errors:")
    print("Shooting method:", bvp.sqr_error(y1, bvp.runge_rombert(y1, y1_2, h, h2, 1)))
    print("Finite difference method:", bvp.sqr_error(y2, bvp.runge_rombert(y2, y2_2, h, h2, 1)))
    print()

    print("Exact solution errors:")
    print("Shooting method:", bvp.sqr_error(y1, y))
    print("Finite difference method:", bvp.sqr_error(y2, y))
    print()

    plt.figure(figsize=(12, 7))
    plt.plot(x, y, label='Exact')
    plt.plot(x, y1, label='Shooting')
    plt.plot(x, y2, label='Finite difference')
    plt.grid()
    plt.title('04-02')
    plt.legend()
    plt.show()

    print('-' * 25, end="\n\n\n")


if __name__ == '__main__':
    if '1' in sys.argv:
        lab04_01()
    if '2' in sys.argv:
        lab04_02()
