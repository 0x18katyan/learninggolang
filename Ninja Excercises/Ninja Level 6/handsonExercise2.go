package main

import (
	"fmt"
)

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func main() {
	j := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("This is the output from foo(): %v\n", foo(j...))
	fmt.Printf("This is the output from bar(): %v\n", bar(j))
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
