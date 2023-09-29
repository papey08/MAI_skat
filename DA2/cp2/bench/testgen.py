import random
import sys

MAX_N = 26

file = open('test.txt', 'w')

to_compress = ''
l = int(sys.argv[1])
for i in range(l):
    to_compress += chr(ord('a') + random.randint(0, MAX_N - 1))
file.write(to_compress + '\n')

file.close()
