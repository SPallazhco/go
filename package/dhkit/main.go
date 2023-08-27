package main

import "fmt"

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Mail      string `json:"mail"`
	Age       string `json:"Age"`
}

func main() {
	u := User{
		FirstName: "Sergio",
		LastName:  "Pallazhco",
		Mail:      "spallazhco@edu.ec",
		Age:       "26",
	}

	fmt.Println(u)
}
