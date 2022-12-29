import random
import sys

file = open('test.txt', 'w')

n = int(sys.argv[1])
m = int(sys.argv[2])

file.write(str(n) + ' ' + str(m) + '\n')

for _ in range(n):
    w = random.randint(1, 10)
    c = random.randint(1, 10)
    file.write(str(w) + ' ' + str(c) + '\n')

file.close()
