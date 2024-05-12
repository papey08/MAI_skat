n, m = 0, 0
cars = []

def solve(x, y, t):
    ans = 0
    for c in cars:
        dist = c[1]-c[0]
        p = t % (dist*2)
        if p > dist:
            p = dist*2 - p
        p += c[0]
        if x <= p <= y:
            ans += 1
    return ans

output = open('knockout.out', 'w')

with open('knockout.in') as file:
    n, m = map(int, file.readline().strip().split())
    for _ in range(n):
        cars.append([*map(int, file.readline().strip().split())])
    for _ in range(m):
        output.write(str(solve(*map(int, file.readline().strip().split()))) + '\n')

output.close()
