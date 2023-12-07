import math

def func(t):
    if t == 1:
        return 1
    
    ans = 0
    sqrt_t = int(math.sqrt(t))
    
    for i in range(1, sqrt_t + 1):
        if t % i == 0:
            ans += 2
    
    if sqrt_t * sqrt_t == t:
        ans -= 1
    
    return ans

q = int(input())
r = int(input())

t = q * (r-1)

print(func(t))
