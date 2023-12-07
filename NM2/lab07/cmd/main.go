package main

import (
	"encoding/json"
	"lab07/internal/config"
	fd "lab07/internal/finite_difference"
	"os"
)

func main() {
	var conf = config.Config{
		XBegin: 0,
		XEnd:   1.05,
		YBegin: 0,
		YEnd:   1.05,
		HX:     0.05,
		HY:     0.05,
		Phi0: func(y float64) float64 {
			return y
		},
		Phi1: func(y float64) float64 {
			return 1. + y
		},
		Psi0: func(x float64) float64 {
			return x
		},
		Psi1: func(x float64) float64 {
			return 1. + x
		},
	}

	resData := fd.FiniteDifference(conf)

	resBytes, _ := json.Marshal(resData)
	_ = os.WriteFile("results.json", resBytes, 0644)
}
