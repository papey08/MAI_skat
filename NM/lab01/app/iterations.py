import numpy as np
import math


def norm(X):
    res = 0
    for x in X:
        res += x ** 2
    return math.sqrt(res)


def solve_system_iterative(a, b, eps):
    """
    Uses iterative method to solve Ax=b
    :param a: system
    :param b: free members
    :param eps:
    :return: x and the number of iterations
    """

    diag_sum = abs(a[0][0]) + abs(a[0][1]) + abs(a[-1][-1]) + abs(a[-1][-2])
    for i in range(1, len(a)-1):
        diag_sum += abs(a[i][i-1]) + abs(a[i][i]) + abs(a[i][i+1])
    full_sum = 0
    for i in range(len(a)):
        for j in range(len(a[i])):
            full_sum += abs(a[i][j])
    if diag_sum <= (full_sum - diag_sum)*2:
        return None, -1
    
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
        cur_x = alpha @ prev_x + beta
        iterations += 1
        if norm(prev_x - cur_x) <= eps and norm(prev_x - cur_x) < 1:
            return cur_x, iterations
    return None, -1
