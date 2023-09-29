import numpy as np
import json
import sys

import app.matrix as matrix
from app.conv import tridiagonal_to_common
from app.iterations import solve_system_iterative
from app.zeidel import solve_system_zeidel
from app.rotation import rotation
from app.qr_decompose import qr_eigen_values


def lab01_01():
    print('-' * 10, '01-01', '-' * 10, sep='', end='\n\n')
    E = matrix.MyMatrix([[1, 0, 0, 0],
                         [0, 1, 0, 0],
                         [0, 0, 1, 0],
                         [0, 0, 0, 1]])
    f = open("tests/01-01.json")
    data = json.loads(f.read())
    X = matrix.MyMatrix(data["A"])
    L, U = X.lu_decompose()
    det = X.det()
    X_inv = X.inversed()
    res = X.solve_system(data["b"])
    b = matrix.MyMatrix([data["b"]]).transposed()
    print("LU decompose:")
    print("L:", L, sep='\n')
    print("U:", U, sep='\n')
    print("\nLU decompose check: ", end='')
    if L * U == X:
        print("OK decompose correct")
    else:
        print("WRONG")
    print("\nDeterminant:", det)
    print("\nInversed matrix:", X_inv, sep='\n')
    print("\nInversion check: ", end='')
    if X_inv * X == E:
        print("OK inversion correct")
    else:
        print("WRONG")
    print("\nSystem solution:\n", res)
    print("\nSolution check: ", end='')
    if X * res == b:
        print("OK solution correct")
    else:
        print("WRONG")
    print('-' * 25, end="\n\n\n")
    f.close()


def lab01_02():
    print('-'*10, '01-02', '-'*10, sep='', end='\n\n')
    f = open("tests/01-02.json")
    data = json.loads(f.read())
    X = matrix.MyMatrix(data["A"])
    # X = tridiagonal_to_common(X)
    res, _ = X.solve_system_tridiagonal(data["b"])
    if res is None:
        print('Unable to use this method')
    else:
        b = matrix.MyMatrix([data["b"]]).transposed()
        print("System solution:\n", res)
        print("\nSolution check: ", end='')
        X = tridiagonal_to_common(X)
        if X * res == b:
            print("OK solution correct")
        else:
            print("WRONG")
    print('-'*25, end="\n\n\n")
    f.close()


def lab01_03():
    print('-' * 10, '01-03', '-' * 10, sep='', end='\n\n')
    f = open("tests/01-03.json")
    data = json.loads(f.read())
    X = np.array(data["A"], dtype='float')
    b = np.array(data["b"], dtype='float')
    eps = data["eps"]
    print("Iteration method:")
    res, n = solve_system_iterative(X, b, eps)
    if n == -1:
        print('Unable to use this method')
    else:
        for x in res:
            print(x)
        print("\nNumber of iterations: ", n, end='\n\n')
        x_matr = matrix.MyMatrix(X.tolist())
        b_matr = matrix.MyMatrix(b.tolist()).transposed()
        res_matr = matrix.MyMatrix(res.tolist()).transposed()
        print("Iterations method check: ", end='')
        if x_matr * res_matr == b_matr:
            print("OK solution correct")
        else:
            print("WRONG")
        print("\nZeidel method:")
        res, n = solve_system_zeidel(X, b, eps)
        for x in res:
            print(x)
        print("\nNumber of iterations: ", n, end='\n\n')
        x_matr = matrix.MyMatrix(X.tolist())
        b_matr = matrix.MyMatrix(b.tolist()).transposed()
        res_matr = matrix.MyMatrix(res.tolist()).transposed()
        print("Zeidel method check: ", end='')
        if x_matr * res_matr == b_matr:
            print("OK solution correct")
        else:
            print("WRONG")
    print('-' * 25, end="\n\n\n")
    f.close()


def lab01_04():
    print('-' * 10, '01-04', '-' * 10, sep='', end='\n\n')
    f = open("tests/01-04.json")
    data = json.loads(f.read())
    X = np.array(data["A"], dtype='float')
    eps = data["eps"]
    values, vectors, iterations = rotation(X, eps)
    if iterations == -1:
        print('Matrix is not symmetrical')
    else:
        values_matr = matrix.MyMatrix(values.tolist())
        vectors_matr = matrix.MyMatrix(vectors.tolist())
        print('Eigen values:')
        print(values_matr)
        print('\nEigen vectors:')
        print(vectors_matr)
        print('\nNumber of iterations:', iterations)
    print('-' * 25, end="\n\n\n")
    f.close()


def lab01_05():
    print('-' * 10, '01-05', '-' * 10, sep='', end='\n\n')
    f = open("tests/01-05.json")
    data = json.loads(f.read())
    X = np.array(data["A"], dtype='float')
    eps = data["eps"]
    values = qr_eigen_values(X, eps)
    values_matr = matrix.MyMatrix(values)
    print('Eigen values:')
    print(values_matr)
    print('-' * 25, end="\n\n\n")
    f.close()


if __name__ == "__main__":
    if '1' in sys.argv:
        lab01_01()
    if '2' in sys.argv:
        lab01_02()
    if '3' in sys.argv:
        lab01_03()
    if '4' in sys.argv:
        lab01_04()
    if '5' in sys.argv:
        lab01_05()
