package function

import (
	"errors"
	"fmt"
)

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Add(x int, y int) int {
	return x + y
}

func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Println(value)
	}
}

func Calculator(operation Operation, x, y float64) (float64, error) {
	switch operation {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("y musn't be zero")
		} else {
			return x / y, nil
		}
	case MUL:
		return x * y, nil
	}
	return 0, errors.New("Has been and error")
}

func Split(value int) (x, y int) {
	x = value * 4 / 9
	y = value - x
	return
}

func MSum(values ...float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum
}

func CalculatorDinamic(operation Operation, values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("There aren't values")
	}
	sum := values[0]
	for _, value := range values[1:] {
		switch operation {
		case SUM:
			sum += value
		case SUB:
			sum -= value
		case DIV:
			if value == 0 {
				return 0, errors.New("y musn't be zero")
			} else {
				sum /= value
			}
		case MUL:
			sum *= value
		}
	}
	return sum, nil

}
