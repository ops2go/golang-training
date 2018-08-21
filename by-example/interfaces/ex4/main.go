package main

import (
	"fmt"
	"math"
)

//struct
type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

//method
func (sq Square) area() float64 {
	return sq.side * sq.side
}

func (ci Circle) area() float64 {
	return math.Pi * ci.radius * ci.radius
}

//interface
type Shape interface {
	area() float64
}

//call interface in function
func info(sq Shape) {
	fmt.Println(sq)
	fmt.Println(sq.area())
}

func info2(ci Circle) {
	fmt.Println(ci)
	fmt.Println(ci.area())
}

//assign a variable to the struct and give struct value(s)
//values pass through info function which calls the interface
func main() {
	square1 := Square{5}
	circle1 := Circle{10}
	info(square1)
	info2(circle1)
}
