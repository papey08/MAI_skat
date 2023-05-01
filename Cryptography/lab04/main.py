import os

from app.compare import compare
from app.generate import generate_random_letters_file
from app.generate import generate_random_words_file

orig_text01 = 'texts/american_psycho.txt'
orig_text02 = 'texts/fight_club.txt'

def lab04_01(n):
    res = compare(orig_text01, orig_text02, n)
    print('Два осмысленных текста'.ljust(50), f'Длина: {n}'.ljust(20), f'Доля совпадений: {res:.5f}', end='\n\n')

def lab04_02(n):
    gen_letters = 'texts/generated_random_letters.txt'
    generate_random_letters_file(gen_letters, n)
    res = compare(orig_text01, gen_letters, n)
    print('Осмысленный текст и текст из случайных букв'.ljust(50), f'Длина: {n}'.ljust(20), f'Доля совпадений: {res:.5f}', end='\n\n')
    os.remove(gen_letters)

def lab04_03(n):
    gen_words = 'texts/generated_random_words.txt'
    generate_random_words_file(gen_words, n)
    res = compare(orig_text02, gen_words, n)
    print('Осмысленный текст и текст из случайных слов'.ljust(50), f'Длина: {n}'.ljust(20), f'Доля совпадений: {res:.5f}', end='\n\n')
    os.remove(gen_words)

def lab04_04(n):
    gen_letters01 = 'texts/generated_random_letters01.txt'
    gen_letters02 = 'texts/generated_random_letters02.txt'
    generate_random_letters_file(gen_letters01, n)
    generate_random_letters_file(gen_letters02, n)
    res = compare(gen_letters01, gen_letters02, n)
    print(f'Два текста из случайных букв'.ljust(50), f'Длина: {n}'.ljust(20), f'Доля совпадений: {res:.5f}', end='\n\n')
    os.remove(gen_letters01)
    os.remove(gen_letters02)

def lab04_05(n):
    gen_words01 = 'texts/generated_random_words01.txt'
    gen_words02 = 'texts/generated_random_words02.txt'
    generate_random_words_file(gen_words01, n)
    generate_random_words_file(gen_words02, n)
    res = compare(gen_words01, gen_words02, n)
    print(f'Два текста из случайных слов'.ljust(50), f'Длина: {n}'.ljust(20), f'Доля совпадений: {res:.5f}', end='\n\n')
    os.remove(gen_words01)
    os.remove(gen_words02)

if __name__ == '__main__':
    print()
    for n in [1000, 5000, 25000, 50000, 100000]:
        lab04_01(n)
        lab04_02(n)
        lab04_03(n)
        lab04_04(n)
        lab04_05(n)
