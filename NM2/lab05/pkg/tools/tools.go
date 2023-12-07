package tools

import "math"

const eps = 1e-9

// compareFloats returns -1 if a < b, 0 if a == b and 1 if a > b
func compareFloats(a, b float64) int {
	if math.Abs(a-b) < eps {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

// Arange is the same function as numpy.arange
func Arange(start float64, stop float64, step float64) []float64 {
	res := make([]float64, 0, int((stop-start)/step))
	x := start
	for compareFloats(x, stop) == -1 {
		res = append(res, x)
		x += step
	}
	return res
}

// Zeros is the same function as numpy.Zeros
func Zeros(lines, cols int) [][]float64 {
	res := make([][]float64, lines)
	for i := range res {
		res[i] = make([]float64, cols)
	}
	return res
}
