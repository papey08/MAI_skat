from enum import Enum, auto
import math

class Mode(Enum):
    DEFAULT = auto()
    MULTI_CLASS = auto()
    REGRESSION = auto()
    BINARY = auto()
    MULTI_LABEL = auto()

class Activation(Enum):
    NONE = auto()
    SIGMOID = auto()
    TANH = auto()
    RELU = auto()
    LINEAR = auto()
    SOFTMAX = auto()

def output_activation(c: Mode) -> Activation:
    if c == Mode.MULTI_CLASS:
        return Activation.SOFTMAX
    elif c == Mode.REGRESSION:
        return Activation.LINEAR
    elif c == Mode.BINARY or c == Mode.MULTI_LABEL:
        return Activation.SIGMOID
    else:
        return Activation.NONE

def get_activation(act: Activation):
    if act == Activation.SIGMOID:
        return sigmoid_f, sigmoid_df
    elif act == Activation.TANH:
        return tanh_f, tanh_df
    elif act == Activation.RELU:
        return relu_f, relu_df
    elif act == Activation.LINEAR:
        return linear_f, linear_df
    else:
        return linear_f, linear_df

def linear_f(x: float) -> float:
    return x

def linear_df(_: float) -> float:
    return 1

def relu_f(x: float) -> float:
    return max(x, 0)

def relu_df(y: float) -> float:
    return 1 if y > 0 else 0

def sigmoid_f(x: float) -> float:
    return 1 / (1 + math.exp(-x))

def sigmoid_df(y: float) -> float:
    return y * (1 - y)

def tanh_f(x: float) -> float:
    return math.tanh(x)

def tanh_df(y: float) -> float:
    return 1 - y ** 2
