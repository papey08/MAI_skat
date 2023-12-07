import numpy as np

from app.sweep import sweep

def finite_differrence(cond1, cond2, equation, borders, h=0.01, accuracy=2):
    x = np.arange(borders[0], borders[1] + h, h)
    N = np.shape(x)[0]

    A = np.zeros((N, N))
    b = np.zeros(N)

    for i in range(1, N - 1):
        A[i][i-1] = 1/h**2 - equation['p'](x[i])/(2*h)
        A[i][i] = -2/h**2 + equation['q'](x[i])
        A[i][i+1] = 1/h**2 + equation['p'](x[i])/(2*h)
        b[i] = equation['f'](x[i])

    if accuracy == 1:
        A[0][0] = cond1['a'] - cond1['b']/h
        A[0][1] = cond1['b']/h
        b[0] = cond1['c']

        A[N-1][N-2] = -cond2['b']/h
        A[N-1][N-1] = cond2['a'] + cond2['b']/h
        b[N-1] = cond2['c']

    elif accuracy == 2:
        p = equation['p']
        q = equation['q']
        a1 = cond1['a']; b1 = cond1['b']; c1 = cond1['c']
        a2 = cond2['a']; b2 = cond2['b']; c2 = cond2['c']

        A[0][0] = a1 - (3*b1)/(2*h) + ((2 - h*p(x[1]))*b1)/((2 + h*p(x[1]))*2*h)
        A[0][1] = 2*b1/h + ((h*h*q(x[1]) - 2)*b1)/((2 + h*p(x[1]))*h)
        A[0][2] = 0
        b[0] = c1
        
        A[N-1][N-3] = 0
        A[N-1][N-2] = -2*b2/h - ((h*h*q(x[N-2]) - 2)*b2)/((2 - h*p(x[N-2]))*h)
        A[N-1][N-1] = a2 + 3*b2/(2*h) - ((2 + h*p(x[N-2]))*b2)/((2 - h*p(x[N-2]))*2*h)
        b[N-1] = c2
    
    return sweep(A, b)
