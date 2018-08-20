package main

import (
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type hero struct {
	person
	Cover      string
	Superpower string
}

func main() {
	p1 := hero{
		person: person{
			First: "Clark",
			Last:  "Kent",
			Age:   30,
		},
		Cover:      "Superman",
		Superpower: "Invincibility",
	}

	fmt.Println(p1.First, p1.Last, p1.Age, p1.Cover, p1.Superpower)
}
