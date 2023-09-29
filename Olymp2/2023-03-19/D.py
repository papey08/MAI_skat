def isDivide(rs, cs, rc, cc):
    return (rs - cs + rc - cc)%2 == 0

r, c = map(int, input().split())
rs, cs, rc, cc = map(int, input().split())
f = input()
s = ''
if f == 'C':
    s = 'S'
else:
    s = 'C'

if not isDivide(rs, cs, rc, cc):
    print(f)
else:
    print(s)
