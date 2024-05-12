def is_valid(name, family):
    for member in family:
        if name == member:
            return False
        if name.endswith("-" + member):
            return False
    return True

n = int(input())
family = set()
for _ in range(n):
    name = input()
    family.add(name)

q = int(input())
for _ in range(q):
    name = input()
    if is_valid(name, family):
        print("Good")
    else:
        print("Bad")
