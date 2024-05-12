_ = input()
a = list(map(int, input().split()))
b = list(map(int, input().split()))

a_max = sum(a)
b_max = sum(b)
a_min = len(a)
b_min = len(b)

if a_max + a_min - b_max - b_min > 0:
    print('ALICE')
elif a_max + a_min - b_max - b_min < 0:
    print('BOB')
else:
    print('TIED')
