import numpy as np

def tridiagonal_solve(A, b):
    v = [0 for _ in range(len(A))]
    u = [0 for _ in range(len(A))]
    v[0] = A[0][1] / -A[0][0]
    u[0] = b[0] / A[0][0]

    for i in range(1, len(A)-1):
        v[i] = A[i][i+1] / (-A[i][i] - A[i][i-1] * v[i-1])
        u[i] = (A[i][i-1] * u[i-1] - b[i]) / (-A[i][i] - A[i][i-1] * v[i-1])
    
    v[-1] = 0
    u[-1] = (A[-1][-2] * u[-2] - b[-1]) / (-A[-1][-1] - A[-1][-2] * v[-2])

    x = [0 for _ in range(len(A))]
    x[-1] = u[-1]

    for i in range(len(A)-1, 0, -1):
        x[i-1] = v[i-1] * x[i] + u[i-1]
    
    return np.array(x)
