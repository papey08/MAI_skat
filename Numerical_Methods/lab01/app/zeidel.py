import numpy as np
from app.iterations import norm


def zeidel_multiplication(alpha, x, beta):
    """
    :return: alhpa * x + beta
    """
    res = np.copy(x)
    for i in range(alpha.shape[0]):
        res[i] = beta[i]
        for j in range(alpha.shape[1]):
            res[i] += alpha[i][j] * res[j]
    return res


def solve_system_zeidel(a, b, eps):
    """
    Uses zeidel method to solve ax=b
    :param a: system
    :param b: free members
    :param eps:
    :return: x and number of iterations
    """
    n = a.shape[0]
    alpha = np.zeros_like(a, dtype='float')
    beta = np.zeros_like(b, dtype='float')
    for i in range(n):
        for j in range(n):
            if i == j:
                alpha[i][j] = 0
            else:
                alpha[i][j] = -a[i][j] / a[i][i]
        beta[i] = b[i] / a[i][i]
    iterations = 0
    cur_x = np.copy(beta)
    while iterations <= 50:
        prev_x = np.copy(cur_x)
        cur_x = zeidel_multiplication(alpha, prev_x, beta)
        iterations += 1
        if norm(prev_x - cur_x) <= eps:
            return cur_x, iterations
    return None, -1
