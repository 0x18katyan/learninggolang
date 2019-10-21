package main

import (
	"fmt"
)

func main() {
	foo(2, 123, 12, 312, 31, 23, 12, 312, 3, 123, 12, 312, 3, 123, 12, 312, 3)
}

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i := range x {
		sum += i
	}
	fmt.Println("sum is:", sum)
}
