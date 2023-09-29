# при запуске указывать количество строк в тесте аргументом командной строки
# Например, введя python3 testGen.py 1000 получите файл test.txt с 1000 пар ключ-значение

import random
import sys

MAX_KEY = 65535
MIN_KEY = 0
VALUE_SIZE = 64
A_LETTER = 97
LETTERS_AMOUNT = 26

file = open('test.txt', 'w')

amount = int(sys.argv[1])

for i in range(0, amount):
    key = random.randrange(MIN_KEY, MAX_KEY)
    value = ''
    for j in range(0, VALUE_SIZE):
        value += chr(random.randrange(A_LETTER, A_LETTER + LETTERS_AMOUNT))
    file.write(str(key) + ' ' + value + '\n')

file.close()
