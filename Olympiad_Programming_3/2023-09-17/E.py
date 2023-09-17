for _ in range(int(input())):
    n = int(input())
    inp = list(map(int, input().split()))
    p = [0 for _ in range(n)]
    p[0] = inp[0]
    for i in range(1, len(p)):
        p[i] = p[i-1] ^ inp[i]
    res_set = set()
    res_set.add(p[0])
    for i in range(n-1, 0, -1):
        res_set.add(p[i])
        for j in range(i):
            res_set.add(p[i] ^ p[j])
    print(len(res_set))
