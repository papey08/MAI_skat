package solver

import "math"

type Adam struct {
	lr  float64
	b   float64
	b2  float64
	eps float64

	v, m []float64
}

func NewAdam(lr, beta, beta2, epsilon float64) *Adam {
	return &Adam{
		lr:  fparam(lr, 0.001),
		b:   fparam(beta, 0.9),
		b2:  fparam(beta2, 0.999),
		eps: fparam(epsilon, 1e-8),
	}
}

func (a *Adam) Init(size int) {
	a.v, a.m = make([]float64, size), make([]float64, size)
}

func (a *Adam) Update(_, gradient float64, t, idx int) float64 {
	lrt := a.lr * (math.Sqrt(1.0 - math.Pow(a.b2, float64(t)))) /
		(1.0 - math.Pow(a.b, float64(t)))
	a.m[idx] = a.b*a.m[idx] + (1.0-a.b)*gradient
	a.v[idx] = a.b2*a.v[idx] + (1.0-a.b2)*math.Pow(gradient, 2.0)

	return -lrt * (a.m[idx] / (math.Sqrt(a.v[idx]) + a.eps))
}
