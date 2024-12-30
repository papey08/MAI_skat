def format_matrix(matrix):
    return '(' + '; '.join(' '.join(map(str, row)) for row in matrix) + ')'

def write_records_to_file(records, filename='output.txt'):
    with open(filename, 'w') as file:
        for i, record in enumerate(records):
            if 'integer_ker' in record:
                file.write(f"integer_ker={format_matrix(record['integer_ker'])}\n")
            if 'normalized_ker' in record:
                file.write(f"normalized_ker={format_matrix(record['normalized_ker'])}\n")
            if i < len(records) - 1:
                file.write('%\n')
