from app.check import check_iterations
from app.check import check_newton


def iterations(f, phi, l, r, eps):
    """
    :param f: f(x) = 0
    :param phi: phi(x) = x
    :param l, r: [l, r]
    :return: x and number of iterations
    """
    done = check_iterations(f, phi, l, r)
    x_prev = (l + r) * 0.5
    i = 0
    while i <= 50:
        i += 1
        x = phi(x_prev)
        if abs(f(x) - f(x_prev)) < eps:
            return x, i
        x_prev = x
    return 0, -1, done


def newton(f, df, l, r, eps):
    """
    :param f: f(x) = 0
    :param df: f'(x)
    :param l, r: [l, r]
    :return: x and number of iterations
    """
    done = check_newton(f, df, l, r)
    x_prev = (l + r) * 0.5
    i = 0
    while i <= 50:
        i += 1
        x = x_prev - f(x_prev) / df(x_prev)
        if abs(f(x) - f(x_prev)) < eps:
            return x, i
        x_prev = x
    return 0, -1, done
