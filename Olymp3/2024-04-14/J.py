n, x, a = map(int, input().split())

ans = n // (a // x)
if n % (a//x) != 0:
    ans += 1
print(ans)
