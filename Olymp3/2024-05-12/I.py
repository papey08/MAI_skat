def is_graph(tree, n, m):
    for row in tree:
        if "graph" in row:
            return "YES"
    for col in range(m):
        column = ""
        for row in range(n):
            column += tree[row][col]
        if "graph" in column:
            return "YES"

    return "NO"

n, m = map(int, input().split())
tree = []
for _ in range(n):
    row = input()
    tree.append(row)

print(is_graph(tree, n, m))
