import app.ker as ker
import app.input as input
import app.output as output

res = []
data = input.read_from_file()
for d in data:
    integer_ker, normalized_ker = ker.solve(d['matrix'])
    res.append(
        {
            'integer_ker': integer_ker,
            'normalized_ker': normalized_ker
        }
    )
output.write_records_to_file(res)
