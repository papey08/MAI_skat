# при запуске указывать размерность квадратной матрицы аргументом командной строки
# Например, введя python3 testgen.py 1000 получите файл test.txt с квадратной матрицей размером 1000 на 1000

import random
import sys

MIN_ELEM = -1000000
MAX_ELEM = 1000000

n = int(sys.argv[1])

file = open('test.txt', 'w')
file.write(str(n) + '\n')

for i in range(n):
    for j in range(n):
        file.write(str(random.randrange(MIN_ELEM, MAX_ELEM)) + ' ')
    file.write('\n')

file.close()
