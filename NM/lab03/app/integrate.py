def rectangle_trapeze(f, l, r, h, is_rectangle=True):
    if l > r:
        return None
    result = 0
    cur_x = l
    while cur_x < r:
        if is_rectangle:
            result += f((cur_x + cur_x + h)*0.5)
        else:
            result += 0.5*(f(cur_x + h) + f(cur_x))
        cur_x += h
    return h*result


def simpson(f, l, r, h):
    if l > r:
        return None
    while ((l-r)//h) % 2 != 0:
        h *= 0.9
    result = 0
    cur_x = l + h
    while cur_x < r:
        result += f(cur_x - h) + 4*f(cur_x) + f(cur_x + h)
        cur_x += 2*h
    return result * h / 3


def runge_rombert(h1, h2, i1, i2, p):
    return i1 + (i1 - i2) / ((h2 / h1)**p - 1)
