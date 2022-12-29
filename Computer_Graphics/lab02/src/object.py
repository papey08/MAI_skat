import numpy as np
import matrix as mtx

class Object:
    def __init__(self, sx, sy, sz):
        self.rotX = 0
        self.rotY = 0
        self.rotZ = 0
        self.dx = 0
        self.dy = 0
        self.dz = 0
        self.scale = 1
        self.startPoints = []
        self.points = []
        self.polygon = []
        self.readModel()
        self.scretchStartModel(sx, sy, sz)

    def readModel(self):
        self.startPoints.append(np.array(list(map(float, [3.0, -3, 0.0])) + [1]))
        self.startPoints.append(np.array(list(map(float, [1.5, -3, 2.59807621135331594])) + [1]))
        self.startPoints.append(np.array(list(map(float, [-1.5, -3, 2.59807621135331594])) + [1]))
        self.startPoints.append(np.array(list(map(float, [-3, -3, 0])) + [1]))
        self.startPoints.append(np.array(list(map(float, [-1.5, -3, -2.59807621135331594])) + [1]))
        self.startPoints.append(np.array(list(map(float, [1.5, -3, -2.59807621135331594])) + [1]))
        self.startPoints.append(np.array(list(map(float, [7.0, 4, 0.0])) + [1]))
        self.startPoints.append(np.array(list(map(float, [3.5, 4, 6.062177826491])) + [1]))
        self.startPoints.append(np.array(list(map(float, [-3.5, 4, 6.062177826491])) + [1]))
        self.startPoints.append(np.array(list(map(float, [-7.0, 4, 0])) + [1]))
        self.startPoints.append(np.array(list(map(float, [-3.5, 4, -6.062177826491])) + [1]))
        self.startPoints.append(np.array(list(map(float, [3.5, 4, -6.062177826491])) + [1]))
        self.polygon.append(list(map(int, [0, 1, 7, 6])))
        self.polygon.append(list(map(int, [1, 2, 8, 7])))
        self.polygon.append(list(map(int, [2, 3, 9, 8])))
        self.polygon.append(list(map(int, [3, 4, 10, 9])))
        self.polygon.append(list(map(int, [4, 5, 11, 10])))
        self.polygon.append(list(map(int, [0, 5, 11, 6])))
        self.polygon.append(list(map(int, [0, 1, 2, 3, 4, 5])))
        self.polygon.append(list(map(int, [6, 7, 8, 9, 10, 11])))

    def scretchStartModel(self, sx, sy, sz):
        self.points = self.startPoints @ mtx.stretchingMatrix(sx, sy, sz)
        