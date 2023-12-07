from math import gcd
from functools import reduce
def find_gcd(l):
    x = reduce(gcd, l)
    return x

for _ in range(int(input())):
    _ = input()
    a = list(map(int, input().split()))
    b = list(map(int, input().split()))
    a_gcd = find_gcd(a)
    b_gcd = find_gcd(b)
    if a_gcd == b_gcd:
        print(0)
    elif max(a_gcd, b_gcd) % min(a_gcd, b_gcd) == 0:
        print(1)
    else:
        print(2)
    
