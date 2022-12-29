import random
import requests
import sys

word_site = "https://www.mit.edu/~ecprice/wordlist.10000"

file = open('test.txt', 'w')

response = requests.get(word_site)
words = [str(x)[2:-1] for x in response.content.splitlines()]

text = ''

while len(text) < int(sys.argv[1]):
    text += random.choice(words)

file.write(text + '\n')

for _ in range(int(sys.argv[2])):
    file.write(random.choice(words) + '\n')

file.close()
