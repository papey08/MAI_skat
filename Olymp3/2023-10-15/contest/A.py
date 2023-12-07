pf = int(input())
pe = int(input())
pd = int(input())
k = int(input())

if k == 0:
    print('F')
elif k == 1:
    print('E')
elif k == pf + pe + pd + 1:
    print('Champ AB')
elif k > pf + pe + (pd + 1) // 2:
    print('C')
elif k == pf + pe + (pd + 1) // 2:
    print('C/D')
elif k > pf + (pe + 1) // 2:
    print('D')
elif k == pf + (pe + 1) // 2:
    print('D/E')
else:
    print('E')
