package main

import (
	"fmt"
)

func foo() func() {
	return func() {
		fmt.Println("This is a returned func().")
	}
}

func main() {
	x := foo()
	x()
}
