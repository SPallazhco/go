package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("GO_ENV")
	s := os.Getenv("MY_ENV1")
	fmt.Println(s)

	godotenv.Load("otherFolder/.var")
	o := os.Getenv("MY_ENV2")
	fmt.Println(o)

	myEnvs, err := godotenv.Read("otherFolder/.var")
	fmt.Println(myEnvs)
	fmt.Println(err)

	if err := godotenv.Overload("newVar/.env2"); err != nil {
		fmt.Println(err)
	}
	oN := os.Getenv("MY_ENV6")
	fmt.Println(oN)

}
