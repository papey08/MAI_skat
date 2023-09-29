import itertools

n, m = (int(x) for x in input().split())
r = sorted(int(x) for x in input().split())
b = sorted(int(x) for x in input().split())
res = 0
for i in range(n):
    for j in range(m):
        for sub_r in itertools.combinations(r, i + 1):
            for sub_b in itertools.combinations(b, j + 1):
                if sum(sub_b) == sum(sub_r):
                    res += 1

print(res)
