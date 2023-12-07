import numpy as np

def mean_square_error(y, y_correct):
    return np.sqrt(np.sum((y - y_correct)**2))
