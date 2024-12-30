from math import sqrt

EPS = 1e-5

def gaussian_elimination(matrix, augmented=False):
    rows = len(matrix)
    cols = len(matrix[0])
    
    for i in range(rows):
        max_row = i
        for j in range(i + 1, rows):
            if abs(matrix[j][i]) > abs(matrix[max_row][i]):
                max_row = j

        matrix[i], matrix[max_row] = matrix[max_row], matrix[i]

        if abs(matrix[i][i]) < EPS:
            continue

        for j in range(i + 1, cols):
            matrix[i][j] /= matrix[i][i]
        matrix[i][i] = 1.0

        for j in range(i + 1, rows):
            factor = matrix[j][i]
            for k in range(i, cols):
                matrix[j][k] -= factor * matrix[i][k]
    for i in range(rows - 1, -1, -1):
        if abs(matrix[i][i]) < EPS:
            continue
        for j in range(i - 1, -1, -1):
            factor = matrix[j][i]
            for k in range(i, cols):
                matrix[j][k] -= factor * matrix[i][k]
    return matrix

def find_null_space(matrix):
    rows = len(matrix)
    cols = len(matrix[0])
    augmented_matrix = [row + [0] for row in matrix]
    reduced_matrix = gaussian_elimination(augmented_matrix)
    pivot_columns = []
    for i in range(rows):
        for j in range(cols):
            if abs(reduced_matrix[i][j]) == 1:
                pivot_columns.append(j)
                break
    null_space = []
    for free_col in set(range(cols)) - set(pivot_columns):
        solution = [0] * cols
        solution[free_col] = 1
        for row in range(rows):
            if reduced_matrix[row][free_col] != 0:
                solution[row] = -reduced_matrix[row][free_col]
        null_space.append(solution)

    return null_space

def check_vectors_if_integer(matrix):
    for v in matrix:
        for n in v:
            if abs(round(n) - n) > EPS:
                return False
    return True

def find_integer_vectors(matrix):
    while not check_vectors_if_integer(matrix):
        for i in range(len(matrix)):
            for j in range(len(matrix)):
                matrix[i][j] *= 10
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            matrix[i][j] = int(round(matrix[i][j]))
    return matrix


def normalize_vectors(vectors):
    normalized_basis = []
    for vector in vectors:
        length = sqrt(sum(coord ** 2 for coord in vector))
        if length != 0:
            normalized_vector = [coord / length for coord in vector]
            normalized_basis.append(normalized_vector)
    return normalized_basis

def solve(matrix):
    null_space = find_null_space(matrix)
    normalized_space = normalize_vectors(null_space)
    integer_space = find_integer_vectors(null_space)
    
    return integer_space, normalized_space
