package main

import (
	"errors"
	"fmt"
)

/* By convention, errors are the last return value and have
type error, a built-in interface.

errors.New constructs a basic error value with the given
 error message.

 A nil value in the error position indicates that there was
 no error. */

func bad1(x int) (int, error) {
	if x == 13 {
		return -1, errors.New("That's an unlucky number")
	}
	return x + 3, nil
}

//struct
type bad1Error struct {
	x    int
	prob string
}

//method
func (e *bad1Error) Error() string {
	return fmt.Sprint("%d - %s, e.x, e.prob")
}

func main() {
	for _, i := range []int{7, 50} {
		if r, e := bad1(i); e != nil {
			fmt.Println("failure", e)
		} else {
			fmt.Println("success", r)
		}
	}

}
