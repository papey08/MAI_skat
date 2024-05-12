output = open('joke.out', 'w')

def solve(n, l, ans, nums, num):
    if l == '' and len(ans) == n:
        return ans
    elif l == '' and len(ans) < n:
        return []
    else:
        num += l[0]
        l = l[1:]
        if num == '0':
            return []
        if len(num) == 2 and int(num) in nums:
            return []
        elif len(num) == 2 and int(num) not in nums:
            if int(num) > n:
                return []
            nums.add(int(num))
            ans.append(int(num))
            num = ''
            return solve(n, l, ans.copy(), nums.copy(), num)
        elif len(num) == 1 and int(num) in nums:
            return solve(n, l, ans.copy(), nums.copy(), num)
        elif len(num) == 1 and int(num) not in nums:
            sol1 = solve(n, l, ans.copy(), nums.copy(), num)
            if int(num) > n:
                return []
            nums.add(int(num))
            ans.append(int(num))
            num = ''
            sol2 = solve(n, l, ans.copy(), nums.copy(), num)
            if len(sol2) == 0:
                return sol1
            else:
                return sol2


with open('joke.in') as f:
    l = f.readline().strip()
    n = len(l) if len(l) <= 9 else 9 + (len(l)-9) // 2
    nums = set()
    ans = []
    num = ''
    ans = solve(n, l, ans, nums, num)
    for x in ans:
        output.write(str(x) + ' ')


output.close()
