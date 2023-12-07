package analytical

import (
	"lab06/internal/config"
	"lab06/pkg/tools"
	"math"
)

func Analytical(c config.Config, solution func(x, t, a float64) float64) [][]float64 {
	tau := math.Sqrt(c.Sigma * math.Pow(c.H, 2))
	x := tools.Arange(c.XBegin, c.XEnd, c.H)
	t := tools.Arange(c.TBegin, c.TEnd, tau)

	res := tools.Zeros(len(t), len(x))
	for i := range x {
		for j := range t {
			res[j][i] = solution(x[i], t[j], c.A)
		}
	}
	return res
}
