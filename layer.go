package main

import (
	"fmt"
	"math/rand"
)

type layer struct {
	// indexed as weights[j, i] where
	// 		i - 1 = index of node in prev layer (incoming),
	//		j = index of node in current layer
	// bias is always the 0th weight for each node, e.g. weights[4, 0] will be the bias weight for node 4 in the current layer.
	weights [][]float64
	activationFn func(float64) float64
	activationFnDiff func(float64) float64
}

func CreateLayer(size int, inputSize int, activationFn func(float64) float64, activationFnDiff func(float64) float64) layer {
	l := layer{
		weights: [][]float64{},
		activationFn: activationFn,
		activationFnDiff: activationFnDiff,
	}

	l.AddMany(size, inputSize)
	return l
}

// Add a fixed number of perceptrons to this layer.
func (l *layer) AddMany(count int, inputSize int) {
	for range count {
		l.Add(inputSize)
	}
}

// Add a perceptron to this layer.
func (l *layer) Add(inputSize int) {

	node_weights := []float64{rand.Float64() * 2 - 1} // add bias weight

	for i := 0; i < inputSize; i++ {
		node_weights = append(node_weights, rand.Float64() * 2 - 1) // [-1, 1]
	}

	l.weights = append(l.weights, node_weights)
}


/*
*

	Given a set of weights, forward the inputs through a single perceptron by computing a weighted sum and then passing that through an activation function.
	- NOTE: The length of the weights array should be N + 1, where N is the size of the previous layer.
	- weights[0] should correspond to the bias weight. This is always multiplied by 1.

*
*/
func ForwardNode(inputs []float64, weights []float64, activationFn func(float64) float64) (output float64, activation float64) {
	assert(len(weights) == len(inputs)+1, "Expected weight array of length "+string(len(inputs)+1)+" (N+1) but got length "+string(len(weights)))
	activation = 0.0

	activation += weights[0]

	for idx, input := range inputs {
		activation += weights[idx+1] * input
	}

	return activationFn(activation), activation
}


// Feed forward through the layer.
// Also returns the activations (weighted sum plus bias) for use in the backpropagation algorithm.
func (l layer) Forward(inputs []float64) (outputs []float64, activations []float64) {
	for _, w := range l.weights {
		o, a := ForwardNode(inputs, w, l.activationFn) 
		outputs = append(outputs, o)
		activations = append(activations, a)
	}
	return outputs, activations
}

// Print out the layer's weights.
func (l layer) Print() {
	for _, weights := range l.weights {
		fmt.Print("\t[")
		for i, w := range weights {
			if i == 0 {
				fmt.Printf("b: %.2f | ", w) // bias
			} else {
				fmt.Printf("%.2f ", w)
			}
		}
		fmt.Print("]\n")
	}
}