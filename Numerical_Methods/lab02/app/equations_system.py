import numpy as MyMatrix
from app.check import check_system


def get_q(l1, r1, l2, r2, dphi1_dx1, dphi1_dx2, dphi2_dx1, dphi2_dx2):
    """
    :return: q coefficient for iteration method
    """
    x1 = (l1 + r1) / 2 + abs(r1 - l1)
    x2 = (l2 + r2) / 2 + abs(r2 - l2)
    max1 = abs(dphi1_dx1([x1, x2])) + abs(dphi1_dx2([x1, x2]))
    max2 = abs(dphi2_dx1([x1, x2])) + abs(dphi2_dx2([x1, x2]))
    return max(max1, max2)


def mul(a, b):
    return a @ b


def l_inf_norm(a):
    abs_a = [abs(i) for i in a]
    return max(abs_a)


def iterations(f1, f2, phi1, phi2, dphi1_dx1, dphi1_dx2, dphi2_dx1, dphi2_dx2, l1, r1, l2, r2, eps):
    if not check_system(dphi1_dx1, dphi1_dx2, dphi2_dx1, dphi2_dx2):
        return 0, -1
    x_prev = [(l1 + r1) * 0.5, (l2 + r2) * 0.5]
    q = get_q(l1, r1, l2, r2, dphi1_dx1, dphi1_dx2, dphi2_dx1, dphi2_dx2)
    if q >= 1:
        return 0, -1
    i = 0
    while i <= 50:
        i += 1
        x = [phi1(x_prev), phi2(x_prev)]
        if q / (1 - q) * l_inf_norm([(x[i] - x_prev[i]) for i in range(len(x))]) < eps:
            return x, i
        x_prev = x
    return 0, -1


def newton(f1, f2, df1_dx1, df1_dx2, df2_dx1, df2_dx2, l1, r1, l2, r2, eps):
    x_prev = MyMatrix.array([(l1 + r1) / 2, (l2 + r2) / 2])
    jacobi = [[df1_dx1(x_prev), df1_dx2(x_prev)],
              [df2_dx1(x_prev), df2_dx2(x_prev)]]
    jacobi_inversed = MyMatrix.linalg.inv(MyMatrix.array(jacobi))
    i = 0
    while i <= 50:
        i += 1
        x = x_prev - mul(jacobi_inversed, MyMatrix.array([f1(x_prev), f2(x_prev)]))
        if l_inf_norm([(x[i] - x_prev[i]) for i in range(len(x))]) < eps:
            return x, i
        x_prev = x

    return 0, -1


def zeidel(x_0, eps, phi1, phi2, dphi1_dx1, dphi1_dx2, dphi2_dx1, dphi2_dx2):
    if not check_system(dphi1_dx1, dphi1_dx2, dphi2_dx1, dphi2_dx2):
        return 0, -1
    phi_x_i = x_0
    i = 0
    while i <= 50:
        x_0 = phi_x_i.copy()
        x_k = phi1(x_0)
        phi_x_i = [x_k, phi2([x_k, x_0[1]])]
        i += 1
        if max(abs(x_0[0] - phi_x_i[0]), abs(x_0[1] - phi_x_i[1])) < eps:
            return phi_x_i, i
    return 0, -1
