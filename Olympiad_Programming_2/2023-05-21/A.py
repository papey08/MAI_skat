a_str = input()
b_str = input()

a = float(a_str)
b = float(b_str)

total = int(a + b)
min_draws = 1 if '.' in a_str or '.' in b_str else 0
max_draws = int(min(a, b) / 0.5)

print(total, min_draws, max_draws, sep='\n')
