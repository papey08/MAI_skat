_, _ = map(int, input().split())
a = list(map(int, input().split()))
kexs = list(map(int, input().split()))

def bin_search(a, n):
    l, r = 0, len(a)
    while l < r:
        m = (l + r) // 2
        if a[m] < n:
            l = m + 1
        else:
            r = m
    return l

for kex in kexs:
    l, r = 0, 10**10
    while l < r:
        m = (l + r) // 2
        if m - bin_search(a, m) < kex:
            l = m + 1
        else:
            r = m
    print(r-1, end=' ')       
print()
