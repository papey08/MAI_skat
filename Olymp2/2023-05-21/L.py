def check(l):
    for i in range(len(l)//2):
        if l[i] != '*' and l[-i-1] != '*':
            if l[i] != l[-i-1]:
                return False
    return True

lg = input()

l = [lg[0], lg[1], lg[5], lg[6]]
g = [lg[2], lg[3], lg[4], lg[7], lg[8]]

res = 1
if not check(l) or not check(g):
    res = 0

if l[0] == '*' and l[-1] == '*':
    res *= 26
if l[1] == '*' and l[-2] == '*':
    res *= 26

if g[0] == '*' and g[-1] == '*':
    res *= 10
if g[1] == '*' and g[-2] == '*':
    res *= 10
if g[2] == '*':
    res *= 10

print(res)
