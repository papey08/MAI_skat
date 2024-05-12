_ = input()
answers = input()

a, b, c, d, e = 0, 0, 0, 0, 0
for ans in answers:
    if ans == 'a':
        a += 1
    elif ans == 'b':
        b += 1
    elif ans == 'c':
        c += 1
    elif ans == 'd':
        d += 1
    else:
        e += 1

print(min(a, b, c, d, e), max(a, b, c, d, e))
