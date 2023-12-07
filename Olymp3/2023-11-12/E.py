n = int(input())

answers = []

for _ in range(n):
    _ = int(input())
    passports = list(map(int, input().split()))

    flag_incorrect = False
    
    # проверка на повторяющиеся идентификационные номера в рамках одной породы
    ids = set()
    for i in range(0, len(passports), 3):
        if passports[i] not in ids:
            ids.add(passports[i])
        else:
            answers.append('INCORRECT')
            flag_incorrect = True
            break

    if flag_incorrect:
        continue

    # проверка 2
    males = set()
    for i in range(1, len(passports), 3):
        males.add(passports[i])
    for i in range(2, len(passports), 3):
        if passports[i] in males:
            answers.append('INCORRECT')
            flag_incorrect = True
            break

    if flag_incorrect:
        continue

    # проверка на телёнок оказался родителем себя
    for i in range(0, len(passports), 3):
        if passports[i] == passports[i+1] or passports[i] == passports[i+2]:
            answers.append('INCORRECT')
            flag_incorrect = True
            break
    
    if flag_incorrect:
        continue

    parents_children = dict()
    for i in range(0, len(passports), 3):
        if passports[i] in parents_children and (passports[i+1] in parents_children[passports[i]] or passports[i+2] in parents_children[passports[i]]):
            answers.append('INCORRECT')
            flag_incorrect = True
            break
        else:
            if passports[i+1] not in parents_children:
                parents_children[passports[i+1]] = set()
                parents_children[passports[i+1]].add(passports[i])
            else:
                parents_children[passports[i+1]].add(passports[i])

            if passports[i+2] not in parents_children:
                parents_children[passports[i+2]] = set()
                parents_children[passports[i+2]].add(passports[i])
            else:
                parents_children[passports[i+2]].add(passports[i])

    if flag_incorrect:
        continue
    else:
        answers.append('CORRECT')

for ans in answers:
    print(ans)
