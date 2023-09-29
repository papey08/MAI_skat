import matplotlib.pyplot as plt

from app.lab01 import lu_decompose as decomposite, \
    solve_system as solve_lu_one_matrix


def least_squares(x, y, n):
    assert len(x) == len(y)
    A = []
    b = []
    for k in range(n + 1):
        A.append([sum(map(lambda x: x**(i + k), x)) for i in range(n + 1)])
        b.append(sum(map(lambda x: x[0] * x[1]**k, zip(y, x))))
    lu = decomposite(A)
    return solve_lu_one_matrix(lu, b)


def p(coefs, x):
    return sum([c * x**i for i, c in enumerate(coefs)])


def sum_squared_errors(x, y, ls_coefs):
    y_ls = [p(ls_coefs, x_i) for x_i in x]
    return sum((y_i - y_ls_i)**2 for y_i, y_ls_i in zip(y, y_ls))


def draw_plot(x_i, y_i, ls1, ls2):
    plt.scatter(x_i, y_i)
    plt.plot(x_i, [p(ls1, i) for i in x_i], label='degree 1')
    plt.plot(x_i, [p(ls2, i) for i in x_i], label='degree 2')
    plt.title('03-03')
    plt.legend()
    plt.show()

