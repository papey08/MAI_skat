with open('e.in', 'r') as f:
    l = int(f.readline())

ans = (l // 9) * 81 + (l % 9) ** 2

with open('e.out', 'w') as f:
    f.write(str(ans))
