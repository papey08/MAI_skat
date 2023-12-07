out = open('output.txt', 'w')

with open('input.txt') as inp:
    n = int(inp.read().strip())
    out.write(str(4**n) + '\n')
