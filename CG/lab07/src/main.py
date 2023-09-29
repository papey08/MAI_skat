import numpy as np
import matplotlib.pyplot as plt
import scipy.interpolate as si
from matplotlib.widgets import TextBox

N = 6

points = [[0, x] for x in range(N)]
points = np.array(points)
x = points[:, 0]
t = points[:, 1]
axcolor = 'white'

def interpol():
    global x
    global t
    ipl_t = np.linspace(min(t), max(t), 100)
    x_tup = si.splrep(t, x, k=4)
    x_list = list(x_tup)
    xl = x.tolist()
    x_list[1] = xl + [0.0, 0.0, 0.0, 0.0]
    x_i = si.splev(ipl_t, x_list)
    return [ipl_t, x_i]

def update0(input):
    global x
    global cords
    amp = float(input)
    x[0] = amp
    cords = interpol()
    l.set_ydata(cords[1])
    a.set_ydata(x)

def update1(input):
    global x
    global cords
    amp = float(input)
    x[1] = amp
    cords = interpol()
    l.set_ydata(cords[1])
    a.set_ydata(x)

def update2(input):
    global x
    global cords
    amp = float(input)
    x[2] = amp
    cords = interpol()
    l.set_ydata(cords[1])
    a.set_ydata(x)

def update3(input):
    global x
    global cords
    amp = float(input)
    x[3] = amp
    cords = interpol()
    l.set_ydata(cords[1])
    a.set_ydata(x)

def update4(input):
    global x
    global cords
    amp = float(input)
    x[4] = amp
    cords = interpol()
    l.set_ydata(cords[1])
    a.set_ydata(x)

def update5(input):
    global x
    global cords
    amp = float(input)
    x[5] = amp
    cords = interpol()
    l.set_ydata(cords[1])
    a.set_ydata(x)

updates = [update0, update1, update2, update3, update4, update5]

fig = plt.figure()
ax = fig.add_subplot(211)
a, = plt.plot(t, x, '-og')
cords = interpol()
l, = plt.plot(cords[0], cords[1])
plt.xlim([min(t), max(t)])
plt.ylim([0, 11])
plt.gcf().canvas.manager.set_window_title("© Попов Матвей М8О-308Б-20")

axes = [None for _ in range(N)]
rect = [0.7, 0.05 * N, 0.2, 0.05]

for i in range(len(axes)):
    axes[i] = plt.axes(rect, facecolor=axcolor)
    rect[1] -= 0.05

fields = [None for _ in range(N)]

for i in range(len(fields)):
    fields[i] = TextBox(axes[i], 'dot ' + str(i) + '    ', '0')
    fields[i].on_submit(updates[i])

plt.show()
