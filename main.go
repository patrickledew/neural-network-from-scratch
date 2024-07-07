package main

import "fmt"

func main() {
	// XOR function
	d := dataset{{{0, 0}, {0}}, {{0, 1}, {1}}, {{1, 0}, {1}}, {{1, 1}, {0}}}


	// Construct network
	var n network
	n.AddLayer(2, 2, sigmoid, sigmoid_diff)
	n.AddLayer(1, 2, sigmoid, sigmoid_diff)


	// Train the network and gather loss values over time
	loss_series := []float64{}
	num_iters := 1000000
	loss_thresh := 0.001
	for i := range num_iters {
		x, y := d.Random()
		train_network(n, x, y, 0.01)
		loss := eval(&n, d)
		loss_series = append(loss_series, loss)
		if (loss < loss_thresh) {
			fmt.Println("Stopping early at loss =", loss, " iterations =", i)
			break
		}
		if i == num_iters - 1 {
			fmt.Println("Training took full", num_iters, "iterations to complete, loss =", loss)
		}
	}
	print_graph(loss_series, 150, 20, "LOSS")

	n.Print()

	for _, dp := range d {
		out, _, _ := n.Forward(dp[0])
		fmt.Printf("%v -> %v (%v) [%.4f]\n", dp[0], dp[1], out, loss(out, dp[1]))
	}
}