package main

import (
	"fmt"
)

//struct
type Friend struct {
	name  string
	phone string
}

//method
func (f Friend) info() (string, string) {
	return f.name, f.phone
}

type Relationship interface {
	info() (string, string)
}

func details(r Relationship) {
	fmt.Println("Contact Info:")
	fmt.Println(r.info())
}

func main() {
	f1 := Friend{"Tyler", "Unknown"}
	f2 := Friend{"Lacee", "970-397-7945"}
	details(f1)
	details(f2)
}
