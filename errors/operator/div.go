package operator

func Div(x, y float64) float64 {
	// defer func() {
	// 	fmt.Println("In my function div defear")
	// }()
	if y == 0 {
		panic("divisor mustn't be zero")
	}

	return x / y
}
