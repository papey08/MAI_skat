import app.matrix as matrix


def tridiagonal_to_common(m: list):
    n = len(m)
    res = [[0]*n for _ in range(n)]
    res[0][0], res[0][1] = m[0][0], m[0][1]
    for i in range(1, n - 1):
        res[i][i-1], res[i][i], res[i][i+1] = m[i][0], m[i][1], m[i][2]
    res[-1][-2], res[-1][-1] = m[-1][0], m[-1][1]
    return matrix.MyMatrix(res)
