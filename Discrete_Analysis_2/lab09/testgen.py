import random
import sys

file = open('test.txt', 'w')

n = int(sys.argv[1])
m = int(sys.argv[2])

start = random.randint(1, n)
finish = random.randint(1, n)

file.write(str(n) + ' ' + str(m) + ' ' + str(start) + ' ' + str(finish) + '\n')

MAX_WEIGHT = 10 ** 9
edges = [set() for _ in range(n)]

for _ in range(m):
    v_from = 0
    v_to = 0
    weight = 0
    while True:
        v_from = random.randint(0, n - 1)
        v_to = random.randint(0, n - 1)
        weight = random.randint(1, MAX_WEIGHT)
        if v_to not in edges[v_from] and v_to != v_from:
            edges[v_from].add(v_to)
            edges[v_to].add(v_from)
            break
    file.write(str(v_from + 1) + ' ' + str(v_to + 1) + ' ' + str(weight) + '\n')

file.close()
