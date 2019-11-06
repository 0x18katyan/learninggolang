package main

import (
	"fmt"
)

func main() {
	s1 := foo()
	fmt.Println(s1)

	s2 := bar()
	fmt.Println(s2)
	fmt.Printf("%T\n", s2)

	x := s2()
	fmt.Printf("%v is of type %T.\n", x, x)

}

func foo() string {
	return "Hello World"
}

func bar() func() int {
	return func() int {
		return 451
	}
}
