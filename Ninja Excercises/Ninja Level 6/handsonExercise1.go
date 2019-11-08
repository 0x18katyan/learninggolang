package main

import (
	"fmt"
)

func foo() int {
	return 420
}

func bar() (int, string) {
	return 42, "forty-two"
}

func main() {
	x := foo()
	y, z := bar()
	fmt.Println(x)
	fmt.Println(y, z)
}
