import numpy as np

def sweep(A, b):
    p = np.zeros(len(b))
    q = np.zeros(len(b))
    p[0] = -A[0][1]/A[0][0]
    q[0] = b[0]/A[0][0]
    
    for i in range(1, len(p)-1):
        p[i] = -A[i][i+1]/(A[i][i] + A[i][i-1]*p[i-1])
        q[i] = (b[i] - A[i][i-1]*q[i-1])/(A[i][i] + A[i][i-1]*p[i-1])
    
    p[-1] = 0
    q[-1] = (b[-1] - A[-1][-2]*q[-2])/(A[-1][-1] + A[-1][-2]*p[-2])
    x = np.zeros(len(b))
    x[-1] = q[-1]
    
    for i in reversed(range(len(b)-1)):
        x[i] = p[i]*x[i+1] + q[i]
    
    return x
