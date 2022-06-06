# При запуске указывать количество добавляемых пар аргументом командной строки.
# Напимер, введя python3 testsGenerator.py 1000 получите файл test.txt размером
# 3000 строк, первые 1000 строк — команды для добавления, затем 1000 строк для поиска
# и в конце 1000 строк для удаления.
# Так же в самом начале файла test.txt будет количество команд для каждого теста. Это
# необходимо для правильной работы бенчмарков.

import random 
import sys    
import os

A_LETTER = 97
LETTERS_AMOUNT = 26
MAX_KEY_LENGTH = 255
MAX_VALUE = (2 ** 64) - 1

pushFile = open('pushTest.txt', 'w')
popFile = open('popTest.txt', 'w')
findFile = open('findTest.txt', 'w')

amount = int(sys.argv[1])
for i in range(0, amount):
    key = ''
    keyLength = random.randrange(1, MAX_KEY_LENGTH)
    value = random.randrange(0, MAX_VALUE)
    for i in range(0, keyLength):
        key += chr(random.randrange(A_LETTER, A_LETTER + LETTERS_AMOUNT))
    pushFile.write(key + ' ' + str(value) + '\n')
    findFile.write(key + '\n')
    popFile.write(key + ' ' + str(value) + '\n')

pushFile.close()
findFile.close()
popFile.close()

filenames = ['pushTest.txt', 'findTest.txt', 'popTest.txt']
with open('test.txt', 'w') as outfile:
    outfile.write(str(amount) + '\n')
    for fname in filenames:
        with open(fname) as infile:
            for line in infile:
                outfile.write(line)

for x in filenames:
    os.remove(x)
