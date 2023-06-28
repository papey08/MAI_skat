a, b = map(int, input().split())
ans = []
res = 2**b

while a != res:
    if a * 2 <= res:
        ans.append('0')
        a *= 2
    else:
        ans.append('1')
        a = res - a
        
ans = ''.join(ans[::-1])
print(ans)