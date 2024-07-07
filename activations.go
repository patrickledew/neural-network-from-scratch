//// Implementations of various activation functions and their derivatives.

package main

import "math"

func relu(in float64) float64 {
	if in > 0 {
		return in
	} else {
		return 0
	}
}

func relu_diff(in float64) float64 {
	if in > 0 {
		return 1.0
	} else {
		return 0.0
	}
}

func leaky_relu(in float64) float64 {
	if in > 0 {
		return in
	} else {
		return 0.2*in
	}
}

func leaky_relu_diff(in float64) float64 {
	if in > 0 {
		return 1.0
	} else {
		return 0.2
	}
}

func sigmoid(in float64) float64 {
	return 1.0 / (1 + math.Exp(-in))
}

func sigmoid_diff(in float64) float64 {
	return sigmoid(in) * (1 - sigmoid(in))
}

