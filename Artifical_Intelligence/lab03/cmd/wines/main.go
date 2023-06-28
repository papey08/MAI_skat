package main

import (
	"ai_lab3/internal/activation"
	"ai_lab3/internal/dataset"
	"ai_lab3/internal/network"
	"ai_lab3/internal/training"
	"ai_lab3/internal/training/solver"
	"ai_lab3/internal/util"
	"ai_lab3/internal/weights"
)

func main() {
	data, err := dataset.Load("./cmd/wines/wine.data")
	if err != nil {
		panic(err)
	}

	for i := range data {
		util.Standardize(data[i].Input)
	}
	data.Shuffle()

	n := network.NewNetwork(&network.Params{
		Inputs:       len(data[0].Input),
		LayoutConfig: []int{6, 3}, // network with 1 hidden layer of 6 nodes and an output layer of 3 nodes
		Activation:   activation.Tanh,
		Mode:         activation.MultiClass,
		Weight:       weights.NewNormal(1, 0),
		Bias:         true,
	})

	trainer := training.NewTrainer(solver.NewSGD(0.005, 0.5, 0.001), 200)
	_ = trainer.Train(n, data, data, 2000)
}
