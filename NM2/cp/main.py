import numpy as np
import matplotlib.pyplot as plt

from app.finite_difference import finite_differrence
from app.tools import mean_square_error

def f(x):
    return (1 + x) * np.exp(-x * x)

equation = {
    'p': lambda x: 4 * x,
    'q': lambda x: 4 * x**2 + 2,
    'f': lambda x: 0,
}

cond1 = {
    'a': 0,
    'b': 1,
    'c': 1,
}

cond2 = {
    'a': 4,
    'b': -1,
    'c': 23 * (np.e ** (-4)),
}

borders = (0, 2)
h = (0.1, 0.01, 0.001)

if __name__ == '__main__':
    x = [np.arange(borders[0], borders[1], hi) for hi in h]
    y_correct = [f(xi) for xi in x]
    y = [finite_differrence(cond1, cond2, equation, borders, hi)[:-1] for hi in h]

    for i in range(len(h)):
        print(f'Mean square error with h = {h[i]}: {mean_square_error(y[i], y_correct[i])}')

    plt.figure(figsize=(12, 7))
    plt.plot(x[2], y_correct[2], label='Correct')
    for i in range(len(x)):
        plt.plot(x[i], y[i], label='h = '+str(h[i]))
    plt.legend()
    plt.show()
