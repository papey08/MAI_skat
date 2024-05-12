amount = int(input())

n = 0
for _ in range(amount):
    n ^= int(input())

print(bin(n)[2:].count('1'))
