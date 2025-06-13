import numpy as np

def standardize(xx):
    m = np.mean(xx)
    s = np.std(xx, ddof=1)

    if s == 0:
        s = 1

    xx[:] = (xx - m) / s

def softmax(xx):
    xx = np.array(xx)
    max_val = np.max(xx)
    exp_vals = np.exp(xx - max_val)
    return exp_vals / np.sum(exp_vals)
