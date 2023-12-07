package explicit_finite_difference

import (
	"lab05/internal/config"
	"lab05/pkg/tools"
	"math"
)

func ExplicitFiniteDifference(c config.Config) [][]float64 {
	tau := c.Sigma * math.Pow(c.H, 2) / c.A
	x := tools.Arange(c.XBegin, c.XEnd, c.H)
	t := tools.Arange(c.TBegin, c.TEnd, tau)

	res := tools.Zeros(len(t), len(x))

	for i := 0; i < len(x); i++ {
		res[0][i] = c.Psi(x[i])
	}

	for i := 1; i < len(t); i++ {
		res[i][0] = c.Phi0(t[i], c.A)
		for j := 1; j < len(x)-1; j++ {
			res[i][j] = c.Sigma*res[i-1][j-1] + (1-2*c.Sigma)*res[i-1][j] + c.Sigma*res[i-1][j+1]
		}
		res[i][len(res[0])-1] = c.Phi1(t[i], c.A)
	}

	return res
}
