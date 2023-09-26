# при запуске указывать количество чисел в массивах аргументом командной строки
# Например, введя python3 testgen.py 1000 получите файл test.txt с двумя массивами по 1000 чисел в каждом

import random
import sys

MIN_NUM = -1000000
MAX_NUM = 1000000

n = int(sys.argv[1])

file = open('test.txt', 'w')
file.write(str(n) + '\n')

for i in range(n):
    file.write(str(random.randrange(MIN_NUM, MAX_NUM)) + ' ')
file.write('\n')

for i in range(n):
    file.write(str(random.randrange(MIN_NUM, MAX_NUM)) + ' ')
file.write('\n')

file.close()
