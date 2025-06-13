import numpy as np
import time
import random
from tabulate import tabulate
from typing import List, Tuple, Optional

from src.activation import Mode
from src.loss import get_loss
from src.network import Network, Layer

class Logger:
    def __init__(self):
        self.headers = ["Epochs", "Elapsed", "Loss", "Accuracy"]
        self.table = []

    def init(self, network: Network):
        if network.config.mode == Mode.MULTI_CLASS:
            self.headers.append("Accuracy")

    def write_log(self, network: Network, validation: List['Pair'], duration: float, epoch: int):
        loss = self.cross_validate(network, validation)
        accuracy = self.format_accuracy(network, validation)
        row = [epoch, duration, loss, accuracy]
        self.table.append(row)
    
    def flush_log(self):
        print(tabulate(self.table, headers=self.headers, tablefmt="pretty"))

    def cross_validate(self, network: Network, validation: List['Pair']) -> float:
        predictions = [network.predict(pair.input) for pair in validation]
        responses = [pair.response for pair in validation]
        f, _ = get_loss(network.config.loss)
        return f(predictions, responses)

    def format_accuracy(self, network: Network, validation: List['Pair']) -> str:
        if network.config.mode == Mode.MULTI_CLASS:
            correct = 0
            for pair in validation:
                est = network.predict(pair.input)
                if np.argmax(pair.response) == np.argmax(est):
                    correct += 1
            return f"{correct / len(validation):.2f}"
        return ""

class Pair:
    def __init__(self, input_data: List[float], response: List[float]):
        self.input = input_data
        self.response = response

class Pairs:
    def __init__(self, pairs: List[Pair]):
        self.pairs = pairs

    def shuffle(self):
        random.shuffle(self.pairs)

    def split(self, ratio: float) -> Tuple['Pairs', 'Pairs']:
        first, second = [], []
        for pair in self.pairs:
            if ratio > random.random():
                first.append(pair)
            else:
                second.append(pair)
        return Pairs(first), Pairs(second)

    def split_size(self, size: int) -> List['Pairs']:
        return [Pairs(self.pairs[i:i + size]) for i in range(0, len(self.pairs), size)]

    def split_n(self, n: int) -> List['Pairs']:
        result = [[] for _ in range(n)]
        for i, pair in enumerate(self.pairs):
            result[i % n].append(pair)
        return [Pairs(chunk) for chunk in result]

class Trainer:
    def __init__(self, solver: 'Solver', verbosity: int):
        self.deltas = []
        self.solver = solver
        self.logger = Logger()
        self.verbosity = verbosity

    def new_training(self, layers: List[Layer]) -> List[List[float]]:
        return [np.zeros(len(layer.neurons)) for layer in layers]

    def calculate_deltas(self, network: Network, ideal: List[float]):
        last_layer = network.layers[-1]
        for i, neuron in enumerate(last_layer.neurons):
            _, df = get_loss(network.config.loss)
            self.deltas[-1][i] = df(neuron.value, ideal[i], neuron.d_activate(neuron.value))

        for i in range(len(network.layers) - 2, -1, -1):
            for j, neuron in enumerate(network.layers[i].neurons):
                total = sum(syn.weight * self.deltas[i + 1][k] for k, syn in enumerate(neuron.out))
                self.deltas[i][j] = neuron.d_activate(neuron.value) * total

    def update(self, network: Network, iteration: int):
        idx = 0
        for i, layer in enumerate(network.layers):
            for j, neuron in enumerate(layer.neurons):
                for k, syn in enumerate(neuron.in_):
                    update = self.solver.update(syn.weight, self.deltas[i][j] * syn.in_, iteration, idx)
                    syn.weight += update
                    idx += 1

    def learn(self, network: Network, example: Pair, iteration: int) -> Optional[str]:
        err = network.forward(example.input)
        if err:
            return err
        self.calculate_deltas(network, example.response)
        self.update(network, iteration)
        return None

    def train(self, network: Network, examples: Pairs, validation: Pairs, iterations: int) -> Optional[Logger]:
        self.deltas = self.new_training(network.layers)
        self.logger.init(network)
        self.solver.init(network.num_weights())
        start_time = time.time()

        for i in range(1, iterations + 1):
            examples.shuffle()
            for example in examples.pairs:
                err = self.learn(network, example, i)
                if err:
                    return err
            if self.verbosity > 0 and i % self.verbosity == 0:
                elapsed = time.time() - start_time
                self.logger.write_log(network, validation.pairs, elapsed, i)

        return self.logger

class Adam:
    def __init__(self, lr: float = 0.001, beta: float = 0.9, beta2: float = 0.999, epsilon: float = 1e-8):
        self.lr = lr
        self.beta = beta
        self.beta2 = beta2
        self.epsilon = epsilon
        self.v = []
        self.m = []

    def init(self, size: int):
        self.v = np.zeros(size)
        self.m = np.zeros(size)

    def update(self, value: float, gradient: float, iteration: int, idx: int) -> float:
        lrt = self.lr * (np.sqrt(1.0 - np.power(self.beta2, iteration))) / (1.0 - np.power(self.beta, iteration))
        self.m[idx] = self.beta * self.m[idx] + (1.0 - self.beta) * gradient
        self.v[idx] = self.beta2 * self.v[idx] + (1.0 - self.beta2) * np.power(gradient, 2.0)
        return -lrt * (self.m[idx] / (np.sqrt(self.v[idx]) + self.epsilon))

class SGD:
    def __init__(self, lr: float = 0.01, momentum: float = 0.0, decay: float = 0.0):
        self.lr = lr
        self.momentum = momentum
        self.decay = decay
        self.moments = []

    def init(self, size: int):
        self.moments = np.zeros(size)

    def update(self, value: float, gradient: float, iteration: int, idx: int) -> float:
        lr = self.lr / (1 + self.decay * iteration)
        self.moments[idx] = self.momentum * self.moments[idx] - lr * gradient
        return self.moments[idx]

class Solver:
    def init(self, size: int):
        pass

    def update(self, value: float, gradient: float, iteration: int, idx: int) -> float:
        pass

def fparam(val: float, fallback: float) -> float:
    return val if val != 0.0 else fallback
