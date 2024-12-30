def format_matrix(matrix):
    return '(' + '; '.join(' '.join(map(str, row)) for row in matrix) + ')'

def format_vector(vector):
    return '(' + ' '.join(map(str, vector[0])) + ')'

def write_records_to_file(records, filename='output.txt'):
    with open(filename, 'w') as file:
        for i, record in enumerate(records):
            if 'det' in record:
                file.write(f"det={record['det']}\n")
            if 'inversed' in record:
                file.write(f"inversed={format_matrix(record['inversed'])}\n")
            if 'rank' in record:
                file.write(f"rank={record['rank']}\n")
            if 'solution' in record:
                file.write(f"solution={format_vector(record['solution'])}\n")
            if i < len(records) - 1:
                file.write('%\n')
