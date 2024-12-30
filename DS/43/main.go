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
	var sumLnX, sumY, sumLnX2, sumLnXY float64

	for i := 0; i < len(x); i++ {
		lnX := math.Log(x[i])
		sumLnX += lnX
		sumY += y[i]
		sumLnX2 += lnX * lnX
		sumLnXY += lnX * y[i]
	}

	n := float64(len(x))
	b1 := (n*sumLnXY - sumLnX*sumY) / (n*sumLnX2 - sumLnX*sumLnX)
	b0 := (sumY - b1*sumLnX) / n

	return b0, b1
}

func main() {
	x, y := getInput(filename)
	b0, b1 := solve(x, y)
	fmt.Println("b0 =", b0)
	fmt.Println("b1 =", b1)
}
