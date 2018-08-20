package main

import (
	"fmt"
)

/* OOP
Encapsulation:
state ("fields")
behavior ("methods")
exported / unexported

Reusability
inheritance (embedded types)

Polymorphism (interfaces)

Overriding
promotion

*/

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"Coleman", "Word", 25}
	p2 := person{"Lacee", "Jansson", 26}

	intro := "The class roster includes:"
	agee := "who is"
	ageee := "years old"

	fmt.Println(intro, p1.first, p1.last, agee, p1.age, ageee)
	fmt.Println(intro, p2.first, p2.last, agee, p2.age, ageee)
}
