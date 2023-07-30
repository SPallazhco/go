package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n*****ARRAYS*****\n")
	m1 := make(map[int]string)
	m1[1] = "A"
	m1[2] = "B"
	m1[3] = "C"
	m1[4] = "D"
	fmt.Println(m1)
	fmt.Println(m1[2])
	delete(m1, 2)
	fmt.Println(m1)
	m1[1] = "a"
	fmt.Println(m1)
	m1[7] = ""
	fmt.Println(m1[7])

	v, ok := m1[7]
	v1, ok1 := m1[111]
	_, ok2 := m1[999] // "_" : It is meaning that I never going to use that value.
	fmt.Println("Value: ", v, "It's present: ", ok)
	fmt.Println("Value: ", v1, "It's present: ", ok1)
	fmt.Println("It's present: ", ok2)

	m2 := map[int]string{
		1: "H",
		2: "I",
		3: "J",
	}
	fmt.Println(m2)

	fmt.Println("\n*****BUCLES*****\n")

	sum := 0

	for i := 0; i < 10; i++ {
		sum += 1 // likewise: sum++
	}

	fmt.Println(sum)

	for sum < 1000 {
		sum += 1
	}
	fmt.Println(sum)

	for {
		if sum == 10000 {
			break
		}
		sum += 1
	}
	fmt.Println(sum)

	arrayVar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range arrayVar {
		fmt.Println("Index: ", i, "\tValue: ", arrayVar[i])
	}

	fmt.Println()
	for i, v := range arrayVar {
		fmt.Println("Index: ", i, "\tValue: ", v)
	}

	fmt.Println()
	for _, v := range arrayVar {
		fmt.Println("\tValue: ", v)
	}
	fmt.Println()

	mapaFloat := make(map[int]float64)
	mapaFloat[0] = 0.2
	mapaFloat[1] = 1.3
	mapaFloat[2] = 2.4
	mapaFloat[3] = 3.5
	for key, value := range mapaFloat {
		fmt.Println("Key: ", key, "\tValue: ", value)
	}
	fmt.Println()

	mapaFloat1 := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
	}
	for key, value := range mapaFloat1 {
		fmt.Println("Key: ", key, "\tValue: ", value)
	}
	fmt.Println()

	mapaArrays := map[string][]int{
		"a": nil,
		"b": {1, 3, 5, 7},
		"c": {2, 4, 6, 8},
	}

	for key, value := range mapaArrays {
		fmt.Println("Key: ", key, "\tValue: ", value)
		for _, val := range value {
			fmt.Println("\tValue: ", val)
		}
	}
}
