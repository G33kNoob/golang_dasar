package calculation

import (
	"errors"
	"fmt"
)

func Add(numberSatu int, numberDua int) int {
	return (numberSatu + numberDua)
}
func Calculate(val1 int, val2 int, art string) (returnVal int, returnErr error) {
	fmt.Println(val1, val2, art)
	// fmt.Println("ok")
	switch art {
	case "+":
		returnVal = val1 + val2
	case "-":
		returnVal = val1 - val2
	case "*":
		returnVal = val1 * val2
	case "/":
		returnVal = val1 / val2
	default:
		returnErr = errors.New("math: aritmatika salah")
	}
	return
}
func Sum(vals ...int) int {
	sum := 0
	for _, val := range vals {
		sum = +val
	}
	return sum
}
