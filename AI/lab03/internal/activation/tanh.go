package activation

import "math"

func TanhF(x float64) float64 {
	return (1 - math.Exp(-2*x)) / (1 + math.Exp(-2*x))
}

func TanhDf(y float64) float64 {
	return 1 - math.Pow(y, 2)
}
