_ = input()

sentence = input().split()

corrects = dict()
incorrects = dict()

for _ in range(int(input())):
    word = input().split()
    if word[2] == 'correct':
        if word[0] in corrects:
            corrects[word[0]].append(word[1])
        else:
            corrects[word[0]] = [word[1]]
    else:
        if word[0] in incorrects:
            incorrects[word[0]].append(word[1])
        else:
            incorrects[word[0]] = [word[1]]

correct_combinations = 1

for word in sentence:
    correct_translates = 0 if word not in corrects else len(corrects[word])
    correct_combinations *= correct_translates

if correct_combinations == 0:
    all_combinations = 1
    for word in sentence:
        incorrect_translates = 0 if word not in incorrects else len(incorrects[word])
        correct_translates = 0 if word not in corrects else len(corrects[word])
        all_combinations *= (correct_translates + incorrect_translates)

    if all_combinations == 1:
        translate = []
        for word in sentence:
            if word in corrects:
                translate.append(corrects[word][0])
            else:
                translate.append(incorrects[word][0])
        print(' '.join(translate))
        print('incorrect')
    else:
        print('0 correct')
        print(all_combinations, 'incorrect')
elif correct_combinations == 1:
    translate = []
    for word in sentence:
        translate.append(corrects[word][0])
    print(' '.join(translate))
    print('correct')
else:
    all_combinations = 1
    for word in sentence:
        incorrect_translates = 0 if word not in incorrects else len(incorrects[word])
        all_combinations *= (len(corrects[word]) + incorrect_translates)
    print(correct_combinations, 'correct')
    print(all_combinations - correct_combinations, 'incorrect')
