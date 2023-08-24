package main

import (
	"errors"
	"fmt"

	"github.com/SPallazhco/go/errors/operator"
)

func main() {
	fmt.Println("*** ERRORS ***")
	var err error
	err = errors.New("My new error")
	fmt.Println(err)
	fmt.Println(err.Error()) // Devuelve en string

	err2 := fmt.Errorf("My format err, string: %s, numbre: %d", "my string", 23)
	fmt.Println(err2)
	fmt.Println(err2.Error()) // Devuelve en string

	defer func() { // sin el defer se ejecuta en el orden que esta y con defer al final del programa
		fmt.Println("IN MAIN DEFER")
		r := recover()
		if r != nil {
			fmt.Println("RECOVER IN ", r)
		}
	}()

	// x := 4
	// y := 2
	// z := x / y
	// fmt.Println("z is", z)

	z := operator.Div(4, 0)
	fmt.Println(z)

}
