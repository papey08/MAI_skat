import numpy as np
import matplotlib.pyplot as plt
from matplotlib.widgets import TextBox

def PlotHalfSphere(radius, precision, transparency):
    precision += 1
    phi = []
    theta = []
    pointTheta = 0
    pointPhi = 0
    dTheta = (0.5 * pi) / (precision - 1)
    dPhi = (2 * pi) / (precision - 1)
    tmp = []
    for _ in range(precision):
        tmp.append(pointTheta)
    theta.append(tmp)
    for _ in range(precision - 1):
        pointTheta = pointTheta + dTheta
        tmp = []
        for _ in range(precision):
            tmp.append(pointTheta)
        theta.append(tmp)
    tmp = []
    tmp.append(pointPhi)
    for _ in range(precision - 1):
        pointPhi = pointPhi + dPhi
        tmp.append(pointPhi)
    for _ in range(precision):
        phi.append(tmp)
    x = radius * np.sin(theta) * np.cos(phi)
    y = radius * np.sin(theta) * np.sin(phi)
    z = radius * np.cos(theta)
    axis.plot_surface(x, y, z, alpha=transparency)
    z = np.zeros((precision, precision))
    axis.plot_surface(x, y, z, alpha=transparency)
    plt.show()


def AxisInstallation():
    axis.set_xlim([-radius - 0.5, radius + 0.5])
    axis.set_ylim([-radius - 0.5, radius + 0.5])
    axis.set_zlim([0, 1.5 * radius])

def ChangePrecision(input):
    axis.clear()
    AxisInstallation()
    PlotHalfSphere(radius, int(input), transparency)


pi = np.pi
radius = 10
transparency = 1

f = plt.figure()
axis = f.add_subplot(111, projection='3d')
AxisInstallation()
precisionField = plt.axes([0.5, 0.05, 0.1, 0.05])
precisionTextBox = TextBox(precisionField, 'Precision: ', '40')
precisionTextBox.on_submit(ChangePrecision)
PlotHalfSphere(radius, 40, transparency)
