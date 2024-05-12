def get_good_pairs(arr, k):
    d = dict()
    for i in range(len(arr)):
        if arr[i] not in d:
            d[arr[i]] = [i]
        else:
            d[arr[i]].append(i)
    
    ans = 0
    for i in range(len(arr)):
        if k - arr[i] in d:
            idxs = d[k-arr[i]]
            for j in idxs:
                if i < j:
                    ans += 1
    return ans

_, k = map(int, input().split())
mahmoud_arr = list(map(int, input().split()))
bashar_arr = list(map(int, input().split()))

mahmoud_ans = get_good_pairs(mahmoud_arr, k)
bashar_ans = get_good_pairs(bashar_arr, k)

if mahmoud_ans > bashar_ans:
    print('MAHMOUD')
elif bashar_ans > mahmoud_ans:
    print('BASHAR')
else:
    print('DRAW')
