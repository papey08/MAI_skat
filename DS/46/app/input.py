def parse_matrix(line):
    return [list(map(int, row.split())) for row in line.strip("() ").split(";")]

def parse_vector(line):
    return list(map(int, line.strip("() ").split()))

def read_from_file(filename='input.txt'):
    with open(filename, 'r') as file:
        data = file.read().split('%')
    
    results = []
    for record in data:
        record = record.strip()
        if not record:
            continue
        matrices = {'matrix': [], 'vector': []}
        for line in record.splitlines():
            line = line.strip().split("=", 1)[-1].strip()
            if ";" in line:
                matrices['matrix'].extend(parse_matrix(line))
            else:
                matrices['vector'].extend(parse_vector(line))
        results.append(matrices)
    
    return results
