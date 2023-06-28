n = int(input())
branches = list(map(int, input().split()))

# ctr = 0

for i in range(1, n+1):
    if i != branches[i-1]:
        print(i, branches[i-1])
    else:
        temp = i
        while i != branches[temp-1]:
            temp += 1
        print(i, branches[temp-1])
