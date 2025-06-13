import numpy as np
from typing import Callable

def new_uniform(std_dev: float, mean: float) -> Callable[[], float]:
    return lambda: uniform(std_dev, mean)

def uniform(std_dev: float, mean: float) -> float:
    return (np.random.rand() - 0.5) * std_dev + mean

def new_normal(std_dev: float, mean: float) -> Callable[[], float]:
    return lambda: normal(std_dev, mean)

def normal(std_dev: float, mean: float) -> float:
    return np.random.normal(mean, std_dev)
