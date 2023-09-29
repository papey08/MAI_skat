package training

import (
	"ai_lab3/internal/loss"
	"ai_lab3/internal/network"
	"ai_lab3/internal/training/solver"
	"time"
)

type Trainer struct {
	deltas    [][]float64
	solver    solver.Solver
	logger    *Logger
	verbosity int
}

func NewTrainer(solver solver.Solver, verbosity int) *Trainer {
	return &Trainer{
		solver:    solver,
		logger:    NewLogger(),
		verbosity: verbosity,
	}
}

func newTraining(layers []*network.Layer) [][]float64 {
	deltas := make([][]float64, len(layers))
	for i, l := range layers {
		deltas[i] = make([]float64, len(l.Neurons))
	}
	return deltas
}

func (t *Trainer) calculateDeltas(n *network.Network, ideal []float64) {
	for i, neuron := range n.Layers[len(n.Layers)-1].Neurons {
		_, df := loss.GetLoss(n.Config.Loss)
		t.deltas[len(n.Layers)-1][i] = df(
			neuron.Value,
			ideal[i],
			neuron.DActivate(neuron.Value))
	}

	for i := len(n.Layers) - 2; i >= 0; i-- {
		for j, neuron := range n.Layers[i].Neurons {
			var sum float64
			for k, s := range neuron.Out {
				sum += s.Weight * t.deltas[i+1][k]
			}
			t.deltas[i][j] = neuron.DActivate(neuron.Value) * sum
		}
	}
}

func (t *Trainer) update(n *network.Network, it int) {
	var idx int
	for i, l := range n.Layers {
		for j := range l.Neurons {
			for k := range l.Neurons[j].In {
				update := t.solver.Update(l.Neurons[j].In[k].Weight,
					t.deltas[i][j]*l.Neurons[j].In[k].In,
					it,
					idx)
				l.Neurons[j].In[k].Weight += update
				idx++
			}
		}
	}
}

func (t *Trainer) learn(n *network.Network, e Pair, it int) error {
	err := n.Forward(e.Input)
	if err != nil {
		return err
	}
	t.calculateDeltas(n, e.Response)
	t.update(n, it)
	return nil
}

func (t *Trainer) Train(n *network.Network, examples, validation Pairs, iterations int) error {
	t.deltas = newTraining(n.Layers)
	train := make(Pairs, len(examples))
	copy(train, examples)
	t.logger.Init(n)
	t.solver.Init(n.NumWeights())
	start := time.Now()
	for i := 1; i <= iterations; i++ {
		examples.Shuffle()
		for j := 0; j < len(examples); j++ {
			err := t.learn(n, examples[j], i)
			if err != nil {
				return err
			}
		}
	}
	t.logger.WriteLog(n, validation, time.Since(start), iterations)
	return nil
}
