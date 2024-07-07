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


func print_graph(series []float64, width int, height int, title string, rng ...float64) {
	var maxVal, minVal float64
	
	if (len(rng) >= 2) {
		minVal = rng[0]
		maxVal = rng[1]
	} else {
		maxVal = max(series)
		minVal = math.Min(0, min(series))
	}

	stepSize := (maxVal - minVal) / float64(height)

	totalDashLength := width - len(title) - 2
	var leftDashLength, rightDashLength int

	if (totalDashLength % 2 == 0) {
		leftDashLength = totalDashLength / 2
		rightDashLength = totalDashLength / 2
	} else {
		leftDashLength = totalDashLength / 2
		rightDashLength = totalDashLength / 2 - 1
	}
	
	if (leftDashLength + rightDashLength) != totalDashLength {
		fmt.Println("AAAAAAAAAAH")
		os.Exit(1)
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

	for y := range height {

		threshold := float64(height - y) * stepSize + minVal
		lowerThreshold := float64(height - y - 1) * stepSize + minVal
		if y == 0 {
			fmt.Printf("%5.2f ", maxVal)
		} else if y == height - 1 {
			fmt.Printf("%5.2f ", minVal)
		} else {
			fmt.Print("   |  ")
		}
		for x := range width {
			idx := int((float64(len(series)) / float64(width)) * float64(x)) // gets index in series
			nextIdx := int((float64(len(series)) / float64(width)) * float64(x + 1))
			prevIdx := int((float64(len(series)) / float64(width)) * float64(x - 1))
			// check if d > threshold value for line
			if (series[idx] >= threshold) {
				fmt.Print("█")
			} else if series[idx] >= lowerThreshold && len(series) > nextIdx && 0 <= prevIdx && series[nextIdx] >= threshold && series[prevIdx] >= threshold {
				
				fmt.Print("▂")
			} else if series[idx] >= lowerThreshold && len(series) > nextIdx && series[nextIdx] >= threshold {
				fmt.Print("▗")
			} else if series[idx] >= lowerThreshold && 0 <= prevIdx && series[prevIdx] >= threshold {
				fmt.Print("▖")
			}else {
				fmt.Print(" ")
			}
		}

		fmt.Println(" |")
	}
	// print x axis
	fmt.Printf("   +- %d ", 0)
	for range width - 8 {
		fmt.Print("-")
	}
	fmt.Printf("%6d +\n", len(series))
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