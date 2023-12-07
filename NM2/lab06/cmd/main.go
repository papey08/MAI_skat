package main

import (
	"encoding/json"
	"lab06/internal/analytical"
	"lab06/internal/config"
	efd "lab06/internal/explicit_finite_difference"
	ifd "lab06/internal/implicit_finite_difference"
	"math"
	"os"
)

type results struct {
	AnalyticRes [][]float64 `json:"analytic_res"`
	IfdRes      [][]float64 `json:"ifd_res"`
	EfdRes      [][]float64 `json:"efd_res"`
}

func main() {
	var conf = config.Config{
		XBegin: 0,
		XEnd:   math.Pi,

		TBegin: 0,
		TEnd:   5,

		H:     0.01,
		Sigma: 0.5,

		A: 1,

		Phi0: func(t, a float64) float64 {
			return -math.Sin(a * t)
		},
		Phi1: func(t, a float64) float64 {
			return math.Sin(a * t)
		},

		Psi0: func(x, a float64) float64 {
			return math.Sin(x)
		},

		Psi1: func(x, a float64) float64 {
			return -a * math.Cos(x)
		},
	}

	solution := func(x, t, a float64) float64 {
		return math.Sin(x - a*t)
	}

	res := results{
		AnalyticRes: analytical.Analytical(conf, solution),
		IfdRes:      ifd.ImplicitFiniteDifference(conf),
		EfdRes:      efd.ExplicitFiniteDifference(conf),
	}

	resBytes, _ := json.Marshal(res)
	_ = os.WriteFile("results.json", resBytes, 0644)
}
