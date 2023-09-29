import numpy as np


def find_max_upper_element(X):
    """
    :return: coords of the max element above the main diagonal
    """
    n = X.shape[0]
    i_max, j_max = 0, 1
    max_elem = abs(X[0][1])
    for i in range(n):
        for j in range(i + 1, n):
            if abs(X[i][j]) > max_elem:
                max_elem = abs(X[i][j])
                i_max = i
                j_max = j
    return i_max, j_max


def matrix_norm(X):
    """
    :return: L2 norm for elements above the main diagonal
    """
    n = X.shape[0]
    norm = 0
    for i in range(n):
        for j in range(i + 1, n):
            norm += X[i][j]**2
    return np.sqrt(norm)


def rotation(A, eps):
    """
    :return: eigen values, eigen vectors, number of iterations
    """
    n = A.shape[0]
    for i in range(n):
        for j in range(i, n):
            if A[i][j] != A[j][i]:
                return None, None, -1

    A_i = np.copy(A)
    eigen_vectors = np.eye(n)
    iterations = 0
    while matrix_norm(A_i) > eps:
        i_max, j_max = find_max_upper_element(A_i)
        if A_i[i_max][i_max] - A_i[j_max][j_max] == 0:
            phi = np.pi / 4
        else:
            phi = 0.5 * np.arctan(2 * A_i[i_max][j_max] / (A_i[i_max][i_max] -
                                                           A_i[j_max][j_max]))
        U = np.eye(n)
        U[i_max][j_max] = -np.sin(phi)
        U[j_max][i_max] = np.sin(phi)
        U[i_max][i_max] = np.cos(phi)
        U[j_max][j_max] = np.cos(phi)
        A_i = U.T @ A_i @ U
        eigen_vectors = eigen_vectors @ U
        iterations += 1
    eigen_values = np.array([A_i[i][i] for i in range(n)])
    return eigen_values, eigen_vectors, iterations
