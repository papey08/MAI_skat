x = int(input())
y = int(input())
z = int(input())

a = int(input())
b = int(input())
c = int(input())

def count(arr):
    if sum(arr) < 5:
        return a
    elif 5 <= sum(arr) < 10:
        return b
    else:
        return c

print(min(
    count([x]) + count([y]) + count([z]),
    count([x, y]) + count([z]),
    count([x, z]) + count([y]),
    count([y, z]) + count([x]),
    count([x, y, z])
    ))
