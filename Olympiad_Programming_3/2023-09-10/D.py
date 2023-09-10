import math

for _ in range(int(input())):
    x, y = map(int, input().split())
    if y == 100:
        print(-1)
    else:
        b = 0
        while b / (x + b) < y / 100:
            b += 1
        print(x + b)
