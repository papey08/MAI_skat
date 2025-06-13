import numpy as np
from enum import Enum

class Loss(Enum):
    NONE = 0
    CROSS_ENTROPY = 1
    BIN_CROSS_ENTROPY = 2
    MEAN_SQUARED = 3

def bin_cross_entropy_f(estimate, ideal):
    epsilon = np.finfo(float).eps
    ce = ideal * np.log(estimate + epsilon) + (1 - ideal) * np.log(1 - estimate + epsilon)
    return -np.mean(ce)

def bin_cross_entropy_df(estimate, ideal, _):
    return estimate - ideal

def cross_entropy_f(estimate, ideal):
    ce = ideal * np.log(estimate)
    return -np.mean(ce)

def cross_entropy_df(estimate, ideal, _):
    return estimate - ideal

def mean_squared_f(estimate, ideal):
    return np.mean(np.square(estimate - ideal))

def mean_squared_df(estimate, ideal, activation):
    return activation * (estimate - ideal)

def get_loss(loss: Loss):
    if loss == Loss.CROSS_ENTROPY:
        return cross_entropy_f, cross_entropy_df
    elif loss == Loss.MEAN_SQUARED:
        return mean_squared_f, mean_squared_df
    elif loss == Loss.BIN_CROSS_ENTROPY:
        return bin_cross_entropy_f, bin_cross_entropy_df
    else:
        return cross_entropy_f, cross_entropy_df

def loss_to_string(loss: Loss):
    if loss == Loss.CROSS_ENTROPY:
        return "CE"
    elif loss == Loss.BIN_CROSS_ENTROPY:
        return "BinCE"
    elif loss == Loss.MEAN_SQUARED:
        return "MSE"
    else:
        return "N/A"
