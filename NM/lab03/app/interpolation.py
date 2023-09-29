from app.format import format_polynom


def lagrange_with_check(f, x, test_point):
    res1, err1 = lagrange(f, x[1:], test_point)
    res2, err2 = lagrange(f, x[:-1], test_point)
    if err1 > err2:
        return res2, err2
    else:
        return res1, err1


def lagrange(f, x, test_point):
    y = [f(t) for t in x]
    assert len(x) == len(y)
    polynom_str = 'L(x) ='
    polynom_test_value = 0
    for i in range(len(x)):
        cur_enum_str = ''
        cur_enum_test = 1
        cur_denom = 1
        for j in range(len(x)):
            if i == j:
                continue
            cur_enum_str += f'(x-{x[j]:.2f})'
            cur_enum_test *= (test_point[0] - x[j])
            cur_denom *= (x[i] - x[j])

        polynom_str += f'+{(y[i] / cur_denom):.2f}' + cur_enum_str
        polynom_test_value += y[i] * cur_enum_test / cur_denom

    return format_polynom(polynom_str), abs(polynom_test_value - test_point[1])


def newton(f, x, test_point):
    y = [f(t) for t in x]
    assert len(x) == len(y)

    n = len(x)
    coefs = [y[i] for i in range(n)]
    for i in range(1, n):
        for j in range(n - 1, i - 1, -1):
            coefs[j] = float(coefs[j] - coefs[j - 1]) / float(x[j] - x[j - i])

    polynom_str = 'P(x) = '
    polynom_test_value = 0

    cur_multipliers_str = ''
    cur_multipliers = 1
    for i in range(n):
        polynom_test_value += cur_multipliers * coefs[i]
        if i == 0:
            polynom_str += f'{coefs[i]:.2f}'
        else:
            polynom_str += '+' + cur_multipliers_str + '*' + f'{coefs[i]:.2f}'

        cur_multipliers *= (test_point[0] - x[i])
        cur_multipliers_str += f'(x-{x[i]:.2f})'

    return format_polynom(polynom_str), abs(polynom_test_value - test_point[1])
