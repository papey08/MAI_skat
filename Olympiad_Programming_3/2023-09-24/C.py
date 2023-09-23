n = int(input())

red = sorted(list(map(int, input().split())))
blue = sorted(list(map(int, input().split())))

distr = [0 for _ in range(n)]

for i in range(n):
    distr[i] = red[i] + blue[-1 - i]

print(max(distr) - min(distr))
