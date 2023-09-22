package main

import "fmt"

func main() {
	var a uint16 = 213
	fmt.Printf("Number %d in binary is %b\n", a, a)
	fmt.Println("\nLeft shift")
	fmt.Printf("Number %d in binary is %b\n", a<<1, a<<1)
	fmt.Printf("Number %d in binary is %b\n", a>>2, a>>2)

	fmt.Println("a and b value: ")
	var b uint16 = 20
	fmt.Printf("a: %.3d - %.10b\n", a, a)
	fmt.Printf("b: %.3d - %.10b\n", b, b)

	fmt.Printf("Bitwise AND: %d - %.10b\n", a&b, a&b)
	fmt.Printf("Bitwise OR: %d - %.10b\n", a|b, a|b)
	fmt.Printf("Bitwise NOT: %d - %.10b\n", ^a, ^a)
	fmt.Printf("Bitwise XNOT: %d - %.10b\n", a^b, a^b)
	fmt.Printf("Bitwise NAND: %d - %.10b\n", ^(a & b), ^(a & b))
	fmt.Printf("Bitwise NOR: %d - %.10b\n", ^(a | b), ^(a | b))
	fmt.Printf("Bitwise XNOR: %d - %.10b\n", ^(a ^ b), ^(a ^ b))

	fmt.Println()

	rol1 := 1
	rol2 := 1 << 1
	rol3 := 1 << 2
	rol4 := 1 << 3

	fmt.Printf("Roles binary: %.4b, %.4b, %.4b & %.4b\n", rol1, rol2, rol3, rol4)
	fmt.Printf("Roles binary: %.4d, %.4d, %.4d & %.4d\n", rol1, rol2, rol3, rol4)

	myprofile := rol1 + rol4
	fmt.Printf("My Profile: (%d, %b) \n", myprofile, myprofile)

	fmt.Println()

	fmt.Println(0 != myprofile&rol1)
}
