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

func solve(xArr []float64, yArr []float64) (float64, float64, float64) {
	var sumX, sumY, sumSin8X, sumX2, sumSin8X2, sumXY, sumYSin8X, sumXSin8X float64

	for i := 0; i < len(xArr); i++ {
		x := xArr[i]
		y := yArr[i]
		sin8X := math.Sin(8 * x)

		sumX += x
		sumY += y
		sumSin8X += sin8X
		sumX2 += x * x
		sumSin8X2 += sin8X * sin8X
		sumXY += x * y
		sumYSin8X += y * sin8X
		sumXSin8X += x * sin8X
	}

	n := float64(len(xArr))
	denominator := n*sumX2*sumSin8X2 + 2*sumX*sumSin8X*sumXSin8X - sumX2*sumSin8X*sumSin8X - n*sumXSin8X*sumXSin8X - sumX*sumX*sumSin8X2
	b0 := (sumY*sumX2*sumSin8X2 + sumX*sumSin8X*sumYSin8X + sumSin8X*sumXSin8X*sumY - sumY*sumXSin8X*sumXSin8X - sumX*sumXSin8X*sumYSin8X - sumSin8X2*sumX*sumY) / denominator
	b1 := (n*sumXY*sumSin8X2 + sumY*sumSin8X*sumXSin8X + sumSin8X*sumYSin8X*sumX - sumY*sumX2*sumSin8X2 - n*sumXSin8X*sumYSin8X - sumSin8X*sumSin8X*sumX) / denominator
	b2 := (n*sumX2*sumYSin8X + sumY*sumX*sumXSin8X + sumY*sumSin8X*sumXY - sumX*sumXSin8X*sumY - sumXY*sumSin8X*sumSin8X - n*sumY*sumXSin8X) / denominator

	return b0, b1, b2
}

func main() {
	x, y := getInput(filename)
	b0, b1, b2 := solve(x, y)
	fmt.Println("b0 =", b0)
	fmt.Println("b1 =", b1)
	fmt.Println("b2 =", b2)
}
