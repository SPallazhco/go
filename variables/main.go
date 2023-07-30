package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var myIntVar int
	myIntVar = -12
	fmt.Println("My Variable int is: ", myIntVar)

	var myUint uint
	myUint = 5
	fmt.Println("My variable uint: ", myUint)

	var myString string
	myString = "Hola mundo..!!!"
	fmt.Println(myString)

	var myBoolean bool
	myBoolean = true
	fmt.Println("My variable bool is: ", myBoolean)

	fmt.Println("My variable address is: ", &myString)

	myIntVar2 := 12
	fmt.Println("My variable is: ", myIntVar2)

	const myStringConst = "Hola soy GO"
	fmt.Println("My const is: ", myStringConst)

	const myIntConst int = 7
	fmt.Println("My const int is", myIntConst)

	// Print with a Format
	fmt.Println()
	var myOIntVar int
	var myEightIntVar int8
	var mySixteenIntVar int16
	var myThirtyTwoIntVar int32

	fmt.Printf("Int default value: %d\n", myEightIntVar)

	fmt.Printf("Type: %T, bytes: %d, bits: %d \n", myEightIntVar, unsafe.Sizeof(myEightIntVar), unsafe.Sizeof(myEightIntVar)*8)
	fmt.Printf("Type: %T, bytes: %d, bits: %d \n", mySixteenIntVar, unsafe.Sizeof(mySixteenIntVar), unsafe.Sizeof(mySixteenIntVar)*8)
	fmt.Printf("Type: %T, bytes: %d, bits: %d \n", myThirtyTwoIntVar, unsafe.Sizeof(myThirtyTwoIntVar), unsafe.Sizeof(myThirtyTwoIntVar)*8)
	fmt.Printf("Type: %T, bytes: %d, bits: %d \n", myOIntVar, unsafe.Sizeof(myOIntVar), unsafe.Sizeof(myOIntVar)*8)

	//Floats
	var myFloat32Var float32
	fmt.Printf("Float Value default: %f \n", myFloat32Var)
	fmt.Printf("Type: %T, bytes: %d, bits: %d \n", myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)
	var myFloat64Var float64
	fmt.Printf("Type: %T, bytes: %d, bits: %d \n", myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)

	var myStringVarInformation string
	fmt.Printf("String Value default: %s \n", myStringVarInformation)

	myStringVarComa := `My string var 
	with multiple
	line`
	fmt.Printf("The variable value is:\t%s", myStringVarComa)

	{
		fmt.Println("// CONVERT //")
		floatVar := 33.11
		fmt.Printf("Type: %T, value: %f\n", floatVar, floatVar)
		floatString := fmt.Sprintf("%.2f", floatVar)
		fmt.Printf("Type: %T, value: %s\n", floatString, floatString)

		intVar := 27
		fmt.Printf("Type: %T, value: %d\n", intVar, intVar)
		intString := fmt.Sprintf("%.2d", intVar)
		fmt.Printf("Type: %T, value: %s\n", intString, intString)

		intVal, error := strconv.ParseInt("243", 0, 64)
		fmt.Printf("Type: %T, value: %d\n", intVal, intVal)
		fmt.Println("Error: \n", error)

		intValErr, error := strconv.ParseInt("24d3", 0, 64)
		fmt.Printf("Type: %T, value: %d\n", intValErr, intValErr)
		fmt.Println("Error: \n", error)

		floatVarConvert, _ := strconv.ParseFloat("-11.27", 64)
		fmt.Printf("Type: %T, Value: %.2f\n", floatVarConvert, floatVarConvert)
	}

	{
		fmt.Println("// Byte //")
		var varByteA byte = 'A'
		var varBytea byte = 'a'
		var varByteR byte = 82
		var varBytes byte = 115
		vector := []byte{23, 65, 34, 90}
		fmt.Println(varByteA)
		fmt.Println(varBytea)
		fmt.Println(varByteR)
		fmt.Println(varBytes)
		fmt.Println(vector)
		fmt.Println("Value in String: ")
		fmt.Println(string(varByteA))
		fmt.Println(string(varBytea))
		fmt.Println(string(varByteR))
		fmt.Println(string(varBytes))
		fmt.Println(string(vector))

		fmt.Println("******ARRAY*****")
		var myArrayVar [6]int
		myArrayVar[1] = 1
		myArrayVar[2] = 2
		myArrayVar[3] = 3
		fmt.Println("******SLICE*****")
		var slice1 []int
		fmt.Printf("Size: %d, Value: %v\n", len(slice1), slice1)
		slice1 = append(slice1, 10, 20, 30, 40, 50)
		fmt.Printf("Size: %d, Value: %v\n", len(slice1), slice1)
		mySlice := myArrayVar[2:5]
		fmt.Printf("Size: %d, Value: %v\n", len(mySlice), mySlice)

		fmt.Println(&myArrayVar[2])
		fmt.Println(&mySlice[0])

		fmt.Println(myArrayVar)

		mySlice[0] = 2
		mySlice[1] = 3
		mySlice[2] = 4

		fmt.Printf("Size: %d, Value: %v\n", len(mySlice), mySlice)
		fmt.Println(myArrayVar[:4])
		fmt.Println(myArrayVar[2:])

		slice := make([]int, 3)
		fmt.Println(slice)

	}
}
