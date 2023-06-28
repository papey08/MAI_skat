import random
import string
import requests

def generate_random_letters_file(filename, n):
    random_string = ''.join(random.choice(string.ascii_letters) for _ in range(n))
    with open(filename, 'w') as file:
        file.write(random_string)

def generate_random_words_file(filename, n):
    word_site = "https://www.mit.edu/~ecprice/wordlist.10000"
    response = requests.get(word_site)
    word_list = [str(x)[2:-1] for x in response.content.splitlines()]
    words_to_use = []
    while len(' '.join(words_to_use)) < n:
        word = random.choice(word_list)
        if not any(c not in string.ascii_letters for c in word):
            words_to_use.append(word)
    random_string = ' '.join(words_to_use)
    while len(random_string) < n:
        random_string += ' ' if random.randint(0, 1) == 0 else '\n'
    with open(filename, 'w') as file:
        file.write(random_string)
