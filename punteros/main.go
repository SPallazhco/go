package main

import (
	"fmt"
)

type MyStruct struct {
	ID   int
	Name string
}

func (m MyStruct) Set(name string) {
	fmt.Printf("address: %p\n", &m)
	m.Name = name
}

func (m *MyStruct) SetP(name string) {
	fmt.Printf("address: %p\n", &m)
	m.Name = name
}

func main() {
	fmt.Println("**** PUNTEROS ****")
	var i int
	var iP *int
	i = 34
	iP = &i
	fmt.Println("DIRECCION DE MEMORIA i: ", &i)
	fmt.Println(i)

	fmt.Println("DIRECCION HACIA DONDE APUNTA: ", *iP)
	fmt.Println(iP)

	*iP = 1
	fmt.Printf("VALUE I: %d, VALUE POINTER i: %d\n", i, *iP)

	myVar := 30
	fmt.Printf("ADDRES: %p, VALUES: %v\n", &myVar, myVar)
	increment(myVar)
	fmt.Println(myVar)

	incrementP(&myVar)
	fmt.Println(myVar)
	fmt.Println()

	var mySlice []int
	mySlice = append(mySlice, 3, 4, 2)
	fmt.Printf("ADDRES: %p, VALUES: %v\n", &mySlice, mySlice)
	fmt.Printf("ADDRES 1: %p, ADDRES 1: %p, ADDRES 1: %p\n", &mySlice[0], &mySlice[1], &mySlice[2])
	updateSlice(mySlice)
	fmt.Println(mySlice)
	fmt.Println()

	myStruct := &MyStruct{ID: 2, Name: "Test"} // Tambien asi se declara los punteros
	fmt.Printf("address: %p\n", myStruct)
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)
	updateStruct(myStruct)
	fmt.Printf("address: %p\n", myStruct)
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)

	fmt.Println()

	myStruct.Set("Metodo set")
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)

	myStruct.SetP("Metodo setP")
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)

}

func increment(val int) {
	fmt.Println(&val)
	val++
}

func incrementP(val *int) {
	fmt.Println(val)
	*val++
}

func updateSlice(v []int) {
	fmt.Printf("address: %p", v)
	fmt.Printf("address 1: %p, address 2: %p, address 3: %p\n", &v[0], &v[1], &v[2])
	v[0] = 12
	v = append(v, 9, 9)
}

func updateStruct(v *MyStruct) {
	fmt.Printf("address: %p\n", v)
	v.ID = 99
	v.Name = "Update"
}
