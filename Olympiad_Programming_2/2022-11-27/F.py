def checker(a, b, c):
    return a + b > c and a + c > b and b + c > a and a ** 2 + b ** 2 == c ** 2


a, b, c = sorted(int(x) for x in input().split())

ans = None
possible = False
x1 = None
x2 = None

for i in range(0, 10001):
    if checker(a + i, b + i, c + i):
        x1 = i

for i in range(-1, -10001, -1):
    if checker(a + i, b + i, c + i):
        x2 = i
    if a + i <= 0:
        break

if x1 is not None and x2 is not None:
    possible = True
    if abs(x1) > abs(x2):
        ans = x1
    elif abs(x1) < abs(x2):
        ans = x2
    else:
        print(max(x1, x2))
elif x1 is not None:
    possible = True
    ans = x1
elif x2 is not None:
    possible = True
    ans = x2

if possible:
    print("Possible")
    print(ans)
else:
    print("Impossible")
