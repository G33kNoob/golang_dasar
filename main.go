package main

import (
	"basic/calculation"
	"fmt"
)

func main() {
	x := []int{10, 20, 30, 4}
	sum := calculation.Sum(x...)
	fmt.Println(sum)
}
