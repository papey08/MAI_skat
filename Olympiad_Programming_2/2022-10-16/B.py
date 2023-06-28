def check_tr(a, b, c) -> bool:
    return a + b > c and b + c > a and a + c > b


def checker_right(a, b, c, d) -> bool:
    if check_tr(a, b, c + d) and ((a / d == b / c) or (a / c == b / d)):
        return True
    else:
        return False

a = int(input())
b = int(input())
c = int(input())
d = int(input())

if checker_right(a, b, c, d) or \
        checker_right(a, c, b, d) or\
        checker_right(a, d, b, c) or\
        checker_right(b, c, a, d) or\
        checker_right(b, d, a, c) or\
        checker_right(c, d, a, b):
    print(1)
else:
    print(0)