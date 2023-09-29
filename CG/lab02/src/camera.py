import math
import numpy as np

class Camera:
    def __init__(self, canvas):
        self.position = [0, 0, 10, 0]
        self.right = [1, 0, 0]
        self.up = [0, 1, 0]
        self.forward = [0, 0, 1]
        self.h_fov = math.pi/3
        self.height = canvas.winfo_height()
        self.width = canvas.winfo_width()
        self.v_fov = self.h_fov * (self.height/self.width)
        self.near_plane = 0.1
        self.far_plan = 100
        self.aspect = self.width/self.height

    def ViewMatrix(self):
        rx, ry, rz = self.right
        ux, uy, uz = self.up
        fx, fy, fz = self.forward
        x, y, z, w = self.position
        matrix = np.array([
            [rx, ux, fx, 0],
            [ry, uy, fy, 0],
            [rz, uz, fz, 0],
            [-x, -y, -z, 1]
        ])
        return matrix

    def projectionMatrix(self):
        a = 1/(math.tan(self.h_fov/2))
        b = (self.far_plan+self.near_plane)/(self.far_plan-self.near_plane)
        c = -2*self.near_plane*self.far_plan /(self.far_plan-self.near_plane)
        matrix = [
            [a/self.aspect, 0, 0, 0],
            [0, a, 0, 0],
            [0, 0, b, 1],
            [0, 0, c, 0]
        ]
        return matrix
        
    def orthogonalMatrix(self):
        matrix = np.eye(4, dtype=float)
        matrix[3, 0] = self.width//2
        matrix[3, 1] = self.height//2
        return matrix

    def toScreenMatrix(self):
        matrix = np.eye(4, dtype=float)
        matrix[0, 0] = self.width//2
        matrix[1, 1] = -self.height//2
        matrix[3, 0] = self.width//2
        matrix[3, 1] = self.height//2
        return matrix
        