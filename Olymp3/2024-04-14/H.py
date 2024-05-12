_ = input()
flags = dict()
for x in list(map(int, input().split())):
    if x not in flags:
        flags[x] = 1
    else:
        flags[x] += 1

ans = 0
for x in flags:
    ans += 2**flags[x] - 1
    ans %= (10**9 + 7)  

print(ans)
