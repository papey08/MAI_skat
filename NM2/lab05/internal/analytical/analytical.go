package analytical

import (
	"lab05/internal/config"
	"lab05/pkg/tools"
	"math"
)

func Analytical(c config.Config, solution func(x, t, a float64) float64) [][]float64 {
	tau := c.Sigma * math.Pow(c.H, 2) / c.A
	x := tools.Arange(c.XBegin, c.XEnd, c.H)
	t := tools.Arange(c.TBegin, c.TEnd, tau)

	res := tools.Zeros(len(t), len(x))
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(t); j++ {
			res[j][i] = solution(x[i], t[j], c.A)
		}
	}
	return res
}
