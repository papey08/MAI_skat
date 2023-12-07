package main

import (
	"encoding/json"
	"lab05/internal/analytical"
	"lab05/internal/config"
	cn "lab05/internal/crank_nicolson"
	efd "lab05/internal/explicit_finite_difference"
	ifd "lab05/internal/implicit_finite_difference"
	"math"
	"os"
)

type results struct {
	AnalyticRes [][]float64 `json:"analytic_res"`
	IfdRes      [][]float64 `json:"ifd_res"`
	EfdRes      [][]float64 `json:"efd_res"`
	CnRes       [][]float64 `json:"cn_res"`
}

func main() {
	var conf = config.Config{
		XBegin: 0,
		XEnd:   1,

		TBegin: 0,
		TEnd:   5,

		H:     0.01,
		Sigma: 0.45,
		A:     1,

		Phi0: func(t, a float64) float64 {
			return 0
		},
		Phi1: func(t, a float64) float64 {
			return 0
		},

		Psi: func(x float64) float64 {
			return math.Sin(2 * math.Pi * x)
		},
	}
	solution := func(x, t, a float64) float64 {
		return math.Exp(-4*math.Pow(math.Pi, 2)*a*t) * math.Sin(2*math.Pi*x)
	}

	res := results{
		AnalyticRes: analytical.Analytical(conf, solution),
		IfdRes:      ifd.ImplicitFiniteDifference(conf),
		EfdRes:      efd.ExplicitFiniteDifference(conf),
		CnRes:       cn.CrankNicolson(conf, 0.5),
	}

	resBytes, _ := json.Marshal(res)
	_ = os.WriteFile("results.json", resBytes, 0644)
}
