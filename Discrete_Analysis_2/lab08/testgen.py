import random
import sys

file = open('test.txt', 'w')

m = int(sys.argv[1])
n = int(sys.argv[2])
MAX_NUM = 50

file.write(str(m) + ' ' + str(n) + '\n')

for _ in range(m):
    for _ in range(n):
        file.write(str(random.randint(0, MAX_NUM)) + ' ')
    file.write(str(random.randint(0, MAX_NUM)) + '\n')

file.close()
