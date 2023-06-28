package loss

import "math"

func CrossEntropyF(estimate, ideal [][]float64) float64 {
	var sum float64
	for i := range estimate {
		ce := 0.0
		for j := range estimate[i] {
			ce += ideal[i][j] * math.Log(estimate[i][j])
		}
		sum -= ce
	}
	return sum / float64(len(estimate))
}

func CrossEntropyDf(estimate, ideal, _ float64) float64 {
	return estimate - ideal
}
