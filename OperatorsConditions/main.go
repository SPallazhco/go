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

}
