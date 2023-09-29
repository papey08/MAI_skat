package network

import (
	"ai_lab3/internal/activation"
	"ai_lab3/internal/loss"
	"ai_lab3/internal/weights"
	"fmt"
)

type Network struct {
	Layers []*Layer
	Biases [][]*Synapse
	Config *Params
}

type Params struct {
	Inputs       int
	LayoutConfig []int
	Activation   activation.Activation
	Mode         activation.Mode
	Weight       weights.WeightInitializer `json:"-"`
	Loss         loss.Loss
	Bias         bool
}

func NewNetwork(c *Params) *Network {

	if c.Weight == nil {
		c.Weight = weights.NewUniform(0.5, 0)
	}
	if c.Activation == activation.None {
		c.Activation = activation.Sigmoid
	}
	if c.Loss == loss.None {
		switch c.Mode {
		case activation.MultiClass, activation.MultiLabel:
			c.Loss = loss.CrossEntropy
		case activation.Binary:
			c.Loss = loss.BinCrossEntropy
		default:
			c.Loss = loss.MeanSquared
		}
	}

	layers := initializeLayers(c)

	var biases [][]*Synapse
	if c.Bias {
		biases = make([][]*Synapse, len(layers))
		for i := 0; i < len(layers); i++ {
			if c.Mode == activation.Regression && i == len(layers)-1 {
				continue
			}
			biases[i] = layers[i].ApplyBias(c.Weight)
		}
	}

	return &Network{
		Layers: layers,
		Biases: biases,
		Config: c,
	}
}

func initializeLayers(c *Params) []*Layer {
	layers := make([]*Layer, len(c.LayoutConfig))
	for i := range layers {
		act := c.Activation
		if i == (len(layers)-1) && c.Mode != activation.Default {
			act = activation.OutputActivation(c.Mode)
		}
		layers[i] = NewLayer(c.LayoutConfig[i], act)
	}

	for i := 0; i < len(layers)-1; i++ {
		layers[i].Connect(layers[i+1], c.Weight)
	}

	for _, neuron := range layers[0].Neurons {
		neuron.In = make([]*Synapse, c.Inputs)
		for i := range neuron.In {
			neuron.In[i] = NewSynapse(c.Weight())
		}
	}

	return layers
}

func (n *Network) fire() {
	for _, b := range n.Biases {
		for _, s := range b {
			s.fire(1)
		}
	}
	for _, l := range n.Layers {
		l.fire()
	}
}

func (n *Network) Forward(input []float64) error {
	if len(input) != n.Config.Inputs {
		return fmt.Errorf("invalid input dimension - expected: %d got: %d", n.Config.Inputs, len(input))
	}
	for _, n := range n.Layers[0].Neurons {
		for i := 0; i < len(input); i++ {
			n.In[i].fire(input[i])
		}
	}
	n.fire()
	return nil
}

func (n *Network) Predict(input []float64) []float64 {
	_ = n.Forward(input)

	outLayer := n.Layers[len(n.Layers)-1]
	out := make([]float64, len(outLayer.Neurons))
	for i, neuron := range outLayer.Neurons {
		out[i] = neuron.Value
	}
	return out
}

func (n *Network) NumWeights() (num int) {
	for _, l := range n.Layers {
		for _, n := range l.Neurons {
			num += len(n.In)
		}
	}
	return
}
