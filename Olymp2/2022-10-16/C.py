a = input()
digits_sum = 0
amount_6 = 0
amount_9 = 0
for x in a:
    digits_sum += int(x)
    if int(x) == 6:
        amount_6 += 1
    if int(x) == 9:
        amount_9 += 1

if digits_sum % 9 == 0:
    print(0)
elif digits_sum % 9 == 3:
    if amount_9 >= 1:
        print(1)
    elif amount_6 >= 2:
        print(2)
    else:
        print(-1)
elif digits_sum % 9 == 6:
    if amount_6 >= 1:
        print(1)
    elif amount_9 >= 2:
        print(2)
    else:
        print(-1)
else:
    print(-1)
