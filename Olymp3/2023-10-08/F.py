import math

print('? 0')
a = float(input())
print('? 180')
b = float(input())
print('? 270')
h = float(input())

c1 = math.sqrt(a**2 + h**2)
c2 = math.sqrt(b**2 + h**2)

s = 0.5 * h * (a + b)

print('!', int(round(((a+b)*c1*c2/(4.0*s)))))
