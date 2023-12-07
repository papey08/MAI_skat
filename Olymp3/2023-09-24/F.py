import math

def get_dist(t_x, t_y, x, y):
    return math.sqrt((t_x - x)**2 + (t_y - y)**2)

n = int(input())

_, _, t_x, t_y = map(int, input().split())

a_nearest_dist = 10**10
a_throws = dict()

for _ in range(n):
    x, y = map(int, input().split())
    dist = get_dist(t_x, t_y, x, y)
    if dist < a_nearest_dist:
        a_nearest_dist = dist
    if dist in a_throws:
        a_throws[dist] += 1
    else:
        a_throws[dist] = 1

b_nearest_dist = 10**10
b_throws = dict()

for _ in range(n):
    x, y = map(int, input().split())
    dist = get_dist(t_x, t_y, x, y)
    if dist < b_nearest_dist:
        b_nearest_dist = dist
    if dist in b_throws:
        b_throws[dist] += 1
    else:
        b_throws[dist] = 1

ans_team = ''
ans_score = 0

if a_nearest_dist < b_nearest_dist:
    ans_team = 'A'
    for x in a_throws:
        if x < b_nearest_dist:
            ans_score += a_throws[x]
else:
    ans_team = 'R'
    for x in b_throws:
        if x < a_nearest_dist:
            ans_score += b_throws[x]

print(ans_team, ans_score)
