package activation

import "math"

func ReLUF(x float64) float64 {
	return math.Max(x, 0)
}

func ReLUDf(y float64) float64 {
	if y > 0 {
		return 1
	}
	return 0
}
