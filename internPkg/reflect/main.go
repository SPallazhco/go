package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID        int64  `myLabel:"lb1" myOtherLabel:"lob2"`
	Email     string `myLabel:"lb2"`
	FirstName string `myLabel:"lb3"`
	LastName  string `myLabel:"lb4"`
	Adress    Adress `myLabel:"lb5"`
}

type Adress struct {
	Country string
	State   string
}

func Examine(data interface{}) {

	if reflect.ValueOf(data).Kind() == reflect.Struct {
		v := reflect.ValueOf(data)
		t := reflect.TypeOf(data)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int, reflect.Int64, reflect.Int32:
				myVar := v.Field(i).Int()
				fmt.Printf("Field: %d Type: %T Value: %v\n", i, myVar, myVar)
			case reflect.String:
				myVar := v.Field(i).String()
				fmt.Printf("Field: %d Type: %T Value: %v\n", i, myVar, myVar)
			case reflect.Ptr:
				fmt.Printf("Field: %d  Value: %v\n", i, v.Field(i))
			case reflect.Struct:
				if v.Field(i).Type() == reflect.TypeOf(Adress{}) {
					myVar := v.Field(i).Interface().(Adress)
					fmt.Println(myVar.Country)
					fmt.Println(myVar.State)
					fmt.Printf("Field: %d  Value: %v\n", i, myVar)
				} else {
					fmt.Println("Unsoported", v.Field(i).Type())
				}
			default:
				fmt.Println("Unsoported", v.Field(i).Type())
			}
			c := t.Field(i).Tag
			fmt.Println("TAG    ", c)
			fmt.Println("TAG    ", c.Get("myLabel"))
			fmt.Println()
		}
	} else {
		t := reflect.TypeOf(data)
		v := reflect.ValueOf(data)
		k := t.Kind()
		fmt.Println("Type ", t)
		fmt.Println("Value", v)
		fmt.Println("Kind", k)
	}
}

func main() {
	fmt.Println("REFLECTION")
	myInt := 5
	myPnt := &myInt
	Examine(myInt)
	Examine(myPnt)
	fmt.Println()
	u := User{
		1,
		"sergiohotmail.com",
		"sergio",
		"costa",
		Adress{"Ecuador", "Cuenca"},
	}
	Examine(u)
}
