package main

import (
	"fmt"
)

func main() {

	greeting := func(message string) {
		fmt.Println(message)
	}
	greeting("Hello World!")
}
