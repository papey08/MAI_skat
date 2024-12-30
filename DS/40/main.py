import numpy as np
import json

ITERATIONS = 25

def get_input(filename='./input.json'):
    f = open(filename)
    data = json.loads(f.read())
    f.close
    return data['a'], data['b']

def solve(a, b):
    a = np.array(a, dtype=float)
    b = np.array(b, dtype=float)
    y = np.zeros(len(b))
    
    def grad(y):
        return 2*a@y - 2*b
    def step(g):
        return (g@g) / (2*g@a@g)

    for _ in range(ITERATIONS):
        g = grad(y)
        y = y - step(g)*g

    return np.linalg.norm(y)

if __name__ == "__main__":
    a, b = get_input()
    print(f'{solve(a, b):.4f}')  
