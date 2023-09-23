A1, P1 = map(int, input().split())  
A2, P2 = map(int, input().split())  

total_A_goals = A1 + A2
total_P_goals = P1 + P2
if total_A_goals > total_P_goals:
    result = "A"  
elif total_P_goals > total_A_goals:
    result = "P"
else:
    result = "D"
print(result)
