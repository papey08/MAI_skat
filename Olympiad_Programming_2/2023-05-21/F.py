w = int(input())

for x in range(31, 100000000):
    y = 2 * x
    z = 6 * x
    if 50 + w <= ((x - 30) * y * z) / 1000:
        print(z, y, x)
        break
