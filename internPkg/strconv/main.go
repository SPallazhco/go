package main

import (
	"fmt"
	"strconv"
)

func main() {
	p := fmt.Println
	s := strconv.Itoa(-42)
	p(s)

	s = strconv.FormatBool(true)
	p(s)
	s = strconv.FormatFloat(3.1415, 'E', -1, 64)
	p(s)
	s = strconv.FormatInt(-42, 16)
	p(s)
	s = strconv.FormatUint(42, 16)
	p(s)

	b, err := strconv.ParseBool("tue")
	fmt.Println(b, err)
	f, err := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f, err)
	in, err := strconv.ParseInt("-42", 10, 64)
	fmt.Println(in, err)
	u, err := strconv.ParseUint("42", 10, 64)
	fmt.Println(u, err)

}
