package implicit_finite_difference

import (
	"lab06/internal/config"
	"lab06/pkg/tools"
	"math"
)

func ImplicitFiniteDifference(c config.Config) [][]float64 {
	tau := math.Sqrt(c.Sigma * math.Pow(c.H, 2))
	x := tools.Arange(c.XBegin, c.XEnd, c.H)
	t := tools.Arange(c.TBegin, c.TEnd, tau)
	res := tools.Zeros(len(t), len(x))

	for i := range x {
		res[0][i] = c.Psi0(x[i], c.A)
		res[1][i] = c.Psi0(x[i], c.A) + tau*c.Psi1(x[i], c.A)
	}

	for i := 2; i < len(t); i++ {
		A := tools.Zeros(len(x)-2, len(x)-2)

		A[0][0] = -(1 + 2*c.Sigma + 3*math.Pow(tau, 2))
		A[0][1] = c.Sigma
		for i := 1; i < len(A)-1; i++ {
			A[i][i-1] = c.Sigma
			A[i][i] = -(1 + 2*c.Sigma + 3*math.Pow(tau, 2))
			A[i][i+1] = c.Sigma
		}
		A[len(A)-1][len(A)-2] = c.Sigma
		A[len(A)-1][len(A)-1] = -(1 + 2*c.Sigma + 3*math.Pow(tau, 2))

		b := make([]float64, len(res[0])-2)
		for j := range b {
			b[j] = -2*res[i-1][j+1] + res[i-2][j+1]
		}
		b[0] -= c.Sigma * c.Phi0(t[i], c.A)
		b[len(b)-1] -= c.Sigma * c.Phi1(t[i], c.A)

		res[i][0] = c.Phi0(t[i], c.A)
		res[i][len(res[i])-1] = c.Phi1(t[i], c.A)
		for j, v := range tools.TridiagSolve(A, b) {
			res[i][j+1] = v
		}
	}
	return res
}
