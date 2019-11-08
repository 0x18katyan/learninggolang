package main

import (
	"fmt"
)

func main() {
	x := foo1()
	fmt.Printf("%T\n", x)
	fmt.Println(x(), x(), x())

	y := foo1()
	fmt.Println(y(), y(), y(), y(), y())

}

func foo1() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
