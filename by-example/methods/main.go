package main

import (
	"fmt"
)

type dog struct {
	first  string
	last   string
	breed  string
	age    int
	weight int
}

// The method

func (d dog) fullname() string {
	return d.first + d.last
}

func main() {
	d1 := dog{"Charlie", "Word", "na", 1, 25}
	d2 := dog{"Rocco", "Jansson", "pitbull", 10, 90}

	fmt.Println(d1.fullname())
	fmt.Println(d2.fullname())

}
