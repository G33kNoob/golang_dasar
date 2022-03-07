package calculation

import "fmt"

func Add(numberSatu int, numberDua int) int {
	return (numberSatu + numberDua)
}
func Calculate(nilai1 int, nilai2 int, aritmatika string) {
	fmt.Println(nilai1, aritmatika, nilai2)
}
func Sum(vals ...int) int {
	sum := 0
	for _, val := range vals {
		sum = +val
	}
	return sum
}
