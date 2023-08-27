package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id1 := uuid.New().String()
	fmt.Println(id1)

	id2 := uuid.New()
	fmt.Println(id2)
	fmt.Println(id2.Version())
}
