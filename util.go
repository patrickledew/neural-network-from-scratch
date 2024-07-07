//// Various utility functions

package main

import (
	"fmt"
	"math"
	"os"
)

func assert(statement bool, message string, params ...any) {
	if !statement {
		fmt.Printf("Assert failed, exiting: %s\n", message)
		fmt.Println(params...)
		os.Exit(1)
	}
}

func print_graph_autoscale(series []float64, width int, height int, title string) {
	maxVal := max(series)
	minVal := math.Min(0, min(series))
	print_graph(series, width, height, title, minVal, maxVal)
}

func print_graph(series []float64, width int, height int, title string, minVal float64, maxVal float64) {
	stepSize := (maxVal - minVal) / float64(height)

	draw_graph_title(width, title)

	var steps []float64 // the lower thresholds for each row. If series[y] > steps[y], that is rendered as a block
	for y := range height {
		lowerBound := float64(height - y - 1) * stepSize + minVal
		steps = append(steps, lowerBound)
	}
	for y := range height {
		draw_yaxis_row(height, y, minVal, maxVal)
		graph_row(series, width, y, steps, stepSize)
		fmt.Println(" |")
	}
	// print x axis
	fmt.Printf("   +- %d ", 0)
	for range width - 8 {
		fmt.Print("-")
	}
	fmt.Printf("%6d +\n", len(series))
}

func draw_graph_title(width int, title string) {
	totalDashLength := width - len(title) - 2
	var leftDashLength, rightDashLength int

	if (totalDashLength % 2 == 0) {
		leftDashLength = totalDashLength / 2
		rightDashLength = totalDashLength / 2
	} else {
		leftDashLength = totalDashLength / 2
		rightDashLength = totalDashLength / 2 - 1
	}

	fmt.Print("   +--")
	for range leftDashLength {
		fmt.Print("-")
	}
	fmt.Printf(" %s ", title)
	for range rightDashLength {
		fmt.Print("-")
	}
	fmt.Println("-+")
}

func draw_yaxis_row(height int, y int, minVal float64, maxVal float64) {
	if y == 0 {
		fmt.Printf("%5.2f ", maxVal)
	} else if y == height - 1 {
		fmt.Printf("%5.2f ", minVal)
	} else {
		fmt.Print("   |  ")
	}
}

func graph_row(series []float64, width int, y int, steps []float64, stepSize float64) {
	
	for x := range width {
		idx := int((float64(len(series)) / float64(width)) * float64(x)) // gets index in series based on x coordinate
		nextIdx := int((float64(len(series)) / float64(width)) * float64(x + 1))
	
		sample := series[idx]
		lower := steps[y]
		upper := lower + stepSize
		fraction := (sample - lower) / (upper - lower) // if between upper lower bounds, [0, 1]. Above upper bound = >1, below lower bound = <0
		if (nextIdx < len(series)) {
			// If this is not the last datapoint, look at the NEXT sample to see what it's fraction is
			// If the slope between these is significant, we will render even smaller characters
			nextFraction := (series[nextIdx] - lower) / (upper - lower)
			// if this slope is large (>1 or <-1) we will render a quadrant character
			slope := (nextFraction - fraction)
			
			if (slope > 0.5) { // graph sloping up
				if (fraction < 0.25 && fraction > 1) {
					fmt.Print("▗")
					continue
				} else if (fraction < 0.75 && fraction > 0.5) {
					fmt.Print("▟")
					continue
				}
			} else if (slope < -0.5) { // graph sloping down
				if (fraction > 1 && fraction < 1.25) { // for this fraction we normally wouldn't render anything in this cell, so left half will be empty
					fmt.Print("▙")
					continue
				} else if (fraction > 0.5 && fraction < 0.75) {
					fmt.Print("▖")
					continue
				}
			}
			
		}
		
		if (fraction > 1) {
			fmt.Print("█")
		} else if (fraction > 0.75) {
			fmt.Print("▆")
		} else if (fraction > 0.5) {
			fmt.Print("▄")
		} else if (fraction > 0.25) {
			fmt.Print("▂")
		} else {
			fmt.Print(" ")
		}


		// if (series[idx] >= steps[y]) {
		// 	fmt.Print("█")
		// } else if series[idx] >= steps[y + 1] && len(series) > nextIdx && 0 <= prevIdx && series[nextIdx] >= steps[y] && series[prevIdx] >= steps[y] {
			
		// 	fmt.Print("▂")
		// } else if series[idx] >= steps[y + 1] && len(series) > nextIdx && series[nextIdx] >= steps[y] {
		// 	fmt.Print("▗")
		// } else if series[idx] >= steps[y + 1] && 0 <= prevIdx && series[prevIdx] >= steps[y] {
		// 	fmt.Print("▖")
		// }else {
		// 	fmt.Print(" ")
		// }
	}
}

func max(arr []float64) float64{
	max := math.Inf(-1)
	for _, x := range arr {
		if x > max {
			max = x
		}
	}
	return max
}
func min(arr []float64) float64{
	min := math.Inf(1)
	for _, x := range arr {
		if x < min {
			min = x
		}
	}
	return min
}

func clear_graph(width int, height int) {
	for range height + 2 {
		fmt.Print("\033[F") // move back
		fmt.Print("\033[2M")
	}
	fmt.Print("\033[2M")
}