import random
import sys

textAmount = int(sys.argv[1])

file = open('test.txt', 'w')

pattern = 'abcabaabcabaabcabaabcabaabcabaabcabaabcabaabcabaabcabaabcabaabcabaab'

wrongPatterns = ['abcabaabcabaabcabaabcabaabcabaabcabaabcabaabcaba',
                 'abcabaabcabaabcabaabcabaabcabaabcabaabcabaabcaba',
                 'abcabaabcabaabcabaabcabaabcabaabcabaabcabaabcaba',
                 'abcabaabcabaabcabaabcabaabcabaabcabaabcabaabcaba',
                 'abcabaabcabaabcabaabcabaabcabaabcabaabcabaabcaba',
                 'abcabaabcabaabcabaabcabaabcabaabcabaabcabaabcaba',
                 'abcabaabcabaabcabaabcabaabcabaabcabaabcabaabcaba',
                 'abcabaabcabaabcabaabcabaabcabaabcabaabcabaabcabaabcabaabcabaabcabaab']

text = ''

for x in range(textAmount):
    text += random.choice(wrongPatterns)

file.write(pattern + '\n' + text + '\n')

file.close()
