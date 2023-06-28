package network

import (
	"ai_lab3/internal/activation"
	"ai_lab3/internal/util"
	"ai_lab3/internal/weights"
)

type Layer struct {
	Neurons []*Neuron
	A       activation.Activation
}

func NewLayer(n int, a activation.Activation) *Layer {
	neurons := make([]*Neuron, n)
	for i := 0; i < n; i++ {
		act := a
		if a == activation.Softmax {
			act = activation.Linear
		}
		neurons[i] = NewNeuron(act)
	}
	return &Layer{
		Neurons: neurons,
		A:       a,
	}
}

func (l *Layer) fire() {
	for _, n := range l.Neurons {
		n.fire()
	}
	if l.A == activation.Softmax {
		outs := make([]float64, len(l.Neurons))
		for i, neuron := range l.Neurons {
			outs[i] = neuron.Value
		}
		sm := util.Softmax(outs)
		for i, neuron := range l.Neurons {
			neuron.Value = sm[i]
		}
	}
}

func (l *Layer) Connect(next *Layer, weight weights.WeightInitializer) {
	for i := range l.Neurons {
		for j := range next.Neurons {
			syn := NewSynapse(weight())
			l.Neurons[i].Out = append(l.Neurons[i].Out, syn)
			next.Neurons[j].In = append(next.Neurons[j].In, syn)
		}
	}
}

func (l *Layer) ApplyBias(weight weights.WeightInitializer) []*Synapse {
	biases := make([]*Synapse, len(l.Neurons))
	for i := range l.Neurons {
		biases[i] = NewSynapse(weight())
		biases[i].IsBias = true
		l.Neurons[i].In = append(l.Neurons[i].In, biases[i])
	}
	return biases
}
