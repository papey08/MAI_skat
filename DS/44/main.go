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

func solve(x []float64, y []float64) (float64, float64) {
	var sumExpX, sumY, sumExpX2, sumExpXY float64

	for i := 0; i < len(x); i++ {
		expX := math.Exp(0.1 * x[i])
		sumExpX += expX
		sumY += y[i]
		sumExpX2 += expX * expX
		sumExpXY += expX * y[i]
	}

	n := float64(len(x))
	b1 := (n*sumExpXY - sumExpX*sumY) / (n*sumExpX2 - sumExpX*sumExpX)
	b0 := (sumY - b1*sumExpX) / n

	return b0, b1
}

func main() {
	x, y := getInput(filename)
	b0, b1 := solve(x, y)
	fmt.Println("b0 =", b0)
	fmt.Println("b1 =", b1)
}
