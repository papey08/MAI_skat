package loss

import "math"

func BinCrossEntropyF(estimate, ideal [][]float64) float64 {
	epsilon := math.SmallestNonzeroFloat64
	var sum float64
	for i := range estimate {
		ce := 0.0
		for j := range estimate[i] {
			ce += ideal[i][j]*math.Log(estimate[i][j]+epsilon) + (1.0-ideal[i][j])*math.Log(1.0-estimate[i][j]+epsilon)
		}
		sum -= ce
	}
	return sum / float64(len(estimate))
}

func BinCrossEntropyDf(estimate, ideal, _ float64) float64 {
	return estimate - ideal
}
