package main

import (
	"fmt"
)

func main() {
	yearsOld := 27
	fmt.Println("OPERATORS")
	fmt.Println(yearsOld > 27)
	fmt.Println(yearsOld < 7)
	fmt.Println(yearsOld >= 27)
	fmt.Println(yearsOld >= 37, "\n")

	fmt.Println(yearsOld > 32 || yearsOld == 27)
	fmt.Println(yearsOld > 32 || yearsOld == 27, "\n")

	fmt.Println(yearsOld < 32 && yearsOld == 27)
	fmt.Println(yearsOld < 32 && yearsOld == 29, "\n")

	fmt.Println(true)
	fmt.Println(!true)
	fmt.Println(!(yearsOld == 27), "\n")

	fmt.Println(yearsOld < 25 && yearsOld == 32 || yearsOld < 40)
	fmt.Println(yearsOld < 25 && (yearsOld == 32 || yearsOld < 40))

	fmt.Println("CONDITIONS")
	var myYearsOld int = 0
	myYearsOld = 28
	if myYearsOld == 27 {
		fmt.Printf("My age is: %d", myYearsOld)
	} else {
		fmt.Println("It isn't my age")
	}

	if value := true; value {
		fmt.Println("Is true")
	}

	// Bad Practice
	number := 3
	if number == 1 {
		fmt.Println("Is one")
	} else if number == 2 {
		fmt.Println("Is two")
	} else if number == 3 {
		fmt.Println("Is three")
	}
	fmt.Println("CASE")
	switch number {
	case 1:
		fmt.Println("Is one")
	case 2:
		fmt.Println("Is two")
	case 3:
		fmt.Println("Is three")
	default:
		fmt.Println("Is undefined")
	}
	fmt.Println("CASE2")
	switch number := 1; number {
	case 1:
		fmt.Println("Is one")
	case 2:
		fmt.Println("Is two")
	case 3:
		fmt.Println("Is three")
	default:
		fmt.Println("Is undefined")
	}
	fmt.Println("CASE3")
	switch {
	case number > 3:
		fmt.Println("Is one")
	case number < 1:
		fmt.Println("Is two")
	case number == 3:
		fmt.Println("Is three")
	default:
		fmt.Println("Is undefined")
	}

}
