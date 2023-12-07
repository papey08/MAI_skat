import numpy as np

def analytial(x_range, y_range, t_range, h_x, h_y, tau, solution):
    x = np.arange(*x_range, h_x)
    y = np.arange(*y_range, h_y)
    t = np.arange(*t_range, tau)

    res = np.zeros((len(t), len(x), len(y)))
    for i in range(len(x)):
        for j in range(len(y)):
            for k in range(len(t)):
                res[k][i][j] = solution(x[i], y[j], t[k])
    
    return res
