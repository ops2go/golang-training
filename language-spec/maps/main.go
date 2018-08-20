package main

import (
	"fmt"
)

// key value pairs
func main() {
	age := make(map[string]int)

	age["Coleman"] = 25

	colemanword := age["Coleman"]

	fmt.Println(colemanword)
}
