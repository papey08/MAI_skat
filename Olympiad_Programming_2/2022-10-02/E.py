n = int(input())
k = int(input())

ans = -1

if 2**k <= n:
    ans = ((n & ((2**63 - 1 >> k) << k)) >> k)
    bin_str = bin(ans)[2:]
    i = 0
    if bin_str[::-1][i] != '1':
        while i < len(bin_str):
            if bin_str[::-1][i] == '1':
                ans ^= 1 << i
                break
            ans |= (1 << i)
            i += 1

if ans << k > 0:
    print(ans << k)
else:
    print(ans)
