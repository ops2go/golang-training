package main

import (
	"fmt"
)

func main() {
	n := average(34, 355, 74, 76)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}
	// return total of array divided by length of array
	return total / float64(len(sf))
}
