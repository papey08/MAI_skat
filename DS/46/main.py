import app.matrix as matrix
import app.input as input
import app.output as output

data = input.read_from_file()

res = []
for d in data:
    if 'vector' in d and d['vector'] != []:
        res.append({'solution' : matrix.Matrix(d['matrix']).solve_system(d['vector'])})
    else:
        m = matrix.Matrix(d['matrix'])
        res.append({
            'det': m.det(),
            'inversed': m.inversed(),
            'rank': m.rank()
        })

output.write_records_to_file(res)
