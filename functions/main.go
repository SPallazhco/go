package main

import (
	"fmt"

	"github.com/SPallazhco/go/functions/function"
)

func main() {
	value := function.Add(2, 3)
	fmt.Println("Respuesta: ", value)

	function.RepeatString(5, "HOLA MUNDO")

	response, err := function.Calculator(function.SUM, 3, 6)
	fmt.Println("Respuesta: ", response, "\tERROR: ", err)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Value: ", response)
	}

	response1, error1 := function.Calculator(function.DIV, 777, 0)
	fmt.Println("Respuesta: ", response1, "\tERROR: ", error1)

	if error1 != nil {
		fmt.Println(error1.Error())
	} else {
		fmt.Println("Value: ", response1)
	}

	x, y := function.Split(5)
	fmt.Println("Value x: ", x, "Value y: ", y)

	responseMSum := function.MSum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("responseMSum: ", responseMSum)

	fmt.Println()

	responseCalculatorDinamic, error4 := function.CalculatorDinamic(function.MUL, 1, 2, 3, 4, 5, 6)
	fmt.Println("Values: ", responseCalculatorDinamic, "Error: ", error4)

	responseCalculatorDinamic, error4 = function.CalculatorDinamic(function.MUL)
	fmt.Println("Values: ", responseCalculatorDinamic, "Error: ", error4)

	responseCalculatorDinamic, error4 = function.CalculatorDinamic(function.DIV, 1, 2, 0, 4, 5, 6)
	fmt.Println("Values: ", responseCalculatorDinamic, "Error: ", error4)

	fmt.Println()
	fn := function.FactoryOperation(function.MUL)
	v := fn(4, 6)
	fmt.Println(v)

	function1 := function.FactoryOperation(function.MUL)
	val1 := function1(2, 3)
	fmt.Println("Response Factory: ", val1)

	function := function.FactoryOperation(function.DIV)
	val := function(2, 3)
	fmt.Println("Response Factory: ", val)

}
