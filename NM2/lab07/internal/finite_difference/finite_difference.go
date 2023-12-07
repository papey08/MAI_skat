package finite_difference

import (
	"lab07/internal/config"
	"lab07/internal/model"
	"lab07/pkg/tools"
	"math"
)

func FiniteDifference(c config.Config) model.ResultData {
	x := tools.Arange(c.XBegin, c.XEnd, c.HX)
	y := tools.Arange(c.YBegin, c.YEnd, c.HY)
	res := tools.Zeros(len(x), len(y))

	for i := range x {
		res[i][0] = c.Psi0(x[i])
		res[i][len(res[i])-1] = c.Psi1(x[i])
	}
	for i := range y {
		res[0][i] = c.Phi0(y[i])
		res[len(res)-1][i] = c.Phi1(y[i])
	}

	mapping := make([][]int, len(x))
	for i := range mapping {
		mapping[i] = make([]int, len(y))
	}
	curId := 0
	for i := 1; i < len(x)-1; i++ {
		for j := 1; j < len(y)-1; j++ {
			mapping[i][j] = curId
			curId++
		}
	}

	eqNums := (len(x) - 2) * (len(y) - 2)
	if eqNums < 0 {
		panic("number of equations less than a zero")
	}

	A := tools.Zeros(eqNums, eqNums)
	b := make([]float64, eqNums)

	for i := 1; i < len(x)-1; i++ {
		for j := 1; j < len(y)-1; j++ {
			curId := mapping[i][j]
			A[curId][curId] = 1

			if j-1 == 0 {
				b[curId] += c.Psi0(x[i]) * math.Pow(c.HX, 2) / (2 * (math.Pow(c.HX, 2) + math.Pow(c.HY, 2)))
			} else {
				A[curId][mapping[i][j-1]] = -math.Pow(c.HX, 2) / (2 * (math.Pow(c.HX, 2) + math.Pow(c.HY, 2)))
			}

			if j+1 == len(y)-1 {
				b[curId] += c.Psi1(x[i]) * math.Pow(c.HX, 2) / (2 * (math.Pow(c.HX, 2) + math.Pow(c.HY, 2)))
			} else {
				A[curId][mapping[i][j+1]] = -math.Pow(c.HX, 2) / (2 * (math.Pow(c.HX, 2) + math.Pow(c.HY, 2)))
			}

			if i-1 == 0 {
				b[curId] += c.Phi0((y[j])) * math.Pow(c.HY, 2) / (2 * (math.Pow(c.HX, 2) + math.Pow(c.HY, 2)))
			} else {
				A[curId][mapping[i-1][j]] = -math.Pow(c.HY, 2) / (2 * (math.Pow(c.HX, 2) + math.Pow(c.HY, 2)))
			}

			if i+1 == len(x)-1 {
				b[curId] += c.Phi1(y[j]) * math.Pow(c.HY, 2) / (2 * (math.Pow(c.HX, 2) + math.Pow(c.HY, 2)))
			} else {
				A[curId][mapping[i+1][j]] = -math.Pow(c.HY, 2) / (2 * (math.Pow(c.HX, 2) + math.Pow(c.HY, 2)))
			}
		}
	}

	return model.ResultData{
		A:       A,
		B:       b,
		Mapping: mapping,
		Res:     res,
		X:       x,
		Y:       y,
	}
}
