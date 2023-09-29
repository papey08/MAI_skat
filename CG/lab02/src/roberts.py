import numpy as np

def RobertsAlgo(modelMatrix, polygon, matrixT, position):
    D = np.array([[-1],[-1],[-1]])
    matrixV = np.zeros((4, len(polygon)))
    for i in range(len(polygon)):
        curMatrix = modelMatrix[polygon[i][:3]][:3, :3]
        matrix = np.linalg.inv(curMatrix) @ D
        matrixV[0][i] = matrix[0]
        matrixV[1][i] = matrix[1]
        matrixV[2][i] = matrix[2]
        matrixV[3][i] = -1

    matrixVT = np.linalg.inv(matrixT) @ matrixV
    final = position @ matrixVT
    return final
