package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	x2 := returnFunc()
	fmt.Printf("%T is the type of x2, and the value of x2 is %v after calling it with x2().\n", x2, x2())
}

func foo() string {
	s := bar()
	return fmt.Sprintf("This is returned from foo(): %v.", s())
}

func bar() func() int {
	return func() int {
		return 420
	}
}

func returnFunc() func() int {
	return func() int {
		return 100
	}
}
