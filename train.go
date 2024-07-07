package main

// mean squared error
func loss(output []float64, expected []float64) float64 {
	mse := 0.0
	for i, yhat := range output {
		y := expected[i]
		mse += loss_single(y, yhat)
	}
	mse /= float64(len(output))
	return mse
}

/**
squared error
y - expected
yhat - actual
**/
func loss_single(y float64, yhat float64) float64 {
	return 1./2. * (y - yhat)*(y - yhat)
}

/////////////// BACKPROP /////////////

func train_network(n network, inputs []float64, expected []float64, lr float64) {
	output, layer_outputs, layer_activations := n.Forward(inputs)
	
	// Create a list for errors in each layer, since it will be filled out back to front
	errs := make([][]float64, len(n)) // errs[k][j] is error of jth node in kth layer

	// calculate errors, backwards from output layer

	// FOR OUTPUT LAYER
	output_layer := n[len(n) - 1]
	for j := range output_layer.weights { // iter over each node in output layer
		y := expected[j]
		yhat := output[j]
		err := (yhat - y) * output_layer.activationFnDiff(layer_activations[len(n)-1][j])
		errs[len(n) - 1] = append(errs[len(n) - 1], err)
	}
	
	// FOR HIDDEN LAYERS
	for k := len(n) - 2; k >= 0; k-- { // kth layer
		lyr := n[k]
		next_lyr := n[k+1]
		for j := range lyr.weights { // jth node in current layer
			node_activation := layer_activations[k][j] // gets the activation of node j in layer k
			err_sum := 0.0
			for l, node_weights := range next_lyr.weights { // iterate over all nodes in NEXT layer and get the weights that connect them to the current node
				w := node_weights[j+1] // j is 0-indexed so add one to skip over bias weight
				err_sum += w * errs[k+1][l] // add to sum
			}
			err := lyr.activationFnDiff(node_activation) * err_sum
			errs[k] = append(errs[k], err)
		}
	}

	// update weights
	for k, lyr := range n {
		for j, node_weights := range lyr.weights {
			for i := range node_weights {
				var input_to_weight float64
				if (i == 0) {
					input_to_weight = 1 // first weight always corresponds to bias
				} else if k == 0 {
					// If first layer, use input rather than prev layer output
					input_to_weight = inputs[i-1]
				} else {
					// Use prev layer output
					input_to_weight = layer_outputs[k-1][i-1]
				}
				// perform gradient descent
				n[k].weights[j][i] -= lr * errs[k][j] * input_to_weight
			}
		}
	}
}

func eval(n *network, d dataset) float64 {
	sum_loss := 0.0
	for i := range d {
		out, _, _ := n.Forward(d[i][0])
		sum_loss += loss(out, d[i][1])
	}
	return sum_loss / float64(len(d))
}