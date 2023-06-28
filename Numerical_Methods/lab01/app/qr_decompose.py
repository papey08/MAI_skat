import numpy as np
from app.iterations import norm


def sign(x):
    if x > 0:
        return 1
    elif x == 0:
        return 0
    else:
        return -1


def get_householder_matrix(A, col_num):
    n = A.shape[0]
    v = np.zeros(n)
    a = A[:, col_num]
    v[col_num] = a[col_num] + sign(a[col_num]) * norm(a[col_num:])
    for i in range(col_num + 1, n):
        v[i] = a[i]
    v = v[:, np.newaxis]
    H = np.eye(n) - (2 / (v.T @ v)) * (v @ v.T)
    return H


def qr_decompose(A):
    """
    A = QR
    :return: Q, R
    """
    n = A.shape[0]
    Q = np.eye(n)
    A_i = np.copy(A)
    for i in range(n - 1):
        H = get_householder_matrix(A_i, i)
        Q = Q @ H
        A_i = H @ A_i
    return Q, A_i


def get_roots(A, i):
    n = A.shape[0]
    a11 = A[i][i]
    a12 = A[i][i + 1] if i + 1 < n else 0
    a21 = A[i + 1][i] if i + 1 < n else 0
    a22 = A[i + 1][i + 1] if i + 1 < n else 0
    return np.roots((1, -a11 - a22, a11 * a22 - a12 * a21))


def is_complex(A, i, eps):
    Q, R = qr_decompose(A)
    A_next = np.dot(R, Q)
    lambda1 = get_roots(A, i)
    lambda2 = get_roots(A_next, i)
    return abs(lambda1[0] - lambda2[0]) <= eps and \
        abs(lambda1[1] - lambda2[1]) <= eps


def get_eigen_value(A, i, eps):
    A_i = np.copy(A)
    while True:
        Q, R = qr_decompose(A_i)
        A_i = R @ Q
        if norm(A_i[i + 1:, i]) <= eps:
            return A_i[i][i], A_i
        elif norm(A_i[i + 2:, i]) <= eps and is_complex(A_i, i, eps):
            return get_roots(A_i, i), A_i


def qr_eigen_values(A, eps):
    n = A.shape[0]
    A_i = np.copy(A)
    eigen_values = []
    i = 0
    while i < n:
        cur_eigen_values, A_i_plus_1 = get_eigen_value(A_i, i, eps)
        if isinstance(cur_eigen_values, np.ndarray):
            eigen_values.extend(cur_eigen_values)
            i += 2
        else:
            eigen_values.append(cur_eigen_values)
            i += 1
        A_i = A_i_plus_1
    return eigen_values
