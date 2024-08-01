package main

import (
	"errors"
	"fmt"
)

func main() {
	operation := "/"
	a := 1
	b := 2

	//1
	res := calculate(a, b, operation)

	//2
	printResult(a, b, res, operation)
}

// 1
func calculate(a, b int, opr string) int {
	switch opr {
	case "/":
		//1.1
		return div(a, b)
	case "*":
		//1.2
		return mult(a, b)
	default:
		return 0
	}
}

// 1.1
func div(a, b int) int {
	//1.1.1
	err := validDiv(a, b)
	if err != nil {
		return 0
	}

	return a / b
}

// 1.1.1
func validDiv(a, b int) error {
	if b == 0 {
		return errors.New("b not be  zero")

	}
	return nil
}
func mult(a, b int) int {
	return a * b
}

func printResult(a, b, res int, operation string) {
	fmt.Printf("%v %v %v = %v", a, operation, b, res)
}

