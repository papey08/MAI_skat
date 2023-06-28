package solver

type Solver interface {
	Init(size int)
	Update(value, gradient float64, iteration, idx int) float64
}

func fparam(val, fallback float64) float64 {
	if val == 0.0 {
		return fallback
	}
	return val
}
