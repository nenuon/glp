package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	fmt.Println(max(1, 2, 3, 4, 5, -1, -2, -3, -4, -5))
	fmt.Println(min(5, 4, 3, 2, 1, -1, -2, -3, -4, -5))
	fmt.Println(min())
}

func max(vals ...int) int {
	if len(vals) == 0 {
		log.Fatalln("max: please at least 1 input")
		os.Exit(1)
	}
	maxVal := math.MinInt32
	for _, val := range vals {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func min(vals ...int) int {
	if len(vals) == 0 {
		log.Fatalln("min: please at least 1 input")
		os.Exit(1)
	}
	minVal := math.MaxInt32
	for _, val := range vals {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}
