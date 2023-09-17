def get_days(t):
    days = [0 for _ in range(t[-1] // 1440 + 1)]
    for x in t:
        days[x // 1440] += 1
    return days

def get_max_seq_of_success_days(days):
    success_days_ctr = 0
    max_success_days = 0
    for d in days:
        if d >= 3:
            success_days_ctr += 1
        else:
            if max_success_days < success_days_ctr:
                max_success_days = success_days_ctr
            success_days_ctr = 0
    if success_days_ctr > max_success_days:
        max_success_days = success_days_ctr
    return max_success_days

_ = input()
t = list(map(int, input().split()))

ans = 0

for time_zone in range(25):
    for i in range(len(t)):
        t[i] += 60
    temp_ans = get_max_seq_of_success_days(get_days(t))
    if temp_ans > ans:
        ans = temp_ans

print(ans)
