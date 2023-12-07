package explicit_finite_difference

import (
	"lab06/internal/config"
	"lab06/pkg/tools"
	"math"
)

func ExplicitFiniteDifference(c config.Config) [][]float64 {
	tau := math.Sqrt(c.Sigma * math.Pow(c.H, 2))
	x := tools.Arange(c.XBegin, c.XEnd, c.H)
	t := tools.Arange(c.TBegin, c.TEnd, tau)

	res := tools.Zeros(len(t), len(x))
	for i := range x {
		res[0][i] = c.Psi0(x[i], c.A)
	}

	for i := range x {
		res[1][i] = c.Psi0(x[i], c.A) + tau*c.Psi1(x[i], c.A)
	}

	for i := 2; i < len(t); i++ {
		res[i][0] = c.Phi0(t[i], c.A)
		for j := 1; j < len(x)-1; j++ {
			res[i][j] = c.Sigma*(res[i-1][j+1]-2*res[i-1][j]+res[i-1][j-1]) + (2-3*math.Pow(tau, 2))*res[i-1][j] - res[i-2][j]
		}
		res[i][len(res[i])-1] = c.Phi1(t[i], c.A)
	}
	return res
}
