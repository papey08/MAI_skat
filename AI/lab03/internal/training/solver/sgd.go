package solver

type SGD struct {
	lr       float64
	decay    float64
	momentum float64
	moments  []float64
}

func NewSGD(lr, momentum, decay float64) *SGD {
	return &SGD{
		lr:       fparam(lr, 0.01),
		decay:    decay,
		momentum: momentum,
	}
}

func (o *SGD) Init(size int) {
	o.moments = make([]float64, size)
}

func (o *SGD) Update(_, gradient float64, iteration, idx int) float64 {
	lr := o.lr / (1 + o.decay*float64(iteration))
	o.moments[idx] = o.momentum*o.moments[idx] - lr*gradient
	return o.moments[idx]
}
