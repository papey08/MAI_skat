import numpy as np
from app.tridiagonal_solve import tridiagonal_solve

def variable_directions(x_range, y_range, t_range, h_x, h_y, tau, a, mu1, mu2, phi_0, phi_1, phi_2, phi_3, psi):
    x = np.arange(*x_range, h_x)
    y = np.arange(*y_range, h_y)
    t = np.arange(*t_range, tau)
    res = np.zeros((len(t), len(x), len(y)))

    for x_id in range(len(x)):
        for y_id in range(len(y)):
            res[0][x_id][y_id] = psi(x[x_id], y[y_id], mu1, mu2)
    
    for t_id in range(1, len(t)):
        U_halftime = np.zeros((len(x), len(y)))
        
        for x_id in range(len(x)):
            res[t_id][x_id][0] = phi_2(x[x_id], t[t_id], a, mu1, mu2)
            res[t_id][x_id][-1] = phi_3(x[x_id], t[t_id], a, mu1, mu2)
            U_halftime[x_id][0] = phi_2(x[x_id], t[t_id] - tau/2, a, mu1, mu2)
            U_halftime[x_id][-1] = phi_3(x[x_id], t[t_id] - tau/2, a, mu1, mu2)
        
        for y_id in range(len(y)):
            res[t_id][0][y_id] = phi_0(y[y_id], t[t_id], a, mu1, mu2)
            res[t_id][-1][y_id] = phi_1(y[y_id], t[t_id], a, mu1, mu2)
            U_halftime[0][y_id] = phi_0(y[y_id], t[t_id] - tau/2, a, mu1, mu2)
            U_halftime[-1][y_id] = phi_1(y[y_id], t[t_id] - tau/2, a, mu1, mu2)
        
        for y_id in range(1, len(y)-1):
            A = np.zeros((len(x)-2, len(x)-2))
            b = np.zeros((len(x)-2))

            A[0][0] = 2 * h_x**2 * h_y**2 + 2 * a * tau * h_y**2
            A[0][1] = -a * tau * h_y**2
            for i in range(1, len(A) - 1):
                A[i][i-1] = -a * tau * h_y**2
                A[i][i] = 2 * h_x**2 * h_y**2 + 2 * a * tau * h_y**2
                A[i][i+1] = -a * tau * h_y**2
            A[-1][-2] = -a * tau * h_y**2
            A[-1][-1] = 2 * h_x**2 * h_y**2 + 2 * a * tau * h_y**2

            for x_id in range(1, len(x)-1):
                b[x_id-1] = (
                    res[t_id-1][x_id][y_id-1] * a * tau * h_x**2
                    + res[t_id-1][x_id][y_id] * (2 * h_x**2 * h_y**2 - 2 * a * tau * h_x**2)
                    + res[t_id-1][x_id][y_id+1] * a * tau * h_x**2
                )
            b[0] -= (-a * tau * h_y**2) * phi_0(y[y_id], t[t_id] - tau/2, a, mu1, mu2)
            b[-1] -= (-a * tau * h_y**2) * phi_1(y[y_id], t[t_id] - tau/2, a, mu1, mu2)
            U_halftime[1:-1, y_id] = np.array(tridiagonal_solve(A, b))
        
        for x_id in range(1, len(x)-1):
            A = np.zeros((len(y)-2, len(y)-2))
            b = np.zeros((len(y)-2))

            A[0][0] = 2 * h_x**2 * h_y**2 + 2 * a * tau * h_x**2
            A[0][1] = -a * tau * h_x**2
            for i in range(1, len(A) - 1):
                A[i][i-1] = -a * tau * h_x**2
                A[i][i] = 2 * h_x**2 * h_y**2 + 2 * a * tau * h_x**2
                A[i][i+1] = -a * tau * h_x**2
            A[-1][-2] = -a * tau * h_x**2
            A[-1][-1] = 2 * h_x**2 * h_y**2 + 2 * a * tau * h_x**2

            for y_id in range(1, len(y)-1):
                b[y_id-1] = (
                    U_halftime[x_id-1][y_id] * a * tau * h_y**2
                    + U_halftime[x_id][y_id] * (2 * h_x**2 * h_y**2 - 2 * a * tau * h_y**2)
                    + U_halftime[x_id+1][y_id] * a * tau * h_y**2
                )
            b[0] -= (-a * tau * h_x**2) * phi_2(x[x_id], t[t_id], a, mu1, mu2)
            b[-1] -= (-a * tau * h_x**2) * phi_3(x[x_id], t[t_id], a, mu1, mu2)
            res[t_id][x_id][1:-1] = tridiagonal_solve(A, b)

    return res
