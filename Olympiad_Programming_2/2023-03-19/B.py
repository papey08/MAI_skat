# import math

n = int(input())
b = list(map(int, input().split()))

scheme = [0] * n
a = []
ans = 10**6
mid = n // 2

for i in range(mid, n):
    scheme[i] = i-mid+1
for i in range(mid):
    scheme[i] = i+mid+1+n%2


for t in b:
    a.append(scheme[t-1])
    if abs(a[-1] - t) < ans:
        ans = abs(a[-1] - t)

print(ans)
print(*a)
