//// Neural Network from Scratch in Go
// This project is meant to demonstrate the fundamentals of deep learning by simulating
// a multilayer network of artificial neurons. It implements the famous backpropagation
// algorithm, and can be trained on various 1-dimensional datasets.
//
// The main function below currently trains a network to replicate the behavior of the
// XOR function. It then plots the loss of the network over time, then print the final
// network's weights and offer example predictions from the network.
// Author: Patrick LeDew

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// XOR function
	// d := dataset{{{0, 0}, {0}}, {{0, 1}, {1}}, {{1, 0}, {1}}, {{1, 1}, {0}}}

	d := dataset{}

	for range 10000 {
		x := rand.Float64() * 10 - 5
		y := rand.Float64() * 10 - 5
		z := sigmoid(3*x/y - 4)
		d.Add([]float64{x, y}, []float64{z})
	}

	// Construct network
	var n network
	// n.AddLayer(2, 2, sigmoid, sigmoid_diff)
	// n.AddLayer(1, 2, sigmoid, sigmoid_diff)
	// n.AddLayer(2, 2, relu, relu_diff)
	// n.AddLayer(1, 2, relu, relu_diff)
	n.AddLayer(4, 2, leaky_relu, leaky_relu_diff)
	n.AddLayer(1, 4, leaky_relu, leaky_relu_diff)


	// Train the network and gather loss values over time
	loss_series := []float64{}
	num_iters := 1000
	loss_thresh := 0.001

	earlyStopping := false

	draw_sem := make(chan struct{}, 1)
	for i := range num_iters {
		x, y := d.Random()
		train_network(n, x, y, 0.01)
		loss := eval(&n, d)
		loss_series = append(loss_series, loss)
		if (i % 10 == 0) {
				go draw_loss_graph(loss_series, draw_sem, i==0)
		}
	

		if (earlyStopping && (loss < loss_thresh)) {
			draw_loss_graph(loss_series, draw_sem, false) // draw graph one last time
			fmt.Println("Stopping early at loss =", loss, " iterations =", i)
			break
		}
		if i == num_iters - 1 {
			draw_loss_graph(loss_series, draw_sem, false) // draw graph one last time
			fmt.Println("Training took full", num_iters, "iterations to complete, loss =", loss)
		}
	}
	// print_graph_autoscale(loss_series, 150, 40, "LOSS")


	n.Print()

	// fmt.Println("in -> out (actual) [loss]")
	// for _, dp := range d {
	// 	out, _, _ := n.Forward(dp[0])
	// 	fmt.Printf("%v -> %v (%v) [%.4f]\n", dp[0], dp[1], out, loss(out, dp[1]))
	// }
}

func draw_loss_graph(series []float64, semaphore chan struct{}, firstDraw bool) {
	semaphore <- struct{}{} // acquire semaphore
	if (!firstDraw) {
		clear_graph(50, 40)
	}
	print_graph_autoscale(series, 150, 20, "LOSS")
	time.Sleep(50 * time.Millisecond)
	<-semaphore // release semaphore
}