package main

import "fmt"

type network []layer

func (n *network) AddLayer(size int, inputSize int, activationFn func(float64) float64, activationFnDiff func(float64) float64) {
	*n = append(*n, CreateLayer(size, inputSize, activationFn, activationFnDiff))
}

/*
*

Feed an input through the network.

Returns:
- output: The final output of the network.
- layer_outputs: The intermediate outputs of each layer of the network. layer_outputs[i,j] is the output of the jth node of the ith layer. Used for backpropagation.
- layer_activations: The intermediate activations (weighted sum + bias) of each layer. Used for backpropagation.
*
*/
func (n network) Forward(input []float64) (output []float64, layer_outputs [][]float64, layer_activations [][]float64) {
	for i, l := range n {
		out, activation := l.Forward(input)
		layer_outputs = append(layer_outputs, out)
		layer_activations = append(layer_activations, activation)
		input = layer_outputs[i]
	}
	return layer_outputs[len(layer_outputs)-1], layer_outputs, layer_activations
}

func (n network) Print() {
	fmt.Println("------------------------------------")
	fmt.Print("layers: ",  )
	for k, l := range n {
		if k == 0 { // first layer
			fmt.Printf("[in %d] -> ", len(l.weights[0]))
			fmt.Printf("%d -> ", len(l.weights))
		} else if k == len(n) - 1 { // last layer
			fmt.Printf("out %d\n", len(l.weights))
		} else {
			fmt.Printf("%d -> ", len(l.weights))
		}
	}

	for k, l := range n {
		fmt.Printf("layer %d weights:\n", k)
		l.Print()
	}

	fmt.Println("------------------------------------")

}