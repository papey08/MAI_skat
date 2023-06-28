

a = ''
b = ''

with open('bulls.in', 'r') as file1:
    a = file1.readline().strip()
    b = file1.readline().strip()

n = 0
d = set()

for i in range(4):
    d.add(a[i])
    if a[i] == b[i]:
        n += 1

m = 0

for i in range(4):
    if b[i] in d:
        m += 1

m -= n
res = str(n) + ' ' + str(m)
# print(n, m)

with open('bulls.out', 'w') as file2:
    file2.write(res)
