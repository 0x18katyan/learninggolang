package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := sum("The sum of xi is ", xi...)
	fmt.Println(x)
}

func sum(y string, x ...int) string {
	fmt.Println(x)
	sum := 0
	for i := range x {
		sum += i
	}
	return fmt.Sprintf("%v%v.", y, sum)
}

func othermain() {
	xi := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(xi)
}
