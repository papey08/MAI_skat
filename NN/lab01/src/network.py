import numpy as np
from typing import List, Callable, Optional

from src.activation import Activation, Mode, get_activation, output_activation
from src.loss import Loss
from src.weights import new_uniform
from src.softmax import softmax

class Synapse:
    def __init__(self, weight: float):
        self.weight = weight
        self.in_ = 0.0
        self.out = 0.0
        self.is_bias = False

    def fire(self, value: float):
        self.in_ = value
        self.out = self.in_ * self.weight

class Layer:
    def __init__(self, n: int, a: Activation):
        self.neurons = [Neuron(a if a != Activation.SOFTMAX else Activation.LINEAR) for _ in range(n)]
        self.activation = a

    def fire(self):
        for neuron in self.neurons:
            neuron.fire()
        
        if self.activation == Activation.SOFTMAX:
            outs = np.array([neuron.value for neuron in self.neurons])
            sm = softmax(outs)
            for i, neuron in enumerate(self.neurons):
                neuron.value = sm[i]

    def connect(self, next_layer: 'Layer', weight: Callable[[], float]):
        for neuron in self.neurons:
            for next_neuron in next_layer.neurons:
                syn = Synapse(weight())
                neuron.out.append(syn)
                next_neuron.in_.append(syn)

    def apply_bias(self, weight: Callable[[], float]) -> List['Synapse']:
        biases = []
        for neuron in self.neurons:
            bias = Synapse(weight())
            bias.is_bias = True
            neuron.in_.append(bias)
            biases.append(bias)
        return biases

class Network:
    def __init__(self, config: 'Params'):
        self.layers = self.initialize_layers(config)
        self.biases = self.initialize_biases(config) if config.bias else []
        self.config = config

    def initialize_layers(self, config: 'Params') -> List[Layer]:
        layers = []
        for i, n_neurons in enumerate(config.layout_config):
            act = config.activation if i != len(config.layout_config) - 1 else output_activation(config.mode)
            layers.append(Layer(n_neurons, act))
        
        for i in range(len(layers) - 1):
            layers[i].connect(layers[i + 1], config.weight)
        
        for neuron in layers[0].neurons:
            neuron.in_ = [Synapse(config.weight()) for _ in range(config.inputs)]
        
        return layers

    def initialize_biases(self, config: 'Params') -> List[List[Synapse]]:
        biases = []
        for i, layer in enumerate(self.layers):
            if config.mode == Mode.REGRESSION and i == len(self.layers) - 1:
                continue
            biases.append(layer.apply_bias(config.weight))
        return biases

    def fire(self):
        for bias_layer in self.biases:
            for bias in bias_layer:
                bias.fire(1.0)
        for layer in self.layers:
            layer.fire()

    def forward(self, input_data: List[float]):
        for neuron in self.layers[0].neurons:
            for i, value in enumerate(input_data):
                neuron.in_[i].fire(value)
        
        self.fire()
        return None

    def predict(self, input_data: List[float]) -> List[float]:
        self.forward(input_data)
        out_layer = self.layers[-1]
        return [neuron.value for neuron in out_layer.neurons]

    def num_weights(self) -> int:
        return sum(len(neuron.in_) for layer in self.layers for neuron in layer.neurons)

class Neuron:
    def __init__(self, activation: Activation):
        self.activation = activation
        self.in_ = []
        self.out = []
        self.value = 0.0

    def fire(self):
        total = sum(syn.out for syn in self.in_)
        self.value = self.activate(total)
        for syn in self.out:
            syn.fire(self.value)

    def activate(self, x: float) -> float:
        f, _ = get_activation(self.activation)
        return f(x)

    def d_activate(self, x: float) -> float:
        _, df = get_activation(self.activation)
        return df(x)

class Params:
    def __init__(self, inputs: int, layout_config: List[int], activation: Activation, mode: Mode,
                 loss: Loss, bias: bool, weight: Optional[Callable[[], float]] = None):
        self.inputs = inputs
        self.layout_config = layout_config
        self.activation = activation
        self.mode = mode
        self.loss = loss
        self.bias = bias
        self.weight = weight if weight is not None else new_uniform(0.5, 0.0)
