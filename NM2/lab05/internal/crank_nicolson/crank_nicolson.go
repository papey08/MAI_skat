package crank_nicolson

import (
	"lab05/internal/config"
	"lab05/pkg/tools"
	"math"
)

func CrankNicolson(c config.Config, theta float64) [][]float64 {
	tau := c.Sigma * math.Pow(c.H, 2) / c.A
	x := tools.Arange(c.XBegin, c.XEnd, c.H)
	t := tools.Arange(c.TBegin, c.TEnd, tau)

	res := tools.Zeros(len(t), len(x))

	for i := 0; i < len(x); i++ {
		res[0][i] = c.Psi(x[i])
	}

	for i := 1; i < len(t); i++ {
		a := tools.Zeros(len(x)-2, len(x)-2)
		a[0][0] = -1 - 2*c.Sigma*theta
		a[0][1] = c.Sigma * theta

		for j := 1; j < len(a)-1; j++ {
			a[j][j-1] = c.Sigma * theta
			a[j][j] = -1 - 2*c.Sigma*theta
			a[j][j+1] = c.Sigma * theta
		}

		a[len(a)-1][len(a[0])-2] = c.Sigma * theta
		a[len(a)-1][len(a[0])-1] = -1 - 2*c.Sigma*theta

		b := make([]float64, len(a))
		for j := 1; j < len(b)-2; j++ {
			b[j] = res[i-1][j] + (1-theta)*c.Sigma*(res[i-1][j-1]-2*res[i-1][j]+res[i-1][j+1])
		}
		b[0] -= c.Sigma * theta * c.Phi0(t[i], c.A)
		b[len(b)-1] -= c.Sigma * theta * c.Phi1(t[i], c.A)

		res[i][0] = c.Phi0(t[i], c.A)
		res[i][len(res[i])-1] = c.Phi1(t[i], c.A)
		sol := tools.TridiagSolve(a, b)
		for j := 1; j < len(res[i])-1; j++ {
			res[i][j] = sol[j-1]
		}
	}
	return res
}
