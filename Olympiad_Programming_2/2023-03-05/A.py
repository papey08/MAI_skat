a = list(map(int, input().split()))
a_i = [0] * 10
for i in range(len(a)):
    a_i[a[i]] = i

coded = input()

coded_2nd = [a_i[int(c)] for c in coded]
res = ['' for _ in range(len(coded) // 2)]

for i in range(0, len(coded), 2):
    j = ((i // 2) // 10) * 10
    res[j + int(coded_2nd[i])] = str(coded_2nd[i+1])

res = ''.join(res)
print(res)
