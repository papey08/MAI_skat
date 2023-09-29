import math

b = [0] * (math.ceil(math.log2(1000000000))+1)

M = 1000000007

n = int(input())
nums = list(map(int, input().split()))
for x in nums:
    c = 0
    while x:
        x //=2
        c += 1
    b[c] += 1
ans = 0
for x in b:
    if x:
        ans = (ans + x*(ans+1)) % M
print(ans)
