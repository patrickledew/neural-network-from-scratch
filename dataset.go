//// Simple class to store examples for training the network.

package main

import "math/rand"

type dataset [][][]float64

func (d dataset) Random() (x []float64, y []float64) {
	i := rand.Intn(len(d))
	return d[i][0], d[i][1]
}

func (d *dataset) Add(x []float64, y []float64) {
	*d = append(*d, [][]float64{x, y})
}