package network

type Synapse struct {
	Weight  float64
	In, Out float64
	IsBias  bool
}

func NewSynapse(weight float64) *Synapse {
	return &Synapse{Weight: weight}
}

func (s *Synapse) fire(value float64) {
	s.In = value
	s.Out = s.In * s.Weight
}
