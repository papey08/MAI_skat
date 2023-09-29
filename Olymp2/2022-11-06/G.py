a = []
with open('g.in', 'r') as f:
    for line in f.readlines():
        a.append(line)

ans = (int(a[0]) * int(a[1]))/(int(a[1]) + int(a[2]))

with open('g.out', 'w') as f:
    f.write(str(ans))