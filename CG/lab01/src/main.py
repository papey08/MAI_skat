import matplotlib.pyplot as plt
import numpy as np

# a = float(input())
# ta = float(input())
# tb = float(input())

a = 20
ta = -100
tb = -2

tlin = np.linspace(ta, tb, 1000)

def x(t, a):
    return 3*a*t / (1 + t**3)

def y(t, a):
    return 3*a*(t**2) / (1 + t**3)

ax = plt.gca()
ax.axhline(y=0, color='k')
ax.axvline(x=0, color='k')

ax.arrow(x=0, y=5, dx=-0.5, dy=-1)
ax.arrow(x=0, y=5, dx=0.5, dy=-1)

ax.arrow(x=5, y=0, dx=-1, dy=1)
ax.arrow(x=5, y=0, dx=-1, dy=-1)

plt.plot(x(tlin, a), y(tlin, a))

plt.show()
