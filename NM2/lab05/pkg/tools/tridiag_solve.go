package tools

// TridiagSolve solves Ax = b and returns x
func TridiagSolve(A [][]float64, b []float64) []float64 {
	n := len(A)
	v := make([]float64, n)
	u := make([]float64, n)
	v[0] = A[0][1] / -A[0][0]
	u[0] = b[0] / A[0][0]

	for i := 1; i < n-1; i++ {
		v[i] = A[i][i+1] / (-A[i][i] - A[i][i-1]*v[i-1])
		u[i] = (A[i][i-1]*u[i-1] - b[i]) / (-A[i][i] - A[i][i-1]*v[i-1])
	}
	v[n-1] = 0
	u[n-1] = (A[n-1][n-2]*u[n-2] - b[n-1]) / (-A[n-1][n-1] - A[n-1][n-2]*v[n-2])

	res := make([]float64, n)
	res[n-1] = u[n-1]
	for i := n - 1; i > 0; i-- {
		res[i-1] = v[i-1]*res[i] + u[i-1]
	}
	return res
}
