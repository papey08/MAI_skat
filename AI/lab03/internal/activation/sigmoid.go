package activation

import "math"

func SigmoidF(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func SigmoidDf(y float64) float64 {
	return y * (1 - y)
}
