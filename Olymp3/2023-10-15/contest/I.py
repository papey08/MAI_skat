def eratosthenes(n):
    sieve = list(range(n + 1))
    sieve[1] = 0
    for i in sieve:
        if i > 1:
            for j in range(2*i, len(sieve), i):
                sieve[j] = 0
    sieve_set = []
    for x in sieve:
        if x != 0:
            sieve_set.append(x)
    return sieve_set

sieve = eratosthenes(1000000)

def check_digits(n, banned):
    for c in str(n):
        if c in banned:
            return False
    return True

def find_smallest(banned):
    for n in sieve:
        if check_digits(n, banned):
            print(n)
            return
    print(-1)

n = int(input())
banned = []
for _ in range(n):
    banned.append(input())

find_smallest(banned)
