package main

import "fmt"

type square struct {
	side float64
}

func (x square) area() float64 {
	return x.side * x.side
}

type shape interface {
	area() float64
}

func info(x shape) {
	fmt.Println(x)
	fmt.Println(x.area())
}

func main() {
	s := square{10}
	fmt.Printf("%T\n", s)
	info(s)
}
