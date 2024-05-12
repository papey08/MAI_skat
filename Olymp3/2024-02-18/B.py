n = int(input())

ans = 1

if n <= 3:
    ans = 1
else:
    bomb = 2
    boss = 0
    while boss%n != (bomb-1)%n and boss%n != (bomb)%n and boss%n != (bomb+1)%n:
        ans += 1
        boss += 1
        bomb += 2


print(ans)    
