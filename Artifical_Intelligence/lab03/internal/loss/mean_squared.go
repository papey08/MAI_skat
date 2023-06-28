package loss

import "math"

func MeanSquaredF(estimate, ideal [][]float64) float64 {
	var sum float64
	for i := 0; i < len(estimate); i++ {
		for j := 0; j < len(estimate[i]); j++ {
			sum += math.Pow(estimate[i][j]-ideal[i][j], 2)
		}
	}
	return sum / float64(len(estimate)*len(estimate[0]))
}

func MeanSquaredDf(estimate, ideal, activation float64) float64 {
	return activation * (estimate - ideal)
}
