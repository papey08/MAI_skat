package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

const filename = "./input.json"

type (
	matrix [][]float64
	vector []float64
	input  struct {
		X vector `json:"x"`
		Y vector `json:"y"`
		M int    `json:"m"`
	}
)

func getInput(filename string) ([]float64, []float64, int) {
	var inp input
	file, _ := os.ReadFile(filename)
	_ = json.Unmarshal(file, &inp)
	return inp.X, inp.Y, inp.M
}

func createVandermondeMatrix(x vector, degree int) matrix {
	n := len(x)
	vandermonde := make(matrix, n)
	for i := range vandermonde {
		vandermonde[i] = make(vector, degree+1)
		for j := 0; j <= degree; j++ {
			vandermonde[i][j] = math.Pow(x[i], float64(j))
		}
	}
	return vandermonde
}

func multiplyMatrixVector(a matrix, b vector) vector {
	result := make(vector, len(a))
	for i := range a {
		for j := range a[i] {
			result[i] += a[i][j] * b[j]
		}
	}
	return result
}

func transposeMatrix(a matrix) matrix {
	m, n := len(a), len(a[0])
	At := make(matrix, n)
	for i := range At {
		At[i] = make(vector, m)
		for j := range a {
			At[i][j] = a[j][i]
		}
	}
	return At
}

func solveGaussian(a matrix, b vector) vector {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ratio := a[j][i] / a[i][i]
			for k := i; k < n; k++ {
				a[j][k] -= ratio * a[i][k]
			}
			b[j] -= ratio * b[i]
		}
	}
	x := make(vector, n)
	for i := n - 1; i >= 0; i-- {
		x[i] = b[i]
		for j := i + 1; j < n; j++ {
			x[i] -= a[i][j] * x[j]
		}
		x[i] /= a[i][i]
	}
	return x
}

func solve(x, Y vector, degree int) vector {
	vandermonde := createVandermondeMatrix(x, degree)
	transposedVandermonde := transposeMatrix(vandermonde)
	a := make(matrix, degree+1)
	for i := range a {
		a[i] = make(vector, degree+1)
		for j := range a[i] {
			for k := range transposedVandermonde[i] {
				a[i][j] += transposedVandermonde[i][k] * vandermonde[k][j]
			}
		}
	}
	b := multiplyMatrixVector(transposedVandermonde, Y)
	return solveGaussian(a, b)
}

func main() {
	x, y, m := getInput(filename)

	coefficients := solve(x, y, m)

	for i, coef := range coefficients {
		fmt.Printf("β%d = %.4f\n", i, coef)
	}
}
