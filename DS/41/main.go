package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

const filename = "./input.json"

type input struct {
	X []float64 `json:"x"`
	Y []float64 `json:"y"`
}

func getInput(filename string) ([]float64, []float64) {
	var inp input
	file, _ := os.ReadFile(filename)
	_ = json.Unmarshal(file, &inp)
	return inp.X, inp.Y
}

func solve(x []float64, y []float64) float64 {
	var sumX, sumY, sumXY, sumX2 float64

	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += math.Pow(sumX2, 2)
	}

	n := float64(len(x))
	c1 := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	c2 := (sumY - c1*sumX) / n

	return math.Abs(c1) + math.Abs(c2)
}

func main() {
	x, y := getInput(filename)
	fmt.Println(solve(x, y))
}
