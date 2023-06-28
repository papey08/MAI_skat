package training

import (
	"math"
	"math/rand"
)

type Pair struct {
	Input    []float64
	Response []float64
}

type Pairs []Pair

func (p Pairs) Shuffle() {
	for i := range p {
		j := rand.Intn(i + 1)
		p[i], p[j] = p[j], p[i]
	}
}

func (p Pairs) Split(n float64) (first, second Pairs) {
	for i := range p {
		if n > rand.Float64() {
			first = append(first, p[i])
		} else {
			second = append(second, p[i])
		}
	}
	return
}

func (p Pairs) SplitSize(size int) []Pairs {
	res := make([]Pairs, 0, len(p))
	for i := 0; i < len(p); i += size {
		res = append(res, p[i:int(math.Min(float64(i+size), float64(len(p))))])
	}
	return res
}

func (p Pairs) SplitN(n int) []Pairs {
	res := make([]Pairs, n)
	for i, el := range p {
		res[i%n] = append(res[i%n], el)
	}
	return res
}
