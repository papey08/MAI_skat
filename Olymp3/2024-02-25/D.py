chain = input()
table = []
for _ in range(len(chain)+1):
    table.append(dict())

table[0]['A'] = 0
table[0]['T'] = 0
table[0]['G'] = 0
table[0]['C'] = 0
# table[0][chain[0]] += 1

for i in range(1, len(chain)+1):
    table[i]['A'] = table[i-1]['A']
    table[i]['T'] = table[i-1]['T']
    table[i]['G'] = table[i-1]['G']
    table[i]['C'] = table[i-1]['C']
    table[i][chain[i-1]] += 1

for _ in range(int(input())):
    s, e = map(int, input().split())
    s -= 1
    a = table[e]['A'] - table[s]['A']
    t = table[e]['T'] - table[s]['T']
    g = table[e]['G'] - table[s]['G']
    c = table[e]['C'] - table[s]['C']
    ans = []
    while (len(ans) < 4):
        m = max(a, t, g, c) 
        if m == a:
            ans.append('A')
            a = -1
        elif m == t:
            ans.append('T')
            t = -1
        elif m == g:
            ans.append('G')
            g = -1
        elif m == c:
            ans.append('C')
            c = -1
    print(''.join(ans))
