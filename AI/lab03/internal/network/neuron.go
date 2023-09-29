package network

import "ai_lab3/internal/activation"

type Neuron struct {
	A     activation.Activation
	In    []*Synapse
	Out   []*Synapse
	Value float64
}

func NewNeuron(activation activation.Activation) *Neuron {
	return &Neuron{
		A: activation,
	}
}

func (n *Neuron) fire() {
	var sum float64
	for _, s := range n.In {
		sum += s.Out
	}
	n.Value = n.Activate(sum)
	nVal := n.Value
	for _, s := range n.Out {
		s.fire(nVal)
	}
}

func (n *Neuron) Activate(x float64) float64 {
	f, _ := activation.GetActivation(n.A)
	return f(x)
}

func (n *Neuron) DActivate(x float64) float64 {
	_, df := activation.GetActivation(n.A)
	return df(x)
}
