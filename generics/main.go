package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {

	v1 := []float64{1.3, 5.45, 12.223, 6.92, 78.102}
	v2 := []int32{9, 23, 1, 23, 8, 90}

	fmt.Println(Sum(v1))
	fmt.Println(Sum(v2))

	fmt.Println(Sum02(v1))
	fmt.Println(Sum02(v2))

	anyType(1, 2)
	anyType("3", "4")
	// anyType("5", 5) Esto esta mal

	anyType2("3", true)
	comparableType("3", "3")
	orderValues(1, 2)

	csInt := CustomeSlice[int]{1, 2, 3, 4, 5}
	csStr := CustomeSlice[string]{"a", "b", "c"}
	fmt.Println(csInt)
	fmt.Println(csStr)
	fmt.Println()

	x, y := Point(5), Point(2)
	vv := MinNumber(x, y)
	fmt.Println(vv)
	fmt.Println()

	fd := MyGenericStruct[MyFirstData]{StrValu: "Test", Data: MyFirstData{}}
	fd.Data.PrintOne()

	fd2 := MyGenericStruct[MySecondData]{StrValu: "Test", Data: MySecondData{}}
	fd2.Data.PrintSecond()

}

func Sum[N int | int32 | int64 | float64 | float32](v []N) N {
	var total N
	for _, tV := range v {
		total += tV
	}
	return total
}

type Number interface {
	int | int32 | int64 | float64 | float32
}

func Sum02[N Number](v []N) N {
	var total N
	for _, tV := range v {
		total += tV
	}
	return total
}

func anyType[N any](v1, v2 N) {
	fmt.Println(v1, v2)
	// fmt.Println(v1 == v2) Si es any no se puede compara
}

func anyType2[N1 any, N2 any](v1 N1, v2 N2) {
	fmt.Println(v1, v2)
}

func comparableType[N comparable](v1, v2 N) {
	fmt.Println(v1 == v2)
	// fmt.Println(v1 > v2) Esto no es valido solo comparar
}

func orderValues[N constraints.Ordered](v1, v2 N) {
	fmt.Println(v1 > v2)
	fmt.Println(v1 < v2)
	fmt.Println(v1 == v2)
}

type CustomeSlice[V int | string] []V

func MinNumber[T N1](x, y T) T {
	if x < y {
		return x
	}
	return y
}

type Point int

type N1 interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type (
	MyGenericStruct[D any] struct {
		StrValu string
		Data    D
	}

	MyFirstData struct{}

	MySecondData struct{}
)

func (d MyFirstData) PrintOne() {
	fmt.Println("First")
}

func (d MySecondData) PrintSecond() {
	fmt.Println("Second")
}
